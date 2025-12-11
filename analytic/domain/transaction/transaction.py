from datetime import datetime
from decimal import Decimal


class Transaction:
    def __init__(self,
                 transaction_id: int,
                 source_account: int | None,
                 target_account: int | None,
                 amount: Decimal,
                 currency: str,
                 created_at: datetime | None,
                 status: str | None):
        self.transaction_id = transaction_id
        self.source_account = source_account
        self.target_account = target_account
        self.amount = amount
        self.currency = currency
        self.created_at = created_at
        self.status = status
