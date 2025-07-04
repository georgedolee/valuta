package validator

func ValidateCurrencies(currencies []string) error {
	for _, currencyCode := range currencies {
		if err := ValidateCurrencyCode(currencyCode); err != nil {
			return err
		}
	}
	return nil
}
