package transaction_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/transaction"
	"time"
)

func ToTransactionDomain(p *pb.TransactionModel) *transaction.Transaction {
	if p == nil {
		return nil
	}
	var createdAt time.Time
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt.AsTime()
	}
	return &transaction.Transaction{
		TransactionID: p.TransactionId,
		SourceAccount: p.SourceAccount,
		TargetAccount: p.TargetAccount,
		Amount:        p.Amount,
		Currency:      p.Currency,
		CreatedAt:     &createdAt,
		Status:        p.Status,
	}
}

func FromTransactionDomain(t *transaction.Transaction) *pb.TransactionModel {
	if t == nil {
		return nil
	}
	return &pb.TransactionModel{
		TransactionId: t.TransactionID,
		SourceAccount: t.SourceAccount,
		TargetAccount: t.TargetAccount,
		Amount:        t.Amount,
		Currency:      t.Currency,
		Status:        t.Status,
	}
}

func ToTransactionCategoryDomain(p *pb.TransactionCategoryModel) *transaction.TransactionCategory {
	if p == nil {
		return nil
	}
	return &transaction.TransactionCategory{
		CategoryID: p.CategoryId,
		Name:       p.Name,
	}
}

func FromTransactionCategoryDomain(tc *transaction.TransactionCategory) *pb.TransactionCategoryModel {
	if tc == nil {
		return nil
	}
	return &pb.TransactionCategoryModel{
		CategoryId: tc.CategoryID,
		Name:       tc.Name,
	}
}

func ToTransactionFeeDomain(p *pb.TransactionFeeModel) *transaction.TransactionFee {
	if p == nil {
		return nil
	}
	return &transaction.TransactionFee{
		ID:            p.Id,
		TransactionID: p.TransactionId,
		FeeID:         p.FeeId,
		Amount:        p.Amount,
	}
}

func FromTransactionFeeDomain(tf *transaction.TransactionFee) *pb.TransactionFeeModel {
	if tf == nil {
		return nil
	}
	return &pb.TransactionFeeModel{
		Id:            tf.ID,
		TransactionId: tf.TransactionID,
		FeeId:         tf.FeeID,
		Amount:        tf.Amount,
	}
}
