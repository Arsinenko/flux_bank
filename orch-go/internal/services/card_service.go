package services

import (
	"context"
	"orch-go/internal/domain/card"
	"orch-go/internal/infrastructure/repository/card_repo"
)

type CardService struct {
	repo card_repo.Repository
}

func NewCardService(repo card_repo.Repository) *CardService {
	return &CardService{repo: repo}
}

func (s *CardService) GetCardById(ctx context.Context, id int32) (*card.Card, error) {
	return s.repo.GetById(ctx, id)
}

func (s *CardService) GetCardsByAccountId(ctx context.Context, accountId int32) ([]*card.Card, error) {
	return s.repo.GetByAccountId(ctx, accountId)
}

func (s *CardService) GetAllCards(ctx context.Context) ([]*card.Card, error) {
	return s.repo.GetAll(ctx)
}

func (s *CardService) CreateCard(ctx context.Context, card *card.Card) (*card.Card, error) {
	return s.repo.Add(ctx, card)
}

func (s *CardService) UpdateCard(ctx context.Context, card *card.Card) error {
	return s.repo.Update(ctx, card)
}

func (s *CardService) DeleteCard(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *CardService) CreateCardBulk(ctx context.Context, cards []*card.Card) error {
	return s.repo.AddBulk(ctx, cards)
}

func (s *CardService) UpdateCardBulk(ctx context.Context, cards []*card.Card) error {
	return s.repo.UpdateBulk(ctx, cards)
}

func (s *CardService) DeleteCardBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
