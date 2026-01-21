from api.generated.card_analytic_pb2_grpc import CardAnalyticServiceServicer
from api.generated.card_pb2 import GetCardByIdRequest, GetCardByIdsRequest, GetAllCardsResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.card.card_repo import CardRepositoryAbc
from mappers.card_mapper import CardMapper
from google.protobuf.empty_pb2 import Empty


class CardAnalyticService(CardAnalyticServiceServicer):
    def __init__(self, card_repo: CardRepositoryAbc):
        self.card_repo = card_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.card_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllCardsResponse(cards=CardMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetCardByIdRequest, context):
        result = await self.card_repo.get_by_id(request.card_id)
        return CardMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetCardByIdsRequest, context):
        result = await self.card_repo.get_by_ids(request.card_ids)
        return GetAllCardsResponse(cards=CardMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.card_repo.get_count()
        return CountResponse(count=result)
