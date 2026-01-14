from typing import List
from api.generated.notification_pb2 import NotificationModel
from domain.notification.notification import Notification

class NotificationMapper:
    @staticmethod
    def to_domain(model: NotificationModel) -> Notification:
        return Notification(
            notification_id=model.notification_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            message=model.message if model.HasField("message") else None,
            created_at=model.created_at.ToDatetime() if model.HasField("created_at") else None,
            is_read=model.is_read if model.HasField("is_read") else None
        )

    @staticmethod
    def to_model(domain: Notification) -> NotificationModel:
        model = NotificationModel(
            notification_id=domain.notification_id,
            customer_id=domain.customer_id,
            message=domain.message,
            is_read=domain.is_read
        )
        if domain.created_at:
            model.created_at.FromDatetime(domain.created_at)
        return model

    @staticmethod
    def to_domain_list(models: List[NotificationModel]) -> List[Notification]:
        return [NotificationMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Notification]) -> List[NotificationModel]:
        return [NotificationMapper.to_model(domain) for domain in domains]
