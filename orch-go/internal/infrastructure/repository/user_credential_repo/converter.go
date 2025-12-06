package user_credential_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/user_credential"
	"time"
)

func ToDomain(model *pb.UserCredentialModel) *user_credential.UserCredential {
	if model == nil {
		return nil
	}

	var updatedAt time.Time
	if model.UpdatedAt != nil {
		updatedAt = model.UpdatedAt.AsTime()

	}
	return &user_credential.UserCredential{
		CustomerId:   model.CustomerId,
		Username:     model.Username,
		PasswordHash: model.PasswordHash,
		UpdatedAt:    updatedAt,
	}
}
