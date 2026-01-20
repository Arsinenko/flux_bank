package services

import (
	"context"
	"orch-go/internal/domain/notification"
	"orch-go/internal/infrastructure/repository/notification_repo"
)

type NotificationService struct {
	repo notification_repo.Repository
}

func NewNotificationService(repo notification_repo.Repository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) GetNotificationById(ctx context.Context, id int32) (*notification.Notification, error) {
	return s.repo.GetById(ctx, id)
}

func (s *NotificationService) GetNotificationsByCustomer(ctx context.Context, customerId int32, isRead bool) ([]*notification.Notification, error) {
	return s.repo.GetByCustomer(ctx, customerId, isRead)
}

func (s *NotificationService) GetNotificationsByDateRange(ctx context.Context, req notification.GetByDateRangeRequest) ([]*notification.Notification, error) {
	return s.repo.GetByDateRange(ctx, req)
}

func (s *NotificationService) GetAllNotifications(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*notification.Notification, error) {
	return s.repo.GetAll(ctx, pageN, pageSize, orderBy, isDesc)
}

func (s *NotificationService) CreateNotification(ctx context.Context, notif *notification.Notification) (*notification.Notification, error) {
	return s.repo.Add(ctx, notif)
}

func (s *NotificationService) UpdateNotification(ctx context.Context, notif *notification.Notification) error {
	return s.repo.Update(ctx, notif)
}

func (s *NotificationService) DeleteNotification(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *NotificationService) CreateNotificationBulk(ctx context.Context, notifs []*notification.Notification) error {
	return s.repo.AddBulk(ctx, notifs)
}

func (s *NotificationService) UpdateNotificationBulk(ctx context.Context, notifs []*notification.Notification) error {
	return s.repo.UpdateBulk(ctx, notifs)
}

func (s *NotificationService) DeleteNotificationBulk(ctx context.Context, ids []int32) error {
	return s.repo.DeleteBulk(ctx, ids)
}
