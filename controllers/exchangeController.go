package controllers

import (
	"currency-converter/helpers"
	"currency-converter/models"
	"currency-converter/utils"
	"net/http"
	"strings"
)

// ConvertCurrency is used to handle request body content and request the external API to convert the currency
func ConvertCurrency(w http.ResponseWriter, r *http.Request) {

	data, err := utils.GetBodyContent(r.Body, &models.CurrencyExchangeRequest{})

	if err != nil {
		msg := "Error on passing request body content"
		handleError(w, http.StatusInternalServerError, msg, err)
		return
	}

	fromCurrency := data.(*models.CurrencyExchangeRequest)
	validations := utils.ValidateRequest(fromCurrency)

	if validations != nil {
		msg := "Currency exchange request failed due to validations: " + strings.Join(validations, ", ")
		handleError(w, http.StatusBadRequest, msg, err)
		return
	}

	toCurrency, err := helpers.GetExchangeData(fromCurrency)

	if err != nil {
		msg := "Error on currency conversion"
		handleError(w, http.StatusInternalServerError, msg, err)
		return
	}

	utils.RespondWithData(w, http.StatusOK, toCurrency)
}

// handleError is used to log errors and send error response to the user
func handleError(w http.ResponseWriter, statusCode int, msg string, err error) {
	utils.AddLog(msg, err)
	response := models.ErrorResponse{
		Status:  false,
		Message: msg,
	}
	utils.RespondWithData(w, statusCode, response)
}
