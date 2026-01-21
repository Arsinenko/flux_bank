from typing import List
from api.generated.user_credential_pb2 import UserCredentialModel
from domain.user_credential.user_credential import UserCredential


class UserCredentialMapper:
    @staticmethod
    def to_domain(model: UserCredentialModel) -> UserCredential:
        return UserCredential(
            customer_id=model.customer_id,
            username=model.username,
            password_hash=model.password_hash,
            updated_at=model.updated_at.ToDatetime() if model.HasField("updated_at") else None
        )

    @staticmethod
    def to_model(domain: UserCredential) -> UserCredentialModel:
        model = UserCredentialModel(
            customer_id=domain.customer_id,
            username=domain.username,
            password_hash=domain.password_hash
        )
        if domain.updated_at:
            model.updated_at.FromDatetime(domain.updated_at)
        return model

    @staticmethod
    def to_domain_list(models: List[UserCredentialModel]) -> List[UserCredential]:
        return [UserCredentialMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[UserCredential]) -> List[UserCredentialModel]:
        return [UserCredentialMapper.to_model(domain) for domain in domains]
