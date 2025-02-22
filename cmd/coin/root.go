/*
Copyright © 2025 Wanderson Lontra wandersonlontra@yahoo.com.br
*/
package coin

import (
	"fmt"
	"os"

	"github.com/WandersonLontra/coin-cli/configs"
	"github.com/WandersonLontra/coin-cli/internal/cache"
	"github.com/WandersonLontra/coin-cli/internal/entity"
	"github.com/WandersonLontra/coin-cli/internal/web"
	"github.com/spf13/cobra"
)

type RunE func(cmd *cobra.Command, args []string) error
type FuncGetCurrencies func(fetcher *web.Fetcher, cacheStored *cache.CacheHandler, forceToFetch bool) (*entity.Currency, error)

var fetcher = web.NewFetcher(configs.BaseUrl, configs.AccessKey)
var cacheStored = cache.NewCacheHandler(configs.CacheDir,configs.CacheFile, configs.TTLCache)

func getCurrencies(fetcher *web.Fetcher, cacheStored *cache.CacheHandler, forceToFetch bool) (*entity.Currency, error) {
	if cacheStored.Exists() && !cacheStored.IsCacheExpired() && !forceToFetch {
		return cacheStored.Get()
	}

	currencies, err := fetcher.GetCurrencies()
	if err != nil {
		return nil, fmt.Errorf("error fetching currencies: %s", err)
	}

	err = cacheStored.Delete()
	if err != nil {
		return nil, fmt.Errorf("error deleting cache file: %s", err)
	}

	err = cacheStored.Set(currencies)
	if err != nil {
		return nil, fmt.Errorf("error setting cache file: %s", err)
	}

	return currencies, nil
}

var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "A currency converter CLI",
	Long: `coin is a CLI application to convert currencies.
	You can convert currencies using the latest exchange rate.`,
}

func Execute() {
	err := coinCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	coinCmd.PersistentFlags().BoolP("force", "F", false, "Force to fetch the latest exchange rate. Default is false - PS: It consumes the API rate limit")
}


