from abc import ABC, abstractmethod
from typing import List

from domain.transaction.transaction_fee import TransactionFee


class TransactionFeeRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[TransactionFee]:
        pass

    @abstractmethod
    async def get_by_id(self, id: int) -> TransactionFee | None:
        pass
