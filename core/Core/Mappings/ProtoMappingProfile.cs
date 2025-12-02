using System;
using System.Globalization;
using AutoMapper;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using ProtoModels = global::Core;
using ProtoDateOnlyMessage = global::Core.DateOnly;
using SystemDateOnly = global::System.DateOnly;
using SystemTimeOnly = global::System.TimeOnly;

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
        CreateMap<Account, ProtoModels.AccountModel>()
            .ForMember(dest => dest.Balance,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Balance)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<ProtoModels.AccountModel, Account>()
            .ForMember(dest => dest.Balance,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.Balance)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.CreatedAt)))
            .ForMember(dest => dest.Cards, opt => opt.Ignore())
            .ForMember(dest => dest.Customer, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionSourceAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionTargetAccountNavigations, opt => opt.Ignore())
            .ForMember(dest => dest.Type, opt => opt.Ignore());
    }

    private void MapAccountType()
    {
        CreateMap<AccountType, ProtoModels.AccountTypeModel>();

        CreateMap<ProtoModels.AccountTypeModel, AccountType>()
            .ForMember(dest => dest.Accounts, opt => opt.Ignore());
    }

    private void MapAtm()
    {
        CreateMap<Atm, ProtoModels.AtmModel>();

        CreateMap<ProtoModels.AtmModel, Atm>()
            .ForMember(dest => dest.Branch, opt => opt.Ignore());
    }

    private void MapBranch()
    {
        CreateMap<Branch, ProtoModels.BranchModel>();

        CreateMap<ProtoModels.BranchModel, Branch>()
            .ForMember(dest => dest.Atms, opt => opt.Ignore());
    }

    private void MapCard()
    {
        CreateMap<Card, ProtoModels.CardModel>()
            .ForMember(dest => dest.ExpiryDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToTimestamp(src.ExpiryDate)));

        CreateMap<ProtoModels.CardModel, Card>()
            .ForMember(dest => dest.ExpiryDate,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateOnly(src.ExpiryDate)))
            .ForMember(dest => dest.Account, opt => opt.Ignore());
    }

    private void MapCustomer()
    {
        CreateMap<Customer, ProtoModels.CustomerModel>()
            .ForMember(dest => dest.BirthDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.BirthDate)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<ProtoModels.CustomerModel, Customer>()
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
    }

    private void MapCustomerAddress()
    {
        CreateMap<CustomerAddress, ProtoModels.CustomerAddressModel>();

        CreateMap<ProtoModels.CustomerAddressModel, CustomerAddress>()
            .ForMember(dest => dest.IsPrimary,
                opt => opt.MapFrom(src => src.HasIsPrimary ? src.IsPrimary : false))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapDeposit()
    {
        CreateMap<Deposit, ProtoModels.DepositModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.StartDate)))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.EndDate)));

        CreateMap<ProtoModels.DepositModel, Deposit>()
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
        CreateMap<ExchangeRate, ProtoModels.ExchangeRateModel>()
            .ForMember(dest => dest.Rate,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Rate)))
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.UpdatedAt)));

        CreateMap<ProtoModels.ExchangeRateModel, ExchangeRate>()
            .ForMember(dest => dest.Rate,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Rate)))
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.UpdatedAt)));
    }

    private void MapFeeType()
    {
        CreateMap<FeeType, ProtoModels.FeeTypeModel>();

        CreateMap<ProtoModels.FeeTypeModel, FeeType>()
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore());
    }

    private void MapLoan()
    {
        CreateMap<Loan, ProtoModels.LoanModel>()
            .ForMember(dest => dest.Principal,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Principal)))
            .ForMember(dest => dest.InterestRate,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.InterestRate)))
            .ForMember(dest => dest.StartDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.StartDate)))
            .ForMember(dest => dest.EndDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.EndDate)));

        CreateMap<ProtoModels.LoanModel, Loan>()
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
        CreateMap<LoanPayment, ProtoModels.LoanPaymentModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)))
            .ForMember(dest => dest.PaymentDate,
                opt => opt.MapFrom(src => MappingConverters.DateOnlyToProto(src.PaymentDate)))
            .ForMember(dest => dest.IsPaid,
                opt => opt.MapFrom(src => src.IsPaid));

        CreateMap<ProtoModels.LoanPaymentModel, LoanPayment>()
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
        CreateMap<LoginLog, ProtoModels.LoginLogModel>()
            .ForMember(dest => dest.LoginTime,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.LoginTime)));

        CreateMap<ProtoModels.LoginLogModel, LoginLog>()
            .ForMember(dest => dest.LoginTime,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.LoginTime) ?? default))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapNotification()
    {
        CreateMap<Notification, ProtoModels.NotificationModel>()
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<ProtoModels.NotificationModel, Notification>()
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.CreatedAt) ?? default))
            .ForMember(dest => dest.IsRead,
                opt => opt.MapFrom(src => src.HasIsRead ? src.IsRead : false))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapPaymentTemplate()
    {
        CreateMap<PaymentTemplate, ProtoModels.PaymentTemplateModel>()
            .ForMember(dest => dest.DefaultAmount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.DefaultAmount)));

        CreateMap<ProtoModels.PaymentTemplateModel, PaymentTemplate>()
            .ForMember(dest => dest.DefaultAmount,
                opt => opt.MapFrom(src => MappingConverters.StringToNullableDecimal(src.DefaultAmount)))
            .ForMember(dest => dest.Customer, opt => opt.Ignore());
    }

    private void MapTransaction()
    {
        CreateMap<Transaction, ProtoModels.TransactionModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.CreatedAt)));

        CreateMap<ProtoModels.TransactionModel, Transaction>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.CreatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.CreatedAt) ?? default))
            .ForMember(dest => dest.SourceAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TargetAccountNavigation, opt => opt.Ignore())
            .ForMember(dest => dest.TransactionFees, opt => opt.Ignore())
            .ForMember(dest => dest.Categories, opt => opt.Ignore());
    }

    private void MapTransactionCategory()
    {
        CreateMap<TransactionCategory, ProtoModels.TransactionCategoryModel>();

        CreateMap<ProtoModels.TransactionCategoryModel, TransactionCategory>()
            .ForMember(dest => dest.Transactions, opt => opt.Ignore());
    }

    private void MapTransactionFee()
    {
        CreateMap<TransactionFee, ProtoModels.TransactionFeeModel>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.DecimalToString(src.Amount)));

        CreateMap<ProtoModels.TransactionFeeModel, TransactionFee>()
            .ForMember(dest => dest.Amount,
                opt => opt.MapFrom(src => MappingConverters.StringToDecimal(src.Amount)))
            .ForMember(dest => dest.Transaction, opt => opt.Ignore())
            .ForMember(dest => dest.Fee, opt => opt.Ignore());
    }

    private void MapUserCredential()
    {
        CreateMap<UserCredential, ProtoModels.UserCredentialModel>()
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.DateTimeToTimestamp(src.UpdatedAt)));

        CreateMap<ProtoModels.UserCredentialModel, UserCredential>()
            .ForMember(dest => dest.UpdatedAt,
                opt => opt.MapFrom(src => MappingConverters.TimestampToDateTime(src.UpdatedAt) ?? default))
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

        return decimal.Parse(value, Culture);
    }

    public static decimal StringToDecimal(string? value)
    {
        if (string.IsNullOrWhiteSpace(value))
        {
            throw new InvalidOperationException("Decimal value must be provided.");
        }

        return decimal.Parse(value, Culture);
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

