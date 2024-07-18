package cmdmanager

import (
	"errors"
	"fmt"
	"regexp"
)

type CMDManager struct{}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmd *CMDManager) ReadLines() ([]string, error) {
	var prices []string

	fmt.Println("Enter prices (`:q` to quit)")
	for {
		var price string
		fmt.Print("> ")
		fmt.Scan(&price)
		// Escape condition
		if price == ":q" {
			break
		}
		// Input validation
		check, _ := regexp.MatchString("[+-]?([0-9]*[.])?[0-9]+", price)
		if !check {
			return nil, errors.New("Invalid input")
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd *CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
