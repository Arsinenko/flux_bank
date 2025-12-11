from datetime import datetime
from decimal import Decimal


class ExchangeRate:
    def __init__(self,
                 rate_id: int,
                 base_currency: str | None,
                 target_currency: str | None,
                 rate: Decimal | None,
                 updated_at: datetime | None):
        self.rate_id = rate_id
        self.base_currency = base_currency
        self.target_currency = target_currency
        self.rate = rate
        self.updated_at = updated_at
