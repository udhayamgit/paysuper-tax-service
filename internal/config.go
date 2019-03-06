package internal

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DSN            string `envconfig:"DSN" required:"false" default:"postgres://postgres:postgres@localhost:5432/tax_service?sslmode=disable"`
	KubernetesHost string `envconfig:"KUBERNETES_SERVICE_HOST" required:"false"`
	MetricsPort    int    `envconfig:"METRICS_PORT" required:"false" default:"8080"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	return config, envconfig.Process("TAX_SERVICE", config)
}
