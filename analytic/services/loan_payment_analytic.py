from api.generated.loan_payment_analytic_pb2_grpc import LoanPaymentAnalyticServiceServicer
from api.generated.loan_payment_pb2 import GetLoanPaymentByIdRequest, GetLoanPaymentByIdsRequest, GetLoanPaymentsByLoanRequest, GetAllLoanPaymentsResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.loan.loan_payment_repo import LoanPaymentRepositoryAbc
from mappers.loan_payment_mapper import LoanPaymentMapper
from google.protobuf.empty_pb2 import Empty


class LoanPaymentAnalyticService(LoanPaymentAnalyticServiceServicer):
    def __init__(self, loan_payment_repo: LoanPaymentRepositoryAbc):
        self.loan_payment_repo = loan_payment_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.loan_payment_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllLoanPaymentsResponse(loan_payments=LoanPaymentMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetLoanPaymentByIdRequest, context):
        result = await self.loan_payment_repo.get_by_id(request.payment_id)
        return LoanPaymentMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetLoanPaymentByIdsRequest, context):
        result = await self.loan_payment_repo.get_by_ids(request.payment_ids)
        return GetAllLoanPaymentsResponse(loan_payments=LoanPaymentMapper.to_model_list(result))

    async def ProcessGetByLoan(self, request: GetLoanPaymentsByLoanRequest, context):
        result = await self.loan_payment_repo.get_by_loan_id(request.loan_id)
        return GetAllLoanPaymentsResponse(loan_payments=LoanPaymentMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.loan_payment_repo.get_count()
        return CountResponse(count=result)
