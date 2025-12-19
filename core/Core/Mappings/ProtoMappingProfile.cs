using System.Globalization;
using AutoMapper;
using Core.Exceptions;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using ProtoDateOnlyMessage = Core.DateOnly;
using SystemDateOnly = System.DateOnly;
using SystemTimeOnly = System.TimeOnly;

namespace Core.Mappings;

public sealed class ProtoMappingProfile : Profile
{
    public ProtoMappingProfile()
    {
        MapAccount();
        MapAccountType();
        MapAtm();
        MapBranch();
        MapCard();
        MapCustomer();
        MapCustomerAddress();
        MapDeposit();
        MapExchangeRate();
        MapFeeType();
        MapLoan();
        MapLoanPayment();
        MapLoginLog();
        MapNotification();
        MapPaymentTemplate();
        MapTransaction();
        MapTransactionCategory();
        MapTransactionFee();
        MapUserCredential();
    }

    private void MapAccount()
    {
        CreateMap<Account, AccountModel>()
            .ForMember(dest => dest.Balance,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Balance)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<AccountModel, Account>()
            .ForMember(dest => dest.Balance,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.Balance)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.CreatedAt)))
            .ForMember(dest => dest.Cards, opt => opt.Ignore())
            .ForMember(dest => dest.Customer, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionSourceAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionTargetAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.Type, opt => opt.Ignore());
        CreateMap<AddAccountRequest, Account>()
            .ForMember(dest => dest.Balance,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.Balance)))
            .ForMember(dest => dest.AccountId, opt => opt.Ignore())
            .ForMember(dest => dest.CreatedAt, opt => opt.Ignore())
            .ForMember(dest => dest.Cards, opt => opt.Ignore())
            .ForMember(dest => dest.Customer, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionSourceAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionTargetAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.Type, opt => opt.Ignore());
        CreateMap<UpdateAccountRequest, Account>()
            .ForMember(dest => dest.Balance,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.Balance)))
            .ForMember(dest => dest.CreatedAt, opt => opt.Ignore())
            .ForMember(dest => dest.Cards, opt => opt.Ignore())
            .ForMember(dest => dest.Customer, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionSourceAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionTargetAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.Type, opt => opt.Ignore());
    }

    private void MapAccountType()
    {
        CreateMap<AccountType, AccountTypeModel>();

        CreateMap<AccountTypeModel, AccountType>()
            .ForMember(dest => dest.Accounts, opt => opt.Ignore());
        CreateMap<AddAccountTypeRequest, AccountType>()
            .ForMember(dest => dest.TypeId, opt => opt.Ignore())
            .ForMember(dest => dest.Accounts, opt => opt.Ignore());
        CreateMap<UpdateAccountTypeRequest, AccountType>()
            .ForMember(dest => dest.Accounts, opt => opt.Ignore());
    }

    private void MapAtm()
    {
        CreateMap<Atm, AtmModel>();

        CreateMap<AtmModel, Atm>()
            .ForMember(dest => dest.Branch, opt => opt.Ignore());
        CreateMap<AddAtmRequest, Atm>()
            .ForMember(dest => dest.AtmId, opt => opt.Ignore())
            .ForMember(dest => dest.Branch, opt => opt.Ignore());
        
        CreateMap<UpdateAtmRequest, Atm>()
            .ForMember(dest => dest.Branch, opt => opt.Ignore());
    }

    private void MapBranch()
    {
        CreateMap<Branch, BranchModel>();

        CreateMap<BranchModel, Branch>()
            .ForMember(dest => dest.Atms, opt => opt.Ignore());
        CreateMap<AddBranchRequest, Branch>()
            .ForMember(dest => dest.BranchId, opt => opt.Ignore())
            .ForMember(dest => dest.Atms, opt => opt.Ignore());
        CreateMap<UpdateBranchRequest, Branch>()
            .ForMember(dest => dest.Atms, opt => opt.Ignore());
    }

    private void MapCard()
    {
        CreateMap<Card, CardModel>()
            .ForMember(dest => dest.ExpiryDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToTimestamp(src.ExpiryDate)));

        CreateMap<CardModel, Card>()
            .ForMember(dest => dest.ExpiryDate,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateOnly(src.ExpiryDate)))
            .ForMember(dest => dest.Account, opt => opt.Ignore());

        CreateMap<AddCardRequest, Card>()
            .ForMember(dest => dest.ExpiryDate,
                opt => opt.MapFrom(src =>
                    MappingConverters.DateOnlyToTimestamp(
                        src.ExpiryDate == null ? null : ProtoDateOnlyToSystemDateOnly(src.ExpiryDate))))
            .ForMember(dest => dest.CardId, opt => opt.Ignore())
            .ForMember(dest => dest.Account, opt => opt.Ignore());

        CreateMap<UpdateCardRequest, Card>()
            .ForMember(dest => dest.ExpiryDate,
                opt => opt.MapFrom(src =>
                    src.ExpiryDate == null
                        ? null
                        : MappingConverters.DateOnlyToTimestamp(ProtoDateOnlyToSystemDateOnly(src.ExpiryDate))))
            .ForMember(dest => dest.Account, opt => opt.Ignore());
    }

    private static SystemDateOnly? ProtoDateOnlyToSystemDateOnly(ProtoDateOnlyMessage? protoDate)
    {
        if (protoDate == null) return null;
        return new SystemDateOnly(protoDate.Year, protoDate.Month, protoDate.Day);
    }

    private void MapCustomer()
    {
        CreateMap<Customer, CustomerModel>()
            .ForMember(dest => dest.BirthDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.BirthDate)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<CustomerModel, Customer>()
            .ForMember(dest => dest.BirthDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.BirthDate) ?? default))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.CreatedAt) ?? default))
            .ForMember(dest => dest.Phone,
                opt => opt.MapFrom(src => src.Phone ?? string.Empty))
            .ForMember(dest => dest.Accounts, opt => opt.Ignore())
            .ForMember(dest => dest.CustomerAddresses, opt => opt.Ignore())
            .ForMember(dest => dest.Deposits, opt => opt.Ignore())
            .ForMember(dest => dest.Loans, opt => opt.Ignore())
            .ForMember(dest => dest.LoginLogs, opt => opt.Ignore())
            .ForMember(dest => dest.Notifications, opt => opt.Ignore())
            .ForMember(dest => dest.PaymentTemplates, opt => opt.Ignore())
            .ForMember(dest => dest.UserCredential, opt => opt.Ignore());
        CreateMap<AddCustomerRequest, Customer>()
            .ForMember(dest => dest.BirthDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.BirthDate) ?? default))
            .ForMember(dest => dest.CustomerId, opt => opt.Ignore())
            .ForMember(dest => dest.CreatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.Phone,
                opt => opt.MapFrom(src => src.Phone ?? string.Empty))
            .ForMember(dest => dest.Accounts, opt => opt.Ignore())
            .ForMember(dest => dest.CustomerAddresses, opt => opt.Ignore())
            .ForMember(dest => dest.Deposits, opt => opt.Ignore())
            .ForMember(dest => dest.Loans, opt => opt.Ignore())
            .ForMember(dest => dest.LoginLogs, opt => opt.Ignore())
            .ForMember(dest => dest.Notifications, opt => opt.Ignore())
            .ForMember(dest => dest.PaymentTemplates, opt => opt.Ignore())
            .ForMember(dest => dest.UserCredential, opt => opt.Ignore());
        CreateMap<UpdateCustomerRequest, Customer>()
            .ForMember(dest => dest.BirthDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.BirthDate) ?? default))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.Phone,
                opt => opt.MapFrom(src => src.Phone ?? string.Empty))
            .ForMember(dest => dest.Accounts, opt => opt.Ignore())
            .ForMember(dest => dest.CustomerAddresses, opt => opt.Ignore())
            .ForMember(dest => dest.Deposits, opt => opt.Ignore())
            .ForMember(dest => dest.Loans, opt => opt.Ignore())
            .ForMember(dest => dest.LoginLogs, opt => opt.Ignore())
            .ForMember(dest => dest.Notifications, opt => opt.Ignore())
            .ForMember(dest => dest.PaymentTemplates, opt => opt.Ignore())
            .ForMember(dest => dest.UserCredential, opt => opt.Ignore());
    }

    private void MapCustomerAddress()
    {
        CreateMap<CustomerAddress, CustomerAddressModel>();

        CreateMap<CustomerAddressModel, CustomerAddress>()
            .ForMember(dest => dest.IsPrimary,
                opt => opt.MapFrom(src => src.HasIsPrimary && src.IsPrimary))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<AddCustomerAddressRequest, CustomerAddress>()
            .ForMember(dest => dest.AddressId, opt => opt.Ignore())
            .ForMember(dest => dest.IsPrimary,
                opt => opt.MapFrom(src => src.HasIsPrimary && src.IsPrimary))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<UpdateCustomerAddressRequest, CustomerAddress>()
            .ForMember(dest => dest.IsPrimary,
                opt => opt.MapFrom(src => src.HasIsPrimary && src.IsPrimary))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapDeposit()
    {
        CreateMap<Deposit, DepositModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.StartDate)))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.EndDate)));

        CreateMap<DepositModel, Deposit>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.StartDate) ?? default))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.EndDate)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<AddDepositRequest, Deposit>()
            .ForMember(dest => dest.DepositId, opt => opt.Ignore())
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.StartDate) ?? default))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.EndDate)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<UpdateDepositRequest, Deposit>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.StartDate) ?? default))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.EndDate)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapExchangeRate()
    {
        CreateMap<ExchangeRate, ExchangeRateModel>()
            .ForMember(dest => dest.Rate,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Rate)))
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.UpdatedAt)));

        CreateMap<ExchangeRateModel, ExchangeRate>()
            .ForMember(dest => dest.Rate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Rate)))
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.UpdatedAt)));
        CreateMap<AddExchangeRateRequest, ExchangeRate>()
            .ForMember(dest => dest.RateId, opt => opt.Ignore())
            .ForMember(dest => dest.Rate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Rate)))
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.Ignore());
        CreateMap<UpdateExchangeRateRequest, ExchangeRate>()
            .ForMember(dest => dest.Rate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Rate)))
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.Ignore());
    }

    private void MapFeeType()
    {
        CreateMap<FeeType, FeeTypeModel>();

        CreateMap<FeeTypeModel, FeeType>()
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore());
        CreateMap<AddFeeTypeRequest, FeeType>()
            .ForMember(dest => dest.FeeId, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore());
        CreateMap<UpdateFeeTypeRequest, FeeType>()
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore());
    }

    private void MapLoan()
    {
        CreateMap<Loan, LoanModel>()
            .ForMember(dest => dest.Principal,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Principal)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.StartDate)))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.EndDate)));

        CreateMap<LoanModel, Loan>()
            .ForMember(dest => dest.Principal,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Principal)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.StartDate) ?? default))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.EndDate)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore())
            .ForMember(dest => dest.LoanPayments, opt => opt.Ignore());
        CreateMap<AddLoanRequest, Loan>()
            .ForMember(dest => dest.LoanId, opt => opt.Ignore())
            .ForMember(dest => dest.Principal,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Principal)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.StartDate) ?? default))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.EndDate)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore())
            .ForMember(dest => dest.LoanPayments, opt => opt.Ignore());
        CreateMap<UpdateLoanRequest, Loan>()
            .ForMember(dest => dest.Principal,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Principal)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.StartDate) ?? default))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.EndDate)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore())
            .ForMember(dest => dest.LoanPayments, opt => opt.Ignore());
    }

    private void MapLoanPayment()
    {
        CreateMap<LoanPayment, LoanPaymentModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)))
            .ForMember(dest => dest.PaymentDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.PaymentDate)))
            .ForMember(dest => dest.IsPaid,
                opt => opt.MapFrom(src => src.IsPaid));

        CreateMap<LoanPaymentModel, LoanPayment>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.PaymentDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.PaymentDate) ?? default))
            .ForMember(dest => dest.IsPaid,
                opt => opt.MapFrom(src => src.HasIsPaid ? src.IsPaid : false))
            .ForMember(dest => dest.Loan, opt => opt.Ignore());
        CreateMap<AddLoanPaymentRequest, LoanPayment>()
            .ForMember(dest => dest.PaymentId, opt => opt.Ignore())
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.PaymentDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.PaymentDate) ?? default))
            .ForMember(dest => dest.IsPaid,
                opt => opt.MapFrom(src => src.HasIsPaid ? src.IsPaid : false))
            .ForMember(dest => dest.Loan, opt => opt.Ignore());
        CreateMap<UpdateLoanPaymentRequest, LoanPayment>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.PaymentDate,
                opt => opt.MapFrom(src => MappingConverters.ProtoToDateOnly(src.PaymentDate) ?? default))
            .ForMember(dest => dest.IsPaid,
                opt => opt.MapFrom(src => src.HasIsPaid ? src.IsPaid : false))
            .ForMember(dest => dest.Loan, opt => opt.Ignore());
    }

    private void MapLoginLog()
    {
        CreateMap<LoginLog, LoginLogModel>()
            .ForMember(dest => dest.LoginTime,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.LoginTime)));

        CreateMap<LoginLogModel, LoginLog>()
            .ForMember(dest => dest.LoginTime,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.LoginTime) ?? default))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<AddLoginLogRequest, LoginLog>()
            .ForMember(dest => dest.LogId, opt => opt.Ignore())
            .ForMember(dest => dest.LoginTime,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.LoginTime) ?? default))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<UpdateLoginLogRequest, LoginLog>()
            .ForMember(dest => dest.LoginTime,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.LoginTime) ?? default))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapNotification()
    {
        CreateMap<Notification, NotificationModel>()
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<NotificationModel, Notification>()
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.CreatedAt) ?? default))
            .ForMember(dest => dest.IsRead,
                opt => opt.MapFrom(src => src.HasIsRead ? src.IsRead : false))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<AddNotificationRequest, Notification>()
            .ForMember(dest => dest.NotificationId, opt => opt.Ignore())
            .ForMember(dest => dest.CreatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.IsRead,
                opt => opt.MapFrom(src => src.HasIsRead ? src.IsRead : false))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<UpdateNotificationRequest, Notification>()
            .ForMember(dest => dest.CreatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.IsRead,
                opt => opt.MapFrom(src => src.HasIsRead ? src.IsRead : false))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapPaymentTemplate()
    {
        CreateMap<PaymentTemplate, PaymentTemplateModel>()
            .ForMember(dest => dest.DefaultAmount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.DefaultAmount)));

        CreateMap<PaymentTemplateModel, PaymentTemplate>()
            .ForMember(dest => dest.DefaultAmount,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.DefaultAmount)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<AddPaymentTemplateRequest, PaymentTemplate>()
            .ForMember(dest => dest.TemplateId, opt =>opt.Ignore())
            .ForMember(dest => dest.DefaultAmount,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.DefaultAmount)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<UpdatePaymentTemplateRequest, PaymentTemplate>()
            .ForMember(dest => dest.DefaultAmount,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.DefaultAmount)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapTransaction()
    {
        CreateMap<Transaction, TransactionModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<TransactionModel, Transaction>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.CreatedAt) ?? default))
            .ForMember(dest => dest.SourceAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TargetAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore())
            .ForMember(dest => dest.Categories, opt => opt.Ignore());
        CreateMap<AddTransactionRequest, Transaction>()
            .ForMember(dest => dest.TransactionId, opt => opt.Ignore())
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.SourceAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TargetAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore())
            .ForMember(dest => dest.Categories, opt => opt.Ignore());
        CreateMap<UpdateTransactionRequest, Transaction>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.SourceAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TargetAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore())
            .ForMember(dest => dest.Categories, opt => opt.Ignore());
    }

    private void MapTransactionCategory()
    {
        CreateMap<TransactionCategory, TransactionCategoryModel>();

        CreateMap<TransactionCategoryModel, TransactionCategory>()
            .ForMember(dest => dest.Transactions, opt => opt.Ignore());
        CreateMap<AddTransactionCategoryRequest, TransactionCategory>()
            .ForMember(dest => dest.CategoryId, opt => opt.Ignore())
            .ForMember(dest => dest.Transactions, opt => opt.Ignore());
        CreateMap<UpdateTransactionCategoryRequest, TransactionCategory>()
            .ForMember(dest => dest.Transactions, opt => opt.Ignore());
    }

    private void MapTransactionFee()
    {
        CreateMap<TransactionFee, TransactionFeeModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)));

        CreateMap<TransactionFeeModel, TransactionFee>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.Transaction, opt => opt.Ignore())
            .ForMember(dest => dest.Fee, opt => opt.Ignore());
        CreateMap<AddTransactionFeeRequest, TransactionFee>()
            .ForMember(dest => dest.Id, opt => opt.Ignore())
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.Transaction, opt => opt.Ignore())
            .ForMember(dest => dest.Fee, opt => opt.Ignore());
        CreateMap<UpdateTransactionFeeRequest, TransactionFee>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.Transaction, opt => opt.Ignore())
            .ForMember(dest => dest.Fee, opt => opt.Ignore());
    }

    private void MapUserCredential()
    {
        CreateMap<UserCredential, UserCredentialModel>()
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.UpdatedAt)));

        CreateMap<UserCredentialModel, UserCredential>()
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.UpdatedAt) ?? default))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<AddUserCredentialRequest, UserCredential>()
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
        CreateMap<UpdateUserCredentialRequest, UserCredential>()
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.Ignore())
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }
}

internal static class MappingConverters
{
    private static readonly CultureInfo Culture = CultureInfo.InvariantCulture;

    public static string? DecimalToString(decimal? value) =>
        value?.ToString(Culture);

    public static string DecimalToString(decimal value) =>
        value.ToString(Culture);

    public static decimal? StringToNullableDecimal(string? value)
    {
        if (string.IsNullOrWhiteSpace(value))
        {
            return null;
        }

        if (decimal.TryParse(value, Culture, out var result))
        {
            return result;
        }

        return null;
    }

    public static decimal StringToDecimal(string? value)
    {
        if (string.IsNullOrWhiteSpace(value))
        {
            throw new InvalidOperationException("Decimal value must be provided.");
        }

        if (decimal.TryParse(value, Culture, out var result))
        {
            return result;
        }

        throw new ValidationException($"Decimal value must be provided and in correct format. Value: '{value}'");
    }

    public static Timestamp? DateTimeToTimestamp(DateTime? value) =>
        value == null ? null : Timestamp.FromDateTime(DateTime.SpecifyKind(value.Value, DateTimeKind.Utc));

    public static DateTime? TimestampToDateTime(Timestamp? value) =>
        value?.ToDateTime();

    public static ProtoDateOnlyMessage? DateOnlyToProto(SystemDateOnly? value)
    {
        if (value == null)
        {
            return null;
        }

        var actual = value.Value;

        return new ProtoDateOnlyMessage
        {
            Year = actual.Year,
            Month = actual.Month,
            Day = actual.Day
        };
    }

    public static SystemDateOnly? ProtoToDateOnly(ProtoDateOnlyMessage? value)
    {
        if (value == null)
        {
            return null;
        }

        return new SystemDateOnly(value.Year, value.Month, value.Day);
    }

    public static Timestamp? DateOnlyToTimestamp(SystemDateOnly? value)
    {
        if (value == null)
        {
            return null;
        }

        var asDateTime = value.Value.ToDateTime(SystemTimeOnly.MinValue);
        return Timestamp.FromDateTime(DateTime.SpecifyKind(asDateTime, DateTimeKind.Utc));
    }

    public static SystemDateOnly? TimestampToDateOnly(Timestamp? value) =>
        value == null ? null : SystemDateOnly.FromDateTime(value.ToDateTime());
}