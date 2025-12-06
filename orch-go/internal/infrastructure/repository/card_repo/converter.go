package card_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/card"
	"time"
)

func ToDomain(p *pb.CardModel) *card.Card {
	if p == nil {
		return nil
	}
	var expiryDate time.Time
	if p.ExpiryDate != nil {
		expiryDate = p.ExpiryDate.AsTime()
	}
	return &card.Card{
		CardID:     p.CardId,
		AccountID:  p.AccountId,
		CardNumber: p.CardNumber,
		CVV:        p.Cvv,
		ExpiryDate: &expiryDate,
		Status:     p.Status,
	}
}

func FromDomain(c *card.Card) *pb.CardModel {
	if c == nil {
		return nil
	}
	return &pb.CardModel{
		CardId:     c.CardID,
		AccountId:  c.AccountID,
		CardNumber: c.CardNumber,
		Cvv:        c.CVV,
		Status:     c.Status,
	}
}

func ToDateOnly(t *time.Time) *pb.DateOnly {
	if t == nil {
		return nil
	}
	return &pb.DateOnly{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}
