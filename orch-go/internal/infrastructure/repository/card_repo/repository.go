package card_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/card"
)

type Repository struct {
	client pb.CardServiceClient
}

func (r Repository) GetByAccountId(ctx context.Context, accountId int32) ([]*card.Card, error) {
	resp, err := r.client.GetByAccount(ctx, &pb.GetCardsByAccountRequest{AccountId: accountId})
	if err != nil {
		return nil, fmt.Errorf("card_repo.GetByAccountId: %w", err)
	}
	result := make([]*card.Card, 0, len(resp.Cards))
	for _, c := range resp.Cards {
		result = append(result, ToDomain(c))
	}
	return result, nil

}

func NewRepository(client pb.CardServiceClient) Repository {
	return Repository{
		client: client,
	}
}

func (r Repository) GetAll(ctx context.Context) ([]*card.Card, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, fmt.Errorf("card_repo.GetAll: %w", err)
	}
	result := make([]*card.Card, 0, len(resp.Cards))
	for _, c := range resp.Cards {
		result = append(result, ToDomain(c))
	}
	return result, nil
}

func (r Repository) GetById(ctx context.Context, id int32) (*card.Card, error) {
	resp, err := r.client.GetById(ctx, &pb.GetCardByIdRequest{CardId: id})
	if err != nil {
		return nil, fmt.Errorf("card_repo.GetById: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Add(ctx context.Context, card *card.Card) (*card.Card, error) {
	req := &pb.AddCardRequest{
		AccountId:  card.AccountID,
		CardNumber: card.CardNumber,
		Cvv:        card.CVV,
		ExpiryDate: ToDateOnly(card.ExpiryDate),
		Status:     card.Status,
	}
	resp, err := r.client.Add(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("card_repo.Add: %w", err)
	}
	return ToDomain(resp), nil
}

func (r Repository) Update(ctx context.Context, card *card.Card) error {
	req := &pb.UpdateCardRequest{
		CardId:     card.CardID,
		AccountId:  card.AccountID,
		CardNumber: card.CardNumber,
		Cvv:        card.CVV,
		ExpiryDate: ToDateOnly(card.ExpiryDate),
		Status:     card.Status,
	}
	_, err := r.client.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("card_repo.Update: %w", err)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteCardRequest{CardId: id})
	if err != nil {
		return fmt.Errorf("card_repo.Delete: %w", err)
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, cards []*card.Card) error {
	var models []*pb.AddCardRequest
	for _, c := range cards {
		models = append(models, &pb.AddCardRequest{
			AccountId:  c.AccountID,
			CardNumber: c.CardNumber,
			Cvv:        c.CVV,
			ExpiryDate: ToDateOnly(c.ExpiryDate),
			Status:     c.Status,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddCardBulkRequest{Cards: models})
	if err != nil {
		return fmt.Errorf("card_repo.AddBulk: %w", err)
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, cards []*card.Card) error {
	var models []*pb.UpdateCardRequest
	for _, c := range cards {
		models = append(models, &pb.UpdateCardRequest{
			CardId:     c.CardID,
			AccountId:  c.AccountID,
			CardNumber: c.CardNumber,
			Cvv:        c.CVV,
			ExpiryDate: ToDateOnly(c.ExpiryDate),
			Status:     c.Status,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateCardBulkRequest{Cards: models})
	if err != nil {
		return fmt.Errorf("card_repo.UpdateBulk: %w", err)
	}
	return nil
}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteCardRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteCardRequest{CardId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteCardBulkRequest{Cards: models})
	if err != nil {
		return fmt.Errorf("card_repo.DeleteBulk: %w", err)
	}
	return nil
}
