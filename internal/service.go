package internal

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/paysuper/paysuper-tax-service/proto"
	"time"
)

type Tax struct {
	ID        uint32  `gorm:"primary_key"`
	Zip       string  `gorm:"type:varchar(5);index"`
	Country   string  `gorm:"type:varchar(2);not null;unique_index:idx_primary"`
	City      string  `gorm:"type:varchar(100);unique_index:idx_primary"`
	State     string  `gorm:"type:varchar(2);unique_index:idx_primary"`
	Rate      float32 `gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
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

func (s *Service) CreateOrUpdate(ctx context.Context, req *tax_service.TaxRate, res *tax_service.TaxRate) error {
	tax := Tax{
		ID:      req.Id,
		Zip:     req.Zip,
		Country: req.Country,
		City:    req.City,
		State:   req.State,
		Rate:    req.Rate,
	}

	if tax.Country == "US" && (tax.Zip == "" || tax.City == "" || tax.State == "") {
		return fmt.Errorf("invalid tax entry for US %v", tax)
	}

	err := s.db.Where(tax).FirstOrCreate(&tax, &tax).Error
	if err != nil {
		return err
	}

	res.Id = tax.ID
	res.Zip = tax.Zip
	res.Country = tax.Country
	res.City = tax.City
	res.State = tax.State
	res.Rate = tax.Rate

	return nil
}

func (s *Service) DeleteRateById(ctx context.Context, req *tax_service.DeleteRateRequest, res *tax_service.DeleteRateResponse) error {
	return s.db.Delete(&Tax{ID: req.Id}).Error
}

func (s *Service) GetRate(ctx context.Context, req *tax_service.GetRateRequest, res *tax_service.GetRateResponse) error {
	panic("")
}

func (s *Service) createGetQuery(zip, country, city string) (string, []interface{}) {
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
	}

	return query, args

}

func (s *Service) GetRates(ctx context.Context, req *tax_service.GetRatesRequest, res *tax_service.GetRatesResponse) error {
	query, args := s.createGetQuery(req.Zip, req.Country, req.City)

	request := s.db.Order("country desc").Where(query, args...)
	if req.Offset > 0 {
		request = request.Limit(req.Limit)
	}

	if req.Limit > 0 {
		request = request.Limit(req.Limit)
	}

	var rates []Tax

	err := request.Find(&rates).Error
	if err != nil {
		return err
	}

	for _, r := range rates {
		res.Rates = append(res.Rates, &tax_service.TaxRate{
			Zip:     r.Zip,
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
