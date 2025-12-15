package services

import (
	"context"
	"orch-go/internal/domain/login_log"
	"orch-go/internal/infrastructure/repository/login_log_repo"
	"time"
)

type LoginLogService struct {
	repo login_log_repo.Repository
}

func NewLoginLogService(repo login_log_repo.Repository) *LoginLogService {
	return &LoginLogService{repo: repo}
}

func (s *LoginLogService) GetLoginLogById(ctx context.Context, id int32) (*login_log.LoginLog, error) {
	return s.repo.GetById(ctx, id)
}

func (s *LoginLogService) GetLoginLogsByCustomer(ctx context.Context, customerId int32) ([]*login_log.LoginLog, error) {
	return s.repo.GetByCustomer(ctx, customerId)
}

func (s *LoginLogService) GetLoginLogsInTimeRange(ctx context.Context, startTime, endTime time.Time) ([]*login_log.LoginLog, error) {
	return s.repo.GetInTimeRange(ctx, startTime, endTime)
}

func (s *LoginLogService) GetAllLoginLogs(ctx context.Context, pageN, pageSize int32) ([]*login_log.LoginLog, error) {
	return s.repo.GetAll(ctx, pageN, pageSize)
}

func (s *LoginLogService) CreateLoginLog(ctx context.Context, log *login_log.LoginLog) (*login_log.LoginLog, error) {
	return s.repo.Create(ctx, log)
}

func (s *LoginLogService) UpdateLoginLog(ctx context.Context, log *login_log.LoginLog) error {
	return s.repo.Update(ctx, log)
}

func (s *LoginLogService) DeleteLoginLog(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
