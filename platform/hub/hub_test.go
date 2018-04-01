package hub

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	matching "github.com/tylertreat/fast-topic-matching"
)

func BenchmarkMultithreaded4Thread9010(b *testing.B) {
	h := Hub{
		matcher: matching.NewCSTrieMatcher(),
	}
	runBenchmark9010(b, 1000, 4, &h)
}

func runBenchmark9010(b *testing.B, numItems, numThreads int, h *Hub) {
	itemsToInsert := make([][]string, 0, numThreads)
	for i := 0; i < numThreads; i++ {
		items := make([]string, 0, numItems)
		for j := 0; j < numItems; j++ {
			topic := strconv.Itoa(j%10) + "." + strconv.Itoa(j%50) + "." + strconv.Itoa(j)
			items = append(items, topic)
		}
		itemsToInsert = append(itemsToInsert, items)
	}

	var wg sync.WaitGroup
	sub := discardSubscriber(0)
	populateHub(h, 1000, 5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(numThreads)
		for j := 0; j < numThreads; j++ {
			go func(j int) {
				if j%10 == 0 {
					for _, key := range itemsToInsert[j] {
						h.Subscribe(key, sub)
					}
				} else {
					for _, key := range itemsToInsert[j] {
						h.Publish(&SimpleEvent{Name: key})
					}
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
}

func populateHub(h *Hub, subscribers, topicSize int) {
	var discard discardSubscriber
	for i := 0; i < subscribers; i++ {
		prefix := ""
		topic := ""
		for j := 0; j < topicSize; j++ {
			topic += prefix + strconv.Itoa(rand.Int())
			prefix = "."
		}
		h.Subscribe(topic, discard)
	}
}
