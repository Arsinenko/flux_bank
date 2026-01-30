package customer_address_repo

import (
	pb "orch-go/api/generated"
	customerAdress "orch-go/internal/domain/customer_address"
)

func ToDomain(p *pb.CustomerAddressModel) *customerAdress.CustomerAddress {
	if p == nil {
		return nil
	}
	return &customerAdress.CustomerAddress{
		Id:         &p.AddressId,
		CustomerId: *p.CustomerId,
		Country:    *p.Country,
		City:       *p.City,
		Street:     *p.Street,
		ZipCode:    *p.ZipCode,
		IsPrimary:  *p.IsPrimary,
	}
}

func ToProto(c *customerAdress.CustomerAddress) *pb.CustomerAddressModel {
	if c == nil {
		return nil
	}
	return &pb.CustomerAddressModel{
		AddressId:  *c.Id,
		CustomerId: &c.CustomerId,
		Country:    &c.Country,
		City:       &c.City,
		Street:     &c.Street,
		ZipCode:    &c.ZipCode,
		IsPrimary:  &c.IsPrimary,
	}
}
