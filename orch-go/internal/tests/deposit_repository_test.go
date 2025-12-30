package tests

import (
	"context"
	"errors"
	"orch-go/internal/domain/deposit"
	"orch-go/internal/infrastructure/repository/deposit_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"
	"time"

	pb "orch-go/api/generated"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestDepositRepositoryGetById_ReturnDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	depositId := int32(1)

	mockResp := &pb.DepositModel{
		DepositId:    depositId,
		CustomerId:   &[]int32{1}[0],
		Amount:       &[]string{"100"}[0],
		InterestRate: &[]string{"str"}[0],
		Status:       &[]string{"status"}[0],
	}

	req := &pb.GetDepositByIdRequest{
		DepositId: depositId,
	}

	mockClient.
		EXPECT().
		GetById(ctx, req).
		Return(mockResp, nil)

	deposit, err := repo.GetById(ctx, depositId)

	require.NoError(t, err)
	require.Equal(t, depositId, deposit.DepositID)
}

func TestDepositRepositoryGetById_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetById(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	deposit, err := repo.GetById(ctx, int32(1))
	require.Error(t, err)
	require.Nil(t, deposit)
	require.Equal(t, "deposit_repo.GetById: grpc error", err.Error())
}

func TestDepositRepositoryGetAll_ReturnDeposits(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	pageN := int32(1)
	pageSize := int32(10)

	mockResp := &pb.GetAllDepositsResponse{
		Deposits: []*pb.DepositModel{
			{
				DepositId:    []int32{1}[0],
				CustomerId:   &[]int32{1}[0],
				Amount:       &[]string{"100"}[0],
				InterestRate: &[]string{"str"}[0],
				Status:       &[]string{"status"}[0],
			},
			{
				DepositId:    []int32{2}[0],
				CustomerId:   &[]int32{2}[0],
				Amount:       &[]string{"100"}[0],
				InterestRate: &[]string{"str"}[0],
				Status:       &[]string{"status"}[0],
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

	deposits, err := repo.GetAll(ctx, pageN, pageSize)

	require.NoError(t, err)
	require.Equal(t, 2, len(deposits))
}

func TestDepositRepositoryGetAll_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetAll(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	deposits, err := repo.GetAll(ctx, int32(1), int32(10))
	require.Error(t, err)
	require.Nil(t, deposits)
	require.Equal(t, "deposit_repo.GetAll: grpc error", err.Error())
}

func TestDepositRepositoryGetByCustomer_ReturnDeposits(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	customerId := int32(1)

	mockResp := &pb.GetAllDepositsResponse{
		Deposits: []*pb.DepositModel{
			{
				DepositId:    []int32{1}[0],
				CustomerId:   &[]int32{1}[0],
				Amount:       &[]string{"100"}[0],
				InterestRate: &[]string{"str"}[0],
				Status:       &[]string{"status"}[0],
			},
			{
				DepositId:    []int32{2}[0],
				CustomerId:   &[]int32{2}[0],
				Amount:       &[]string{"100"}[0],
				InterestRate: &[]string{"str"}[0],
				Status:       &[]string{"status"}[0],
			},
		},
	}

	req := &pb.GetDepositsByCustomerRequest{
		CustomerId: customerId,
	}

	mockClient.
		EXPECT().
		GetByCustomer(ctx, req).
		Return(mockResp, nil)

	deposits, err := repo.GetByCustomer(ctx, customerId)

	require.NoError(t, err)
	require.Equal(t, 2, len(deposits))
}

func TestDepositRepositoryGetByCustomer_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	customerId := int32(1)

	mockClient.
		EXPECT().
		GetByCustomer(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	deposits, err := repo.GetByCustomer(ctx, customerId)
	require.Error(t, err)
	require.Nil(t, deposits)
	require.Equal(t, "deposit_repo.GetByCustomer: grpc error", err.Error())
}

func TestDepositRepositoryAdd_ReturnDeposit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	deposit := deposit.Deposit{
		DepositID:    1,
		CustomerID:   1,
		Amount:       "100",
		InterestRate: "str",
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(1, 0, 0),
		Status:       "status",
	}

	mockResp := &pb.DepositModel{
		DepositId:    deposit.DepositID,
		CustomerId:   &deposit.CustomerID,
		Amount:       &deposit.Amount,
		InterestRate: &deposit.InterestRate,
		StartDate:    deposit_repo.ToDateOnly(&deposit.StartDate),
		EndDate:      deposit_repo.ToDateOnly(&deposit.EndDate),
		Status:       &deposit.Status,
	}

	req := &pb.AddDepositRequest{
		CustomerId:   &deposit.CustomerID,
		Amount:       &deposit.Amount,
		InterestRate: &deposit.InterestRate,
		StartDate:    deposit_repo.ToDateOnly(&deposit.StartDate),
		EndDate:      deposit_repo.ToDateOnly(&deposit.EndDate),
		Status:       &deposit.Status,
	}

	mockClient.
		EXPECT().
		Add(ctx, req).
		Return(mockResp, nil)

	result, err := repo.Add(ctx, &deposit)

	require.NoError(t, err)
	require.Equal(t, deposit.DepositID, result.DepositID)
}

func TestDepositRepositoryAdd_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Add(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	deposit := deposit.Deposit{
		DepositID:    1,
		CustomerID:   1,
		Amount:       "100",
		InterestRate: "str",
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(1, 0, 0),
		Status:       "status",
	}

	_, err := repo.Add(ctx, &deposit)
	require.Error(t, err)
	require.Equal(t, "deposit_repo.Add: grpc error", err.Error())
}

func TestDepositRepositoryUpdate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	deposit := deposit.Deposit{
		DepositID:    1,
		CustomerID:   1,
		Amount:       "100",
		InterestRate: "str",
		StartDate:    time.Now(),
		EndDate:      time.Now().AddDate(1, 0, 0),
		Status:       "status",
	}

	mockClient.
		EXPECT().
		Update(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.Update(ctx, &deposit)
	require.Error(t, err)
	require.Equal(t, "deposit_repo.Update: grpc error", err.Error())
}

func TestDepositRepositoryDelete_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	id := int32(1)

	mockClient.
		EXPECT().
		Delete(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.Delete(ctx, id)
	require.Error(t, err)
	require.Equal(t, "deposit_repo.Delete: grpc error", err.Error())
}

func TestDepositRepositoryAddBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	deposits := []*deposit.Deposit{
		{
			DepositID:    1,
			CustomerID:   1,
			Amount:       "100",
			InterestRate: "str",
			StartDate:    time.Now(),
			EndDate:      time.Now().AddDate(1, 0, 0),
			Status:       "status",
		},
		{
			DepositID:    2,
			CustomerID:   2,
			Amount:       "100",
			InterestRate: "str",
			StartDate:    time.Now(),
			EndDate:      time.Now().AddDate(1, 0, 0),
			Status:       "status",
		},
	}

	mockClient.
		EXPECT().
		AddBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.AddBulk(ctx, deposits)
	require.Error(t, err)
	require.Equal(t, "deposit_repo.AddBulk: grpc error", err.Error())
}

func TestDepositRepositoryUpdateBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	deposits := []*deposit.Deposit{
		{
			DepositID:    1,
			CustomerID:   1,
			Amount:       "100",
			InterestRate: "str",
			StartDate:    time.Now(),
			EndDate:      time.Now().AddDate(1, 0, 0),
			Status:       "status",
		},
		{
			DepositID:    2,
			CustomerID:   2,
			Amount:       "100",
			InterestRate: "str",
			StartDate:    time.Now(),
			EndDate:      time.Now().AddDate(1, 0, 0),
			Status:       "status",
		},
	}

	mockClient.
		EXPECT().
		UpdateBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.UpdateBulk(ctx, deposits)
	require.Error(t, err)
	require.Equal(t, "deposit_repo.UpdateBulk: grpc error", err.Error())
}

func TestDepositRepositoryDeleteBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockDepositServiceClient(ctrl)
	repo := deposit_repo.NewRepository(mockClient)

	ctx := context.Background()

	ids := []int32{1, 2}

	mockClient.
		EXPECT().
		DeleteBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.DeleteBulk(ctx, ids)
	require.Error(t, err)
	require.Equal(t, "deposit_repo.DeleteBulk: grpc error", err.Error())
}
