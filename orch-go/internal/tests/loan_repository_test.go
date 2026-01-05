package tests

import (
	"context"
	"errors"
	protos "orch-go/api/generated"
	"orch-go/internal/domain/loan"
	"orch-go/internal/infrastructure/repository/loan_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestLoanRepositoryGetAll_ReturnsLoans(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 1)
	mockResp := &protos.GetAllLoansResponse{
		Loans: []*protos.LoanModel{
			{
				LoanId:       1,
				CustomerId:   &[]int32{1}[0],
				Principal:    &[]string{"1000"}[0],
				InterestRate: &[]string{"5.0"}[0],
				Status:       &[]string{"Active"}[0],
				StartDate:    loan_repo.ToDateOnly(&startDate),
				EndDate:      loan_repo.ToDateOnly(&endDate),
			},
		},
	}

	mockClient.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(mockResp, nil)

	loans, err := repo.GetAll(ctx, 1, 1)
	require.NoError(t, err)
	require.Equal(t, 1, len(loans))
}

func TestLoanRepositoryGetAll_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	mockClient.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

	loans, err := repo.GetAll(ctx, 1, 1)
	require.Error(t, err)
	require.Nil(t, loans)
	require.Equal(t, "loan_repo.GetAll: error", err.Error())
}

func TestLoanRepositoryGetById_ReturnsLoan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 1)
	mockResp := &protos.LoanModel{
		LoanId:       1,
		CustomerId:   &[]int32{1}[0],
		Principal:    &[]string{"1000"}[0],
		InterestRate: &[]string{"5.0"}[0],
		Status:       &[]string{"Active"}[0],
		StartDate:    loan_repo.ToDateOnly(&startDate),
		EndDate:      loan_repo.ToDateOnly(&endDate),
	}

	mockClient.EXPECT().GetById(gomock.Any(), &protos.GetLoanByIdRequest{LoanId: 1}).Return(mockResp, nil)

	loan, err := repo.GetById(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, int32(1), loan.LoanID)
}

func TestLoanRepositoryGetById_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	mockClient.EXPECT().GetById(gomock.Any(), &protos.GetLoanByIdRequest{LoanId: 1}).Return(nil, errors.New("error"))

	loan, err := repo.GetById(ctx, 1)
	require.Error(t, err)
	require.Nil(t, loan)
	require.Equal(t, "loan_repo.GetById: error", err.Error())
}

func TestLoanRepositoryAdd_ReturnsLoan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 1)
	mockResp := &protos.LoanModel{
		LoanId:       1,
		CustomerId:   &[]int32{1}[0],
		Principal:    &[]string{"1000"}[0],
		InterestRate: &[]string{"5.0"}[0],
		Status:       &[]string{"Active"}[0],
		StartDate:    loan_repo.ToDateOnly(&startDate),
		EndDate:      loan_repo.ToDateOnly(&endDate),
	}

	mockClient.EXPECT().Add(gomock.Any(), &protos.AddLoanRequest{
		CustomerId:   &[]int32{1}[0],
		Principal:    &[]string{"1000"}[0],
		InterestRate: &[]string{"5.0"}[0],
		Status:       &[]string{"Active"}[0],
		StartDate:    loan_repo.ToDateOnly(&startDate),
		EndDate:      loan_repo.ToDateOnly(&endDate),
	}).Return(mockResp, nil)

	loan, err := repo.Add(ctx, &loan.Loan{
		CustomerID:   &[]int32{1}[0],
		Principal:    &[]string{"1000"}[0],
		InterestRate: &[]string{"5.0"}[0],
		Status:       &[]string{"Active"}[0],
		StartDate:    &startDate,
		EndDate:      &endDate,
	})
	require.NoError(t, err)
	require.Equal(t, int32(1), loan.LoanID)
}

func TestLoanRepositoryAdd_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 1)

	mockClient.EXPECT().Add(gomock.Any(), &protos.AddLoanRequest{
		CustomerId:   &[]int32{1}[0],
		Principal:    &[]string{"1000"}[0],
		InterestRate: &[]string{"5.0"}[0],
		Status:       &[]string{"Active"}[0],
		StartDate:    loan_repo.ToDateOnly(&startDate),
		EndDate:      loan_repo.ToDateOnly(&endDate),
	}).Return(nil, errors.New("error"))

	loan, err := repo.Add(ctx, &loan.Loan{
		CustomerID:   &[]int32{1}[0],
		Principal:    &[]string{"1000"}[0],
		InterestRate: &[]string{"5.0"}[0],
		Status:       &[]string{"Active"}[0],
		StartDate:    &startDate,
		EndDate:      &endDate,
	})
	require.Error(t, err)
	require.Nil(t, loan)
	require.Equal(t, "loan_repo.Add: error", err.Error())
}

func TestLoanRepositoryUpdate_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 1)

	mockClient.EXPECT().Update(ctx, gomock.Any()).Return(nil, errors.New("error"))

	err := repo.Update(ctx, &loan.Loan{
		LoanID:       1,
		CustomerID:   &[]int32{1}[0],
		Principal:    &[]string{"1000"}[0],
		InterestRate: &[]string{"5.0"}[0],
		Status:       &[]string{"Active"}[0],
		StartDate:    &startDate,
		EndDate:      &endDate,
	})
	require.Error(t, err)
	require.Equal(t, "loan_repo.Update: error", err.Error())
}

func TestLoanRepositoryDelete_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	mockClient.EXPECT().Delete(ctx, gomock.Any()).Return(nil, errors.New("error"))

	err := repo.Delete(ctx, 1)
	require.Error(t, err)
	require.Equal(t, "loan_repo.Delete: error", err.Error())
}

func TestLoanRepositoryDeleteBulk_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockLoanServiceClient(ctrl)
	repo := loan_repo.NewLoanRepository(mockClient)

	ctx := context.Background()

	mockClient.EXPECT().DeleteBulk(ctx, gomock.Any()).Return(nil, errors.New("error"))

	err := repo.DeleteBulk(ctx, []int32{1})
	require.Error(t, err)
	require.Equal(t, "loan_repo.DeleteBulk: error", err.Error())
}
