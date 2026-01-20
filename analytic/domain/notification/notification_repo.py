from abc import ABC, abstractmethod
from typing import List

from domain.notification.notification import Notification


class NotificationRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Notification]:
        pass
    @abstractmethod
    async def get_by_id(self, notification_id: int) -> Notification | None:
        pass

    @abstractmethod
    async def get_by_date_range(self, start_date: str, end_date: str) -> List[Notification]:
        pass
    @abstractmethod
    async def get_by_customer(self, customer_id: int) -> List[Notification]:
        pass