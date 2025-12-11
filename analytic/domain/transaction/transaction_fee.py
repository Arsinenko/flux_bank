from decimal import Decimal


class TransactionFee:
    def __init__(self,
                 id: int,
                 transaction_id: int | None,
                 fee_id: int | None,
                 amount: Decimal | None):
        self.id = id
        self.transaction_id = transaction_id
        self.fee_id = fee_id
        self.amount = amount
