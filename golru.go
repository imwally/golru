package main

import (
	"fmt"
	"time"
)

type Item struct {
	Name     string
	Accessed time.Time
}

type Cache struct {
	Data [5]Item
}

func (c *Cache) Add(name string) {
	// The LRU item found
	var lru *Item

	// Temp variable to hold the last checked LRU access time
	tmp := time.Now().UnixNano()

	// Loop through items in cache updating tmp and lru if the access time was
	// lower than the last checked item
	for i := range c.Data {
		accessed := c.Data[i].Accessed
		if accessed.UnixNano() < tmp {
			tmp = accessed.UnixNano()
			lru = &c.Data[i]
		}
	}

	// Add item to cache in the LRU slot
	lru.Name = name
	lru.Accessed = time.Now()
}

func (c *Cache) Get(name string) {
	// Search for item in cache
	for i := range c.Data {
		if c.Data[i].Name == name {
			// Update access time if found
			c.Data[i].Accessed = time.Now()
			return
		}
	}

	// Otherwise add item if not found
	c.Add(name)
}

func (c *Cache) Print() {
	for _, i := range c.Data {
		if i.Name != "" {
			fmt.Println(i.Name, "\t", i.Accessed)
		}
	}
	fmt.Println()
}

func main() {
	var c Cache

	c.Get("a")
	c.Get("b")
	c.Get("c")
	fmt.Println("Get: a, b, c")
	c.Print()

	time.Sleep(2 * time.Second)
	c.Get("a")
	fmt.Println("Get: a")
	c.Print()

	time.Sleep(2 * time.Second)
	c.Get("d")
	fmt.Println("Get: d")
	c.Print()

	time.Sleep(2 * time.Second)
	c.Get("a")
	fmt.Println("Get: a")
	c.Print()

	fmt.Println()
	time.Sleep(2 * time.Second)
	c.Get("e")
	fmt.Println("Get: e")
	c.Print()

	fmt.Println()
	time.Sleep(2 * time.Second)
	c.Get("f")
	fmt.Println("Get: f")
	c.Print()

	fmt.Println()
	time.Sleep(2 * time.Second)
	c.Get("g")
	fmt.Println("Get: g")
	c.Print()
}
