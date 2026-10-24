package internal

import (
	"image"
	"sync"
	"time"
)

type targetImage struct {
	image.Image

	pixelWeights map[uint16]float64
	totalWeight  float64
	lastAccessed time.Time
}

type targetCache struct {
	mu             sync.Mutex
	cache          map[string]*targetImage
	evictOlderThan time.Duration
}

// newTargetCache returns a new targetCache. If evictOlderThan <= 0, cache entries
// are never evicted. 'Clean' method becomes a no-op.
func newTargetCache(evictOlderThan time.Duration) *targetCache {
	if evictOlderThan < 0 {
		evictOlderThan = 0
	}

	return &targetCache{
		cache:          make(map[string]*targetImage),
		evictOlderThan: evictOlderThan,
	}
}

func (c *targetCache) Get(key string) (*targetImage, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	img, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	now := time.Now()
	if c.evictOlderThan > 0 && img.lastAccessed.Before(now.Add(-c.evictOlderThan)) {
		return nil, false
	}

	c.cache[key].lastAccessed = now
	return img, true
}

func (c *targetCache) Set(key string, img image.Image) {
	c.mu.Lock()
	defer c.mu.Unlock()

	weights, totalWeight := getWeights(img)
	c.cache[key] = &targetImage{
		Image:        img,
		pixelWeights: weights,
		totalWeight:  totalWeight,
		lastAccessed: time.Now(),
	}
}

func (c *targetCache) Clean() {
	if c.evictOlderThan <= 0 {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	cache := make(map[string]*targetImage)
	cutoff := time.Now().Add(-c.evictOlderThan)
	for key, tm := range c.cache {
		if tm.lastAccessed.After(cutoff) {
			cache[key] = tm
		}
	}

	c.cache = cache
}
