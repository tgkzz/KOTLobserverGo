package main

import "fmt"

type StockObserver struct {
	name string
}

func (s *StockObserver) Update(t string) {
	fmt.Println("Stock Observer:", s.name, "has been updated, received subject string:", t)
}
