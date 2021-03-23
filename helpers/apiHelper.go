package helpers

import (
	"currency-converter/models"
	"currency-converter/utils"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// GetExchangeData is used to get exchanged currency data by provided data from external API
func GetExchangeData(fromCurrency *models.CurrencyExchangeRequest) (*models.CurrencyExchangeSuccessResponse, error) {

	urlData := utils.RootParamAPIKey + os.Getenv("FIXER_API_KEY") +
		utils.ParamFrom + fromCurrency.FromCurrency +
		utils.ParamTo + fromCurrency.ToCurrency +
		utils.ParamAmount + fmt.Sprint(fromCurrency.Amount)

	request, err := http.NewRequest(utils.GET, os.Getenv("FIXER_CONVERT_URL")+urlData, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(utils.ContentType, utils.ApplicationJSON)
	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	content, err := utils.GetBodyContent(response.Body, &models.CurrencyExchangeSuccessResponse{})
	if err != nil {
		return nil, err
	}

	apiResponse := content.(*models.CurrencyExchangeSuccessResponse)
	if !apiResponse.Success {
		msg := "error on API end point"
		return nil, errors.New(msg)
	}

	return apiResponse, nil
}
