package main

import (
	"sync"
	"time"
)

type sp interface {
	write(key string, val interface{})                  // 存入key/val,如果该key读取的goroutine挂起,则唤醒。此方法不会阻塞,时刻都可以立即执行并返回
	read(key string, timeout time.Duration) interface{} // 读取一个key,如果key不存在阻塞,等待key存在或者超时
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func main() {

}

func (m *Map) write(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{value: val, isExist: true}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}
