module github.com/paysuper/paysuper-tax-service

require (
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/InVisionApp/go-logger v1.0.1 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/micro/go-micro v1.8.0
	github.com/micro/go-plugins v1.2.0
	github.com/prometheus/client_golang v1.0.0
	github.com/stretchr/testify v1.3.0
	go.uber.org/zap v1.10.0
)

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
