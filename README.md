[![Go Report Card](https://goreportcard.com/badge/github.com/alvjtc/gcp-pricing-info)](https://goreportcard.com/report/github.com/alvjtc/gcp-pricing-info)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](http://www.apache.org/licenses/LICENSE-2.0)

# GCP Pricing Info
API to get pricing information about Google Cloud Platform resources as a calculator.

## Build locally

To build the project just use Go standard compiler with the following commands:

```
go mod tidy
go mod download
go mod verify
go build -o bin\gcp-pricing-info[.exe] cmd\server\main.go
```

And then run the server with the command:

```
.\bin\gcp-pricing-info[.exe]
```

Now you can test the API in your host calling the Health Check endpoint http://localhost:8080/healthcheck

## Build Docker container

To build the Docker container you need to run the following Docker command:

```
docker build -t gcp-pricing-info .
```

And then run the container and forward the port with the command:

```
docker run -p 8080:8080 gcp-pricing-info
```

Now you can test the API in your host calling the Health Check endpoint http://localhost:8080/healthcheck
