package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func InitializeCache() {
	Cache = cache.New(cache.NoExpiration, time.Minute*10)
}
