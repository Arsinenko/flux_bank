package tests

import (
	"context"
	"errors"
	"orch-go/internal/domain/account"
	"orch-go/internal/infrastructure/repository/account/account_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"

	pb "orch-go/api/generated"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAccountTypeRepositoryGetAll_ReturnTypes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)

	ctx := context.Background()
	pageN := int32(1)
	pageSize := int32(10)

	description := "test"

	mockResp := &pb.GetAllAccountTypesResponse{
		AccountTypes: []*pb.AccountTypeModel{
			{
				TypeId:      1,
				Name:        "test",
				Description: &description,
			},
		},
	}

	mockClient.
		EXPECT().
		GetAll(ctx, &pb.GetAllRequest{
			PageN:    pageN,
			PageSize: pageSize,
		}).
		Return(mockResp, nil)

	types, err := repo.GetAll(ctx, pageN, pageSize)
	require.NoError(t, err)
	require.Len(t, types, 1)
}

func TestAccountTypeRepositoryGetAll_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)
	ctx := context.Background()
	mockClient.
		EXPECT().
		GetAll(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	types, err := repo.GetAll(ctx, 1, 1)
	require.Error(t, err)
	require.Nil(t, types)
	require.Equal(t, "account_repo.GetAll: grpc error", err.Error())
}

func TestAccountTypeRepositoryGetById_ReturnAccountType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)
	typeId := 1

	ctx := context.Background()
	description := "test"
	mockResp := pb.AccountTypeModel{
		TypeId:      1,
		Name:        "test",
		Description: &description,
	}
	mockClient.
		EXPECT().
		GetById(ctx, &pb.GetAccountTypeByIdRequest{
			TypeId: int32(typeId),
		}).
		Return(&mockResp, nil)

	accountType, err := repo.GetById(ctx, int32(typeId))
	require.NoError(t, err)
	require.Equal(t, "test", accountType.Name)
}

func TestAccountTypeRepositoryGetById_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		GetById(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	accountType, err := repo.GetById(ctx, 1)
	require.Error(t, err)
	require.Nil(t, accountType)
	require.Equal(t, "account_repo.GetById: grpc error", err.Error())

}

func TestAccountTypeRepositoryCreate_ReturnAccountType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)
	ctx := context.Background()

	req := account.AccountType{
		Id:          1,
		Name:        "test",
		Description: nil,
	}

	mockResp := pb.AccountTypeModel{
		TypeId:      1,
		Name:        "test",
		Description: nil,
	}

	mockClient.
		EXPECT().
		Add(ctx, &pb.AddAccountTypeRequest{
			Name:        req.Name,
			Description: req.Description,
		}).
		Return(&mockResp, nil)

	accountType, err := repo.Create(ctx, &req)
	require.NoError(t, err)
	require.Equal(t, "test", accountType.Name)

}

func TestAccountTypeRepositoryCreate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)
	ctx := context.Background()

	mockClient.
		EXPECT().
		Add(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	accountType, err := repo.Create(ctx, &account.AccountType{})
	require.Nil(t, accountType)
	require.Equal(t, "account_repo.Create: grpc error", err.Error())
}

func TestAccountTypeRepositoryUpdate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Update(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.Update(ctx, &account.AccountType{
		Id:          1,
		Name:        "test",
		Description: nil,
	})

	require.Equal(t, "account_repo.Update: grpc error", err.Error())
}

func TestAccountTypeRepositoryDelete_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		Delete(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.Delete(ctx, 1)
	require.Equal(t, "account_repo.Delete: grpc error", err.Error())
}

func TestAccountTypeRepositoryAddBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		AddBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.AddBulk(ctx, []account.AccountType{})
	require.Equal(t, "account_repo.AddBulk: grpc error", err.Error())
}

func TestAccountTypeRepositoryUpdateBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		UpdateBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.UpdateBulk(ctx, []account.AccountType{})
	require.Equal(t, "account_repo.UpdateBulk: grpc error", err.Error())
}

func TestAccountTypeRepositoryDeleteBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockAccountTypeServiceClient(ctrl)
	repo := account_repo.NewAccountTypeRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		DeleteBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.DeleteBulk(ctx, []int32{})
	require.Equal(t, "account_repo.DeleteBulk: grpc error", err.Error())
}
