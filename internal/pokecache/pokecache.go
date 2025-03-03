package pokecache

import (
	"sync"
	"time"

	"github.com/semidesnatada/pokedex/internal/pokeapi"
)

type Cache struct {
	Content map[string]cacheEntry
	Interval time.Duration
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val pokeapi.LocationListResponseFormat
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Content: map[string]cacheEntry{},
		Interval: interval,
		mu: &sync.Mutex{},
	}

	ticker := time.NewTicker(5 * time.Minute)
	go cache.reapLoop(ticker)
	return cache
}

func (c *Cache) Add(key string, item pokeapi.LocationListResponseFormat) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Content[key] = cacheEntry{
		createdAt: time.Now(),
		val:item,
	}
}

func (c *Cache) Get(key string) (pokeapi.LocationListResponseFormat, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.Content[key]
	if !ok {
		return pokeapi.LocationListResponseFormat{}, false //errors.New("no item with this name in cache")
	} else {
		return item.val, true
	}
}

func (c *Cache) reapLoop(ticker *time.Ticker) {
	for {
		c.mu.Lock()
		for key, value := range c.Content {
			if time.Since(value.createdAt) > c.Interval {
				delete(c.Content, key)
			}
		}
		c.mu.Unlock()
		<- ticker.C
		
	}
}