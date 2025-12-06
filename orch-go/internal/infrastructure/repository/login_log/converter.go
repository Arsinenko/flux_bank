package login_log

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/login_log"
	"time"
)

func ToDomain(loginLog *pb.LoginLogModel) *login_log.LoginLog {
	if loginLog == nil {
		return nil
	}

	var loginTime *time.Time
	if loginLog.LoginTime != nil {
		t := loginLog.LoginTime.AsTime()
		loginTime = &t
	}

	return &login_log.LoginLog{
		LogID:      loginLog.LogId,
		CustomerID: loginLog.CustomerId,
		LoginTime:  loginTime,
		IpAddress:  loginLog.IpAddress,
		DeviceInfo: loginLog.DeviceInfo,
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
