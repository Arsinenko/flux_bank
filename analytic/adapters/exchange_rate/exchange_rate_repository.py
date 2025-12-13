from decimal import Decimal
from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.exchange_rate_pb2 import *
from api.generated.exchange_rate_pb2_grpc import ExchangeRateServiceStub
from domain.exchange_rate.exchange_rate import ExchangeRate
from domain.exchange_rate.exchange_rate_repo import ExchangeRateRepositoryAbc


class ExchangeRateRepository(ExchangeRateRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = ExchangeRateServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

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
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, rate_id: int) -> ExchangeRate | None:
        try:
            result = await self.stub.GetById(GetExchangeRateByIdRequest(rate_id=rate_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

    async def get_by_base_currency(self, base_currency: str) -> List[ExchangeRate]:
        try:
            result = await self.stub.GetByBaseCurrency(
                GetExchangeRateByBaseCurrencyRequest(base_currency=base_currency)
            )
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByBaseCurrency: {err}")
            return []
