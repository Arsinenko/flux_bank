from api.generated.loan_analytic_pb2_grpc import LoanAnalyticServiceServicer
from api.generated.loan_pb2 import GetLoanByIdRequest, GetLoanByIdsRequest, GetAllLoansResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.loan.loan_repo import LoanRepositoryAbc
from mappers.loan_mapper import LoanMapper
from google.protobuf.empty_pb2 import Empty


class LoanAnalyticService(LoanAnalyticServiceServicer):
    def __init__(self, loan_repo: LoanRepositoryAbc):
        self.loan_repo = loan_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.loan_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllLoansResponse(loans=LoanMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetLoanByIdRequest, context):
        result = await self.loan_repo.get_by_id(request.loan_id)
        return LoanMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetLoanByIdsRequest, context):
        result = await self.loan_repo.get_by_ids(request.loan_ids)
        return GetAllLoansResponse(loans=LoanMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.loan_repo.get_count()
        return CountResponse(count=result)

    async def ProcessGetCountByStatus(self, request, context):
        result = await self.loan_repo.get_count_by_status(request.status)
        return CountResponse(count=result)
