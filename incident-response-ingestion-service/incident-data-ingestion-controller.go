package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type IncidentDataIngestionController struct {
	analyticsServiceClient AnalyticsServiceClient
}

func (controller *IncidentDataIngestionController) ProcessIncidentData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	bodyBytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var incidentData IncidentData
	err = json.Unmarshal(bodyBytes, &incidentData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Incident data to be enriched", incidentData.Id)
	err = validateIncidentData(incidentData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := controller.analyticsServiceClient.enrichIncidentData(incidentData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var analyticsServiceRequest AnalyticsServiceRequest = AnalyticsServiceRequest{
		Id:            incidentData.Id,
		Asset:         incidentData.Asset,
		Ip:            incidentData.Ip,
		Asn:           response.Asn,
		Category:      response.Category,
		CorrelationId: response.CorrelationId,
	}

	err = controller.analyticsServiceClient.ingestIncidentData(analyticsServiceRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Incident data processed successfully", incidentData.Id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Incident data processed successfully"))
}

func validateIncidentData(incidentData IncidentData) error {
	if incidentData.Id == "" {
		return errors.New("id is required")
	}
	if incidentData.Asset == "" {
		return errors.New("asset name is required")
	}
	if incidentData.Ip == "" {
		return errors.New("ip is required")
	}
	if incidentData.Category == "" {
		return errors.New("category is required")
	}
	return nil
}
