package tests

import (
	"context"
	mock_protos "orch-go/internal/mocks"
	"testing"

	pb "orch-go/api/generated"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetAll_ReturnsAccountTypes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock_protos.NewMockAccountTypeServiceClient(ctrl)

	ctx := context.Background()
	req := pb.GetAllRequest{
		PageN:    1,
		PageSize: 1,
	}
	mockResp := &pb.GetAllAccountTypesResponse{
		AccountTypes: []*pb.AccountTypeModel{
			{
				TypeId:      1,
				Name:        "name",
				Description: &[]string{"description"}[0],
			},
		},
	}

	repo.
		EXPECT().
		GetAll(ctx, &req).
		Return(mockResp, nil)

	accountTypes, err := repo.GetAll(ctx, &pb.GetAllRequest{
		PageN:    1,
		PageSize: 1,
	})

	require.NoError(t, err)
	require.Len(t, accountTypes.AccountTypes, 1)

}
