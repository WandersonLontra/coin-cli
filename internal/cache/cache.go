package cache

import (
	"encoding/json"
	"os"
	"time"
)

var cacheFile = "./cache_file/cache.json"

type Currency struct {
	Timestamp int               	`json:"timestamp"`
	Base      string            	`json:"base"`
	Date      string            	`json:"date"`
	Rates     map[string]float64 	`json:"rates"`
}

func (c *Currency) Set() error {
	f, err := os.Create(cacheFile)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	encoder.SetIndent("", "  ")
	if err := encoder.Encode(c); err != nil {
		return err 
	}
	return nil
}

func (c *Currency) Get() error {
	if !c.Exists() {
		return nil
	}
	dir, err := os.Open(cacheFile)
	if err != nil {
		return err
	}
	defer dir.Close()

	json.NewDecoder(dir).Decode(&c)
	return nil
}

func (c *Currency) Delete() error {
	if !c.Exists() {
		return nil
	}
	err := os.Remove(cacheFile)
	if err != nil {
		return err 
	}
	return nil
}

func (c *Currency) Exists() bool {
	_, err := os.Stat(cacheFile)
	return !os.IsNotExist(err)
}

func (c *Currency) IsTodaysCache() bool {
	today := time.Now().UTC().Format("2006-01-02")
	
	return today == c.Date
}