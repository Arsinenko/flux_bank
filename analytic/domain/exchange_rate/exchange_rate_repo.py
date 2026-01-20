from abc import ABC, abstractmethod
from typing import List

from domain.exchange_rate.exchange_rate import ExchangeRate


class ExchangeRateRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[ExchangeRate]:
        pass

    @abstractmethod
    async def get_by_id(self, rate_id: int) -> ExchangeRate | None:
        pass

    @abstractmethod
    async def get_by_base_currency(self, base_currency: str) -> List[ExchangeRate]:
        pass
