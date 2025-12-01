using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Context;

public partial class MyDbContext : DbContext
{
    public MyDbContext()
    {
    }

    public MyDbContext(DbContextOptions<MyDbContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Account> Accounts { get; set; }

    public virtual DbSet<AccountType> AccountTypes { get; set; }

    public virtual DbSet<Atm> Atms { get; set; }

    public virtual DbSet<Branch> Branches { get; set; }

    public virtual DbSet<Card> Cards { get; set; }

    public virtual DbSet<Customer> Customers { get; set; }

    public virtual DbSet<CustomerAddress> CustomerAddresses { get; set; }

    public virtual DbSet<Deposit> Deposits { get; set; }

    public virtual DbSet<ExchangeRate> ExchangeRates { get; set; }

    public virtual DbSet<FeeType> FeeTypes { get; set; }

    public virtual DbSet<Loan> Loans { get; set; }

    public virtual DbSet<LoanPayment> LoanPayments { get; set; }

    public virtual DbSet<LoginLog> LoginLogs { get; set; }

    public virtual DbSet<Notification> Notifications { get; set; }

    public virtual DbSet<PaymentTemplate> PaymentTemplates { get; set; }

    public virtual DbSet<Transaction> Transactions { get; set; }

    public virtual DbSet<TransactionCategory> TransactionCategories { get; set; }

    public virtual DbSet<TransactionFee> TransactionFees { get; set; }

    public virtual DbSet<UserCredential> UserCredentials { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {
        
    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Account>(entity =>
        {
            entity.HasKey(e => e.AccountId).HasName("accounts_pkey");

            entity.ToTable("accounts");

            entity.HasIndex(e => e.Iban, "accounts_iban_key").IsUnique();

            entity.Property(e => e.AccountId).HasColumnName("account_id");
            entity.Property(e => e.Balance)
                .HasPrecision(18, 2)
                .HasDefaultValue(0m)
                .HasColumnName("balance");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.Iban)
                .HasMaxLength(34)
                .HasColumnName("iban");
            entity.Property(e => e.IsActive)
                .HasDefaultValue(true)
                .HasColumnName("is_active");
            entity.Property(e => e.TypeId).HasColumnName("type_id");

            entity.HasOne(d => d.Customer).WithMany(p => p.Accounts)
                .HasForeignKey(d => d.CustomerId)
                .HasConstraintName("accounts_customer_id_fkey");

            entity.HasOne(d => d.Type).WithMany(p => p.Accounts)
                .HasForeignKey(d => d.TypeId)
                .HasConstraintName("accounts_type_id_fkey");
        });

        modelBuilder.Entity<AccountType>(entity =>
        {
            entity.HasKey(e => e.TypeId).HasName("account_types_pkey");

            entity.ToTable("account_types");

            entity.HasIndex(e => e.Name, "account_types_name_key").IsUnique();

            entity.Property(e => e.TypeId).HasColumnName("type_id");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Name)
                .HasMaxLength(100)
                .HasColumnName("name");
        });

        modelBuilder.Entity<Atm>(entity =>
        {
            entity.HasKey(e => e.AtmId).HasName("atms_pkey");

            entity.ToTable("atms");

            entity.Property(e => e.AtmId).HasColumnName("atm_id");
            entity.Property(e => e.BranchId).HasColumnName("branch_id");
            entity.Property(e => e.Location)
                .HasMaxLength(255)
                .HasColumnName("location");
            entity.Property(e => e.Status)
                .HasMaxLength(50)
                .HasColumnName("status");

            entity.HasOne(d => d.Branch).WithMany(p => p.Atms)
                .HasForeignKey(d => d.BranchId)
                .HasConstraintName("atms_branch_id_fkey");
        });

        modelBuilder.Entity<Branch>(entity =>
        {
            entity.HasKey(e => e.BranchId).HasName("branches_pkey");

            entity.ToTable("branches");

            entity.Property(e => e.BranchId).HasColumnName("branch_id");
            entity.Property(e => e.Address)
                .HasMaxLength(255)
                .HasColumnName("address");
            entity.Property(e => e.City)
                .HasMaxLength(100)
                .HasColumnName("city");
            entity.Property(e => e.Name)
                .HasMaxLength(100)
                .HasColumnName("name");
            entity.Property(e => e.Phone)
                .HasMaxLength(50)
                .HasColumnName("phone");
        });

        modelBuilder.Entity<Card>(entity =>
        {
            entity.HasKey(e => e.CardId).HasName("cards_pkey");

            entity.ToTable("cards");

            entity.HasIndex(e => e.CardNumber, "cards_card_number_key").IsUnique();

            entity.Property(e => e.CardId).HasColumnName("card_id");
            entity.Property(e => e.AccountId).HasColumnName("account_id");
            entity.Property(e => e.CardNumber)
                .HasMaxLength(16)
                .HasColumnName("card_number");
            entity.Property(e => e.Cvv)
                .HasMaxLength(4)
                .HasColumnName("cvv");
            entity.Property(e => e.ExpiryDate).HasColumnName("expiry_date");
            entity.Property(e => e.Status)
                .HasMaxLength(50)
                .HasDefaultValueSql("'active'::character varying")
                .HasColumnName("status");

            entity.HasOne(d => d.Account).WithMany(p => p.Cards)
                .HasForeignKey(d => d.AccountId)
                .HasConstraintName("cards_account_id_fkey");
        });

        modelBuilder.Entity<Customer>(entity =>
        {
            entity.HasKey(e => e.CustomerId).HasName("customers_pkey");

            entity.ToTable("customers");

            entity.HasIndex(e => e.Email, "customers_email_key").IsUnique();

            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.BirthDate).HasColumnName("birth_date");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.Email)
                .HasMaxLength(255)
                .HasColumnName("email");
            entity.Property(e => e.FirstName)
                .HasMaxLength(100)
                .HasColumnName("first_name");
            entity.Property(e => e.LastName)
                .HasMaxLength(100)
                .HasColumnName("last_name");
            entity.Property(e => e.Phone)
                .HasMaxLength(50)
                .HasColumnName("phone");
        });

        modelBuilder.Entity<CustomerAddress>(entity =>
        {
            entity.HasKey(e => e.AddressId).HasName("customer_addresses_pkey");

            entity.ToTable("customer_addresses");

            entity.Property(e => e.AddressId).HasColumnName("address_id");
            entity.Property(e => e.City)
                .HasMaxLength(100)
                .HasColumnName("city");
            entity.Property(e => e.Country)
                .HasMaxLength(100)
                .HasColumnName("country");
            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.IsPrimary)
                .HasDefaultValue(false)
                .HasColumnName("is_primary");
            entity.Property(e => e.Street)
                .HasMaxLength(255)
                .HasColumnName("street");
            entity.Property(e => e.ZipCode)
                .HasMaxLength(20)
                .HasColumnName("zip_code");

            entity.HasOne(d => d.Customer).WithMany(p => p.CustomerAddresses)
                .HasForeignKey(d => d.CustomerId)
                .HasConstraintName("customer_addresses_customer_id_fkey");
        });

        modelBuilder.Entity<Deposit>(entity =>
        {
            entity.HasKey(e => e.DepositId).HasName("deposits_pkey");

            entity.ToTable("deposits");

            entity.Property(e => e.DepositId).HasColumnName("deposit_id");
            entity.Property(e => e.Amount)
                .HasPrecision(18, 2)
                .HasColumnName("amount");
            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.EndDate).HasColumnName("end_date");
            entity.Property(e => e.InterestRate)
                .HasPrecision(5, 2)
                .HasColumnName("interest_rate");
            entity.Property(e => e.StartDate).HasColumnName("start_date");
            entity.Property(e => e.Status)
                .HasMaxLength(50)
                .HasColumnName("status");

            entity.HasOne(d => d.Customer).WithMany(p => p.Deposits)
                .HasForeignKey(d => d.CustomerId)
                .HasConstraintName("deposits_customer_id_fkey");
        });

        modelBuilder.Entity<ExchangeRate>(entity =>
        {
            entity.HasKey(e => e.RateId).HasName("exchange_rates_pkey");

            entity.ToTable("exchange_rates");

            entity.Property(e => e.RateId).HasColumnName("rate_id");
            entity.Property(e => e.BaseCurrency)
                .HasMaxLength(10)
                .HasColumnName("base_currency");
            entity.Property(e => e.Rate)
                .HasPrecision(18, 6)
                .HasColumnName("rate");
            entity.Property(e => e.TargetCurrency)
                .HasMaxLength(10)
                .HasColumnName("target_currency");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");
        });

        modelBuilder.Entity<FeeType>(entity =>
        {
            entity.HasKey(e => e.FeeId).HasName("fee_types_pkey");

            entity.ToTable("fee_types");

            entity.Property(e => e.FeeId).HasColumnName("fee_id");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Name)
                .HasMaxLength(100)
                .HasColumnName("name");
        });

        modelBuilder.Entity<Loan>(entity =>
        {
            entity.HasKey(e => e.LoanId).HasName("loans_pkey");

            entity.ToTable("loans");

            entity.Property(e => e.LoanId).HasColumnName("loan_id");
            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.EndDate).HasColumnName("end_date");
            entity.Property(e => e.InterestRate)
                .HasPrecision(5, 2)
                .HasColumnName("interest_rate");
            entity.Property(e => e.Principal)
                .HasPrecision(18, 2)
                .HasColumnName("principal");
            entity.Property(e => e.StartDate).HasColumnName("start_date");
            entity.Property(e => e.Status)
                .HasMaxLength(50)
                .HasColumnName("status");

            entity.HasOne(d => d.Customer).WithMany(p => p.Loans)
                .HasForeignKey(d => d.CustomerId)
                .HasConstraintName("loans_customer_id_fkey");
        });

        modelBuilder.Entity<LoanPayment>(entity =>
        {
            entity.HasKey(e => e.PaymentId).HasName("loan_payments_pkey");

            entity.ToTable("loan_payments");

            entity.Property(e => e.PaymentId).HasColumnName("payment_id");
            entity.Property(e => e.Amount)
                .HasPrecision(18, 2)
                .HasColumnName("amount");
            entity.Property(e => e.IsPaid)
                .HasDefaultValue(false)
                .HasColumnName("is_paid");
            entity.Property(e => e.LoanId).HasColumnName("loan_id");
            entity.Property(e => e.PaymentDate).HasColumnName("payment_date");

            entity.HasOne(d => d.Loan).WithMany(p => p.LoanPayments)
                .HasForeignKey(d => d.LoanId)
                .HasConstraintName("loan_payments_loan_id_fkey");
        });

        modelBuilder.Entity<LoginLog>(entity =>
        {
            entity.HasKey(e => e.LogId).HasName("login_logs_pkey");

            entity.ToTable("login_logs");

            entity.Property(e => e.LogId).HasColumnName("log_id");
            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.DeviceInfo).HasColumnName("device_info");
            entity.Property(e => e.IpAddress)
                .HasMaxLength(50)
                .HasColumnName("ip_address");
            entity.Property(e => e.LoginTime)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("login_time");

            entity.HasOne(d => d.Customer).WithMany(p => p.LoginLogs)
                .HasForeignKey(d => d.CustomerId)
                .HasConstraintName("login_logs_customer_id_fkey");
        });

        modelBuilder.Entity<Notification>(entity =>
        {
            entity.HasKey(e => e.NotificationId).HasName("notifications_pkey");

            entity.ToTable("notifications");

            entity.Property(e => e.NotificationId).HasColumnName("notification_id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.IsRead)
                .HasDefaultValue(false)
                .HasColumnName("is_read");
            entity.Property(e => e.Message).HasColumnName("message");

            entity.HasOne(d => d.Customer).WithMany(p => p.Notifications)
                .HasForeignKey(d => d.CustomerId)
                .HasConstraintName("notifications_customer_id_fkey");
        });

        modelBuilder.Entity<PaymentTemplate>(entity =>
        {
            entity.HasKey(e => e.TemplateId).HasName("payment_templates_pkey");

            entity.ToTable("payment_templates");

            entity.Property(e => e.TemplateId).HasColumnName("template_id");
            entity.Property(e => e.CustomerId).HasColumnName("customer_id");
            entity.Property(e => e.DefaultAmount)
                .HasPrecision(18, 2)
                .HasColumnName("default_amount");
            entity.Property(e => e.Name)
                .HasMaxLength(100)
                .HasColumnName("name");
            entity.Property(e => e.TargetIban)
                .HasMaxLength(34)
                .HasColumnName("target_iban");

            entity.HasOne(d => d.Customer).WithMany(p => p.PaymentTemplates)
                .HasForeignKey(d => d.CustomerId)
                .HasConstraintName("payment_templates_customer_id_fkey");
        });

        modelBuilder.Entity<Transaction>(entity =>
        {
            entity.HasKey(e => e.TransactionId).HasName("transactions_pkey");

            entity.ToTable("transactions");

            entity.Property(e => e.TransactionId).HasColumnName("transaction_id");
            entity.Property(e => e.Amount)
                .HasPrecision(18, 2)
                .HasColumnName("amount");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.Currency)
                .HasMaxLength(10)
                .HasColumnName("currency");
            entity.Property(e => e.SourceAccount).HasColumnName("source_account");
            entity.Property(e => e.Status)
                .HasMaxLength(50)
                .HasDefaultValueSql("'completed'::character varying")
                .HasColumnName("status");
            entity.Property(e => e.TargetAccount).HasColumnName("target_account");

            entity.HasOne(d => d.SourceAccountNavigation).WithMany(p => p.TransactionSourceAccountNavigations)
                .HasForeignKey(d => d.SourceAccount)
                .HasConstraintName("transactions_source_account_fkey");

            entity.HasOne(d => d.TargetAccountNavigation).WithMany(p => p.TransactionTargetAccountNavigations)
                .HasForeignKey(d => d.TargetAccount)
                .HasConstraintName("transactions_target_account_fkey");

            entity.HasMany(d => d.Categories).WithMany(p => p.Transactions)
                .UsingEntity<Dictionary<string, object>>(
                    "TransactionCategoryMap",
                    r => r.HasOne<TransactionCategory>().WithMany()
                        .HasForeignKey("CategoryId")
                        .OnDelete(DeleteBehavior.ClientSetNull)
                        .HasConstraintName("transaction_category_map_category_id_fkey"),
                    l => l.HasOne<Transaction>().WithMany()
                        .HasForeignKey("TransactionId")
                        .OnDelete(DeleteBehavior.ClientSetNull)
                        .HasConstraintName("transaction_category_map_transaction_id_fkey"),
                    j =>
                    {
                        j.HasKey("TransactionId", "CategoryId").HasName("transaction_category_map_pkey");
                        j.ToTable("transaction_category_map");
                        j.IndexerProperty<int>("TransactionId").HasColumnName("transaction_id");
                        j.IndexerProperty<int>("CategoryId").HasColumnName("category_id");
                    });
        });

        modelBuilder.Entity<TransactionCategory>(entity =>
        {
            entity.HasKey(e => e.CategoryId).HasName("transaction_categories_pkey");

            entity.ToTable("transaction_categories");

            entity.HasIndex(e => e.Name, "transaction_categories_name_key").IsUnique();

            entity.Property(e => e.CategoryId).HasColumnName("category_id");
            entity.Property(e => e.Name)
                .HasMaxLength(100)
                .HasColumnName("name");
        });

        modelBuilder.Entity<TransactionFee>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("transaction_fees_pkey");

            entity.ToTable("transaction_fees");

            entity.Property(e => e.Id).HasColumnName("id");
            entity.Property(e => e.Amount)
                .HasPrecision(18, 2)
                .HasColumnName("amount");
            entity.Property(e => e.FeeId).HasColumnName("fee_id");
            entity.Property(e => e.TransactionId).HasColumnName("transaction_id");

            entity.HasOne(d => d.Fee).WithMany(p => p.TransactionFees)
                .HasForeignKey(d => d.FeeId)
                .HasConstraintName("transaction_fees_fee_id_fkey");

            entity.HasOne(d => d.Transaction).WithMany(p => p.TransactionFees)
                .HasForeignKey(d => d.TransactionId)
                .HasConstraintName("transaction_fees_transaction_id_fkey");
        });

        modelBuilder.Entity<UserCredential>(entity =>
        {
            entity.HasKey(e => e.CustomerId).HasName("user_credentials_pkey");

            entity.ToTable("user_credentials");

            entity.HasIndex(e => e.Username, "user_credentials_username_key").IsUnique();

            entity.Property(e => e.CustomerId)
                .ValueGeneratedNever()
                .HasColumnName("customer_id");
            entity.Property(e => e.PasswordHash).HasColumnName("password_hash");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");
            entity.Property(e => e.Username)
                .HasMaxLength(100)
                .HasColumnName("username");

            entity.HasOne(d => d.Customer).WithOne(p => p.UserCredential)
                .HasForeignKey<UserCredential>(d => d.CustomerId)
                .OnDelete(DeleteBehavior.ClientSetNull)
                .HasConstraintName("user_credentials_customer_id_fkey");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
}
