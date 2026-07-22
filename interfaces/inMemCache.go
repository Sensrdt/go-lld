package main

import "fmt"

// Question 2: In-Memory Cache (Maps + Interface)

type CacheLogic interface {
	Set(key, value string)
	Get(key string) string
	Delete(key string)
}

type Cache struct {
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	if key == "" || value == "" {
		return
	}
	c.data[key] = value

}

func (c *Cache) Get(key string) string {
	val, ok := c.data[key]
	if !ok {
		return ""
	}
	return val
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func InMemCacheHandler() {
	cache := NewCache()
	fmt.Println(cache.Get("1"))

	cache.Set("1", "one")
	cache.Set("2", "two")

	fmt.Println(cache.Get("2"))
}
