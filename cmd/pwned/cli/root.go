package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

// Cmd is the root command.
var Cmd = kingpin.New("pwned", "")

func init() {

	workdir := Cmd.Flag("chdir", "Change working directory.").Default(".").Short('C').String()
	verbose := Cmd.Flag("verbose", "Enable verbose log output.").Short('v').Bool()
	format := Cmd.Flag("format", "Output formatter.").Default("text").String()

	Cmd.PreAction(func(ctx *kingpin.ParseContext) error {
		os.Chdir(*workdir)

		if *verbose {
			log.SetHandler(delta.Default)
			log.SetLevel(log.DebugLevel)
			log.Debugf("up version %s", Cmd.GetVersion())
		}

		Init = func() (*up.Config, *up.Project, error) {
			c, err := up.ReadConfig("up.json")
			if err != nil {
				return nil, nil, errors.Wrap(err, "reading config")
			}

			events := make(event.Events)
			p := up.New(c, events).WithPlatform(lambda.New(c, events))

			switch {
			case *verbose:
				go reporter.Discard(events)
			case *format == "plain" || util.IsCI():
				go reporter.Plain(events)
			default:
				go reporter.Text(events)
			}

			return c, p, nil
		}

		return nil
	})
}
