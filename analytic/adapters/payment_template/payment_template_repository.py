from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.payment_template_pb2 import GetPaymentTemplateByIdRequest, PaymentTemplateModel, GetAllPaymentTemplatesResponse
from api.generated.payment_template_pb2_grpc import PaymentTemplateServiceStub
from domain.payment_template.payment_template import PaymentTemplate
from domain.payment_template.payment_template_repo import PaymentTemplateRepositoryAbc


class PaymentTemplateRepository(PaymentTemplateRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = PaymentTemplateServiceStub(self.chanel)

    @staticmethod
    def to_domain(model: PaymentTemplateModel) -> PaymentTemplate:
        return PaymentTemplate(
            template_id=model.template_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            name=model.name if model.HasField("name") else None,
            target_iban=model.target_iban if model.HasField("target_iban") else None,
            default_amount=model.default_amount if model.HasField("default_amount") else None
        )

    @staticmethod
    def response_to_list(response: GetAllPaymentTemplatesResponse) -> List[PaymentTemplate]:
        return [PaymentTemplateRepository.to_domain(model) for model in response.payment_templates]

    async def get_all(self, page_n: int, page_size: int) -> List[PaymentTemplate]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, template_id: int) -> PaymentTemplate | None:
        request = GetPaymentTemplateByIdRequest(template_id=template_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None
