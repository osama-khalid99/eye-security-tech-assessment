package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	mux := setupEndpoints()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("\nServer is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func setupEndpoints() *http.ServeMux {
	mux := http.NewServeMux()
	controller := IncidentDataIngestionController{
		analyticsServiceClient: AnalyticsServiceClient{},
	}

	mux.HandleFunc("/incident-data-ingestion", controller.ProcessIncidentData)

	return mux
}
