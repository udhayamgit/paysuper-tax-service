# PaySuper Tax Service

[![License: GPL 3.0](https://img.shields.io/badge/License-GPL3.0-green.svg)](https://opensource.org/licenses/Gpl3.0)
[![Build Status](https://travis-ci.org/paysuper/paysuper-tax-service.svg?branch=master)](https://travis-ci.org/paysuper/paysuper-tax-service) 
[![codecov](https://codecov.io/gh/paysuper/paysuper-tax-service/branch/master/graph/badge.svg)](https://codecov.io/gh/paysuper/paysuper-tax-service) 
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/paysuper-tax-service)](https://goreportcard.com/report/github.com/paysuper/paysuper-tax-service)

# Motivation



# Usage

Application designed to be launched with Kubernetes and handle all configuration from env variables:

| Variable                            | Description                                                                                            |
|-------------------------------------|-------------------------------------------------------------------------|
| TAX_SERVICE_DSN                     | Postgress DSN string                                                    |
| TAX_SERVICE_METRICS_PORT            | Http port to expose /health endpoint. Default is `8080`.                |
| TAX_SERVICE_KUBERNETES_SERVICE_HOST | Should be `true` to laucnh in k8 mode. Default is `false` .             |