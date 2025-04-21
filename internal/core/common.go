package core

import (
	"os"
	"sync"
)

type ProcessStore struct {
	store map[string]*os.Process
	mu    sync.RWMutex
}

func NewProcessStore() *ProcessStore {
	return &ProcessStore{
		store: make(map[string]*os.Process),
	}
}

func (p *ProcessStore) Set(key string, process *os.Process) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.store[key] = process
}

func (p *ProcessStore) Get(key string) *os.Process {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.store[key]
}

var globalProcessStore *ProcessStore = NewProcessStore()
