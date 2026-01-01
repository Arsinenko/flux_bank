package tests

import (
	"context"
	"errors"
	pb "orch-go/api/generated"
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
