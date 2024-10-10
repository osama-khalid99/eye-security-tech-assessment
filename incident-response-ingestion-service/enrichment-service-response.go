package main

type EnrichmentServiceResponse struct {
	Asn           string `json:"asn"`
	Category      string `json:"category"`
	CorrelationId string `json:"correlationId"`
}
