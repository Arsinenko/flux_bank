package branch_repo

import (
	pb "orch-go/api/generated"
	"orch-go/internal/domain/branch"
)

func ToDomain(p *pb.BranchModel) *branch.Branch {
	if p == nil {
		return nil
	}
	return &branch.Branch{
		BranchID: &p.BranchId,
		Name:     p.Name,
		City:     p.City,
		Address:  p.Address,
		Phone:    p.Phone,
	}
}

func FromDomain(b *branch.Branch) *pb.BranchModel {
	if b == nil {
		return nil
	}
	return &pb.BranchModel{
		BranchId: *b.BranchID,
		Name:     b.Name,
		City:     b.City,
		Address:  b.Address,
		Phone:    b.Phone,
	}
}
