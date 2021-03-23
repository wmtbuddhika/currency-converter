package models

import "time"

// CurrencyExchangeRequest is used to set user requested exchange data in to a model
type CurrencyExchangeRequest struct {
	FromCurrency string  `json:"fromCurrency"`
	Amount       float32 `json:"amount"`
	ToCurrency   string  `json:"toCurrency"`
}

// CurrencyExchangeResponse is used to send response to the user with exchanged currency data
type CurrencyExchangeResponse struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

// ErrorResponse is used to send error data to user
type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// CurrencyExchangeSuccessResponse is used to set external API success response to a model
type CurrencyExchangeSuccessResponse struct {
	Success bool  `json:"success"`
	Query   Query `json:"query"`
	Info    Info  `json:"info"`
}

// Query is used to set external API success response query data to a model
type Query struct {
	From       string  `json:"from"`
	To         string  `json:"to"`
	Amount     float32 `json:"amount"`
	Historical string  `json:"historical"`
	Date       string  `json:"date"`
	Result     float32 `json:"result"`
}

// Info is used to set external API success response info data to a model
type Info struct {
	Timestamp time.Time `json:"timestamp"`
	Rate      float32   `json:"rate"`
}

// CurrencyExchangeErrorResponse is used to set external API error response to a model
type CurrencyExchangeErrorResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

// Error is used to set external API error response error data to a model
type Error struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}
