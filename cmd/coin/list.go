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
		Use:   "list [symbols...] [base] [force]",
		Short: "List currencies",
		Long: "",
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
		for currency,_ := range currencies.Rates {
			if len(symbols) > 0 {
				if !slices.Contains(symbols, currency) {
					continue
				}
				value, err := converter.Convert(1, base, currency, currencies.Rates)
				if err != nil {
					cmd.Help()
					return err
				}
				cmd.Printf("%s = %.2f\n", currency, value)
				continue
			}
			value, err := converter.Convert(1, base, currency, currencies.Rates)
			if err != nil {
				cmd.Help()
				return err
			}
			cmd.Printf("%s = %.2f\n", currency, value)
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
