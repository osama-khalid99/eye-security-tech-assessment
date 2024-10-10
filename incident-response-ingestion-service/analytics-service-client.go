package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type AnalyticsServiceClient struct {
}

func (analyticsServiceClient *AnalyticsServiceClient) enrichIncidentData(incidentData IncidentData) (*EnrichmentServiceResponse, error) {
	body, err := json.Marshal(incidentData)
	if err != nil {
		fmt.Println("Error enriching incident data", err.Error())
		return nil, err
	}
	requestUrl := viper.GetString("analytics-service.base-url") + "/enrichment"
	fmt.Println("Request URL", requestUrl)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error enriching incident data", err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", viper.GetString("analytics-service.api-key"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error enriching incident data", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errorResponse, _ := io.ReadAll(resp.Body)
		fmt.Println("Error enriching incident data", resp.StatusCode, string(errorResponse))
		return nil, errors.New("failed to enrich data" + string(errorResponse))
	}
	var enrichmentServiceResponse EnrichmentServiceResponse
	err = json.NewDecoder(resp.Body).Decode(&enrichmentServiceResponse)
	if err != nil {
		fmt.Println("Error enriching incident data", err.Error())
		return nil, err
	}
	return &enrichmentServiceResponse, nil
}

func (analyticsServiceClient *AnalyticsServiceClient) ingestIncidentData(analyticsServiceRequest AnalyticsServiceRequest) error {
	body, err := json.Marshal(analyticsServiceRequest)
	if err != nil {
		fmt.Println("Error ingesting incident data", err.Error())
		return err
	}
	requestUrl := viper.GetString("analytics-service.base-url") + "/analytics"
	fmt.Println("Request URL", requestUrl)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error ingesting incident data", err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", viper.GetString("analytics-service.api-key"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error ingesting incident data", err.Error())
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errorResponse, _ := io.ReadAll(resp.Body)
		fmt.Println("Error ingesting incident data", resp.StatusCode, string(errorResponse))
		return errors.New("failed to ingest data to analytics service" + string(errorResponse))
	}
	return nil
}
