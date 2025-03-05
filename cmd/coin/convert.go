/*
Copyright © 2025 Wanderson Lontra wandersonlontra@yahoo.com.br
*/
package coin

import (
	"github.com/WandersonLontra/coin-cli/internal/converter"
	"github.com/spf13/cobra"
)

func newConvertCmd(getCurrencies FuncGetCurrencies) *cobra.Command {
	return &cobra.Command{
		Use:   "convert --from <currency> --to <currency> [--amount <value>] [--force] ",
		Short: "A simple currency converter that fetches exchange rates and converts an amount from one currency to another.",
		Long: "Usage:\n"+  
				"\tconvert --from <currency> --to <currency> [--amount <value>] [--force]\n"+ 
				"\tconvert -f <currency> -t <currency> [-a <value>] [-F]\n\n"+
				"Description:\n"+
				"\tA simple currency converter that fetches exchange rates and converts an amount from one currency to another.\n\n"+
				"Flags:\n"+
				"\t-f, --from     Required. The currency code to convert from (e.g., USD, EUR, BRL).\n"+
				"\t-t, --to       Required. The currency code to convert to (e.g., USD, EUR, BRL).\n"+
				"\t-a, --amount   Optional. The amount to be converted. Default is 1.\n"+
				"\t-F, --force    Optional. Forces fetching the latest exchange rates instead of using cached data.\n\n"+
				"Examples:\n"+
				"\tConvert 1 USD to EUR:\n"+
					"\t convert --from USD --to EUR\n"+
					"\t convert -f USD -t EUR\n\n"+
				"\tConvert 100 GBP to JPY:\n"+
					"\t convert --from GBP --to JPY --amount 100\n"+
					"\t convert -f GBP -t JPY -a 100\n\n"+
				"\tConvert 50 CAD to AUD and force refresh rates:\n"+
					"\t convert --from CAD --to AUD --amount 50 --force\n"+
					"\t convert -f CAD -t AUD -a 50 -F\n\n"+
				"Notes:\n"+
				"\t- Currency codes must follow the ISO 4217 standard (e.g., USD for US Dollar, EUR for Euro).\n"+
				"\t- If no amount is provided, the default value of 1 is used.\n"+
				"\t- Using `--force` (`-F`) may result in slower response times due to fetching fresh data.",
		RunE: runConvert(getCurrencies),
	}
}

func runConvert(getCurrencies FuncGetCurrencies) RunE {
	return func(cmd *cobra.Command, args []string) error {
		amount, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			cmd.Help()
			return err
		}
		from, err := cmd.Flags().GetString("from")
		if err != nil {
			cmd.Help()
			return err
		}
		to, err := cmd.Flags().GetString("to")
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
			cmd.PrintErr(err)
			return err
		}
		conv, err := converter.Convert(amount, from, to, currencies.Rates)
		if err != nil {
			cmd.PrintErr(err)
			return err
		}
		baseRate, err := converter.Convert(1, from, to, currencies.Rates)
		if err != nil {
			cmd.PrintErr(err)
			return err
		}
		cmd.Printf("\033[0;32m%.2f \033[1;36m%s \033[0;37m= \033[0;32m%.2f \033[1;35m%s\n", amount, from, conv, to)
		cmd.Printf("\033[0;37mRate: \033[0;32m%.2f \033[1;36m%s \033[0;37m= \033[0;32m%.2f \033[1;35m%s\n", 1.00, from, baseRate, to)
		return nil
	}
}
func init() {
	convertCmd := newConvertCmd(getCurrencies)
	coinCmd.AddCommand(convertCmd)

	convertCmd.Flags().Float64P("amount", "a", 1, "Amount to convert")
	convertCmd.Flags().StringP("from", "f", "", "Currency to convert from")
	convertCmd.Flags().StringP("to", "t", "", "Currency to convert to")

	convertCmd.MarkFlagsRequiredTogether("from", "to")
}
