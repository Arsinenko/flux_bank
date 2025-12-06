package login_log

import "time"

type LoginLog struct {
	LogID      int32
	CustomerID *int32
	LoginTime  *time.Time
	IpAddress  *string
	DeviceInfo *string
}
