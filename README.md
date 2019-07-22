# PaySuper Tax Service

[![License: GPL 3.0](https://img.shields.io/badge/License-GPL3.0-green.svg)](https://opensource.org/licenses/Gpl3.0)
[![Build Status](https://travis-ci.org/paysuper/paysuper-tax-service.svg?branch=master)](https://travis-ci.org/paysuper/paysuper-tax-service) 
[![codecov](https://codecov.io/gh/paysuper/paysuper-tax-service/branch/master/graph/badge.svg)](https://codecov.io/gh/paysuper/paysuper-tax-service) 
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/paysuper-tax-service)](https://goreportcard.com/report/github.com/paysuper/paysuper-tax-service)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=paysuper_paysuper-tax-service&metric=alert_status)](https://sonarcloud.io/dashboard?id=paysuper_paysuper-tax-service)

# Motivation

PaySuper works as a tax agent for its customers. This service helps us to determine tax rates for different
regions. Tax legislation in different regions is quite complex. For example, the EU rules require at least 
two unambiguous factors in determining the tax rate, Russia requires to pay taxes if the BIN card is issued
in Russia, in the US taxes vary not only at the state and city levels, but even, in some cases, on 
different streets in the same city.

How this service works. We use two blocks of information in the request - information received from the 
user's IP address and provided by the user himself. Priority is always given to information independently 
specified by the user. If it is not enough or for some reason we cannot get tax rates-IP address data is 
used.

Additional factors such as the user's language, time zone location, and other criteria must be calculated
before requests to this service.

# Usage

Application designed to be launched with Kubernetes and handle all configuration from env variables:

| Variable                            | Description                                                                                            |
|-------------------------------------|-------------------------------------------------------------------------|
| TAX_SERVICE_DSN                     | Postgress DSN string                                                    |
| TAX_SERVICE_METRICS_PORT            | Http port to expose /health endpoint. Default is `8080`.                |
| TAX_SERVICE_KUBERNETES_SERVICE_HOST | Should be `true` to laucnh in k8 mode. Default is `false` .             |


## Contributing
We feel that a welcoming community is important and we ask that you follow PaySuper's [Open Source Code of Conduct](https://github.com/paysuper/code-of-conduct/blob/master/README.md) in all interactions with the community.

PaySuper welcomes contributions from anyone and everyone. Please refer to each project's style and contribution guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.

The master branch of this repository contains the latest stable release of this component.
