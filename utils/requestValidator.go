package utils

import "currency-converter/models"

// ValidateRequest is used to validate whether user request meets the requirements
func ValidateRequest(fromCurrency *models.CurrencyExchangeRequest) []string {
	var validations []string

	if fromCurrency.FromCurrency == "" {
		validations = append(validations, EmptyFromCurrency)
	}
	if fromCurrency.ToCurrency == "" {
		validations = append(validations, EmptyToCurrency)
	}
	return validations
}
