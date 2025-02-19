/*
Copyright © 2025 Wanderson Lontra wandersonlontra@yahoo.com.br
*/
package coin

import (
	"fmt"

	"github.com/WandersonLontra/coin-cli/internal/converter"
	"github.com/spf13/cobra"
)

func newConvertCmd(getCurrencies FuncGetCurrencies) *cobra.Command {
	return &cobra.Command{
		Use:   "convert [from] [to] [amount] [force]",
		Short: "Convert currencies",
		Long: "",
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
		fmt.Printf("Converting %.2f from %s to %s\n", amount, from, to)
		currencies, err := getCurrencies(force)
		if err != nil {
			cmd.PrintErr(err)
			return err
		}
		conv, err := converter.Convert(amount, from, to, currencies.Rates)
		if err != nil {
			cmd.PrintErr(err)
			return err
		}
		fmt.Printf("%.2f %s = %.2f %s\n", amount, from, conv, to)
		return nil
	}
}
func init() {
	convertCmd := newConvertCmd(getCurrencies)
	coinCmd.AddCommand(convertCmd)

	convertCmd.Flags().Float64P("amount", "a", 1, "Amount to convert")
	convertCmd.Flags().StringP("from", "f", "", "Currency to convert from")
	convertCmd.Flags().StringP("to", "t", "", "Currency to convert to")
	convertCmd.Flags().BoolP("force", "F", false, "Force to fetch the latest exchange rate. Default is false - PS: It consumes the API rate limit")

	convertCmd.MarkFlagsRequiredTogether("from", "to")
}
