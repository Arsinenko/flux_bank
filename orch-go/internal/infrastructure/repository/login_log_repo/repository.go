package login_log_repo

import (
	"context"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/login_log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Repository struct {
	client pb.LoginLogServiceClient
}

func (l Repository) GetByCustomer(ctx context.Context, customerId int32) ([]*login_log.LoginLog, error) {
	resp, err := l.client.GetByCustomer(ctx, &pb.GetLoginLogsByCustomerRequest{CustomerId: customerId})
	if err != nil {
		return nil, err
	}

	var loginLogs []*login_log.LoginLog
	for _, log := range resp.LoginLogs {
		loginLogs = append(loginLogs, ToDomain(log))
	}

	return loginLogs, nil
}

func (l Repository) GetInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*login_log.LoginLog, error) {
	resp, err := l.client.GetInTimeRange(ctx, &pb.GetLoginLogsInTimeRangeRequest{
		StartTime: timestamppb.New(startTime),
		EndTime:   timestamppb.New(endTime),
	})
	if err != nil {
		return nil, err
	}

	var loginLogs []*login_log.LoginLog
	for _, log := range resp.LoginLogs {
		loginLogs = append(loginLogs, ToDomain(log))
	}

	return loginLogs, nil

}

func NewRepository(client pb.LoginLogServiceClient) Repository {
	return Repository{client: client}
}

func (l Repository) GetAll(ctx context.Context, pageN, pageSize int32, orderBy string, isDesc bool) ([]*login_log.LoginLog, error) {
	resp, err := l.client.GetAll(ctx, &pb.GetAllRequest{
		PageN:    pageN,
		PageSize: pageSize,
		OrderBy:  &wrapperspb.StringValue{Value: orderBy},
		IsDesc:   &wrapperspb.BoolValue{Value: isDesc},
	})
	if err != nil {
		return nil, err
	}

	var loginLogs []*login_log.LoginLog
	for _, log := range resp.LoginLogs {
		loginLogs = append(loginLogs, ToDomain(log))
	}

	return loginLogs, nil

}

func (l Repository) GetById(ctx context.Context, id int32) (*login_log.LoginLog, error) {
	resp, err := l.client.GetById(ctx, &pb.GetLoginLogByIdRequest{LogId: id})
	if err != nil {
		return nil, err
	} else {
		return ToDomain(resp), nil
	}
}

func (l Repository) Create(ctx context.Context, loginLog *login_log.LoginLog) (*login_log.LoginLog, error) {
	req := pb.AddLoginLogRequest{
		CustomerId: loginLog.CustomerID,
		LoginTime:  timestamppb.New(*loginLog.LoginTime),
		IpAddress:  loginLog.IpAddress,
		DeviceInfo: loginLog.DeviceInfo,
	}
	resp, err := l.client.Add(ctx, &req)
	if err != nil {
		return nil, err
	} else {
		return ToDomain(resp), nil
	}
}

func (l Repository) Update(ctx context.Context, loginLog *login_log.LoginLog) error {
	req := pb.UpdateLoginLogRequest{
		LogId:      loginLog.LogID,
		CustomerId: loginLog.CustomerID,
		LoginTime:  timestamppb.New(*loginLog.LoginTime),
		IpAddress:  loginLog.IpAddress,
		DeviceInfo: loginLog.DeviceInfo,
	}
	_, err := l.client.Update(ctx, &req)
	if err != nil {
		return err
	}
	return nil
}

func (l Repository) Delete(ctx context.Context, id int32) error {
	_, err := l.client.Delete(ctx, &pb.DeleteLoginLogRequest{LogId: id})
	if err != nil {
		return err
	}
	return nil
}
