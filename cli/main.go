package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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
	processIncidentData()
}

func processIncidentData() {
	file, err := os.Open("example_data_2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var incidents []IncidentDataIngestionRequest
	firstRecord := true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if record[0] == "" {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if firstRecord {
			firstRecord = false
			continue
		}
		fmt.Println(record)

		splitRecord := strings.Split(record[0], ";")
		incident := IncidentDataIngestionRequest{
			Id:       getAtIndex(splitRecord, 0),
			Asset:    getAtIndex(splitRecord, 1),
			Ip:       getAtIndex(splitRecord, 2),
			Category: getAtIndex(splitRecord, 5),
		}
		incidents = append(incidents, incident)
	}

	analyticsServiceClient := AnalyticsServiceClient{}
	analyticsServiceClient.IngestIncidents(incidents)
}

func getAtIndex(s []string, index int) string {
	if index >= len(s) {
		return ""
	}
	return s[index]
}
