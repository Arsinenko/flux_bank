from typing import List
from api.generated.branch_pb2 import BranchModel
from domain.branch.branch import Branch

class BranchMapper:
    @staticmethod
    def to_domain(model: BranchModel) -> Branch:
        return Branch(
            branch_id=model.branch_id,
            name=model.name,
            city=model.city,
            address=model.address,
            phone=model.phone
        )

    @staticmethod
    def to_model(domain: Branch) -> BranchModel:
        return BranchModel(
            branch_id=domain.branch_id,
            name=domain.name,
            city=domain.city,
            address=domain.address,
            phone=domain.phone
        )

    @staticmethod
    def to_domain_list(models: List[BranchModel]) -> List[Branch]:
        return [BranchMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Branch]) -> List[BranchModel]:
        return [BranchMapper.to_model(domain) for domain in domains]
