from typing import List
from api.generated.payment_template_pb2 import PaymentTemplateModel
from domain.payment_template.payment_template import PaymentTemplate

class PaymentTemplateMapper:
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
    def to_model(domain: PaymentTemplate) -> PaymentTemplateModel:
        return PaymentTemplateModel(
            template_id=domain.template_id,
            customer_id=domain.customer_id,
            name=domain.name,
            target_iban=domain.target_iban,
            default_amount=domain.default_amount
        )

    @staticmethod
    def to_domain_list(models: List[PaymentTemplateModel]) -> List[PaymentTemplate]:
        return [PaymentTemplateMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[PaymentTemplate]) -> List[PaymentTemplateModel]:
        return [PaymentTemplateMapper.to_model(domain) for domain in domains]
