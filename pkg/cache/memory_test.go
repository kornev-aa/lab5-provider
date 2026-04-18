package cache

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestMemoryCacheSetGet(t *testing.T) {
    cache := NewMemoryCache()
    key := "test-key"
    value := []byte("test-value")

    cache.Set(key, value, 1*time.Minute)

    got, found := cache.Get(key)
    assert.True(t, found)
    assert.Equal(t, value, got)
}

func TestMemoryCacheExpired(t *testing.T) {
    cache := NewMemoryCache()
    key := "expiring-key"
    value := []byte("value")

    cache.Set(key, value, 1*time.Millisecond)
    time.Sleep(2 * time.Millisecond)

    _, found := cache.Get(key)
    assert.False(t, found)
}

func TestMemoryCacheDelete(t *testing.T) {
    cache := NewMemoryCache()
    key := "delete-key"

    cache.Set(key, []byte("value"), 1*time.Minute)
    cache.Delete(key)

    _, found := cache.Get(key)
    assert.False(t, found)
}
