package customer_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/customer"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToDomain(p *pb.CustomerModel) *customer.Customer {
	if p == nil {
		return nil
	}
	var createdAt time.Time
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt.AsTime()
	}

	var birthDay *time.Time
	if p.BirthDate != nil {
		t := time.Date(int(p.BirthDate.Year), time.Month(p.BirthDate.Month), int(p.BirthDate.Day), 0, 0, 0, 0, time.UTC)
		birthDay = &t
	}

	return &customer.Customer{
		Id:        p.CustomerId,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Phone:     p.Phone,
		BirthDate: birthDay,
		CreatedAt: &createdAt,
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

func ToProto(c *customer.Customer) *pb.CustomerModel {
	if c == nil {
		return nil
	}
	return &pb.CustomerModel{
		CustomerId: c.Id,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
		Phone:      c.Phone,
		BirthDate:  ToDateOnly(c.BirthDate),
		CreatedAt:  timestamppb.New(*c.CreatedAt),
	}

}
