package notification_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/notification"
	"time"
)

func ToDomain(model *pb.NotificationModel) *notification.Notification {
	if model == nil {
		return nil
	}

	var createdAt time.Time
	if model.CreatedAt != nil {
		createdAt = model.CreatedAt.AsTime()
	}
	return &notification.Notification{
		Id:         model.NotificationId,
		CustomerId: *model.CustomerId,
		Message:    *model.Message,
		CreatedAt:  createdAt,
		IsRead:     *model.IsRead,
	}
}
