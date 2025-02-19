package cache

import (
	"encoding/json"
	"os"
	"time"

	"github.com/WandersonLontra/coin-cli/internal/entity"
)

type CacheHandler struct {
	CacheFile string
}

func NewCacheHandler(cacheFile string) *CacheHandler {
	return &CacheHandler{
		CacheFile: cacheFile,
	}
}
func (c *CacheHandler) Set(currencies *entity.Currency) error {
	f, err := os.Create(c.CacheFile)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	encoder.SetIndent("", "  ")
	if err := encoder.Encode(currencies); err != nil {
		return err 
	}
	return nil
}

func (c *CacheHandler) Get() (*entity.Currency, error) {
	if !c.Exists() {
		return nil, nil
	}
	dir, err := os.Open(c.CacheFile)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	var currencies entity.Currency
	json.NewDecoder(dir).Decode(&currencies)
	return &currencies, nil
}

func (c *CacheHandler) Delete() error {
	if !c.Exists() {
		return nil
	}
	err := os.Remove(c.CacheFile)
	if err != nil {
		return err 
	}
	return nil
}

func (c *CacheHandler) Exists() bool {
	_, err := os.Stat(c.CacheFile)
	return !os.IsNotExist(err)
}

func (c *CacheHandler) IsTodaysCache() bool {
	today := time.Now().UTC().Format("2006-01-02")
	cacheDate, err := c.Get()
	if err != nil {
		return false
	}
	
	return today == cacheDate.Date
}