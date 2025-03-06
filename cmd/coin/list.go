/*
Copyright © 2025 Wanderson Lontra wandersonlontra@yahoo.com.br
*/
package coin

import (
	"slices"

	"github.com/WandersonLontra/coin-cli/internal/converter"
	"github.com/spf13/cobra"
)

func newListCmd(getCurrencies FuncGetCurrencies) *cobra.Command {
	return &cobra.Command{
		Use:   "list [--symbols <currency1,currency2,...>] [--base <currency>] [--force]",
		Short: "Lists exchange rates for various currencies based on a specified base currency",
		Long: "Usage:\n"+  
				"\tlist [--symbols <currency1,currency2,...>] [--base <currency>] [--force]\n"+
				"\tlist [-s <currency1,currency2,...>] [-b <currency>] [-F]\n\n"+
				"Description:\n"+
				"\tLists exchange rates for various currencies based on a specified base currency.\n\n"+
				"Flags:\n"+
				"\t-s, --symbols\tOptional. A comma-separated list of currency codes to fetch exchange rates for. Default is all currencies.\n"+
				"\t-b, --base\tOptional. The base currency for conversion. Default is USD.\n"+
				"\t-F, --force\tOptional. Forces fetching the latest exchange rates instead of using cached data.\n\n"+
				"Examples:\n"+
				"\tList exchange rates for all currencies using USD as the base:\n"+
				"\t list\n\n"+
				"\tList exchange rates for EUR, GBP, and JPY using USD as the base:\n"+
				"\t list --symbols EUR,GBP,JPY\n"+
				"\t list -s EUR,GBP,JPY\n\n"+
				"\tList exchange rates for all currencies using EUR as the base:\n"+
				"\t list --base EUR\n"+
				"\t list -b EUR\n\n"+
				"\tList exchange rates for CAD and AUD using GBP as the base and force refresh rates:\n"+
				"\t list --symbols CAD,AUD --base GBP --force\n"+
				"\t list -s CAD,AUD -b GBP -F\n\n"+
				"Notes:\n"+
				"\t- Currency codes must follow the ISO 4217 standard (e.g., USD for US Dollar, EUR for Euro).\n"+
				"\t- If no `--symbols` flag is provided, all available currencies will be listed.\n"+
				"\t- If no `--base` flag is provided, USD is used as the default base currency.\n"+
				"\t- Using `--force` (`-F`) may result in slower response times due to fetching fresh data.\n",
		RunE: runList(getCurrencies),
	}
}

func runList(getCurrencies FuncGetCurrencies) RunE {
	return func(cmd *cobra.Command, args []string) error {
		symbols, err := cmd.Flags().GetStringSlice("symbols")
		if err != nil {
			cmd.Help()
			return err
		}
		base, err := cmd.Flags().GetString("base")
		if err != nil {
			cmd.Help()
			return err
		}
		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			cmd.Help()
			return err
		}
		currencies, err := getCurrencies(fetcher, cacheStored, force)
		if err != nil {
			cmd.Help()
			return err
		}
		for currency := range currencies.Rates {
			if len(symbols) > 0 {
				if !slices.Contains(symbols, currency) {
					continue
				}
				value, err := converter.Convert(1, base, currency, currencies.Rates)
				if err != nil {
					cmd.Help()
					return err
				}
				cmd.Printf("\033[1;36m%s \033[0;37m= \033[0;32m%.2f \033[0;37m\n", currency, value)
				continue
			}
			value, err := converter.Convert(1, base, currency, currencies.Rates)
			if err != nil {
				cmd.Help()
				return err
			}
			cmd.Printf("\033[1;36m%s \033[0;37m= \033[0;32m%.2f \033[0;37m\n", currency, value)
		}
		return nil
	}
}
func init() {
	listCmd := newListCmd(getCurrencies)
	coinCmd.AddCommand(listCmd)


	listCmd.Flags().StringSliceP("symbols", "s", []string{}, "List of symbols to fetch the exchange rate. Default is all currencies")
	listCmd.Flags().StringP("base", "b", "USD", "Base currency to fetch the exchange rate. Default is USD")
}
