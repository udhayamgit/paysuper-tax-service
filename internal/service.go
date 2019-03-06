package internal

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/paysuper/paysuper-tax-service/proto"
)

type Tax struct {
	gorm.Model
	Zip     string  `gorm:"type:varchar(5);index"`
	Country string  `gorm:"type:varchar(2);not null;unique_index:idx_primary"`
	City    string  `gorm:"type:varchar(100);unique_index:idx_primary"`
	State   string  `gorm:"type:varchar(2);unique_index:idx_primary"`
	Rate    float32 `gorm:"type:decimal(10,2);not null"`
}

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) (*Service, error) {
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(
		&Tax{},
	)

	return &Service{db: db}, nil
}

func (s *Service) SetRate(ctx context.Context, req *tax_service.TaxRate, res *tax_service.EmptyResponse) error {
	tax := Tax{
		Zip:     req.ZipCode,
		Country: req.Country,
		City:    req.City,
		State:   req.State,
		Rate:    req.Rate,
	}

	if tax.Country == "US" && (tax.Zip == "" || tax.City == "" || tax.State == "") {
		return fmt.Errorf("invalid tax entry for US %v", tax)
	}

	return s.db.Where(tax).FirstOrCreate(&tax).Error
}

func (s *Service) DeleteRate(ctx context.Context, req *tax_service.RateLookupQuery, res *tax_service.EmptyResponse) error {
	tax := Tax{
		Zip:     req.ZipCode,
		Country: req.Country,
		City:    req.City,
		State:   req.State,
	}
	return s.db.Delete(tax).Error
}

func (s *Service) GetSingleRate(ctx context.Context, req *tax_service.GetRateRequest, res *tax_service.GetRateResponse) error {
	panic("")
}

func (s *Service) GetRates(ctx context.Context, req *tax_service.GetRatesRequest, res *tax_service.GetRatesResponse) error {
	var query string
	var args []interface{}

	if req.ZipCode != "" {
		query = "zip = ?"
		args = append(args, req.ZipCode)
	} else {
		query = "true"
		if req.Country != "" {
			query = query + " AND country = ?"
			args = append(args, req.Country)
		}

		if req.City != "" {
			query = query + " AND city = ?"
			args = append(args, req.City)
		}
	}

	var limit int32
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = -1
	}

	var offset int32
	if req.Offset > 0 {
		offset = req.Offset
	} else {
		offset = -1
	}

	var rates []Tax
	err := s.db.Order("country desc").Where(query, args...).Limit(limit).Offset(offset).Find(&rates).Error
	if err != nil {
		return err
	}

	for _, r := range rates {
		res.Rates = append(res.Rates, &tax_service.TaxRate{
			ZipCode: r.Zip,
			Country: r.Country,
			City:    r.City,
			State:   r.State,
			Rate:    r.Rate,
		})
	}

	return nil
}

func (s *Service) Status() (interface{}, error) {
	err := s.db.DB().Ping()
	if err != nil {
		return "fail", err
	}

	return "ok", nil
}
