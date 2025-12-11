from abc import ABC, abstractmethod
from typing import List

from domain.payment_template.payment_template import PaymentTemplate


class PaymentTemplateRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[PaymentTemplate]:
        pass
    @abstractmethod
    async def get_by_id(self, template_id: int) -> PaymentTemplate | None:
        pass