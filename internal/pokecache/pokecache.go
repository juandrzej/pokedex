package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cashe struct {
	list map[string]cacheEntry
	mu   sync.Mutex
}
