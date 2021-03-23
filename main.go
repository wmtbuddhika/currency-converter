package main

import (
	"currency-converter/controllers"
	"currency-converter/utils"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

// init is used to load env values from .env file
func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("Error on loading .env file") // use logs
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/convert", controllers.ConvertCurrency).Methods(utils.POST)

	_ = http.ListenAndServe(":"+os.Getenv("PORT"), router)

}
