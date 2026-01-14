from typing import List
from api.generated.atm_pb2 import AtmModel
from domain.atm.atm import Atm

class AtmMapper:
    @staticmethod
    def to_domain(model: AtmModel) -> Atm:
        return Atm(
            atm_id=model.atm_id,
            location=model.location,
            status=model.status,
            branch_id=model.branch_id
        )

    @staticmethod
    def to_model(domain: Atm) -> AtmModel:
        return AtmModel(
            atm_id=domain.atm_id,
            location=domain.location,
            status=domain.status,
            branch_id=domain.branch_id
        )

    @staticmethod
    def to_domain_list(models: List[AtmModel]) -> List[Atm]:
        return [AtmMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Atm]) -> List[AtmModel]:
        return [AtmMapper.to_model(domain) for domain in domains]
