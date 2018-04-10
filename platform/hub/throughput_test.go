package hub

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/tylertreat/fast-topic-matching"
)

const (
	numSubs = 1000
	numMsgs = 1000000
)

var (
	topics = make([]string, numSubs)
	msgs   = make([]SimpleEvent, numMsgs)
)

func init() {
	for i := 0; i < numSubs; i++ {
		if i%10 == 0 {
			topics[i] = fmt.Sprintf("*.%d.%d", rand.Intn(10), rand.Intn(10))
		} else if i%25 == 0 {
			topics[i] = fmt.Sprintf("%d.*.%d", rand.Intn(10), rand.Intn(10))
		} else if i%45 == 0 {
			topics[i] = fmt.Sprintf("%d.%d.*", rand.Intn(10), rand.Intn(10))
		} else {
			topics[i] = fmt.Sprintf("%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10))
		}
	}
	for i := 0; i < numMsgs; i++ {
		topic := topics[i%numSubs]
		msgs[i] = SimpleEvent{
			Name: strings.Replace(topic, "*", strconv.Itoa(rand.Intn(10)), -1),
		}
	}
}

func TestThroughput(t *testing.T) {
	h := Hub{
		matcher: matching.NewCSTrieMatcher(),
	}

	for _, topic := range topics {
		sub := NewBlockingSubscriber()
		if _, err := h.Subscribe(topic, sub); err != nil {
			t.Fatal(err)
		}
		go func(s Subscriber) {
			for {
				_, ok := s.Next()
				if !ok {
					return
				}
			}
		}(sub)
	}

	before := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for _, msg := range msgs {
				msgv := msg
				h.Publish(&msgv)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	dur := time.Since(before)
	throughput := numMsgs * 10 / dur.Seconds()
	fmt.Printf("%f msg/sec\n", throughput)
}

// 262372.533926 msg/sec 1 producer
// 830439.529838 msg/sec 20 producers
// 843773.151216 msg/sec 10 producer
