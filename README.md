# CLI

## Description
CLI application fetches incident activity logs from csv file and forwards it to incident response ingestion service for further processing.

# Incident Response Ingestion Service

## Description
Incident response ingestion service is a microservice which expose endpoints to process incident activity data.

## Prerequisites

- Go

## Run CLI

```bash
$ cd cli
$ go run .
```

## Run Incident Response Ingestion Service

```bash
$ cd incident-response-ingestion-service
$ go run .
```