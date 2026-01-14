from decimal import Decimal
from typing import List
from api.generated.loan_pb2 import LoanModel
from domain.loan.loan import Loan

class LoanMapper:
    @staticmethod
    def to_domain(model: LoanModel) -> Loan:
        return Loan(
            loan_id=model.loan_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            principal=Decimal(model.principal) if model.HasField("principal") else None,
            interest_rate=Decimal(model.interest_rate) if model.HasField("interest_rate") else None,
            start_date=model.start_date.ToDatetime() if model.HasField("start_date") else None,
            end_date=model.end_date.ToDatetime() if model.HasField("end_date") else None,
            status=model.status if model.HasField("status") else None
        )

    @staticmethod
    def to_model(domain: Loan) -> LoanModel:
        model = LoanModel(
            loan_id=domain.loan_id,
            customer_id=domain.customer_id,
            principal=str(domain.principal) if domain.principal is not None else None,
            interest_rate=str(domain.interest_rate) if domain.interest_rate is not None else None,
            status=domain.status
        )
        if domain.start_date:
            model.start_date.FromDatetime(domain.start_date)
        if domain.end_date:
            model.end_date.FromDatetime(domain.end_date)
        return model

    @staticmethod
    def to_domain_list(models: List[LoanModel]) -> List[Loan]:
        return [LoanMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Loan]) -> List[LoanModel]:
        return [LoanMapper.to_model(domain) for domain in domains]
