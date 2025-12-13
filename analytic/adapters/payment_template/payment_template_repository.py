from typing import List

import grpc

from api.generated.payment_template_pb2_grpc import PaymentTemplateServiceStub
from domain.payment_template.payment_template import PaymentTemplate
from domain.payment_template.payment_template_repo import PaymentTemplateRepositoryAbc


class PaymentTemplateRepository(PaymentTemplateRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = PaymentTemplateServiceStub(self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model) -> PaymentTemplate:
        return PaymentTemplate(
            template_id=model.template_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            name=model.name if model.HasField("name") else None,
            target_iban=model.target_iban if model.HasField("target_iban") else None,
            default_amount=model.default_amount if model.HasField("default_amount") else None
        )

    @staticmethod
    def response_to_list(response) -> List[PaymentTemplate]:
        return [PaymentTemplateRepository.to_domain(model) for model in response.payment_templates]

    async def get_all(self, page_n: int, page_size: int) -> List[PaymentTemplate]:
        try:
            from api.generated.custom_types_pb2 import GetAllRequest
            result = await self.stub.GetAll(GetAllRequest(pageN=page_n, pageSize=page_size))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []


    async def get_by_id(self, template_id: int) -> PaymentTemplate | None:
        try:
            from api.generated.payment_template_pb2 import GetPaymentTemplateByIdRequest
            result = await self.stub.GetById(GetPaymentTemplateByIdRequest(template_id=template_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None
