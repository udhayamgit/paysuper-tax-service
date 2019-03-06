package internal

import (
	"context"
	"github.com/paysuper/paysuper-tax-service/proto"
)

type Service struct {
}

func (s *Service) SetRate(ctx context.Context, req *tax_service.TaxEntry, res *tax_service.EmptyResponse) error {
	panic("")
}

func (s *Service) DeleteRate(ctx context.Context, req *tax_service.RateLookupQuery, res *tax_service.EmptyResponse) error {
	panic("")
}

func (s *Service) GetSingleRate(ctx context.Context, req *tax_service.RateLookupQuery, res *tax_service.GetSingleRateResponse) error {
	panic("")
}

func (s *Service) GetRates(ctx context.Context, req *tax_service.RateLookupQuery, res *tax_service.GetRatesResponse) error {
	panic("")
}

func (s *Service) Status() (interface{}, error) {
	return "ok", nil
}
