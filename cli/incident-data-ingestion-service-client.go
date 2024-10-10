package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

type AnalyticsServiceClient struct {
}

func (analyticsServiceClient *AnalyticsServiceClient) IngestIncidents(incidents []IncidentDataIngestionRequest) error {
	for _, incident := range incidents {
		err := ingestIncident(incident)
		if err != nil {
			fmt.Printf("Error ingesting incident %s data: %v\n", incident.Id, err.Error())
		}
	}
	return nil
}

func ingestIncident(incidentData IncidentDataIngestionRequest) error {
	body, err := json.Marshal(incidentData)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST",
		viper.GetString("incident-response-ingestion-service.base-url")+"/incident-data-ingestion", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", viper.GetString("analytics-service.api-key"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println("Incident data ingested successfully", incidentData.Id)
	return nil
}
