//Copyright (C) 2018 Tyler Treat <https://github.com/tylertreat>
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

//	http://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

//Modifications copyright (C) 2018 Leandro Lugaresi

package hub

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEqual(assert *assert.Assertions, expected, actual []Subscriber) {
	assert.Len(actual, len(expected))
	for _, sub := range expected {
		assert.Contains(actual, sub)
	}
}

func populateMatcher(m Matcher, num, topicSize int) {
	for i := 0; i < num; i++ {
		prefix := ""
		topic := ""
		for j := 0; j < topicSize; j++ {
			topic += prefix + strconv.Itoa(rand.Int())
			prefix = "."
		}
		m.Subscribe(topic, Subscriber(topic))
	}
}

func benchmark5050(b *testing.B, numItems, numThreads int, factory func([][]string) Matcher) {
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
	sub := Subscriber("abc")
	m := factory(itemsToInsert)
	populateMatcher(m, 1000, 5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(numThreads)
		for j := 0; j < numThreads; j++ {
			go func(j int) {
				if j%2 != 0 {
					for _, key := range itemsToInsert[j] {
						m.Subscribe(key, sub)
					}
				} else {
					for _, key := range itemsToInsert[j] {
						m.Lookup(key)
					}
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
}

func benchmark9010(b *testing.B, numItems, numThreads int, factory func([][]string) Matcher) {
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
	sub := Subscriber("abc")
	m := factory(itemsToInsert)
	populateMatcher(m, 1000, 5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(numThreads)
		for j := 0; j < numThreads; j++ {
			go func(j int) {
				if j%10 == 0 {
					for _, key := range itemsToInsert[j] {
						m.Subscribe(key, sub)
					}
				} else {
					for _, key := range itemsToInsert[j] {
						m.Lookup(key)
					}
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
}
