package cmd

import (
	"fmt"
	"strconv"

	"github.com/georgedolee/valuta/internal/apiclient"
	"github.com/georgedolee/valuta/internal/formatter"
	"github.com/georgedolee/valuta/internal/validator"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert <amount> <from> <to>",
	Short: "Convert an amount from one currency to another",
	Long: `Convert an amount from one currency to another.

Arguments:
  amount   The amount to convert (float, e.g. 12.50)
  from     Source currency code (3-letter ISO code, e.g. usd)
  to       Target currency code (3-letter ISO code, e.g. eur)
`,
	Args:                  validator.ValidateConvertArgs,
	RunE:                  runConvert,
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(convertCmd)
}

func runConvert(cmd *cobra.Command, args []string) error {
	amountStr, from, to := args[0], args[1], args[2]
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %v", err)
	}

	convertedAmount, err := apiclient.FetchConversionApi(amount, from, to)
	if err != nil {
		return err
	}

	formatter.PrintConversion(from, to, amount, convertedAmount)
	return nil
}
