package converter

import "fmt"

func Convert(amount float64, from string, to string, rates map[string]float64) (float64, error) {
	fromRate, exists := rates[from]
	if !exists {
		return 0, fmt.Errorf("currency %s is not supported", from)
	}

	toRate, exists := rates[to]
	if !exists {
		return 0, fmt.Errorf("currency %s is not supported", to)
	}

	return (amount / fromRate) * toRate, nil
}
