package converter

import (
	"testing"
)

type ConvertTestBase struct {
	name     string
	amount   float64
	from     string
	to       string
	rates    map[string]float64
	expected float64
	expectedErr string
}

var tests []ConvertTestBase = []ConvertTestBase{
	{
		name:     "successful conversion",
		amount:   100,
		from:     "USD",
		to:       "EUR",
		rates:    map[string]float64{"USD": 1, "EUR": 0.88},
		expected: 88,
	},
	{
		name:     "from currency not supported",
		amount:   100,
		from:     "GBP",
		to:       "EUR",
		rates:    map[string]float64{"USD": 1, "EUR": 0.88},
		expectedErr: "currency GBP is not supported",
	},
	{
		name:     "to currency not supported",
		amount:   100,
		from:     "USD",
		to:       "JPY",
		rates:    map[string]float64{"USD": 1, "EUR": 0.88},
		expectedErr: "currency JPY is not supported",
	},
	{
		name:     "both currencies not supported",
		amount:   100,
		from:     "GBP",
		to:       "JPY",
		rates:    map[string]float64{"USD": 1, "EUR": 0.88},
		expectedErr: "currency GBP is not supported",
	},
	{
		name:     "zero amount conversion",
		amount:   0,
		from:     "USD",
		to:       "EUR",
		rates:    map[string]float64{"USD": 1, "EUR": 0.88},
		expected: 0,
	},
	{
		name:     "same from and to currencies",
		amount:   100,
		from:     "USD",
		to:       "USD",
		rates:    map[string]float64{"USD": 1, "EUR": 0.88},
		expected: 100,
	},
}

func TestConvert(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.amount, tt.from, tt.to, tt.rates)
			if err != nil {
				if tt.expectedErr == "" {
					t.Errorf("unexpected error: %v", err)
				} else if err.Error() != tt.expectedErr {
					t.Errorf("expected error %q, got %q", tt.expectedErr, err.Error())
				}
			} else {
				if result != tt.expected {
					t.Errorf("expected result %f, got %f", tt.expected, result)
				}
			}
		})
	}
}