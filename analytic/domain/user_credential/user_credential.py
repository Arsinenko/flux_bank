from datetime import datetime


class UserCredential:
    def __init__(self,
                 customer_id: int,
                 username: str,
                 password_hash: str,
                 updated_at: datetime | None):
        self.customer_id = customer_id
        self.username = username
        self.password_hash = password_hash
        self.updated_at = updated_at
