package main

import (
	"fmt"
	"github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/handlers"
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	k8s "github.com/micro/kubernetes/go/micro"
	"github.com/paysuper/paysuper-tax-service/internal"
	"github.com/paysuper/paysuper-tax-service/pkg"
	"github.com/paysuper/paysuper-tax-service/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Config struct {
	KubernetesHost string `envconfig:"KUBERNETES_SERVICE_HOST" required:"false"`
	MetricsPort    int    `envconfig:"METRICS_PORT" required:"false" default:"8080"`
}

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		logger.Fatal("Config init failed with error", zap.Error(err))
	}

	var service micro.Service
	options := []micro.Option{
		micro.Name(pkg.ServiceName),
		micro.Version(pkg.Version),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	}

	if config.KubernetesHost == "" {
		service = micro.NewService(options...)
		logger.Info("Initialize micro service")
	} else {
		service = k8s.NewService(options...)
		logger.Info("Initialize k8s service")
	}

	service.Init()

	taxService := &internal.Service{}
	err := tax_service.RegisterTaxServiceHandler(service.Server(), taxService)
	if err != nil {
		logger.Fatal("Can`t register service in micro", zap.Error(err))
	}

	initHealth(taxService, config.MetricsPort)
	initPrometheus()

	go func() {
		if err = http.ListenAndServe(fmt.Sprintf(":%d", config.MetricsPort), nil); err != nil {
			logger.Fatal("Metrics listen failed", zap.Error(err))
		}
	}()

	if err := service.Run(); err != nil {
		logger.Fatal("Can`t run service", zap.Error(err))
	}
}

func initHealth(checker health.ICheckable, port int) {
	h := health.New()
	err := h.AddChecks([]*health.Config{
		{
			Name:     "health-check",
			Checker:  checker,
			Interval: time.Duration(1) * time.Second,
			Fatal:    true,
		},
	})

	if err != nil {
		zap.L().Fatal("Health check register failed", zap.Error(err))
	}

	zap.L().Info("Health check listening on :%d", zap.Int("port", port))

	if err = h.Start(); err != nil {
		zap.L().Fatal("Health check start failed", zap.Error(err))
	}

	http.HandleFunc("/health", handlers.NewJSONHandlerFunc(h, nil))
}

func initPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
}
