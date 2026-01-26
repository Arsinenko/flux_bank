using Core;
using Core.Context;
using Core.Exceptions;
using Core.Interfaces;
using Core.Mappings;
using Core.Repositories;
using Core.Services;
using Microsoft.EntityFrameworkCore;

using AccountService = Core.Services.AccountService;
using AccountTypeService = Core.Services.AccountTypeService;
using AtmService = Core.Services.AtmService;
using BranchService = Core.Services.BranchService;
using CardService = Core.Services.CardService;
using CustomerAddressService = Core.Services.CustomerAddressService;
using CustomerService = Core.Services.CustomerService;
using DepositService = Core.Services.DepositService;
using ExchangeRateService = Core.Services.ExchangeRateService;
using FeeTypeService = Core.Services.FeeTypeService;
using LoanPaymentService = Core.Services.LoanPaymentService;
using LoanService = Core.Services.LoanService;
using LoginLogService = Core.Services.LoginLogService;
using NotificationService = Core.Services.NotificationService;
using TransactionCategoryService = Core.Services.TransactionCategoryService;
using TransactionFeeService = Core.Services.TransactionFeeService;
using TransactionService = Core.Services.TransactionService;
using UserCredentialService = Core.Services.UserCredentialService;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddGrpc(options =>
{
    options.Interceptors.Add<GrpcExceptionInterceptor>();
});

builder.Services.AddAutoMapper(typeof(ProtoMappingProfile));

builder.Services.AddDbContext<MyDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection")));

builder.Services.AddMemoryCache();
builder.Services.AddSingleton<ICacheService, CacheService>();
builder.Services.AddScoped<IStatsService, StatsService>();

builder.Services.AddScoped<IAccountRepository, AccountRepository>();
builder.Services.AddScoped<IAccountTypeRepository, AccountTypeRepository>();
builder.Services.AddScoped<IAtmRepository, AtmRepository>();
builder.Services.AddScoped<IBranchRepository, BranchRepository>();
builder.Services.AddScoped<ICardRepository, CardRepository>();
builder.Services.AddScoped<ICustomerAddressRepository, CustomerAddressRepository>();
builder.Services.AddScoped<ICustomerRepository, CustomerRepository>();
builder.Services.AddScoped<IDepositRepository, DepositRepository>();
builder.Services.AddScoped<IExchangeRateRepository, ExchangeRateRepository>();
builder.Services.AddScoped<IFeeTypeRepository, FeeTypeRepository>();
builder.Services.AddScoped<ILoanRepository, LoanRepository>();
builder.Services.AddScoped<ILoanPaymentRepository, LoanPaymentRepository>();
builder.Services.AddScoped<ILoginLogRepository, LoginLogRepository>();
builder.Services.AddScoped<INotificationRepository, NotificationRepository>();
builder.Services.AddScoped<IPaymentTemplateRepository, PaymentTemplateRepository>();
builder.Services.AddScoped<ITransactionRepository, TransactionRepository>();
builder.Services.AddScoped<ITransactionCategoryRepository, TransactionCategoryRepository>();
builder.Services.AddScoped<ITransactionFeeRepository, TransactionFeeRepository>();
builder.Services.AddScoped<IUserCredentialRepository, UserCredentialRepository>();
builder.Services.AddGrpcReflection();

var app = builder.Build();

app.MapGrpcService<AccountService>();
app.MapGrpcService<AccountTypeService>();
app.MapGrpcService<AtmService>();
app.MapGrpcService<BranchService>();
app.MapGrpcService<CardService>();
app.MapGrpcService<CustomerAddressService>();
app.MapGrpcService<CustomerService>();
app.MapGrpcService<DepositService>();
app.MapGrpcService<ExchangeRateService>();
app.MapGrpcService<FeeTypeService>();
app.MapGrpcService<LoanService>();
app.MapGrpcService<LoanPaymentService>();
app.MapGrpcService<LoginLogService>();
app.MapGrpcService<NotificationService>();
app.MapGrpcService<TransactionCategoryService>();
app.MapGrpcService<TransactionFeeService>();
app.MapGrpcService<TransactionService>();
app.MapGrpcService<UserCredentialService>();    

app.MapGrpcReflectionService();
app.MapGet("/",
    () =>
        "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");


app.Run(url:"http://localhost:8080");
