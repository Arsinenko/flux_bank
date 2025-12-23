package tests

import (
	"context"
	"errors"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/account"
	"orch-go/internal/infrastructure/repository/account/account_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.uber.org/mock/gomock"
)

func TestRepository_GetByCustomerId_ReturnsAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)

	repo := account_repo.NewRepository(mockClient)

	ctx := context.Background()
	customerId := int32(1)
	balance := "1000"
	isActive := true
	typeId := int32(1)
	accountModel := &pb.AccountModel{
		AccountId:  1,
		CustomerId: &customerId,
		TypeId:     &typeId,
		Iban:       "DE123",
		Balance:    &balance,
		CreatedAt:  timestamppb.New(time.Now()),
		IsActive:   &isActive,
	}

	mockResp := &pb.GetAllAccountsResponse{
		Accounts: []*pb.AccountModel{
			accountModel,
		},
	}

	mockClient.
		EXPECT().
		GetByCustomerId(ctx, &pb.GetAccountByCustomerIdRequest{
			CustomerId: customerId,
		}).
		Return(mockResp, nil)

	accounts, err := repo.GetByCustomerId(ctx, customerId)
	require.NoError(t, err)
	require.Len(t, accounts, 1)
	require.Equal(t, int32(1), accounts[0].Id)
	require.Equal(t, "DE123", accounts[0].Iban)
}

func TestRepository_GetByCustomerId_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)

	repo := account_repo.NewRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		GetById(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	accounts, err := repo.GetById(ctx, int32(1))
	require.Error(t, err)
	require.Nil(t, accounts)
	require.Equal(t, "account_repo.GetById: grpc error", err.Error())
}

func TestRepository_GetByDateRange_ReturnsAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)
	repo := account_repo.NewRepository(mockClient)

	ctx := context.Background()
	req := pb.GetByDateRangeRequest{
		FromDate: timestamppb.New(time.Now()),
		ToDate:   timestamppb.New(time.Now().Add(time.Hour * 24)),
		PageN:    &[]int32{1}[0],
		PageSize: &[]int32{1}[0],
	}
	mockResp := &pb.GetAllAccountsResponse{
		Accounts: []*pb.AccountModel{
			{
				AccountId:  1,
				CustomerId: &[]int32{1}[0],
				TypeId:     &[]int32{1}[0],
				Iban:       "iban",
				Balance:    &[]string{"1000"}[0],
				CreatedAt:  timestamppb.New(time.Now()),
				IsActive:   &[]bool{true}[0],
			},
		},
	}

	mockClient.
		EXPECT().
		GetByDateRange(ctx, &req).
		Return(mockResp, nil)

	accounts, err := repo.GetByDateRange(ctx, account.GetByDateRange{
		From:     time.Now(),
		To:       time.Now().Add(time.Hour * 24),
		PageN:    1,
		PageSize: 1,
	})

	require.NoError(t, err)
	require.Len(t, accounts, 1)
}

func TestRepository_GetByDateRange_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)

	repo := account_repo.NewRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		GetByDateRange(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	accounts, err := repo.GetByDateRange(ctx, account.GetByDateRange{})
	require.Error(t, err)
	require.Nil(t, accounts)
	require.Equal(t, "account_repo.GetByDateRange: grpc error", err.Error())
}

func TestRepository_GetAll_ReturnsAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)
	repo := account_repo.NewRepository(mockClient)

	ctx := context.Background()
	req := pb.GetAllRequest{
		PageN:    1,
		PageSize: 1,
	}
	mockResp := &pb.GetAllAccountsResponse{
		Accounts: []*pb.AccountModel{
			{
				AccountId:  1,
				CustomerId: &[]int32{1}[0],
				TypeId:     &[]int32{1}[0],
				Iban:       "iban",
				Balance:    &[]string{"1000"}[0],
				CreatedAt:  timestamppb.New(time.Now()),
				IsActive:   &[]bool{true}[0],
			},
		},
	}
	mockClient.
		EXPECT().
		GetAll(ctx, &req).
		Return(mockResp, nil)

	accounts, err := repo.GetAll(ctx, 1, 1)
	require.NoError(t, err)
	require.Len(t, accounts, 1)
}

func TestRepository_GetAll_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)

	repo := account_repo.NewRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		GetAll(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	accounts, err := repo.GetAll(ctx, 1, 1)
	require.Error(t, err)
	require.Nil(t, accounts)
	require.Equal(t, "account_repo.GetAll: grpc error", err.Error())
}

func TestRepository_GetById_ReturnsAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)

	repo := account_repo.NewRepository(mockClient)

	ctx := context.Background()
	accountId := int32(1)
	typeId := int32(1)
	mockResp := &pb.AccountModel{
		AccountId:  1,
		CustomerId: &[]int32{1}[0],
		TypeId:     &typeId,
		Iban:       "DE123",
		Balance:    &[]string{"1000"}[0],
		CreatedAt:  timestamppb.New(time.Now()),
		IsActive:   &[]bool{true}[0],
	}

	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetAccountByIdRequest{
			AccountId: accountId,
		}).
		Return(mockResp, nil)

	account, err := repo.GetById(ctx, accountId)
	require.NoError(t, err)
	require.Equal(t, int32(1), account.Id)
	require.Equal(t, "DE123", account.Iban)
}

func TestGetByIdNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)
	repo := account_repo.NewRepository(mockClient)

	ctx := context.Background()
	accountId := int32(1)

	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetAccountByIdRequest{
			AccountId: accountId,
		}).
		Return(nil, errors.New("account not found"))

	account, err := repo.GetById(ctx, accountId)
	require.Error(t, err)
	require.Nil(t, account)
}

func TestCreate_ReturnsAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)
	repo := account_repo.NewRepository(mockClient)

	ctx := context.Background()
	req := pb.AddAccountRequest{
		CustomerId: &[]int32{1}[0],
		TypeId:     &[]int32{1}[0],
		Iban:       "DE123",
		Balance:    &[]string{"1000"}[0],
		IsActive:   &[]bool{true}[0],
	}
	mockResp := &pb.AccountModel{
		AccountId:  1,
		CustomerId: &[]int32{1}[0],
		TypeId:     &[]int32{1}[0],
		Iban:       "DE123",
		Balance:    &[]string{"1000"}[0],
		CreatedAt:  timestamppb.New(time.Now()),
		IsActive:   &[]bool{true}[0],
	}

	mockClient.
		EXPECT().
		Add(ctx, &req).
		Return(mockResp, nil)

	accountModel := account_repo.AccountToDomain(mockResp)
	account, err := repo.Create(ctx, accountModel)

	require.NoError(t, err)
	require.Equal(t, int32(1), account.Id)
	require.Equal(t, "DE123", account.Iban)
}

func TestCreate_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)
	repo := account_repo.NewRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		Add(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	accountModel := account.Account{
		CustomerId: 1,
		TypeId:     1,
		Iban:       "DE123",
		Balance:    "1000",
		IsActive:   true,
	}
	account, err := repo.Create(ctx, &accountModel)

	require.Error(t, err)
	require.Nil(t, account)
}

func TestUpdate_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)
	repo := account_repo.NewRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		Update(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	accountModel := account.Account{
		Id:         1,
		CustomerId: 1,
		TypeId:     1,
		Iban:       "DE123",
		Balance:    "1000",
		IsActive:   true,
	}
	err := repo.Update(ctx, &accountModel)

	require.Error(t, err)
	require.Equal(t, "account_repo.Update: grpc error", err.Error())
}

func TestDelete_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountServiceClient(ctrl)
	repo := account_repo.NewRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		Delete(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.Delete(ctx, 1)

	require.Error(t, err)
	require.Equal(t, "account_repo.Delete: grpc error", err.Error())
}
