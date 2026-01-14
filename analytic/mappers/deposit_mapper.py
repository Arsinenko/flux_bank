from decimal import Decimal
from typing import List
from api.generated.deposit_pb2 import DepositModel
from domain.deposit.deposit import Deposit

class DepositMapper:
    @staticmethod
    def to_domain(model: DepositModel) -> Deposit:
        return Deposit(
            deposit_id=model.deposit_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            amount=Decimal(model.amount) if model.HasField("amount") else None,
            interest_rate=Decimal(model.interest_rate) if model.HasField("interest_rate") else None,
            start_date=model.start_date.ToDatetime() if model.HasField("start_date") else None,
            end_date=model.end_date.ToDatetime() if model.HasField("end_date") else None,
            status=model.status if model.HasField("status") else None
        )

    @staticmethod
    def to_model(domain: Deposit) -> DepositModel:
        model = DepositModel(
            deposit_id=domain.deposit_id,
            customer_id=domain.customer_id,
            amount=str(domain.amount) if domain.amount is not None else None,
            interest_rate=str(domain.interest_rate) if domain.interest_rate is not None else None,
            status=domain.status
        )
        if domain.start_date:
            model.start_date.FromDatetime(domain.start_date)
        if domain.end_date:
            model.end_date.FromDatetime(domain.end_date)
        return model

    @staticmethod
    def to_domain_list(models: List[DepositModel]) -> List[Deposit]:
        return [DepositMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Deposit]) -> List[DepositModel]:
        return [DepositMapper.to_model(domain) for domain in domains]
