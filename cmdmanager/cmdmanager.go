package cmdmanager

import "fmt"

type Command struct {
}

func (cmd Command) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd Command) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() Command {
	return Command{}
}

func (Command) GetCapitalFrom() string {
	return ""
}

func (Command) GetCapitalTo() string {
	return ""
}