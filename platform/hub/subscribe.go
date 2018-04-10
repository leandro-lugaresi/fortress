package hub

import (
	"context"

	gendiodes "code.cloudfoundry.org/go-diodes"
)

type (
	// nonBlockingSubscriber uses an diode and is optiomal for many writes and a single reader
	// This subscriber is used when need high throughput and losing data is acceptable.
	nonBlockingSubscriber struct {
		d *gendiodes.Poller
	}

	// blockingSubscriber uses an channel to receive events.
	blockingSubscriber struct {
		ch chan Event
	}

	discardSubscriber int
)

// NewNonBlockingSubscriber returns a new NonBlockingSubscriber diode to be used
// with many writers and a single reader.
func NewNonBlockingSubscriber(ctx context.Context, size int, alerter gendiodes.Alerter) Subscriber {
	return &nonBlockingSubscriber{
		d: gendiodes.NewPoller(
			gendiodes.NewManyToOne(size, alerter),
			gendiodes.WithPollingContext(ctx)),
	}
}

// Set inserts the given Event into the diode.
func (d *nonBlockingSubscriber) Set(data Event) {
	d.d.Set(gendiodes.GenericDataType(&data))
}

// Next will return the next Event. If the
// diode is empty this method will block until a Event is available to be
// read or context is done. In case of context done we will return true on the second return param.
func (d *nonBlockingSubscriber) Next() (Event, bool) {
	data := d.d.Next()
	if data == nil {
		return nil, true
	}
	return *(*Event)(data), false
}

// NewBlockingSubscriber returns a new blocking subscriber using chanels imternally.
func NewBlockingSubscriber() Subscriber {
	return &blockingSubscriber{
		ch: make(chan Event),
	}
}

// Set inserts the given Event into the diode.
func (s *blockingSubscriber) Set(event Event) {
	s.ch <- event
}

// Next will return the next Event. If the
// diode is empty this method will block until a Event is available to be
// read or context is done. In case of context done we will return true on the second return param.
func (s *blockingSubscriber) Next() (Event, bool) {
	e, ok := <-s.ch
	return e, ok
}

func (d discardSubscriber) Set(event Event)     {}
func (d discardSubscriber) Next() (Event, bool) { return nil, false }
