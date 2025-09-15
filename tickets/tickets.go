package tickets

import (
	"fmt"
	"www.TrainStation.com/tickets_application/conversion"
	"www.TrainStation.com/tickets_application/supplier"
)

type Ticket struct {
	Supplier supplier.Supplier `json:"-"`
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	From string `json:"start_location"`
	To string `json:"destination"`
}

func (job *Ticket) LoadData() error {
	lines, err := job.Supplier.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func NewTaxIncludedPriceJob(iom supplier.Supplier, taxRate float64) *Ticket {
	return &Ticket{
		Supplier: iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		From: iom.GetCapitalFrom(),
		To: iom.GetCapitalTo(),
	}
}

func (job *Ticket) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	if err != nil {
		// return err
		errorChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	err = job.Supplier.WriteResult(job)
	if err != nil {
		errorChan <- err
		return
	}
	doneChan <- true
}