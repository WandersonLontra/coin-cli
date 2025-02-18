package entity

type Currency struct {
	Success   bool              	`json:"success"`
	Timestamp int               	`json:"timestamp"`
	Base      string            	`json:"base"`
	Date      string            	`json:"date"`
	Rates     map[string]float64 	`json:"rates"`
}
