package main

import (
	"errors"
	"strconv"
)

type StockMonitor struct {
	ticker string
	price  float64

	observers []Observer
}

func (s *StockMonitor) Suscribe(o Observer) (bool, error) {
	for _, observer := range s.observers {
		if observer == o {
			return false, errors.New("Observer alreay exist")
		}
	}
	s.observers = append(s.observers, o)
	return true, nil
}

func (s *StockMonitor) Unsuscribe(o Observer) (bool, error) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("Observer not found")
}

func (s *StockMonitor) Notify() (bool, error) {
	for _, observer := range s.observers {
		observer.Update(s.String())
	}
	return true, nil
}

func (s *StockMonitor) SetPrice(price float64) {
	s.price = price
	s.Notify()
}

func (s *StockMonitor) String() string {
	convertFloatToString := strconv.FormatFloat(s.price, 'f', 2, 64)
	return "Stock Monitor: " + s.ticker + " $" + convertFloatToString
}
