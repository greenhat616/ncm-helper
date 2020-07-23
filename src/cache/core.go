// Package cache is a in memory, key-value, cache impl based on go-cache(github.com/patrickmn/go-cache)
package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// C is the cache instance which is exported to serve
var C = cache.New(5*time.Minute, 10*time.Minute)
