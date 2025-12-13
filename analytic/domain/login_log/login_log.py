from datetime import datetime


class LoginLog:
    def __init__(self,
                 log_id: int,
                 customer_id: int,
                 login_time: datetime,
                 ip_address: str,
                 device_info: str):
        self.log_id = log_id
        self.customer_id = customer_id
        self.login_time = login_time
        self.ip_address = ip_address
        self.device_info = device_info