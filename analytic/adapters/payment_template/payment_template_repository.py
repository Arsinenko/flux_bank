from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.payment_template_pb2 import GetPaymentTemplateByIdRequest, PaymentTemplateModel, GetAllPaymentTemplatesResponse
from api.generated.payment_template_pb2_grpc import PaymentTemplateServiceStub
from domain.payment_template.payment_template import PaymentTemplate
from domain.payment_template.payment_template_repo import PaymentTemplateRepositoryAbc


from mappers.payment_template_mapper import PaymentTemplateMapper


class PaymentTemplateRepository(PaymentTemplateRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = PaymentTemplateServiceStub(self.chanel)

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(GetAllRequest()))
        return result.count

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[PaymentTemplate]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return PaymentTemplateMapper.to_domain_list(result.payment_templates)
        return []

    async def get_by_id(self, template_id: int) -> PaymentTemplate | None:
        request = GetPaymentTemplateByIdRequest(template_id=template_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return PaymentTemplateMapper.to_domain(result)
        return None
