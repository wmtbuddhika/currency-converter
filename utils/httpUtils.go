package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// GetBodyContent is used to transfer body content to the supplied interface
func GetBodyContent(body io.Reader, model interface{}) (interface{}, error) {
	data, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, model)
	return model, err
}

// RespondWithData is used to send the response with provided interface data
func RespondWithData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Add(ContentType, ApplicationJSON)

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		msg := "Error on response JSON encoding"
		AddLog(msg, err)
		fmt.Print(msg)
	}
}
