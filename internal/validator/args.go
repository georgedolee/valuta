package validator

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func ValidateConvertArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("expected 3 arguments: <amount> <from> <to>")
	}

	if _, err := strconv.ParseFloat(args[0], 64); err != nil {
		return fmt.Errorf("invalid amount: %s\nexpected float", args[0])
	}

	if err := ValidateCurrencyCode(args[1]); err != nil {
		return err
	}

	if err := ValidateCurrencyCode(args[2]); err != nil {
		return err
	}

	return nil
}
