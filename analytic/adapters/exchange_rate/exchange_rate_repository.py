from decimal import Decimal
from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.exchange_rate_pb2 import *
from api.generated.exchange_rate_pb2_grpc import ExchangeRateServiceStub
from domain.exchange_rate.exchange_rate import ExchangeRate
from domain.exchange_rate.exchange_rate_repo import ExchangeRateRepositoryAbc


from mappers.exchange_rate_mapper import ExchangeRateMapper


class ExchangeRateRepository(ExchangeRateRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = ExchangeRateServiceStub(channel=self.chanel)

    async def get_all(self, page_n: int, page_size: int) -> List[ExchangeRate]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return ExchangeRateMapper.to_domain_list(result.exchange_rates)
        return []

    async def get_by_id(self, rate_id: int) -> ExchangeRate | None:
        request = GetExchangeRateByIdRequest(rate_id=rate_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return ExchangeRateMapper.to_domain(result)
        return None

    async def get_by_base_currency(self, base_currency: str) -> List[ExchangeRate]:
        request = GetExchangeRateByBaseCurrencyRequest(base_currency=base_currency)
        result = await self._execute(self.stub.GetByBaseCurrency(request))
        if result:
            return ExchangeRateMapper.to_domain_list(result.exchange_rates)
        return []
