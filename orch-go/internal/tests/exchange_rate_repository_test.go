package tests

import (
	"context"
	"errors"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/exchange_rate"
	"orch-go/internal/infrastructure/repository/exchange_rate_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestExchangeRateRepositoryGetAll_ReturnExchangeRates(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	pageN := int32(1)
	pageSize := int32(10)

	mockResp := &pb.GetAllExchangeRatesResponse{
		ExchangeRates: []*pb.ExchangeRateModel{
			{
				RateId:         []int32{1}[0],
				BaseCurrency:   &[]string{"USD"}[0],
				TargetCurrency: &[]string{"EUR"}[0],
				Rate:           &[]string{"1.1"}[0],
			},
			{
				RateId:         []int32{2}[0],
				BaseCurrency:   &[]string{"EUR"}[0],
				TargetCurrency: &[]string{"USD"}[0],
				Rate:           &[]string{"0.9"}[0],
			},
		},
	}

	req := &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
	}

	mockClient.
		EXPECT().
		GetAll(ctx, req).
		Return(mockResp, nil)

	exchangeRates, err := repo.GetAll(ctx, pageN, pageSize)

	require.NoError(t, err)
	require.Equal(t, 2, len(exchangeRates))
}

func TestExchangeRateRepositoryGetAll_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	pageN := int32(1)
	pageSize := int32(10)

	mockClient.
		EXPECT().
		GetAll(ctx, &pb.GetAllRequest{
			PageN:    pageN,
			PageSize: pageSize,
		}).
		Return(nil, errors.New("error"))

	exchangeRates, err := repo.GetAll(ctx, pageN, pageSize)

	require.Error(t, err)
	require.Nil(t, exchangeRates)
	require.Equal(t, "exchange_rate_repo.GetAll: error", err.Error())
}

func TestExchangeRateRepositoryGetById_ReturnExchangeRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	id := int32(1)

	mockResp := &pb.ExchangeRateModel{
		RateId:         id,
		BaseCurrency:   &[]string{"USD"}[0],
		TargetCurrency: &[]string{"EUR"}[0],
		Rate:           &[]string{"1.1"}[0],
	}

	req := &pb.GetExchangeRateByIdRequest{
		RateId: id,
	}

	mockClient.
		EXPECT().
		GetById(ctx, req).
		Return(mockResp, nil)

	exchangeRate, err := repo.GetById(ctx, id)

	require.NoError(t, err)
	require.Equal(t, id, exchangeRate.RateID)
}

func TestExchangeRateRepositoryGetById_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	id := int32(1)

	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetExchangeRateByIdRequest{RateId: id}).
		Return(nil, errors.New("error"))

	exchangeRate, err := repo.GetById(ctx, id)

	require.Error(t, err)
	require.Nil(t, exchangeRate)
	require.Equal(t, "exchange_rate_repo.GetById: error", err.Error())
}

func TestExchangeRateRepositoryAdd_ReturnExchangeRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	exchangeRate := &pb.ExchangeRateModel{
		RateId:         1,
		BaseCurrency:   &[]string{"USD"}[0],
		TargetCurrency: &[]string{"EUR"}[0],
		Rate:           &[]string{"1.1"}[0],
	}

	req := &pb.AddExchangeRateRequest{
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}

	addReq := exchange_rate.ExchangeRate{
		RateID:         exchangeRate.RateId,
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}

	mockClient.
		EXPECT().
		Add(ctx, req).
		Return(exchangeRate, nil)

	rate, err := repo.Add(ctx, &addReq)

	require.NoError(t, err)
	require.Equal(t, rate.RateID, addReq.RateID)
}

func TestExchangeRateRepositoryAdd_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	exchangeRate := &pb.ExchangeRateModel{
		RateId:         1,
		BaseCurrency:   &[]string{"USD"}[0],
		TargetCurrency: &[]string{"EUR"}[0],
		Rate:           &[]string{"1.1"}[0],
	}

	req := &pb.AddExchangeRateRequest{
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}

	addReq := exchange_rate.ExchangeRate{
		RateID:         exchangeRate.RateId,
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}

	mockClient.
		EXPECT().
		Add(ctx, req).
		Return(nil, errors.New("error"))

	rate, err := repo.Add(ctx, &addReq)

	require.Error(t, err)
	require.Nil(t, rate)
	require.Equal(t, "exchange_rate_repo.Add: error", err.Error())
}

func TestExchangeRateRepositoryUpdate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	exchangeRate := &pb.ExchangeRateModel{
		RateId:         1,
		BaseCurrency:   &[]string{"USD"}[0],
		TargetCurrency: &[]string{"EUR"}[0],
		Rate:           &[]string{"1.1"}[0],
	}

	req := &pb.UpdateExchangeRateRequest{
		RateId:         exchangeRate.RateId,
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}

	updateReq := exchange_rate.ExchangeRate{
		RateID:         exchangeRate.RateId,
		BaseCurrency:   exchangeRate.BaseCurrency,
		TargetCurrency: exchangeRate.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}

	mockClient.
		EXPECT().
		Update(ctx, req).
		Return(nil, errors.New("error"))

	err := repo.Update(ctx, &updateReq)

	require.Error(t, err)
	require.Equal(t, "exchange_rate_repo.Update: error", err.Error())
}

func TestExchangeRateRepositoryDelete_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	id := int32(1)

	mockClient.
		EXPECT().
		Delete(ctx, &pb.DeleteExchangeRateRequest{RateId: id}).
		Return(nil, errors.New("error"))

	err := repo.Delete(ctx, id)

	require.Error(t, err)
	require.Equal(t, "exchange_rate_repo.Delete: error", err.Error())
}

func TestExchangeRateRepositoryAddBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	rates := []*exchange_rate.ExchangeRate{
		{
			RateID:         1,
			BaseCurrency:   &[]string{"USD"}[0],
			TargetCurrency: &[]string{"EUR"}[0],
			Rate:           &[]string{"1.1"}[0],
		},
	}

	mockClient.
		EXPECT().
		AddBulk(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.AddBulk(ctx, rates)

	require.Error(t, err)
	require.Equal(t, "exchange_rate_repo.AddBulk: error", err.Error())
}

func TestExchangeRateRepositoryUpdateBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	rates := []*exchange_rate.ExchangeRate{
		{
			RateID:         1,
			BaseCurrency:   &[]string{"USD"}[0],
			TargetCurrency: &[]string{"EUR"}[0],
			Rate:           &[]string{"1.1"}[0],
		},
	}

	mockClient.
		EXPECT().
		UpdateBulk(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.UpdateBulk(ctx, rates)

	require.Error(t, err)
	require.Equal(t, "exchange_rate_repo.UpdateBulk: error", err.Error())
}

func TestExchangeRateRepositoryDeleteBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockExchangeRateServiceClient(ctrl)
	repo := exchange_rate_repo.NewRepository(mockClient)

	ctx := context.Background()

	ids := []int32{1}

	mockClient.
		EXPECT().
		DeleteBulk(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.DeleteBulk(ctx, ids)

	require.Error(t, err)
	require.Equal(t, "exchange_rate_repo.DeleteBulk: error", err.Error())
}
