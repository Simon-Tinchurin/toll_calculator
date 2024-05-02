package main

import "tolling/ctypes"

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Insert(d ctypes.Distance) error {
	return nil
}
