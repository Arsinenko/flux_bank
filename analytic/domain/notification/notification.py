from datetime import datetime

class Notification:
    def __init__(self,
                 notification_id: int,
                 customer_id: int | None,
                 message: str | None,
                 created_at: datetime | None,
                 is_read: bool | None):
        self.notification_id = notification_id
        self.customer_id = customer_id
        self.message = message
        self.created_at = created_at
        self.is_read = is_read