package main

import (
	"fmt"
	"tolling/ctypes"
)

type Aggregator interface {
	AggregateDistance(ctypes.Distance) error
}

type Storer interface {
	Insert(ctypes.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance ctypes.Distance) error {
	fmt.Println("processing and inserting distance in the storage:", distance)
	return i.store.Insert(distance)
}
