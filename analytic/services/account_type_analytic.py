from api.generated.account_type_analytic_pb2_grpc import AccountTypeAnalyticServiceServicer
from api.generated.custom_types_pb2 import CountResponse, GetAllRequest
from domain.account.account_type_repo import AccountTypeRepositoryAbc
from mappers.account_type_mapper import AccountTypeMapper


class AccountTypeAnalytic(AccountTypeAnalyticServiceServicer):
    def __init__(self, account_type_repo: AccountTypeRepositoryAbc):
        self.account_type_repo = account_type_repo

    async def ProcessGetCount(self, request, context):
        result = await self.account_type_repo.get_count()
        return CountResponse(count=result)
    async def ProcessGetByIds(self, request, context):
        result = await self.account_type_repo.get_by_ids(request.ids)
        return AccountTypeMapper.to_model_list(result)

    async def ProcessGetById(self, request, context):
        result = await self.account_type_repo.get_by_id(request.id)
        return AccountTypeMapper.to_model(result) if result else None

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.account_type_repo.get_all(page_n=request.pageN,
                                                      page_size=request.pageSize,
                                                      order_by=request.order_by,
                                                      is_desc=request.is_desc)
        return AccountTypeMapper.to_model_list(result)
