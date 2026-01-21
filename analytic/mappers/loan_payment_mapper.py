from decimal import Decimal
from typing import List
from api.generated.loan_payment_pb2 import LoanPaymentModel
from api.generated.custom_types_pb2 import DateOnly
from domain.loan.loan_payment import LoanPayment
from datetime import date


class LoanPaymentMapper:
    @staticmethod
    def to_domain(model: LoanPaymentModel) -> LoanPayment:
        return LoanPayment(
            payment_id=model.payment_id,
            loan_id=model.loan_id if model.HasField("loan_id") else None,
            amount=Decimal(model.amount) if model.HasField("amount") else None,
            payment_date=date(model.payment_date.year, model.payment_date.month, model.payment_date.day) if model.HasField("payment_date") else None,
            is_paid=model.is_paid if model.HasField("is_paid") else None
        )

    @staticmethod
    def to_model(domain: LoanPayment) -> LoanPaymentModel:
        model = LoanPaymentModel(
            payment_id=domain.payment_id,
            loan_id=domain.loan_id,
            amount=str(domain.amount) if domain.amount is not None else None,
            is_paid=domain.is_paid
        )
        if domain.payment_date:
            model.payment_date.year = domain.payment_date.year
            model.payment_date.month = domain.payment_date.month
            model.payment_date.day = domain.payment_date.day
        return model

    @staticmethod
    def to_domain_list(models: List[LoanPaymentModel]) -> List[LoanPayment]:
        return [LoanPaymentMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[LoanPayment]) -> List[LoanPaymentModel]:
        return [LoanPaymentMapper.to_model(domain) for domain in domains]
