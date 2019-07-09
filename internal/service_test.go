package internal

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/paysuper/paysuper-tax-service/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

var (
	logger *zap.Logger
	db     *gorm.DB
)

func init() {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zap.FatalLevel)

	logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom,
	))

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

func Test_SetRate(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()

	rate := &tax_service.TaxRate{}
	assert.NoError(t, service.CreateOrUpdate(ctx, &tax_service.TaxRate{Country: "RU", City: "Any", Rate: 0.15}, rate))

	rate.Rate = 0.20
	assert.NoError(t, service.CreateOrUpdate(ctx, rate, rate))
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

	createTaxRecord(t, ctx, service, "EN", "Any", "", "", 0.15)
	createTaxRecord(t, ctx, service, "EN", "Any2", "", "", 0.25)
	createTaxRecord(t, ctx, service, "EN", "Any3", "TS", "", 0.35)

	res := &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "EN"}, res)
	assert.Len(t, res.Rates, 3)

	res = &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{City: "Any2"}, res)
	assert.Len(t, res.Rates, 1)

	res = &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "EN", City: "Any3"}, res)
	assert.Len(t, res.Rates, 1)

	res = &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{State: "TS"}, res)
	assert.Len(t, res.Rates, 1)

	res = &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "EN", Offset: 1, Limit: 1}, res)
	assert.Len(t, res.Rates, 1)
	assert.EqualValues(t, 0.25, res.Rates[0].Rate)

	res = &tax_service.GetRatesResponse{}
	service.GetRates(ctx, &tax_service.GetRatesRequest{Country: "US", City: "Any3"}, res)
	assert.Empty(t, res.Rates)
}

func TestService_GetRate_Fail(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()

	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00001", 0.1)
	createTaxRecord(t, ctx, service, "GN", "London", "", "", 0.1)

	res := &tax_service.GetRateResponse{
		Rate: &tax_service.TaxRate{},
	}

	req := &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{},
		IpData: &tax_service.GeoIdentity{
			Country: "RU",
			City:    "Moscow",
		},
	}

	assert.Error(t, service.GetRate(ctx, req, res))

	req = &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{},
		IpData: &tax_service.GeoIdentity{
			Country: "EN",
		},
	}
	assert.Error(t, service.GetRate(ctx, req, res))

	req = &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{},
		IpData: &tax_service.GeoIdentity{
			Country: "US",
			City:    "City",
		},
	}
	assert.Error(t, service.GetRate(ctx, req, res))

	req = &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{},
		IpData: &tax_service.GeoIdentity{
			Country: "US",
			City:    "City",
			State:   "",
		},
	}
	assert.Error(t, service.GetRate(ctx, req, res))

	req = &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{},
		IpData: &tax_service.GeoIdentity{
			Zip: "00000",
		},
	}
	assert.Error(t, service.GetRate(ctx, req, res))
}

func TestService_GetRate_NonUserData(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()

	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00001", 0.1)
	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00002", 0.2)
	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00003", 0.3)
	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00004", 0.4)
	createTaxRecord(t, ctx, service, "GN", "London", "", "", 0.1)

	req := &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{},
		IpData: &tax_service.GeoIdentity{
			Country: "GN",
			City:    "London",
		},
	}

	res := &tax_service.GetRateResponse{
		Rate: &tax_service.TaxRate{},
	}

	assert.NoError(t, service.GetRate(ctx, req, res))
	assert.False(t, res.UserDataPriority)
	assert.EqualValues(t, 0.1, res.Rate.Rate)

	req = &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{},
		IpData: &tax_service.GeoIdentity{
			Country: "US",
			City:    "New York",
			State:   "NY",
		},
	}

	assert.NoError(t, service.GetRate(ctx, req, res))
	assert.False(t, res.UserDataPriority)
	assert.EqualValues(t, 0.4, res.Rate.Rate)

	req.IpData.Zip = "00001"

	assert.NoError(t, service.GetRate(ctx, req, res))
	assert.False(t, res.UserDataPriority)
	assert.EqualValues(t, 0.1, res.Rate.Rate)

	req.IpData.Zip = "00000"
	assert.NoError(t, service.GetRate(ctx, req, res))
	assert.False(t, res.UserDataPriority)
	assert.EqualValues(t, 0.4, res.Rate.Rate)
}

func TestService_GetRate_WithUserData(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	ctx := context.Background()

	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00001", 0.1)
	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00002", 0.2)
	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00003", 0.3)
	createTaxRecord(t, ctx, service, "US", "New York", "NY", "00004", 0.4)
	createTaxRecord(t, ctx, service, "GN", "London", "", "", 0.1)

	req := &tax_service.GetRateRequest{
		UserData: &tax_service.GeoIdentity{
			Country: "US",
			City:    "New York",
		},
		IpData: &tax_service.GeoIdentity{
			Country: "GN",
			City:    "London",
		},
	}

	res := &tax_service.GetRateResponse{
		Rate: &tax_service.TaxRate{},
	}

	assert.NoError(t, service.GetRate(ctx, req, res))
	assert.False(t, res.UserDataPriority)
	assert.EqualValues(t, 0.1, res.Rate.Rate)

	req.UserData.State = "NY"
	assert.NoError(t, service.GetRate(ctx, req, res))
	assert.True(t, res.UserDataPriority)
	assert.EqualValues(t, 0.4, res.Rate.Rate)
}

func TestService_Status(t *testing.T) {
	service, teardown := setupTest(t)
	defer teardown(t)

	status, err := service.Status()
	assert.EqualValues(t, "ok", status)
	assert.NoError(t, err)

	db.Close()

	status, err = service.Status()
	assert.EqualValues(t, "fail", status)
	assert.Error(t, err)
}

func createTaxRecord(t *testing.T, ctx context.Context, service *Service, country, city, state, zip string, rate float64) {
	t.Helper()

	stubResp := &tax_service.TaxRate{}
	tax := &tax_service.TaxRate{
		Zip:     zip,
		Country: country,
		City:    city,
		State:   state,
		Rate:    rate,
	}

	assert.NoError(t, service.CreateOrUpdate(ctx, tax, stubResp))
}
