package main

import "tolling/ctypes"

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStore) Insert(d ctypes.Distance) error {
	m.data[d.OBUID] += d.Value
	return nil
}
