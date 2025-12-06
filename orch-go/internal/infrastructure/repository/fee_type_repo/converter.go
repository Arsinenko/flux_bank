package fee_type_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/fee_type"
)

func ToDomain(p *pb.FeeTypeModel) *fee_type.FeeType {
	if p == nil {
		return nil
	}
	return &fee_type.FeeType{
		FeeID:       p.FeeId,
		Name:        p.Name,
		Description: p.Description,
	}
}

func FromDomain(f *fee_type.FeeType) *pb.FeeTypeModel {
	if f == nil {
		return nil
	}
	return &pb.FeeTypeModel{
		FeeId:       f.FeeID,
		Name:        f.Name,
		Description: f.Description,
	}
}
