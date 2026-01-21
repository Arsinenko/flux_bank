from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from api.generated.transaction_fee_analytic_pb2_grpc import TransactionFeeAnalyticServiceServicer
from api.generated.transaction_fee_pb2 import GetTotalFeeRequest, TotalFeeResponse, GetAllTransactionFeesResponse, \
    GetTransactionFeeByIdRequest
from domain.transaction.transaction_fee_repo import TransactionFeeRepositoryAbc
from mappers.transaction_fee_mapper import TransactionFeeMapper


class TransactionFeeAnalyticService(TransactionFeeAnalyticServiceServicer):
    def __init__(self, transaction_fee_repo: TransactionFeeRepositoryAbc):
        self.transaction_fee_repo = transaction_fee_repo


    async def ProcessGetTotalFee(self, request: GetTotalFeeRequest, context):
        total_fee = await self.transaction_fee_repo.get_total_fee()
        return TotalFeeResponse(total_fee=str(total_fee))

    async def ProcessGetCount(self, request, context):
        c = await self.transaction_fee_repo.get_count()
        return CountResponse(count=c)

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.transaction_fee_repo.get_all(page_n=request.pageN,
                                                         page_size=request.pageSize,
                                                         order_by=request.order_by,
                                                         is_desc=request.is_desc)
        return GetAllTransactionFeesResponse(transaction_fees=TransactionFeeMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetTransactionFeeByIdRequest, context):
        result = await self.transaction_fee_repo.get_by_id(request.id)
        return TransactionFeeMapper.to_model(result) if result else None