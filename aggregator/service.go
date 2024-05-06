package main

import (
	"fmt"
	"tolling/ctypes"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(ctypes.Distance) error
	CalculateInvoice(int) (*ctypes.Invoice, error)
}

type Storer interface {
	Insert(ctypes.Distance) error
	Get(int) (float64, error)
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

func (i *InvoiceAggregator) CalculateInvoice(obuID int) (*ctypes.Invoice, error) {
	dist, err := i.store.Get(obuID)
	if err != nil {
		return nil, err
	}
	inv := &ctypes.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		TotalAmount:   basePrice * dist,
	}
	return inv, nil
}
