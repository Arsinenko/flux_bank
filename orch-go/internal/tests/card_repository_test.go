package tests

import (
	"context"
	"errors"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/card"
	"orch-go/internal/infrastructure/repository/card_repo"
	mock_protos "orch-go/internal/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCardRepositoryGetByAccountId_ReturnCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := int32(1)

	mockResp := &pb.GetAllCardsResponse{
		Cards: []*pb.CardModel{
			{
				CardId:    1,
				AccountId: &accountId,
			},
		},
	}

	req := &pb.GetCardsByAccountRequest{AccountId: accountId}

	mockClient.
		EXPECT().
		GetByAccount(ctx, req).
		Return(mockResp, nil)

	cards, err := repo.GetByAccountId(ctx, accountId)

	require.NoError(t, err)
	require.Len(t, cards, 1)

}

func TestCardRepositoryGetByAccountId_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := 1

	mockClient.
		EXPECT().
		GetByAccount(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	cards, err := repo.GetByAccountId(ctx, int32(accountId))

	require.Error(t, err)
	require.Equal(t, "card_repo.GetByAccountId: grpc error", err.Error())
	require.Nil(t, cards)
}

func TestCardRepositoryGetAll_ReturnCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockResp := &pb.GetAllCardsResponse{
		Cards: []*pb.CardModel{
			{
				CardId:    1,
				AccountId: &[]int32{1}[0],
			},
		},
	}

	req := &pb.GetAllRequest{PageN: 1, PageSize: 1}

	mockClient.
		EXPECT().
		GetAll(ctx, req).
		Return(mockResp, nil)

	cards, err := repo.GetAll(ctx, 1, 1)

	require.NoError(t, err)
	require.Len(t, cards, 1)

}

func TestCardRepositoryGetAll_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		GetAll(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	cards, err := repo.GetAll(ctx, 1, 1)

	require.Error(t, err)
	require.Equal(t, "card_repo.GetAll: grpc error", err.Error())
	require.Nil(t, cards)
}

func TestCardRepositoryGetById_ReturnCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	cardId := int32(1)
	accountId := int32(10)

	mockResp := &pb.CardModel{
		CardId:    cardId,
		AccountId: &accountId,
	}

	req := &pb.GetCardByIdRequest{CardId: cardId}

	mockClient.
		EXPECT().
		GetById(ctx, req).
		Return(mockResp, nil)

	card, err := repo.GetById(ctx, cardId)

	require.NoError(t, err)
	require.Equal(t, cardId, card.CardID)
}

func TestCardRepositoryGetById_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	cardId := int32(1)

	mockClient.
		EXPECT().
		GetById(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	card, err := repo.GetById(ctx, cardId)

	require.Error(t, err)
	require.Equal(t, "card_repo.GetById: grpc error", err.Error())
	require.Nil(t, card)
}

func TestCardRepositoryAdd_ReturnCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := int32(10)
	cardNumber := "1234-5678-90"
	exp_date := time.Now().AddDate(2, 0, 0)

	card := &card.Card{
		CardNumber: cardNumber,
		AccountID:  &accountId,
		CardID:     1,
		CVV:        "cvv",
		Status:     &[]string{"active"}[0],
		ExpiryDate: &exp_date,
	}

	mockResp := &pb.CardModel{
		CardId:     card.CardID,
		AccountId:  card.AccountID,
		CardNumber: card.CardNumber,
		Cvv:        card.CVV,
	}
	req := &pb.AddCardRequest{
		AccountId:  card.AccountID,
		CardNumber: card.CardNumber,
		Cvv:        card.CVV,
		ExpiryDate: timestamppb.New(*card.ExpiryDate),
		Status:     card.Status,
	}

	mockClient.
		EXPECT().
		Add(ctx, req).
		Return(mockResp, nil)

	resultCard, err := repo.Add(ctx, card)

	require.NoError(t, err)
	require.NotNil(t, resultCard)
}

func TestCardRepositoryAdd_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := int32(10)
	cardNumber := "1234-5678-90"
	exp_date := time.Now().AddDate(2, 0, 0)

	card := &card.Card{
		CardNumber: cardNumber,
		AccountID:  &accountId,
		CardID:     1,
		CVV:        "cvv",
		Status:     &[]string{"active"}[0],
		ExpiryDate: &exp_date,
	}

	mockClient.
		EXPECT().
		Add(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	resultCard, err := repo.Add(ctx, card)

	require.Error(t, err)
	require.Nil(t, resultCard)
	require.Equal(t, "card_repo.Add: grpc error", err.Error())
}

func TestCardRepositoryUpdate_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := int32(10)
	cardNumber := "1234-5678-90"
	exp_date := time.Now().AddDate(2, 0, 0)

	card := &card.Card{
		CardNumber: cardNumber,
		AccountID:  &accountId,
		CardID:     1,
		CVV:        "cvv",
		Status:     &[]string{"active"}[0],
		ExpiryDate: &exp_date,
	}

	mockClient.
		EXPECT().
		Update(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.Update(ctx, card)

	require.Error(t, err)
	require.Equal(t, "card_repo.Update: grpc error", err.Error())
}

func TestCardRepositoryDelete_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := int32(10)
	cardNumber := "1234-5678-90"
	exp_date := time.Now().AddDate(2, 0, 0)

	card := &card.Card{
		CardNumber: cardNumber,
		AccountID:  &accountId,
		CardID:     1,
		CVV:        "cvv",
		Status:     &[]string{"active"}[0],
		ExpiryDate: &exp_date,
	}

	mockClient.
		EXPECT().
		Delete(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.Delete(ctx, card.CardID)

	require.Error(t, err)
	require.Equal(t, "card_repo.Delete: grpc error", err.Error())
}

func TestCardRepositoryAddBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := int32(10)
	cardNumber := "1234-5678-90"
	exp_date := time.Now().AddDate(2, 0, 0)

	cardData := &card.Card{
		CardNumber: cardNumber,
		AccountID:  &accountId,
		CardID:     1,
		CVV:        "cvv",
		Status:     &[]string{"active"}[0],
		ExpiryDate: &exp_date,
	}

	mockClient.
		EXPECT().
		AddBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.AddBulk(ctx, []*card.Card{
		cardData,
	})

	require.Error(t, err)
	require.Equal(t, "card_repo.AddBulk: grpc error", err.Error())
}

func TestCardRepositoryUpdateBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	accountId := int32(10)
	cardNumber := "1234-5678-90"
	exp_date := time.Now().AddDate(2, 0, 0)

	cardData := &card.Card{
		CardNumber: cardNumber,
		AccountID:  &accountId,
		CardID:     1,
		CVV:        "cvv",
		Status:     &[]string{"active"}[0],
		ExpiryDate: &exp_date,
	}

	mockClient.
		EXPECT().
		UpdateBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.UpdateBulk(ctx, []*card.Card{cardData})

	require.Error(t, err)
	require.Equal(t, "card_repo.UpdateBulk: grpc error", err.Error())
}

func TestCardRepositoryDeleteBulk_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_protos.NewMockCardServiceClient(ctrl)
	repo := card_repo.NewRepository(mockClient)

	ctx := context.Background()

	mockClient.
		EXPECT().
		DeleteBulk(ctx, gomock.Any()).
		Return(nil, errors.New("grpc error"))

	err := repo.DeleteBulk(ctx, []int32{1})

	require.Error(t, err)
	require.Equal(t, "card_repo.DeleteBulk: grpc error", err.Error())
}
