/*
Copyright © 2025 Wanderson Lontra
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/WandersonLontra/coin-cli/cmd/coin"
	"github.com/WandersonLontra/coin-cli/configs"
	"github.com/WandersonLontra/coin-cli/internal/cache"
	"github.com/WandersonLontra/coin-cli/internal/converter"
	"github.com/WandersonLontra/coin-cli/internal/entity"
	"github.com/WandersonLontra/coin-cli/internal/web"
)

func main() {
	var currencies *entity.Currency

	currencyStored := cache.Currency{}
	err := currencyStored.Get()
	if err != nil {
		log.Printf("error getting cache file: %s", err)
	}
	if !currencyStored.Exists() || !currencyStored.IsTodaysCache() {
		fmt.Println("CACHE NOT FOUND")

		fetcher := web.NewFetcher(configs.BaseUrl, "/latest", configs.AccessKey, configs.BaseCurrency)
	
		currencies, err = fetcher.GetCurrencies()
	
		if err != nil {
			log.Printf("error fetching currencies: %s", err)
			os.Exit(1)
		}
	
		currencyStored = cache.Currency{
			Timestamp: currencies.Timestamp,
			Base:      currencies.Base,
			Date:      currencies.Date,
			Rates:     currencies.Rates,
		}
		err = currencyStored.Delete()
		if err != nil {
			log.Printf("error deleting cache file: %s", err)
			os.Exit(1)
		}
		err = currencyStored.Set()
		if err != nil {
			log.Printf("error setting cache file: %s", err)
			os.Exit(1)
		}
	}

	currencies = &entity.Currency{
		Success:  true,
		Timestamp: currencyStored.Timestamp,
		Base:      currencyStored.Base,
		Date:      currencyStored.Date,
		Rates:     currencyStored.Rates,
	}
	fmt.Println(currencies.Rates["BRL"])
	conv, _ := converter.Convert(1, "USD", "BRL", currencies.Rates)
	fmt.Printf("converted: %f\n", conv)

	coin.Execute()
}
