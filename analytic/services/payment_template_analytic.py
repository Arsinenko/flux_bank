from api.generated.payment_template_analytic_pb2_grpc import PaymentTemplateAnalyticServiceServicer
from api.generated.payment_template_pb2 import GetPaymentTemplateByIdRequest, GetPaymentTemplateByIdsRequest, GetAllPaymentTemplatesResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.payment_template.payment_template_repo import PaymentTemplateRepositoryAbc
from mappers.payment_template_mapper import PaymentTemplateMapper
from google.protobuf.empty_pb2 import Empty


class PaymentTemplateAnalyticService(PaymentTemplateAnalyticServiceServicer):
    def __init__(self, payment_template_repo: PaymentTemplateRepositoryAbc):
        self.payment_template_repo = payment_template_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.payment_template_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllPaymentTemplatesResponse(payment_templates=PaymentTemplateMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetPaymentTemplateByIdRequest, context):
        result = await self.payment_template_repo.get_by_id(request.template_id)
        return PaymentTemplateMapper.to_model(result) if result else None

    # async def ProcessGetByIds(self, request: GetPaymentTemplateByIdsRequest, context):
    #     result = await self.payment_template_repo.get_by_ids(request.template_ids)
    #     return GetAllPaymentTemplatesResponse(payment_templates=PaymentTemplateMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.payment_template_repo.get_count()
        return CountResponse(count=result)
