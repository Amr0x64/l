package main

import "sync/atomic"

// счетчик с использованием атомарных операций
type atomicInc struct {
	counter int64
}

func newAtomic() *atomicInc {
	return &atomicInc{
		counter: 0,
	}
}

func (a *atomicInc) Add() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *atomicInc) Get() int64 {
	return atomic.LoadInt64(&a.counter)
}