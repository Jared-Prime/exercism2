package letter

import "sync"

// FreqMap contains the frequency of runes in a string
type FreqMap map[rune]int

// Frequency counts how often each rune occurs in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts how often each rune occurs in a string
func ConcurrentFrequency(strings []string) FreqMap {
	var lock sync.Mutex
	var wg sync.WaitGroup
	m := FreqMap{}

	for _, s := range strings {
		wg.Add(1)
		go func(s string) {
			for r, count := range Frequency(s) {
				lock.Lock()
				m[r] += count
				lock.Unlock()
			}
			wg.Done()
		}(s)
	}

	wg.Wait()

	return m
}
