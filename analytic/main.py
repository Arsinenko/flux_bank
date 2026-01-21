import os
import punq

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
    grpc_target = os.environ.get("GRPC_TARGET", "localhost:50051")
    container = punq.Container()

    # Repositories
    container.register(AccountRepositoryAbc, AccountRepository, instance=AccountRepository(grpc_target))
    container.register(AccountTypeRepositoryAbc, AccountTypeRepository, instance=AccountTypeRepository(grpc_target))
    container.register(AtmRepositoryAbc, AtmRepository, instance=AtmRepository(grpc_target))
    container.register(BranchRepositoryAbc, BranchRepository, instance=BranchRepository(grpc_target))
    container.register(CardRepositoryAbc, CardRepository, instance=CardRepository(grpc_target))
    container.register(CustomerAddressRepositoryAbc, CustomerAddressRepository, instance=CustomerAddressRepository(grpc_target))
    container.register(CustomerRepositoryAbc, CustomerRepository, instance=CustomerRepository(grpc_target))
    container.register(DepositRepositoryAbc, DepositRepository, instance=DepositRepository(grpc_target))
    container.register(ExchangeRateRepositoryAbc, ExchangeRateRepository, instance=ExchangeRateRepository(grpc_target))
    container.register(FeeTypeRepositoryAbc, FeeTypeRepository, instance=FeeTypeRepository(grpc_target))
    container.register(LoanPaymentRepositoryAbc, LoanPaymentRepository, instance=LoanPaymentRepository(grpc_target))
    container.register(LoanRepositoryAbc, LoanRepository, instance=LoanRepository(grpc_target))
    container.register(LoginLogRepositoryAbc, LoginLogRepository, instance=LoginLogRepository(grpc_target))
    container.register(NotificationRepositoryAbc, NotificationRepository, instance=NotificationRepository(grpc_target))
    container.register(PaymentTemplateRepositoryAbc, PaymentTemplateRepository, instance=PaymentTemplateRepository(grpc_target))
    container.register(TransactionCategoryRepositoryAbc, TransactionCategoryRepository, instance=TransactionCategoryRepository(grpc_target))
    container.register(TransactionFeeRepositoryAbc, TransactionFeeRepository, instance=TransactionFeeRepository(grpc_target))
    container.register(TransactionRepositoryAbc, TransactionRepository, instance=TransactionRepository(grpc_target))
    container.register(UserCredentialRepositoryAbc, UserCredentialRepository, instance=UserCredentialRepository(grpc_target))

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

if __name__ == "__main__":
    container = get_container()
    # Example of resolving a service
    # general_analytic = container.resolve(GeneralAnalyticService)
    # print(general_analytic)
