package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/paysuper/paysuper-tax-service/proto"
	"go.uber.org/zap"
	"time"
)

// Tax defines one tax record in service database.
type Tax struct {
	ID        uint32  `gorm:"primary_key"`
	Zip       string  `gorm:"type:varchar(5);unique_index:idx_primary"`
	Country   string  `gorm:"type:varchar(2);not null;unique_index:idx_primary"`
	City      string  `gorm:"type:varchar(100);unique_index:idx_primary"`
	State     string  `gorm:"type:varchar(2);unique_index:idx_primary"`
	Rate      float32 `gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Service is application entry point.s
type Service struct {
	db *gorm.DB
}

// NewService create new Service.
func NewService(db *gorm.DB) (*Service, error) {
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(
		&Tax{},
	)

	return &Service{db: db}, nil
}

// CreateOrUpdate used to create new or updated existing Tax record in database.
func (s *Service) CreateOrUpdate(ctx context.Context, req *tax_service.TaxRate, res *tax_service.TaxRate) error {
	if req.Country == "US" && (req.Zip == "" || req.State == "") {
		zap.S().Error("invalid CreateOrUpdate request for US", "taxRate", req)
		return fmt.Errorf("invalid tax entry for US %v", req)
	}

	tax := fromTaxRate(req)

	var err error
	if req.Id != 0 {
		err = s.db.First(tax).Error
	} else {
		query, args := s.createGetRatesQuery(req.Zip, req.Country, req.City, req.State)
		err = s.db.Where(query, args...).First(tax).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.S().Error("fail to query rate in CreateOrUpdate", "taxRate", req, "error", err)
		return err
	}

	err = s.db.Save(tax).Error
	if err != nil {
		zap.S().Error("fail to save rate in CreateOrUpdate", "tax", tax, "error", err)
		return err
	}

	copyToTaxRate(res, tax)
	return nil
}

// DeleteRateById mark Tax record as deleted in database.
func (s *Service) DeleteRateById(ctx context.Context, req *tax_service.DeleteRateRequest, res *tax_service.DeleteRateResponse) error {
	err := s.db.Delete(&Tax{ID: req.Id}).Error
	if err != nil {
		zap.S().Error("fail to delete rate", "request", req, "error", err)
	}

	return err
}

// GetRate returns one single Tax record by given user provided and ip based geo information.
func (s *Service) GetRate(ctx context.Context, req *tax_service.GetRateRequest, res *tax_service.GetRateResponse) error {
	res.UserDataPriority = req.UserData.Country != "" && req.UserData.City != ""
	if !res.UserDataPriority {
		return s.getRate(req.IpData, res)
	}
	var err error

	err = s.getRate(req.UserData, res)
	if err != nil {
		res.UserDataPriority = false
		return s.getRateByCountry(req.IpData, res)
	}

	return nil
}

func (s *Service) getRate(req *tax_service.GeoIdentity, res *tax_service.GetRateResponse) error {
	if req.Zip == "" {
		return s.getRateByCountry(req, res)
	}

	err := s.getRateByZip(req, res)
	if err == nil {
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		zap.S().Warn("fail to get rates for zip", "geo_identity", req)
		return s.getRateByCountry(req, res)
	}
	return err
}

func (s *Service) getRateByCountry(req *tax_service.GeoIdentity, res *tax_service.GetRateResponse) error {
	if req.Country == "" {
		zap.S().Error("country empty in getRateByCountry", "geo_identity", req)
		return errors.New("country empty in getRateByCountry")
	}

	if req.City == "" {
		zap.S().Error("city empty in getRateByCountry", "geo_identity", req)
		return errors.New("city empty in getRateByCountry")
	}

	if req.Country == "US" && req.State == "" {
		zap.S().Error("state empty for US in getRateByCountry", "geo_identity", req)
		return errors.New("state empty for US in getRateByCountry")
	}

	rate := &Tax{}

	err := s.db.Where(
		"(country = ? AND city = ? AND state = ?) OR (country = ? AND state = ?)",
		req.Country,
		req.City,
		req.State,
		req.Country,
		req.State,
	).Order("rate desc").First(rate).Error

	if err == nil {
		res.Rate = toTaxRate(rate)
		return nil
	}

	if err == gorm.ErrRecordNotFound {
		zap.S().Error("tax rates for given request unavailable", "get_identity", req, "error", err)
		return errors.New("tax rates for given request unavailable")
	}

	return err
}

func (s *Service) getRateByZip(req *tax_service.GeoIdentity, res *tax_service.GetRateResponse) error {
	rate := &Tax{}

	err := s.db.Where("zip = ?", req.Zip).First(rate).Error
	if err != nil {
		zap.S().Warn("tax rates by zip for given request unavailable", "get_identity", req, "error", err)
		return err
	}

	res.Rate = toTaxRate(rate)
	return nil
}

// GetRates return list of Tax from database based on provided filters.
func (s *Service) GetRates(ctx context.Context, req *tax_service.GetRatesRequest, res *tax_service.GetRatesResponse) error {
	query, args := s.createGetRatesQuery(req.Zip, req.Country, req.City, req.State)

	request := s.db.Order("country asc, rate desc").Where(query, args...)
	if req.Offset > 0 {
		request = request.Offset(req.Offset)
	}

	if req.Limit > 0 {
		request = request.Limit(req.Limit)
	}

	var rates []Tax

	err := request.Find(&rates).Error
	if err != nil {
		zap.S().Warn("get rates search failed", "request", req, "error", err)
		return err
	}

	for _, r := range rates {
		res.Rates = append(res.Rates, toTaxRate(&r))
	}

	return nil
}

func (s *Service) createGetRatesQuery(zip, country, city, state string) (string, []interface{}) {
	var query string
	var args []interface{}

	if zip != "" {
		query = "zip = ?"
		args = append(args, zip)
	} else {
		query = "true"
		if country != "" {
			query = query + " AND country = ?"
			args = append(args, country)
		}

		if city != "" {
			query = query + " AND city = ?"
			args = append(args, city)
		}

		if state != "" {
			query = query + " AND state = ?"
			args = append(args, state)
		}
	}

	return query, args
}

// Status used to return micro service health.
func (s *Service) Status() (interface{}, error) {
	err := s.db.DB().Ping()
	if err != nil {
		return "fail", err
	}

	return "ok", nil
}

func fromTaxRate(req *tax_service.TaxRate) *Tax {
	return &Tax{
		ID:      req.Id,
		Zip:     req.Zip,
		Country: req.Country,
		City:    req.City,
		State:   req.State,
		Rate:    req.Rate,
	}
}

func toTaxRate(source *Tax) *tax_service.TaxRate {
	return &tax_service.TaxRate{
		Id:      source.ID,
		Zip:     source.Zip,
		Country: source.Country,
		City:    source.City,
		State:   source.State,
		Rate:    source.Rate,
	}

}
func copyToTaxRate(target *tax_service.TaxRate, source *Tax) {
	target.Id = source.ID
	target.Zip = source.Zip
	target.Country = source.Country
	target.City = source.City
	target.State = source.State
	target.Rate = source.Rate
}
