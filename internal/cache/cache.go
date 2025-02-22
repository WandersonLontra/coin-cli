package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/WandersonLontra/coin-cli/internal/entity"
)

type CacheHandler struct {
	CacheFile string
	CacheDir  string
	TTLCache  float64
}
func NewCacheHandler(cacheDir, cacheFile string, ttl float64) *CacheHandler {
	return &CacheHandler{
		CacheFile: cacheFile,
		CacheDir:  cacheDir,
		TTLCache: ttl,
	}
}
func (c *CacheHandler) cachePath() string {
	return fmt.Sprintf("%s/%s", c.CacheDir, c.CacheFile)
}

func (c *CacheHandler) Set(currencies *entity.Currency) error {
	err := os.MkdirAll(c.CacheDir, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(c.cachePath())
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

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
	file, err := os.Open(c.cachePath())
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var currencies entity.Currency
	json.NewDecoder(file).Decode(&currencies)
	return &currencies, nil
}

func (c *CacheHandler) Delete() error {
	if !c.Exists() {
		return nil
	}
	err := os.Remove(c.cachePath())
	if err != nil {
		return err 
	}
	return nil
}

func (c *CacheHandler) Exists() bool {
	_, err := os.Stat(c.cachePath())
	return !os.IsNotExist(err)
}

func (c *CacheHandler) IsCacheExpired() bool {
	now := time.Now().UTC()
	cacheDate, err := c.Get()
	if err != nil {
		return false
	}
	from := time.UnixMilli(cacheDate.Timestamp)
	diff := now.Sub(from).Hours()
	return diff >= c.TTLCache
}

func (c *CacheHandler) IsTodaysCache() bool {
	today := time.Now().UTC().Format("2006-01-02")
	cacheDate, err := c.Get()
	if err != nil {
		return false
	}
	
	return today == cacheDate.Date
}