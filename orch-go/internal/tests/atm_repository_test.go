package tests

import (
	"context"
	"errors"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/atm"
	"orch-go/internal/infrastructure/repository/atm_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAtmRepositoryGetByStatus_ReturnAtms(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.GetAllAtmsResponse{Atms: []*pb.AtmModel{
		{
			AtmId:    1,
			BranchId: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		},
	}}

	mockClient.
		EXPECT().
		GetByStatus(ctx, &pb.GetAtmsByStatusRequest{Status: "test"}).
		Return(mockResp, nil)

	atms, err := repo.GetByStatus(ctx, "test")
	require.NoError(t, err)
	require.Len(t, atms, 1)
}

func TestAtmRepositoryGetByStatus_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetByStatus(ctx, &pb.GetAtmsByStatusRequest{Status: "test"}).
		Return(nil, errors.New("grpc error"))

	atms, err := repo.GetByStatus(ctx, "test")
	require.Error(t, err)
	require.Nil(t, atms)
	require.Equal(t, "atm_repo.GetByStatus: grpc error", err.Error())
}

func TestAtmRepositoryGetByLocationSubStr_ReturnAtms(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.GetAllAtmsResponse{Atms: []*pb.AtmModel{
		{
			AtmId:    1,
			BranchId: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		},
	}}

	mockClient.
		EXPECT().
		GetByLocationSubStr(ctx, &pb.GetAtmsByLocationSubStrRequest{SubStr: "test"}).
		Return(mockResp, nil)

	atms, err := repo.GetByLocationSubStr(ctx, "test")
	require.NoError(t, err)
	require.Len(t, atms, 1)
}

func TestAtmRepositoryGetByLocationSubStr_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetByLocationSubStr(ctx, &pb.GetAtmsByLocationSubStrRequest{SubStr: "test"}).
		Return(nil, errors.New("grpc error"))

	atms, err := repo.GetByLocationSubStr(ctx, "test")
	require.Error(t, err)
	require.Nil(t, atms)
	require.Equal(t, "atm_repo.GetByLocationSubStr: grpc error", err.Error())
}

func TestAtmRepositoryGetByBranch_ReturnAtms(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.GetAllAtmsResponse{Atms: []*pb.AtmModel{
		{
			AtmId:    1,
			BranchId: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		},
	}}

	mockClient.
		EXPECT().
		GetByBranch(ctx, &pb.GetAtmsByBranchRequest{BranchId: 1}).
		Return(mockResp, nil)

	atms, err := repo.GetByBranch(ctx, 1)
	require.NoError(t, err)
	require.Len(t, atms, 1)
}

func TestAtmRepositoryGetByBranch_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetByBranch(ctx, &pb.GetAtmsByBranchRequest{BranchId: 1}).
		Return(nil, errors.New("grpc error"))

	atms, err := repo.GetByBranch(ctx, 1)
	require.Error(t, err)
	require.Nil(t, atms)
	require.Equal(t, "atm_repo.GetByBranch: grpc error", err.Error())
}

func TestAtmRepositoryGetAll_ReturnAtms(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.GetAllAtmsResponse{Atms: []*pb.AtmModel{
		{
			AtmId:    1,
			BranchId: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		},
	}}

	mockClient.
		EXPECT().
		GetAll(ctx, &pb.GetAllRequest{}).
		Return(mockResp, nil)

	atms, err := repo.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, atms, 1)
}

func TestAtmRepositoryGetAll_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetAll(ctx, &pb.GetAllRequest{}).
		Return(nil, errors.New("grpc error"))

	atms, err := repo.GetAll(ctx)
	require.Error(t, err)
	require.Nil(t, atms)
	require.Equal(t, "atm_repo.GetAll: grpc error", err.Error())
}

func TestAtmRepositoryGetById_ReturnAtm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.AtmModel{
		AtmId:    1,
		BranchId: &[]int32{1}[0],
		Location: &[]string{"test"}[0],
		Status:   &[]string{"test"}[0],
	}

	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetAtmByIdRequest{AtmId: 1}).
		Return(mockResp, nil)

	atm, err := repo.GetById(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, int32(1), atm.AtmID)
}

func TestAtmRepositoryGetById_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetAtmByIdRequest{AtmId: 1}).
		Return(nil, errors.New("grpc error"))

	atm, err := repo.GetById(ctx, 1)
	require.Error(t, err)
	require.Nil(t, atm)
	require.Equal(t, "atm_repo.GetById: grpc error", err.Error())
}

func TestAtmRepositoryAdd_ReturnAtm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.AtmModel{
		AtmId:    1,
		BranchId: &[]int32{1}[0],
		Location: &[]string{"test"}[0],
		Status:   &[]string{"test"}[0],
	}

	mockClient.
		EXPECT().
		Add(ctx, &pb.AddAtmRequest{
			BranchId: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		}).
		Return(mockResp, nil)

	atm, err := repo.Add(ctx, &atm.Atm{
		AtmID:    1,
		BranchID: &[]int32{1}[0],
		Location: &[]string{"test"}[0],
		Status:   &[]string{"test"}[0],
	})
	require.NoError(t, err)
	require.Equal(t, int32(1), atm.AtmID)
}

func TestAtmRepositoryAdd_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Add(ctx, &pb.AddAtmRequest{
			BranchId: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		}).
		Return(nil, errors.New("grpc error"))

	atm, err := repo.Add(ctx, &atm.Atm{
		AtmID:    1,
		BranchID: &[]int32{1}[0],
		Location: &[]string{"test"}[0],
		Status:   &[]string{"test"}[0],
	})
	require.Error(t, err)
	require.Nil(t, atm)
	require.Equal(t, "atm_repo.Add: grpc error", err.Error())
}

func TestAtmRepositoryUpdate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Update(ctx, &pb.UpdateAtmRequest{
			AtmId:    1,
			BranchId: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		}).
		Return(nil, errors.New("grpc error"))

	err := repo.Update(ctx, &atm.Atm{
		AtmID:    1,
		BranchID: &[]int32{1}[0],
		Location: &[]string{"test"}[0],
		Status:   &[]string{"test"}[0],
	})
	require.Error(t, err)
	require.Equal(t, "atm_repo.Update: grpc error", err.Error())
}

func TestAtmRepositoryDelete_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Delete(ctx, &pb.DeleteAtmRequest{AtmId: 1}).
		Return(nil, errors.New("grpc error"))

	err := repo.Delete(ctx, 1)
	require.Error(t, err)
	require.Equal(t, "atm_repo.Delete: grpc error", err.Error())
}

func TestAtmRepositoryAddBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		AddBulk(ctx, &pb.AddAtmBulkRequest{Atms: []*pb.AddAtmRequest{
			{
				BranchId: &[]int32{1}[0],
				Location: &[]string{"test"}[0],
				Status:   &[]string{"test"}[0],
			},
		}}).
		Return(nil, errors.New("grpc error"))

	err := repo.AddBulk(ctx, []*atm.Atm{
		{
			AtmID:    1,
			BranchID: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		},
	})
	require.Error(t, err)
	require.Equal(t, "atm_repo.AddBulk: grpc error", err.Error())
}

func TestAtmRepositoryUpdateBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		UpdateBulk(ctx, &pb.UpdateAtmBulkRequest{Atms: []*pb.UpdateAtmRequest{
			{
				AtmId:    1,
				BranchId: &[]int32{1}[0],
				Location: &[]string{"test"}[0],
				Status:   &[]string{"test"}[0],
			},
		}}).
		Return(nil, errors.New("grpc error"))

	err := repo.UpdateBulk(ctx, []*atm.Atm{
		{
			AtmID:    1,
			BranchID: &[]int32{1}[0],
			Location: &[]string{"test"}[0],
			Status:   &[]string{"test"}[0],
		},
	})
	require.Error(t, err)
	require.Equal(t, "atm_repo.UpdateBulk: grpc error", err.Error())
}

func TestAtmRepositoryDeleteBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAtmServiceClient(ctrl)
	repo := atm_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		DeleteBulk(ctx, &pb.DeleteAtmBulkRequest{Atms: []*pb.DeleteAtmRequest{
			{
				AtmId: 1,
			},
		}}).
		Return(nil, errors.New("grpc error"))

	err := repo.DeleteBulk(ctx, []int32{1})
	require.Error(t, err)
	require.Equal(t, "atm_repo.DeleteBulk: grpc error", err.Error())
}
