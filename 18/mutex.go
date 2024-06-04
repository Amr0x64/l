package main

import "sync"

// счетчик с использование мьютексов
type mutexInc struct {
	counter int64
	mutex   *sync.RWMutex
}

func newMutex() *mutexInc {
	return &mutexInc{
		counter: 0,
		mutex:   &sync.RWMutex{},
	}
}

func (m *mutexInc) Add() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.counter++
}

func (m *mutexInc) Get() int64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.counter
}