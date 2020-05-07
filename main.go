package main

import (
	"container/list"
	"fmt"
)

func main() {
	c := NewLRUCache(3)

	for i := 0; i < 10; i++ {
		c.Put(i, 0)
		c.PrintQueue()
		c.PrintMap()
	}

	fmt.Println(c.Get(1))
	fmt.Println(c.Get(9))
}

type LRUCache struct {
	queue *list.List
	cache map[int]*list.Element
	len   int
}

type QueueItem struct {
	Key   int
	Value int
}

func NewLRUCache(len int) *LRUCache {
	return &LRUCache{
		queue: list.New(),
		cache: make(map[int]*list.Element, len),
		len:   len,
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	val, ok := c.cache[key]
	if !ok {
		return 0, false
	}

	c.queue.MoveToFront(val)
	return val.Value.(QueueItem).Value, true
}

func (c *LRUCache) Put(key, val int) {
	v, ok := c.cache[key]
	if !ok {
		if c.len == c.queue.Len() {
			last := c.queue.Back()
			c.queue.Remove(last)
			qi := last.Value.(QueueItem)
			delete(c.cache, qi.Key)
		}
	} else {
		c.queue.Remove(v)
	}

	e := c.queue.PushFront(QueueItem{Key: key, Value: val})
	c.cache[key] = e
}

func (c *LRUCache) PrintQueue() {
	fmt.Println("=========QUEUE===========")
	for e := c.queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("********************")
}

func (c *LRUCache) PrintMap() {
	fmt.Println("========MAP============")
	for k, v := range c.cache {
		fmt.Println("key:", k, "val:", v.Value.(QueueItem).Value)
	}
	fmt.Println("********************")
}
