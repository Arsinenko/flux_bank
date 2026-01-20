package notification_repo

import (
	"context"
	"fmt"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/notification"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.NotificationServiceClient
}

func (r Repository) GetByCustomer(ctx context.Context, customerId int32, isRead bool) ([]*notification.Notification, error) {
	resp, err := r.client.GetByCustomer(ctx, &pb.GetNotificationsByCustomerRequest{
		CustomerId: customerId,
		IsRead:     &isRead,
	})
	if err != nil {
		return nil, fmt.Errorf("notification_repo.GetByCustomer: %w", err)
	}
	results := make([]*notification.Notification, 0, len(resp.Notifications))
	for _, notif := range resp.Notifications {
		results = append(results, ToDomain(notif))
	}
	return results, nil
}

func (r Repository) GetByDateRange(ctx context.Context, request notification.GetByDateRangeRequest) ([]*notification.Notification, error) {
	resp, err := r.client.GetByDateRange(ctx, &pb.GetByDateRangeRequest{
		FromDate: timestamppb.New(request.From),
		ToDate:   timestamppb.New(request.To),
		PageN:    &request.PageN,
		PageSize: &request.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var notifications []*notification.Notification
	for _, notif := range resp.Notifications {
		notifications = append(notifications, ToDomain(notif))
	}

	return notifications, nil

}

func NewRepository(client pb.NotificationServiceClient) Repository {
	return Repository{client: client}
}

func (r Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*notification.Notification, error) {
	resp, err := r.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, err
	}

	var notifications []*notification.Notification
	for _, notif := range resp.Notifications {
		notifications = append(notifications, ToDomain(notif))
	}

	return notifications, nil

}

func (r Repository) GetById(ctx context.Context, id int32) (*notification.Notification, error) {
	resp, err := r.client.GetById(ctx, &pb.GetNotificationByIdRequest{NotificationId: id})
	if err != nil {
		return nil, err
	} else {
		return ToDomain(resp), nil
	}
}

func (r Repository) Add(ctx context.Context, notification *notification.Notification) (*notification.Notification, error) {
	req := pb.AddNotificationRequest{
		CustomerId: &notification.CustomerId,
		Message:    &notification.Message,
		IsRead:     &notification.IsRead,
	}
	resp, err := r.client.Add(ctx, &req)
	if err != nil {
		return nil, err
	} else {
		return ToDomain(resp), nil
	}

}

func (r Repository) Update(ctx context.Context, notification *notification.Notification) error {
	req := pb.UpdateNotificationRequest{
		NotificationId: notification.Id,
		CustomerId:     &notification.CustomerId,
		Message:        &notification.Message,
		IsRead:         &notification.IsRead,
	}
	_, err := r.client.Update(ctx, &req)
	if err != nil {
		return err
	}
	return nil

}

func (r Repository) Delete(ctx context.Context, id int32) error {
	_, err := r.client.Delete(ctx, &pb.DeleteNotificationRequest{NotificationId: id})
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) AddBulk(ctx context.Context, notifications []*notification.Notification) error {
	var models []*pb.AddNotificationRequest
	for _, n := range notifications {
		models = append(models, &pb.AddNotificationRequest{
			CustomerId: &n.CustomerId,
			Message:    &n.Message,
			IsRead:     &n.IsRead,
		})
	}
	_, err := r.client.AddBulk(ctx, &pb.AddNotificationBulkRequest{Notifications: models})
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateBulk(ctx context.Context, notifications []*notification.Notification) error {
	var models []*pb.UpdateNotificationRequest
	for _, n := range notifications {
		models = append(models, &pb.UpdateNotificationRequest{
			NotificationId: n.Id,
			CustomerId:     &n.CustomerId,
			Message:        &n.Message,
			IsRead:         &n.IsRead,
		})
	}
	_, err := r.client.UpdateBulk(ctx, &pb.UpdateNotificationBulkRequest{Notifications: models})
	if err != nil {
		return err
	}
	return nil

}

func (r Repository) DeleteBulk(ctx context.Context, ids []int32) error {
	var models []*pb.DeleteNotificationRequest
	for _, id := range ids {
		models = append(models, &pb.DeleteNotificationRequest{NotificationId: id})
	}
	_, err := r.client.DeleteBulk(ctx, &pb.DeleteNotificationBulkRequest{Notifications: models})
	if err != nil {
		return err
	}
	return nil

}
