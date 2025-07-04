package validator

import "fmt"

func ValidateCurrencyCode(code string) error {
	if len(code) != 3 {
		return fmt.Errorf("invalid currency: %s\nexpected 3 letter ISO code", code)
	}

	return nil
}
