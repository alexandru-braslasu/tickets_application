package main

import (
	"fmt"
	"www.TrainStation.com/tickets_application/cmdmanager"
	"www.TrainStation.com/tickets_application/tickets"
	"www.TrainStation.com/tickets_application/utilities"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	var cmd string
	fmt.Println("Choose command: ")
	fmt.Scan(&cmd)

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		var priceJob *tickets.Ticket
		switch(cmd) {
		case "ticket":
			fm := utilities.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
			priceJob = tickets.NewTaxIncludedPriceJob(fm, taxRate)
		case "command":
			cm := cmdmanager.New()
			priceJob = tickets.NewTaxIncludedPriceJob(cm, taxRate)
		default:
			panic("wrong command")
		}
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}