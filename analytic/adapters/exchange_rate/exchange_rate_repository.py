from decimal import Decimal
from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.exchange_rate_pb2 import *
from api.generated.exchange_rate_pb2_grpc import ExchangeRateServiceStub
from domain.exchange_rate.exchange_rate import ExchangeRate
from domain.exchange_rate.exchange_rate_repo import ExchangeRateRepositoryAbc


class ExchangeRateRepository(ExchangeRateRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = ExchangeRateServiceStub(channel=self.chanel)

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
    def response_to_list(response: GetAllExchangeRatesResponse) -> List[ExchangeRate]:
        return [ExchangeRateRepository.to_domain(model) for model in response.exchange_rates]

    async def get_all(self, page_n: int, page_size: int) -> List[ExchangeRate]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, rate_id: int) -> ExchangeRate | None:
        request = GetExchangeRateByIdRequest(rate_id=rate_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

    async def get_by_base_currency(self, base_currency: str) -> List[ExchangeRate]:
        request = GetExchangeRateByBaseCurrencyRequest(base_currency=base_currency)
        result = await self._execute(self.stub.GetByBaseCurrency(request))
        if result:
            return self.response_to_list(result)
        return []
