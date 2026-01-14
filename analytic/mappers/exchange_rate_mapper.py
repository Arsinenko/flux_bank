from decimal import Decimal
from typing import List
from api.generated.exchange_rate_pb2 import ExchangeRateModel
from domain.exchange_rate.exchange_rate import ExchangeRate

class ExchangeRateMapper:
    @staticmethod
    def to_domain(model: ExchangeRateModel) -> ExchangeRate:
        return ExchangeRate(
            rate_id=model.rate_id,
            base_currency=model.base_currency if model.HasField("base_currency") else None,
            target_currency=model.target_currency if model.HasField("target_currency") else None,
            rate=Decimal(model.rate) if model.HasField("rate") else None,
            updated_at=model.updated_at.ToDatetime() if model.HasField("updated_at") else None
        )

    @staticmethod
    def to_model(domain: ExchangeRate) -> ExchangeRateModel:
        model = ExchangeRateModel(
            rate_id=domain.rate_id,
            base_currency=domain.base_currency,
            target_currency=domain.target_currency,
            rate=str(domain.rate) if domain.rate is not None else None
        )
        if domain.updated_at:
            model.updated_at.FromDatetime(domain.updated_at)
        return model

    @staticmethod
    def to_domain_list(models: List[ExchangeRateModel]) -> List[ExchangeRate]:
        return [ExchangeRateMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[ExchangeRate]) -> List[ExchangeRateModel]:
        return [ExchangeRateMapper.to_model(domain) for domain in domains]
