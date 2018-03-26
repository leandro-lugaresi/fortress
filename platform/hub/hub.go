package hub

import matching "github.com/tylertreat/fast-topic-matching"

type (
	// Hub is a component that provide publish and subscribe for events using topics (like rabbitMQ topic exchanges).
	Hub struct {
		matcher matching.Matcher
	}

	Publisher interface {
		// Set send the given Event to be processed by the subscriber
		Set(Event)
	}

	Receiver interface {
		// Next will return the next Event to be processed.
		// This method will block until a Event is available to be
		// readed or context is done. In case of context done we will return true on the second return param.
		Next() (Event, bool)
	}

	Subscriber interface {
		Receiver
		Publisher
	}
)

// Publish will send an event to all the subscribers matching the event name
func (h *Hub) Publish(e Event) {
	for _, sub := range h.matcher.Lookup(e.Topic()) {
		s := sub.(Publisher)
		s.Set(e)
	}
}

// Subscribe create a subscription to receive events.
func (h *Hub) Subscribe(topic string, sub Subscriber) (*matching.Subscription, error) {
	return h.matcher.Subscribe(topic, sub)
}

// Unsubscribe remove and close the Subscription.
func (h *Hub) Unsubscribe(sub *matching.Subscription) error {
	return nil
}

// Close will remove and unsubcribe all the subscriptions.
func (*Hub) Close() error {
	return nil
}
