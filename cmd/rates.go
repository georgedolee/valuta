package cmd

import (
	"strings"

	"github.com/georgedolee/valuta/internal/apiclient"
	"github.com/georgedolee/valuta/internal/formatter"
	"github.com/georgedolee/valuta/internal/validator"
	"github.com/spf13/cobra"
)

var ratesCmd = &cobra.Command{
	Use:   "rates",
	Short: "List exchange rates for Georgian Lari",
	RunE:  runRates,
}

func init() {
	rootCmd.AddCommand(ratesCmd)
	ratesCmd.Flags().StringP("currencies", "c", "USD,EUR,RUB,CNY", "Comma-separated list of currency codes")
}

func runRates(cmd *cobra.Command, args []string) error {
	currenciesStr, _ := cmd.Flags().GetString("currencies")
	currencies := strings.Split(currenciesStr, ",")

	if err := validator.ValidateCurrencies(currencies); err != nil {
		return err
	}

	rates, err := apiclient.FetchRatesApi(currencies)
	if err != nil {
		return err
	}

	formatter.PrintRates(rates)
	return nil
}
