from abc import ABC, abstractmethod
from typing import List
from datetime import datetime

from domain.login_log.login_log import LoginLog
class LoginLogRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[LoginLog]:
        pass

    @abstractmethod
    async def get_by_id(self, log_id: int) -> LoginLog | None:
        pass
    @abstractmethod
    async def get_by_customer(self, customer_id: int) -> List[LoginLog]:
        pass

    @abstractmethod
    async def get_in_time_range(self, start_time: datetime, end_time: datetime) -> List[LoginLog]:
        pass