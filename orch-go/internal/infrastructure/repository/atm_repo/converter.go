package atm_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/atm"
)

func ToDomain(p *pb.AtmModel) *atm.Atm {
	if p == nil {
		return nil
	}
	return &atm.Atm{
		AtmID:    p.AtmId,
		BranchID: p.BranchId,
		Location: p.Location,
		Status:   p.Status,
	}
}

func FromDomain(a *atm.Atm) *pb.AtmModel {
	if a == nil {
		return nil
	}
	return &pb.AtmModel{
		AtmId:    a.AtmID,
		BranchId: a.BranchID,
		Location: a.Location,
		Status:   a.Status,
	}
}
