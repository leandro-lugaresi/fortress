package event

import (
	"context"
	matching "github.com/tylertreat/fast-topic-matching"
)

type (
	// Hub is a component that provide publish and subscribe to messages or events using one topic (like rabbitMQ topic exchanges).
	Hub struct {
		matcher matching.Matcher
	}

	// UnsubscribeFunc tells the hub to unsubscribe the function and abandon its work.
	// A UnsubscribeFunc does not wait for the work to stop.
	UnsubscribeFunc func()

	// SubscribeFunc is 
	SubscribeFunc func(ctx context.Context, e Event)
)

func (h *Hub) Publish(topic string, e Event) {

}

func (h *Hub) On(topic string, fn func(ctx context.Context, e Event)) UnsubscribeFunc {

}

// receiver is used to receive Events by some way.
// The receiver can be async, sync and can drop events.
type receiver interface {
	Next() (Event, ok bool)
}
