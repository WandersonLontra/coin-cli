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
type FuncGetCurrencies func(forceToFetch bool) (*entity.Currency, error)

func getCurrencies(forceToFetch bool) (*entity.Currency, error) {
	var currencyStored cache.Currency
	err := currencyStored.Get()
	if err != nil {
		fmt.Println("cache file not found")
	}

	if currencyStored.Exists() && currencyStored.IsTodaysCache() && !forceToFetch {
		return &entity.Currency{
			Success:  true,
			Timestamp: currencyStored.Timestamp,
			Base:      currencyStored.Base,
			Date:      currencyStored.Date,
			Rates:     currencyStored.Rates,
		}, nil
	}

	fetcher := web.NewFetcher(configs.BaseUrl, "/latest", configs.AccessKey, configs.BaseCurrency)
	currencies, err := fetcher.GetCurrencies()
	if err != nil {
		return nil, fmt.Errorf("error fetching currencies: %s", err)
	}
	err = currencyStored.Delete()
	if err != nil {
		return nil, fmt.Errorf("error deleting cache file: %s", err)
	}
	currencyStored = cache.Currency{
		Timestamp: currencies.Timestamp,
		Base:      currencies.Base,
		Date:      currencies.Date,
		Rates:     currencies.Rates,
	}
	err = currencyStored.Set()
	if err != nil {
		return nil, fmt.Errorf("error setting cache file: %s", err)
	}

	return currencies, nil
}

// coinCmd represents the base command when called without any subcommands
var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "A currency converter CLI",
	Long: `coin is a CLI application to convert currencies.
	You can convert currencies using the latest exchange rate.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the coinCmd.
func Execute() {
	err := coinCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// coinCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coin-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	coinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


