package globalstore

import (
	"sync"
)

 
type GlobalStore struct {
    data map[string]interface{}
    mu   sync.RWMutex
}

var store *GlobalStore
var once sync.Once

 
func GetGlobalStore() *GlobalStore {
    once.Do(func() {
        store = &GlobalStore{
            data: make(map[string]interface{}),
        }
    })
    return store
}

 
func (gs *GlobalStore) Set(key string, value interface{}) {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    gs.data[key] = value
}

 
func (gs *GlobalStore) Get(key string) interface{} {
    gs.mu.RLock()
    defer gs.mu.RUnlock()
    return gs.data[key]
}
