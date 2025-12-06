package login_log_repo

import (
	"context"
	pb "orch-go/api/generated"
	"orch-go/internal/domain/login_log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Repository struct {
	client pb.LoginLogServiceClient
}

func NewRepository(client pb.LoginLogServiceClient) Repository {
	return Repository{client: client}
}

// TODO pagination
func (l Repository) GetAll(ctx context.Context, pageN, pageSize int32) ([]*login_log.LoginLog, error) {
	resp, err := l.client.GetAll(ctx, &emptypb.Empty{})
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
