package tests

import (
	"context"
	"errors"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/loan"
	"orch-go/internal/infrastructure/repository/loan_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestLoanPaymentRepositoryGetAll_ReturnsPayments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockClient := mock_protos.NewMockLoanPaymentServiceClient(ctrl)
	repo := loan_repo.NewLoanPaymentRepository(mockClient)

	paymentDate := time.Now()
	mockResp := &pb.GetAllLoanPaymentsResponse{
		LoanPayments: []*pb.LoanPaymentModel{
			{
				PaymentId:   1,
				LoanId:      &[]int32{1}[0],
				Amount:      &[]string{"1000"}[0],
				PaymentDate: loan_repo.ToDateOnly(&paymentDate),
				IsPaid:      &[]bool{true}[0],
			},
		},
	}

	mockClient.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(mockResp, nil)

	loans, err := repo.GetAll(ctx, 1, 1)
	require.NoError(t, err)
	require.Equal(t, 1, len(loans))
}

func TestLoanPaymentRepositoryGetAll_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockClient := mock_protos.NewMockLoanPaymentServiceClient(ctrl)
	repo := loan_repo.NewLoanPaymentRepository(mockClient)

	mockClient.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

	loans, err := repo.GetAll(ctx, 1, 1)

	require.Nil(t, loans)
	require.Error(t, err)
	require.Equal(t, "loan_payment_repo.GetAll: error", err.Error())
}

func TestLoanPaymentRepositoryGetById_ReturnsPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockClient := mock_protos.NewMockLoanPaymentServiceClient(ctrl)
	repo := loan_repo.NewLoanPaymentRepository(mockClient)

	paymentDate := time.Now()

	mockResp := &pb.LoanPaymentModel{
		PaymentId:   1,
		LoanId:      &[]int32{1}[0],
		Amount:      &[]string{"1000"}[0],
		PaymentDate: loan_repo.ToDateOnly(&paymentDate),
		IsPaid:      &[]bool{true}[0],
	}

	mockClient.EXPECT().GetById(gomock.Any(), &pb.GetLoanPaymentByIdRequest{PaymentId: 1}).Return(mockResp, nil)

	payment, err := repo.GetById(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, int32(1), payment.PaymentID)
}

func TestLoanPaymentRepositoryGetById_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockClient := mock_protos.NewMockLoanPaymentServiceClient(ctrl)
	repo := loan_repo.NewLoanPaymentRepository(mockClient)

	mockClient.EXPECT().GetById(gomock.Any(), &pb.GetLoanPaymentByIdRequest{PaymentId: 1}).Return(nil, errors.New("error"))

	payment, err := repo.GetById(ctx, 1)
	require.Nil(t, payment)
	require.Error(t, err)
	require.Equal(t, "loan_payment_repo.GetById: error", err.Error())
}

func TestLoanPaymentRepositoryAdd_ReturnsPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockClient := mock_protos.NewMockLoanPaymentServiceClient(ctrl)
	repo := loan_repo.NewLoanPaymentRepository(mockClient)

	paymentDate := time.Now()

	mockResp := &pb.LoanPaymentModel{
		PaymentId:   1,
		LoanId:      &[]int32{1}[0],
		Amount:      &[]string{"1000"}[0],
		PaymentDate: loan_repo.ToDateOnly(&paymentDate),
		IsPaid:      &[]bool{true}[0],
	}

	mockClient.EXPECT().Add(gomock.Any(), &pb.AddLoanPaymentRequest{
		LoanId:      &[]int32{1}[0],
		Amount:      &[]string{"1000"}[0],
		PaymentDate: loan_repo.ToDateOnly(&paymentDate),
		IsPaid:      &[]bool{true}[0],
	}).Return(mockResp, nil)

	payment, err := repo.Add(ctx, &loan.LoanPayment{
		PaymentID:   1,
		LoanID:      &[]int32{1}[0],
		Amount:      &[]string{"1000"}[0],
		PaymentDate: &paymentDate,
		IsPaid:      &[]bool{true}[0],
	})
	require.NoError(t, err)
	require.Equal(t, int32(1), payment.PaymentID)
}
