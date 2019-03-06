package internal

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/paysuper/paysuper-tax-service/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

var (
	logger *zap.Logger
	db     *gorm.DB
)

func init() {
	logger, _ = zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

func setupTest(t *testing.T) (*Service, func(t *testing.T)) {
	config, err := LoadConfig()
	if err != nil {
		t.Fatal("config init failed with error", err)
	}

	db, err = gorm.Open("postgres", config.DSN)
	if err != nil {
		t.Fatal("failed to make Postgres connection", err)
	}

	service, err := NewService(db)
	if err != nil {
		t.Fatal("failed to create service", err)
	}

	return service, func(t *testing.T) {
		db.DropTableIfExists(&Tax{})
		db.Close()
	}
}

func Test_SetRateForUsFail(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()

	assert.Error(t, service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "US", Rate: 0.15}, &tax_service.TaxRate{}))
	assert.Error(t, service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "US", City: "Any", Rate: 0.15}, &tax_service.TaxRate{}))
	assert.Error(t, service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "US", City: "Any", State: "AL", Rate: 0.15}, &tax_service.TaxRate{}))
}

func Test_SetRateForUsSuccess(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()
	resRate := &tax_service.TaxRate{}
	assert.NoError(t, service.CreateOrUpdate(ctx, &tax_service.TaxRate{Zip: "12345", Country: "US", City: "Any", State: "AL", Rate: 0.15}, resRate))
	assert.NotZero(t, resRate.Id)

	res := &tax_service.GetRatesResponse{}
	assert.NoError(t, service.GetRates(ctx, &tax_service.GetRatesRequest{Zip: "12345"}, res))
	assert.Equal(t, 1, len(res.Rates))

	assert.Equal(t, "12345", res.Rates[0].Zip)
	assert.Equal(t, "US", res.Rates[0].Country)
	assert.Equal(t, "Any", res.Rates[0].City)
	assert.Equal(t, "AL", res.Rates[0].State)
	assert.EqualValues(t, 0.15, res.Rates[0].Rate)
}

func TestService_DeleteRate(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()

	rate := &tax_service.TaxRate{}
	service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "EN", City: "Any", Rate: 0.15}, rate)

	res := &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "EN"}, res)

	assert.NotEmpty(t, res.Rates)
	assert.NoError(t, service.DeleteRateById(ctx, &tax_service.DeleteRateRequest{Id: rate.Id}, &tax_service.DeleteRateResponse{}))

	res2 := &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "EN"}, res2)
	assert.Empty(t, res2.Rates)
}

func TestService_GetRates(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()

	service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "EN", City: "Any", Rate: 0.15}, &tax_service.TaxRate{})
	service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "EN", City: "Any2", Rate: 0.15}, &tax_service.TaxRate{})
	service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "EN", City: "Any3", Rate: 0.15}, &tax_service.TaxRate{})

	res := &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "EN"}, res)

	assert.Len(t, res.Rates, 3)

	res2 := &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{City: "Any2"}, res2)

	assert.Len(t, res2.Rates, 1)

	res3 := &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "EN", City: "Any3"}, res3)

	assert.Len(t, res3.Rates, 1)

	res4 := &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "US", City: "Any3"}, res4)

	assert.Empty(t, res4.Rates)
}
