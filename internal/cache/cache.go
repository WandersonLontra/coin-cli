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
}
func NewCacheHandler(cacheDir, cacheFile string) *CacheHandler {
	return &CacheHandler{
		CacheFile: cacheFile,
		CacheDir:  cacheDir,
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
	f, err := os.Create(c.cachePath())
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
	dir, err := os.Open(c.cachePath())
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

func (c *CacheHandler) IsTodaysCache() bool {
	today := time.Now().UTC().Format("2006-01-02")
	cacheDate, err := c.Get()
	if err != nil {
		return false
	}
	
	return today == cacheDate.Date
}