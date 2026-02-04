import asyncio
import os

import grpc
import punq
from grpc_reflection.v1alpha import reflection

from adapters.account.account_repository import AccountRepository
from adapters.account.account_type_repository import AccountTypeRepository
from adapters.atm.atm_repository import AtmRepository
from adapters.branch.branch_repository import BranchRepository
from adapters.card.card_repository import CardRepository
from adapters.customer.customer_address_repository import CustomerAddressRepository
from adapters.customer.customer_repository import CustomerRepository
from adapters.deposit.deposit_repository import DepositRepository
from adapters.exchange_rate.exchange_rate_repository import ExchangeRateRepository
from adapters.fee_type.fee_type_repository import FeeTypeRepository
from adapters.loan.loan_payment_repository import LoanPaymentRepository
from adapters.loan.loan_repository import LoanRepository
from adapters.login_log.login_log_repository import LoginLogRepository
from adapters.notification.notification_repository import NotificationRepository
from adapters.payment_template.payment_template_repository import PaymentTemplateRepository
from adapters.transaction.transaction_category_repository import TransactionCategoryRepository
from adapters.transaction.transaction_fee_repository import TransactionFeeRepository
from adapters.transaction.transaction_repository import TransactionRepository
from adapters.user_credential.user_credential_repository import UserCredentialRepository

from domain.account.account_repo import AccountRepositoryAbc
from domain.account.account_type_repo import AccountTypeRepositoryAbc
from domain.atm.atm_repo import AtmRepositoryAbc
from domain.branch.branch_repo import BranchRepositoryAbc
from domain.card.card_repo import CardRepositoryAbc
from domain.customer.customer_address_repo import CustomerAddressRepositoryAbc
from domain.customer.customer_repo import CustomerRepositoryAbc
from domain.deposit.deposit_repo import DepositRepositoryAbc
from domain.exchange_rate.exchange_rate_repo import ExchangeRateRepositoryAbc
from domain.fee_type.fee_type_repo import FeeTypeRepositoryAbc
from domain.loan.loan_payment_repo import LoanPaymentRepositoryAbc
from domain.loan.loan_repo import LoanRepositoryAbc
from domain.login_log.login_log_repo import LoginLogRepositoryAbc
from domain.notification.notification_repo import NotificationRepositoryAbc
from domain.payment_template.payment_template_repo import PaymentTemplateRepositoryAbc
from domain.transaction.transaction_category_repo import TransactionCategoryRepositoryAbc
from domain.transaction.transaction_fee_repo import TransactionFeeRepositoryAbc
from domain.transaction.transaction_repo import TransactionRepositoryAbc
from domain.user_credential.user_credential_repo import UserCredentialRepositoryAbc

from services.account_analytic import AccountAnalyticService
from services.account_type_analytic import AccountTypeAnalytic
from services.atm_analytic import AtmAnalyticService
from services.branch_analytic import BranchAnalyticService
from services.card_analytic import CardAnalyticService
from services.customer_address_analytic import CustomerAddressAnalyticService
from services.customer_analytic import CustomerAnalyticService
from services.deposit_analytic import DepositAnalyticService
from services.exchange_rate_analytic import ExchangeRateAnalyticService
from services.fee_type_analytic import FeeTypeAnalyticService
from services.general_analytic import GeneralAnalyticService
from services.loan_analytic import LoanAnalyticService
from services.loan_payment_analytic import LoanPaymentAnalyticService
from services.login_log_analytic import LoginLogAnalyticService
from services.notification_analytic import NotificationAnalyticService
from services.payment_template_analytic import PaymentTemplateAnalyticService
from services.transaction_analytic import TransactionAnalyticService
from services.transaction_category_analytic import TransactionCategoryAnalyticService
from services.transaction_fee_analytic import TransactionFeeAnalyticService
from services.user_credential_analytic import UserCredentialAnalyticService


def get_container() -> punq.Container:
    grpc_target = os.environ.get("GRPC_TARGET", "localhost:8080")
    channel = grpc.aio.insecure_channel(grpc_target)
    container = punq.Container()

    # Repositories
    container.register(AccountRepositoryAbc, AccountRepository, instance=AccountRepository(channel))
    container.register(AccountTypeRepositoryAbc, AccountTypeRepository, instance=AccountTypeRepository(channel))
    container.register(AtmRepositoryAbc, AtmRepository, instance=AtmRepository(channel))
    container.register(BranchRepositoryAbc, BranchRepository, instance=BranchRepository(channel))
    container.register(CardRepositoryAbc, CardRepository, instance=CardRepository(channel))
    container.register(CustomerAddressRepositoryAbc, CustomerAddressRepository, instance=CustomerAddressRepository(channel))
    container.register(CustomerRepositoryAbc, CustomerRepository, instance=CustomerRepository(channel))
    container.register(DepositRepositoryAbc, DepositRepository, instance=DepositRepository(channel))
    container.register(ExchangeRateRepositoryAbc, ExchangeRateRepository, instance=ExchangeRateRepository(channel))
    container.register(FeeTypeRepositoryAbc, FeeTypeRepository, instance=FeeTypeRepository(channel))
    container.register(LoanPaymentRepositoryAbc, LoanPaymentRepository, instance=LoanPaymentRepository(channel))
    container.register(LoanRepositoryAbc, LoanRepository, instance=LoanRepository(channel))
    container.register(LoginLogRepositoryAbc, LoginLogRepository, instance=LoginLogRepository(channel))
    container.register(NotificationRepositoryAbc, NotificationRepository, instance=NotificationRepository(channel))
    container.register(PaymentTemplateRepositoryAbc, PaymentTemplateRepository, instance=PaymentTemplateRepository(channel))
    container.register(TransactionCategoryRepositoryAbc, TransactionCategoryRepository, instance=TransactionCategoryRepository(channel))
    container.register(TransactionFeeRepositoryAbc, TransactionFeeRepository, instance=TransactionFeeRepository(channel))
    container.register(TransactionRepositoryAbc, TransactionRepository, instance=TransactionRepository(channel))
    container.register(UserCredentialRepositoryAbc, UserCredentialRepository, instance=UserCredentialRepository(channel))

    # Services
    container.register(AccountAnalyticService)
    container.register(AccountTypeAnalytic)
    container.register(AtmAnalyticService)
    container.register(BranchAnalyticService)
    container.register(CardAnalyticService)
    container.register(CustomerAddressAnalyticService)
    container.register(CustomerAnalyticService)
    container.register(DepositAnalyticService)
    container.register(ExchangeRateAnalyticService)
    container.register(FeeTypeAnalyticService)
    container.register(GeneralAnalyticService)
    container.register(LoanAnalyticService)
    container.register(LoanPaymentAnalyticService)
    container.register(LoginLogAnalyticService)
    container.register(NotificationAnalyticService)
    container.register(PaymentTemplateAnalyticService)
    container.register(TransactionAnalyticService)
    container.register(TransactionCategoryAnalyticService)
    container.register(TransactionFeeAnalyticService)
    container.register(UserCredentialAnalyticService)

    return container


async def serve():
    container = get_container()

    server = grpc.aio.server()

    account_analytic_service = container.resolve(AccountAnalyticService)
    account_type_analytic_service = container.resolve(AccountTypeAnalytic)
    atm_analytic_service = container.resolve(AtmAnalyticService)
    branch_analytic_service = container.resolve(BranchAnalyticService)
    card_analytic_service = container.resolve(CardAnalyticService)
    customer_address_analytic_service = container.resolve(CustomerAddressAnalyticService)
    customer_analytic_service = container.resolve(CustomerAnalyticService)
    deposit_analytic_service = container.resolve(DepositAnalyticService)
    exchange_rate_analytic_service = container.resolve(ExchangeRateAnalyticService)
    fee_type_analytic_service = container.resolve(FeeTypeAnalyticService)
    general_analytic_service = container.resolve(GeneralAnalyticService)
    loan_analytic_service = container.resolve(LoanAnalyticService)
    loan_payment_analytic_service = container.resolve(LoanPaymentAnalyticService)
    login_log_analytic_service = container.resolve(LoginLogAnalyticService)
    notification_analytic_service = container.resolve(NotificationAnalyticService)
    payment_template_analytic_service = container.resolve(PaymentTemplateAnalyticService)
    transaction_analytic_service = container.resolve(TransactionAnalyticService)
    transaction_category_analytic_service = container.resolve(TransactionCategoryAnalyticService)
    transaction_fee_analytic_service = container.resolve(TransactionFeeAnalyticService)
    user_credential_analytic_service = container.resolve(UserCredentialAnalyticService)

    from api.generated.account_analytic_pb2_grpc import add_AccountAnalyticServiceServicer_to_server
    from api.generated.account_type_analytic_pb2_grpc import add_AccountTypeAnalyticServiceServicer_to_server
    from api.generated.atm_analytic_pb2_grpc import add_AtmAnalyticServiceServicer_to_server
    from api.generated.branch_analytic_pb2_grpc import add_BranchAnalyticServiceServicer_to_server
    from api.generated.card_analytic_pb2_grpc import add_CardAnalyticServiceServicer_to_server
    from api.generated.customer_address_analytic_pb2_grpc import add_CustomerAddressAnalyticServiceServicer_to_server
    from api.generated.customer_analytic_pb2_grpc import add_CustomerAnalyticServiceServicer_to_server
    from api.generated.deposit_analytic_pb2_grpc import add_DepositAnalyticServiceServicer_to_server
    from api.generated.exchange_rate_analytic_pb2_grpc import add_ExchangeRateAnalyticServiceServicer_to_server
    from api.generated.fee_type_analytic_pb2_grpc import add_FeeTypeAnalyticServiceServicer_to_server
    from api.generated.general_analytic_pb2_grpc import add_GeneralAnalyticServiceServicer_to_server
    from api.generated.loan_analytic_pb2_grpc import add_LoanAnalyticServiceServicer_to_server
    from api.generated.loan_payment_analytic_pb2_grpc import add_LoanPaymentAnalyticServiceServicer_to_server
    from api.generated.login_log_analytic_pb2_grpc import add_LoginLogAnalyticServiceServicer_to_server
    from api.generated.notification_analytic_pb2_grpc import add_NotificationAnalyticServiceServicer_to_server
    from api.generated.payment_template_analytic_pb2_grpc import add_PaymentTemplateAnalyticServiceServicer_to_server
    from api.generated.transaction_analitic_pb2_grpc import add_TransactionAnalyticServiceServicer_to_server
    from api.generated.transaction_category_analytic_pb2_grpc import add_TransactionCategoryAnalyticServiceServicer_to_server
    from api.generated.transaction_fee_analytic_pb2_grpc import add_TransactionFeeAnalyticServiceServicer_to_server
    from api.generated.user_credential_analytic_pb2_grpc import add_UserCredentialAnalyticServiceServicer_to_server

    add_AccountAnalyticServiceServicer_to_server(account_analytic_service, server)
    add_AccountTypeAnalyticServiceServicer_to_server(account_type_analytic_service, server)
    add_AtmAnalyticServiceServicer_to_server(atm_analytic_service, server)
    add_BranchAnalyticServiceServicer_to_server(branch_analytic_service, server)
    add_CardAnalyticServiceServicer_to_server(card_analytic_service, server)
    add_CustomerAddressAnalyticServiceServicer_to_server(customer_address_analytic_service, server)
    add_CustomerAnalyticServiceServicer_to_server(customer_analytic_service, server)
    add_DepositAnalyticServiceServicer_to_server(deposit_analytic_service, server)
    add_ExchangeRateAnalyticServiceServicer_to_server(exchange_rate_analytic_service, server)
    add_FeeTypeAnalyticServiceServicer_to_server(fee_type_analytic_service, server)
    add_GeneralAnalyticServiceServicer_to_server(general_analytic_service, server)
    add_LoanAnalyticServiceServicer_to_server(loan_analytic_service, server)
    add_LoanPaymentAnalyticServiceServicer_to_server(loan_payment_analytic_service, server)
    add_LoginLogAnalyticServiceServicer_to_server(login_log_analytic_service, server)
    add_NotificationAnalyticServiceServicer_to_server(notification_analytic_service, server)
    add_PaymentTemplateAnalyticServiceServicer_to_server(payment_template_analytic_service, server)
    add_TransactionAnalyticServiceServicer_to_server(transaction_analytic_service, server)
    add_TransactionCategoryAnalyticServiceServicer_to_server(transaction_category_analytic_service, server)
    add_TransactionFeeAnalyticServiceServicer_to_server(transaction_fee_analytic_service, server)
    add_UserCredentialAnalyticServiceServicer_to_server(user_credential_analytic_service, server)

    SERVICE_NAMES = [
        'account_analytic.AccountAnalyticService',
        'account_type_analytic.AccountTypeAnalyticService',
        'atm_analytic.AtmAnalyticService',
        'branch_analytic.BranchAnalyticService',
        'card_analytic.CardAnalyticService',
        'customer_address_analytic.CustomerAddressAnalyticService',
        'customer_analytic.CustomerAnalyticService',
        'deposit_analytic.DepositAnalyticService',
        'exchange_rate_analytic.ExchangeRateAnalyticService',
        'fee_type_analytic.FeeTypeAnalyticService',
        'general_analytic.GeneralAnalyticService',
        'loan_analytic.LoanAnalyticService',
        'loan_payment_analytic.LoanPaymentAnalyticService',
        'login_log_analytic.LoginLogAnalyticService',
        'notification_analytic.NotificationAnalyticService',
        'payment_template_analytic.PaymentTemplateAnalyticService',
        'transaction_analytic.TransactionAnalyticService',
        'transaction_category_analytic.TransactionCategoryAnalyticService',
        'transaction_fee_analytic.TransactionFeeAnalyticService',
        'user_credential_analytic.UserCredentialAnalyticService',
        # Добавляем стандартный reflection сервис
        reflection.SERVICE_NAME,
    ]

    # Включаем reflection
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    listen_addr = "50051"

    server.add_insecure_port(f"[::]:{listen_addr}")


    print(f"Starting server on {listen_addr}")
    print("Server reflection is ENABLED")
    try:
        await server.start()
        await server.wait_for_termination()
    except KeyboardInterrupt:
        await server.stop(5)



async def main():
    # container = get_container()
    # customer_analytic_service: CustomerAnalyticService = container.resolve(CustomerAnalyticService)
    #
    # result = await customer_analytic_service.customer_repo.get_all(0, 0, "", False)
    # print(result)
    await serve()



if __name__ == '__main__':
    asyncio.run(main())

