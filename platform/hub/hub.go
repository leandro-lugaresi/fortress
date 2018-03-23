package hub

import (
	"context"
	matching "github.com/tylertreat/fast-topic-matching"
)

type (
	// Hub is a component that provide publish and subscribe for events using topics (like rabbitMQ topic exchanges).
	Hub struct {
		matcher matching.Matcher
	}

	Publisher interface {
		// Set send the given Event to be processed by the subscriber
		Set(Event)
	}

	Subscriber interface {
		// Next will return the next Event to be processed. If the
		// diode is empty this method will block until a Event is available to be
		// readed or context is done. In case of context done we will return true on the second return param.
		Next() (Event, bool)
	}

	// UnsubscribeFunc tells the hub to unsubscribe the function and abandon its work.
	// A UnsubscribeFunc does not wait for the work to stop.
	UnsubscribeFunc func()

	// SubscribeFunc is the function executed when an subscriber receive one message.
	SubscribeFunc func(ctx context.Context, e Event)
)

// Publish will send an event to all the subscribers matching the event name
func (h *Hub) Publish(e Event) {
	for _, sub := range h.matcher.Lookup(e.Name) {
		s := sub.(Publisher)
		s.Set(e)
	}
}

// On create a subscription to receive events.
func (h *Hub) On(topic string, sub Subscriber) matching.Subscription {
	h.matcher.Subscribe(topic, sub)
}

// receiver is used to receive Events by some way.
// The receiver can be async, sync and can drop events.
type receiver interface {
	Next() (Event, ok bool)
}
