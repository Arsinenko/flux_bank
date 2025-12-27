package tests

import (
	"context"
	"errors"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/branch"
	"orch-go/internal/infrastructure/repository/branch_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestBranchRepositoryGetAll_ReturnBranches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.GetAllBranchesResponse{Branches: []*pb.BranchModel{
		{
			BranchId: 1,
			Name:     &[]string{"test"}[0],
		},
	}}

	mockClient.
		EXPECT().
		GetAll(ctx, &pb.GetAllRequest{PageN: 1, PageSize: 1}).
		Return(mockResp, nil)

	branches, err := repo.GetAll(ctx, 1, 1)
	require.NoError(t, err)
	require.Len(t, branches, 1)
}

func TestBranchRepositoryGetAll_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetAll(ctx, &pb.GetAllRequest{PageN: 1, PageSize: 1}).
		Return(nil, errors.New("grpc error"))

	branches, err := repo.GetAll(ctx, 1, 1)
	require.Error(t, err)
	require.Nil(t, branches)
	require.Equal(t, "branch_repo.GetAll: grpc error", err.Error())
}

func TestBranchRepositoryGetById_ReturnBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.BranchModel{
		BranchId: 1,
		Name:     &[]string{"test"}[0],
	}

	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetBranchByIdRequest{BranchId: 1}).
		Return(mockResp, nil)

	branch, err := repo.GetById(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, int32(1), branch.BranchID)
}

func TestBranchRepositoryGetById_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetBranchByIdRequest{BranchId: 1}).
		Return(nil, errors.New("grpc error"))

	branch, err := repo.GetById(ctx, 1)
	require.Error(t, err)
	require.Nil(t, branch)
	require.Equal(t, "branch_repo.GetById: grpc error", err.Error())
}

func TestBranchRepositoryCreate_ReturnBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.BranchModel{
		BranchId: 1,
		Name:     &[]string{"test"}[0],
	}

	mockClient.
		EXPECT().
		Add(ctx, &pb.AddBranchRequest{
			Name: &[]string{"test"}[0],
		}).
		Return(mockResp, nil)

	branch, err := repo.Add(ctx, &branch.Branch{
		Name: &[]string{"test"}[0],
	})
	require.NoError(t, err)
	require.Equal(t, int32(1), branch.BranchID)
}

func TestBranchRepositoryCreate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Add(ctx, &pb.AddBranchRequest{
			Name: &[]string{"test"}[0],
		}).
		Return(nil, errors.New("grpc error"))

	branch, err := repo.Add(ctx, &branch.Branch{
		Name: &[]string{"test"}[0],
	})
	require.Error(t, err)
	require.Nil(t, branch)
	require.Equal(t, "branch_repo.Add: grpc error", err.Error())
}

func TestBranchRepositoryUpdate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Update(ctx, &pb.UpdateBranchRequest{
			BranchId: 1,
			Name:     &[]string{"test"}[0],
		}).
		Return(nil, errors.New("grpc error"))

	err := repo.Update(ctx, &branch.Branch{
		BranchID: 1,
		Name:     &[]string{"test"}[0],
	})
	require.Error(t, err)
	require.Equal(t, "branch_repo.Update: grpc error", err.Error())
}

func TestBranchRepositoryAddBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		AddBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.AddBulk(ctx, []*branch.Branch{})
	require.Error(t, err)
	require.Equal(t, "branch_repo.AddBulk: grpc error", err.Error())
}

func TestBranchRepositoryUpdateBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		UpdateBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.UpdateBulk(ctx, []*branch.Branch{})
	require.Error(t, err)
	require.Equal(t, "branch_repo.UpdateBulk: grpc error", err.Error())
}

func TestBranchRepositoryDeleteBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockBranchServiceClient(ctrl)
	repo := branch_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		DeleteBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.DeleteBulk(ctx, []int32{})
	require.Error(t, err)
	require.Equal(t, "branch_repo.DeleteBulk: grpc error", err.Error())
}
