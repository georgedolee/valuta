package formatter

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/georgedolee/valuta/internal/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const baseCurrency = "GEL"

func PrintConversion(from, to string, amount, converted float64) {
	amountStr, fromStr := formatAmountAndCode(amount, from)
	convertedStr, toStr := formatAmountAndCode(converted, to)

	fmt.Printf("%s %s --> %s %s", amountStr, fromStr, convertedStr, toStr)
}

func PrintRates(currencies []model.Currencies) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.Style().Options.SeparateRows = true

	for _, currency := range currencies {
		amountStr, fromStr := formatAmountAndCode(float64(currency.Quantity), currency.Code)
		convertedStr, toStr := formatAmountAndCode(currency.Rate, baseCurrency)

		t.AppendRow([]interface{}{fromStr, amountStr, toStr, convertedStr})
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft},
		{Number: 2, Align: text.AlignLeft},
		{Number: 3, Align: text.AlignLeft},
		{Number: 4, Align: text.AlignLeft},
	})

	t.Render()
}

func formatAmountAndCode(amount float64, code string) (string, string) {
	amountStr := color.New(color.FgHiGreen, color.Bold).Sprintf("%.2f", amount)
	codeStr := color.New(color.FgHiBlue, color.Bold).Sprintf("%s", strings.ToUpper(code))
	return amountStr, codeStr
}
