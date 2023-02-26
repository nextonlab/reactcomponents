package crud

import (
	"math/rand"
	"sync"
)

type idList struct {
	sync.RWMutex
	ids       []int64
	positions map[int64]int
	maxID     int64
}

func newIDList() *idList {
	return &idList{
		positions: make(map[int64]int),
	}
}

func (l *idList) allocID() int64 {
	l.Lock()
	defer l.Unlock()

	l.maxID++
	return l.maxID
}

func (l *idList) len() int {
	l.RLock()
	defer l.RUnlock()

	return len(l.ids)
}

func (l *idList) pushID(id int64) {
	l.Lock()
	defer l.Unlock()

	l.positions[id] = len(l.ids)
	l.ids = append(l.ids, id)
}

func (l *idList) popID(id int64) {
	l.Lock()
	defer l.Unlock()

	pos, ok := l.po