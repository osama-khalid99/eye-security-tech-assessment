package main

type IncidentDataIngestionRequest struct {
	Id       string `json:"id"`
	Asset    string `json:"asset"`
	Ip       string `json:"ip"`
	Category string `json:"category"`
}
