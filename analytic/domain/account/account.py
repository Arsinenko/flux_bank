from datetime import datetime


class Account:
    def __init__(self, account_id: int, customer_id: int, type_id: int, iban: str, balance: str, created_at: datetime, is_active: bool):
        self.account_id: int = account_id
        self.customer_id: int = customer_id
        self.type_id: int = type_id
        self.iban: str = iban
        self.balance: str = balance
        self.created_at: datetime = created_at
        self.is_active: bool = is_active
