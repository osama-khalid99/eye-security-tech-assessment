package main

type AnalyticsServiceRequest struct {
	Id            string `json:"id"`
	Asset         string `json:"asset"`
	Ip            string `json:"ip"`
	Asn           string `json:"asn"`
	Category      string `json:"category"`
	CorrelationId string `json:"correlationId"`
}
