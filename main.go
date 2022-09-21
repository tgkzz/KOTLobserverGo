package main

func main() {
	// Creation of a ner Stock for monitoring
	stockMonitor := &StockMonitor{
		ticker: "APPL",
		price:  0.0,
	}

	observerA := &StockObserver{
		name: "Stock Observer asd",
	}

	observerB := &StockObserver{
		name: "Stock Observer qwe",
	}

	stockMonitor.Suscribe(observerA)
	stockMonitor.Suscribe(observerB)

	stockMonitor.SetPrice(2389)

	stockMonitor.Unsuscribe(observerA)

	stockMonitor.SetPrice(324981)

}
