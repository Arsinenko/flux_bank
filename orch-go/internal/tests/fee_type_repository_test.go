package tests

import (
	"context"
	"errors"
	protos "orch-go/api/generated"
	"orch-go/internal/domain/fee_type"
	"orch-go/internal/infrastructure/repository/fee_type_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestFeeTypeRepositoryGetAll_ReturnsAllFeeTypes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	typeModels := []*protos.FeeTypeModel{
		{
			FeeId:       1,
			Name:        &[]string{"test"}[0],
			Description: &[]string{"test"}[0],
		},
		{
			FeeId:       2,
			Name:        &[]string{"test"}[0],
			Description: &[]string{"test"}[0],
		},
	}

	mockClient.
		EXPECT().
		GetAll(ctx, gomock.Any()).
		Return(&protos.GetAllFeeTypesResponse{FeeTypes: typeModels}, nil)

	types, err := repo.GetAll(ctx, 1, 2)

	require.NoError(t, err)
	require.Len(t, types, 2)
}

func TestFeeTypeRepositoryGetAll_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetAll(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	_, err := repo.GetAll(ctx, 1, 2)

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.GetAll: error", err.Error())
}

func TestFeeTypeRepositoryGetById_ReturnsFeeType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	typeModel := &protos.FeeTypeModel{
		FeeId:       1,
		Name:        &[]string{"test"}[0],
		Description: &[]string{"test"}[0],
	}

	mockClient.
		EXPECT().
		GetById(ctx, gomock.Any()).
		Return(typeModel, nil)

	feeType, err := repo.GetById(ctx, 1)

	require.NoError(t, err)
	require.Equal(t, feeType.FeeID, typeModel.FeeId)
}

func TestFeeTypeRepositoryGetById_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetById(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	_, err := repo.GetById(ctx, 1)

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.GetById: error", err.Error())
}

func TestFeeTypeRepositoryAdd_ReturnsFeeType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	typeModel := &protos.FeeTypeModel{
		FeeId:       1,
		Name:        &[]string{"test"}[0],
		Description: &[]string{"test"}[0],
	}

	mockClient.
		EXPECT().
		Add(ctx, gomock.Any()).
		Return(typeModel, nil)

	feeType, err := repo.Add(ctx, &fee_type.FeeType{
		FeeID:       1,
		Name:        &[]string{"test"}[0],
		Description: &[]string{"test"}[0],
	})

	require.NoError(t, err)
	require.Equal(t, feeType.FeeID, typeModel.FeeId)
}

func TestFeeTypeRepositoryAdd_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Add(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	_, err := repo.Add(ctx, &fee_type.FeeType{
		FeeID:       1,
		Name:        &[]string{"test"}[0],
		Description: &[]string{"test"}[0],
	})

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.Add: error", err.Error())
}

func TestFeeTypeRepositoryUpdate_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Update(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.Update(ctx, &fee_type.FeeType{
		FeeID:       1,
		Name:        &[]string{"test"}[0],
		Description: &[]string{"test"}[0],
	})

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.Update: error", err.Error())
}

func TestFeeTypeRepositoryDelete_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Delete(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.Delete(ctx, 1)

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.Delete: error", err.Error())
}

func TestFeeTypeRepositoryAddBulk_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		AddBulk(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.AddBulk(ctx, []*fee_type.FeeType{
		{
			FeeID:       1,
			Name:        &[]string{"test"}[0],
			Description: &[]string{"test"}[0],
		},
	})

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.AddBulk: error", err.Error())
}

func TestFeeTypeRepositoryUpdateBulk_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		UpdateBulk(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.UpdateBulk(ctx, []*fee_type.FeeType{
		{
			FeeID:       1,
			Name:        &[]string{"test"}[0],
			Description: &[]string{"test"}[0],
		},
	})

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.UpdateBulk: error", err.Error())
}

func TestFeeTypeRepositoryDeleteBulk_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockFeeTypeServiceClient(ctrl)
	repo := fee_type_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		DeleteBulk(ctx, gomock.Any()).
		Return(nil, errors.New("error"))

	err := repo.DeleteBulk(ctx, []int32{1})

	require.Error(t, err)
	require.Equal(t, "fee_type_repo.DeleteBulk: error", err.Error())
}
