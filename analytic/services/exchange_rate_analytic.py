from api.generated.exchange_rate_analytic_pb2_grpc import ExchangeRateAnalyticServiceServicer
from api.generated.exchange_rate_pb2 import GetExchangeRateByIdRequest, GetExchangeRateByIdsRequest, GetAllExchangeRatesResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.exchange_rate.exchange_rate_repo import ExchangeRateRepositoryAbc
from mappers.exchange_rate_mapper import ExchangeRateMapper
from google.protobuf.empty_pb2 import Empty


class ExchangeRateAnalyticService(ExchangeRateAnalyticServiceServicer):
    def __init__(self, exchange_rate_repo: ExchangeRateRepositoryAbc):
        self.exchange_rate_repo = exchange_rate_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.exchange_rate_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllExchangeRatesResponse(exchange_rates=ExchangeRateMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetExchangeRateByIdRequest, context):
        result = await self.exchange_rate_repo.get_by_id(request.exchange_rate_id)
        return ExchangeRateMapper.to_model(result) if result else None

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.exchange_rate_repo.get_count()
        return CountResponse(count=result)
