package goro

import (
	"net/http"
)

// RouteCache - temporary storage for routes
type RouteCache struct {
	Entries map[string]cacheEntry
}

type cacheEntry struct {
	Params  map[string]interface{}
	Handler http.Handler
	Route   Route
}

// NewRouteCache - creates a new default RouteCache
func NewRouteCache() *RouteCache {
	return &RouteCache{
		Entries: make(map[string]cacheEntry),
	}
}

// Put - add an item to the route cache
func (rc *RouteCache) Put(path string, entry cacheEntry) {
	rc.Entries[path] = entry
}

// Clear - reset the cache
func (rc *RouteCache) Clear() {
	rc.Entries = make(map[string]cacheEntry)
}
