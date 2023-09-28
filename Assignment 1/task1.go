package main

import "fmt"

type EvictionAlgo interface {
	Evict(c *Cache)
}

type Fifo struct{}

func (f *Fifo) Evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
	delete(c.storage, c.keys[0])
	c.keys = c.keys[1:]
}

type Lru struct{}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
	delete(c.storage, c.keys[len(c.keys)-1])
	c.keys = c.keys[:len(c.keys)-1]
}

type Lfu struct{}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
	var leastFreq int = c.counts[c.keys[0]]
	var keyToEvict string = c.keys[0]
	for _, key := range c.keys {
		if c.counts[key] < leastFreq {
			leastFreq = c.counts[key]
			keyToEvict = key
		}
	}
	delete(c.storage, keyToEvict)
	c.keys = removeElement(c.keys, keyToEvict)
	c.counts[keyToEvict] = 0
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
	keys         []string
	counts       map[string]int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	keys := make([]string, 0)
	counts := make(map[string]int)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
		keys:         keys,
		counts:       counts,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
	c.keys = append(c.keys, key)
	c.counts[key] = 1
}

func (c *Cache) get(key string) (string, bool) {
	if value, ok := c.storage[key]; ok {
	  delete(c.storage, key)
	  c.keys = removeElement(c.keys, key)
	  c.counts[key] = 0
	  return value, true
	}
	return "", false
  }

func (c *Cache) evict() {
	c.evictionAlgo.Evict(c)
	c.capacity--
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
}

func removeElement(s []string, element string) []string {
	index := -1
	for i, v := range s {
		if v == element {
			index = i
			break
		}
	}
	if index == -1 {
		return s
	}
	return append(s[:index], s[index+1:]...)
}