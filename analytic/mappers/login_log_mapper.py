from typing import List
from api.generated.login_log_pb2 import LoginLogModel
from domain.login_log.login_log import LoginLog

class LoginLogMapper:
    @staticmethod
    def to_domain(model: LoginLogModel) -> LoginLog:
        return LoginLog(
            log_id=model.log_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            login_time=model.login_time.ToDatetime() if model.HasField("login_time") else None,
            ip_address=model.ip_address if model.HasField("ip_address") else None,
            device_info=model.device_info if model.HasField("device_info") else None
        )

    @staticmethod
    def to_model(domain: LoginLog) -> LoginLogModel:
        model = LoginLogModel(
            log_id=domain.log_id,
            customer_id=domain.customer_id,
            ip_address=domain.ip_address,
            device_info=domain.device_info
        )
        if domain.login_time:
            model.login_time.FromDatetime(domain.login_time)
        return model

    @staticmethod
    def to_domain_list(models: List[LoginLogModel]) -> List[LoginLog]:
        return [LoginLogMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[LoginLog]) -> List[LoginLogModel]:
        return [LoginLogMapper.to_model(domain) for domain in domains]
