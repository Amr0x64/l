package main

import (
	"fmt"
	"math/rand"
	"sync"
)
func main() {
	c := Cache{
		Data: make(map[int]int),
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			c.Set(rand.Int(), rand.Int())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.Data)
}

// Возьмем как  пример кеш и используем в ней map
func (c *Cache) Set(key int, value int) {
	c.Lock()
	defer c.Unlock()
	c.Data[key] = value
}

//ИЗ за частой записи используем Mutex, вместо RWmutex(который бы решил также нашу задачу, но будет излишен) 
type Cache struct {
	sync.Mutex             
	Data       map[int]int 
}