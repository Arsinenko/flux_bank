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
func ToAtmsDomain(resp *pb.GetAllAtmsResponse) []*atm.Atm {
	result := make([]*atm.Atm, 0, len(resp.Atms))
	for _, a := range resp.Atms {
		result = append(result, ToDomain(a))
	}
	return result
}
