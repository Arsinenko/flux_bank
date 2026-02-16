using System;
using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace Core.Models
{
    /// <inheritdoc />
    public partial class Initial : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "account_types",
                columns: table => new
                {
                    type_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    name = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    description = table.Column<string>(type: "text", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("account_types_pkey", x => x.type_id);
                });

            migrationBuilder.CreateTable(
                name: "branches",
                columns: table => new
                {
                    branch_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    name = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    city = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    address = table.Column<string>(type: "character varying(255)", maxLength: 255, nullable: false),
                    phone = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("branches_pkey", x => x.branch_id);
                });

            migrationBuilder.CreateTable(
                name: "customers",
                columns: table => new
                {
                    customer_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    first_name = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    last_name = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    email = table.Column<string>(type: "character varying(255)", maxLength: 255, nullable: false),
                    phone = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false),
                    birth_date = table.Column<DateOnly>(type: "date", nullable: false),
                    created_at = table.Column<DateTime>(type: "timestamp without time zone", nullable: false, defaultValueSql: "now()")
                },
                constraints: table =>
                {
                    table.PrimaryKey("customers_pkey", x => x.customer_id);
                });

            migrationBuilder.CreateTable(
                name: "exchange_rates",
                columns: table => new
                {
                    rate_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    base_currency = table.Column<string>(type: "character varying(10)", maxLength: 10, nullable: false),
                    target_currency = table.Column<string>(type: "character varying(10)", maxLength: 10, nullable: false),
                    rate = table.Column<decimal>(type: "numeric(18,6)", precision: 18, scale: 6, nullable: false),
                    updated_at = table.Column<DateTime>(type: "timestamp without time zone", nullable: true, defaultValueSql: "now()")
                },
                constraints: table =>
                {
                    table.PrimaryKey("exchange_rates_pkey", x => x.rate_id);
                });

            migrationBuilder.CreateTable(
                name: "fee_types",
                columns: table => new
                {
                    fee_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    name = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    description = table.Column<string>(type: "text", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("fee_types_pkey", x => x.fee_id);
                });

            migrationBuilder.CreateTable(
                name: "transaction_categories",
                columns: table => new
                {
                    category_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    name = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("transaction_categories_pkey", x => x.category_id);
                });

            migrationBuilder.CreateTable(
                name: "atms",
                columns: table => new
                {
                    atm_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    branch_id = table.Column<int>(type: "integer", nullable: true),
                    location = table.Column<string>(type: "character varying(255)", maxLength: 255, nullable: false),
                    status = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("atms_pkey", x => x.atm_id);
                    table.ForeignKey(
                        name: "atms_branch_id_fkey",
                        column: x => x.branch_id,
                        principalTable: "branches",
                        principalColumn: "branch_id");
                });

            migrationBuilder.CreateTable(
                name: "accounts",
                columns: table => new
                {
                    account_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    type_id = table.Column<int>(type: "integer", nullable: false),
                    iban = table.Column<string>(type: "character varying(34)", maxLength: 34, nullable: false),
                    balance = table.Column<decimal>(type: "numeric(18,2)", precision: 18, scale: 2, nullable: true, defaultValue: 0m),
                    created_at = table.Column<DateTime>(type: "timestamp without time zone", nullable: true, defaultValueSql: "now()"),
                    is_active = table.Column<bool>(type: "boolean", nullable: false, defaultValue: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("accounts_pkey", x => x.account_id);
                    table.ForeignKey(
                        name: "accounts_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "accounts_type_id_fkey",
                        column: x => x.type_id,
                        principalTable: "account_types",
                        principalColumn: "type_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "customer_addresses",
                columns: table => new
                {
                    address_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    country = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    city = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    street = table.Column<string>(type: "character varying(255)", maxLength: 255, nullable: false),
                    zip_code = table.Column<string>(type: "character varying(20)", maxLength: 20, nullable: false),
                    is_primary = table.Column<bool>(type: "boolean", nullable: false, defaultValue: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("customer_addresses_pkey", x => x.address_id);
                    table.ForeignKey(
                        name: "customer_addresses_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "deposits",
                columns: table => new
                {
                    deposit_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    amount = table.Column<decimal>(type: "numeric(18,2)", precision: 18, scale: 2, nullable: false),
                    interest_rate = table.Column<decimal>(type: "numeric(5,2)", precision: 5, scale: 2, nullable: false),
                    start_date = table.Column<DateOnly>(type: "date", nullable: false),
                    end_date = table.Column<DateOnly>(type: "date", nullable: true),
                    status = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("deposits_pkey", x => x.deposit_id);
                    table.ForeignKey(
                        name: "deposits_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "loans",
                columns: table => new
                {
                    loan_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    principal = table.Column<decimal>(type: "numeric(18,2)", precision: 18, scale: 2, nullable: false),
                    interest_rate = table.Column<decimal>(type: "numeric(5,2)", precision: 5, scale: 2, nullable: false),
                    start_date = table.Column<DateOnly>(type: "date", nullable: false),
                    end_date = table.Column<DateOnly>(type: "date", nullable: true),
                    status = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("loans_pkey", x => x.loan_id);
                    table.ForeignKey(
                        name: "loans_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "login_logs",
                columns: table => new
                {
                    log_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    login_time = table.Column<DateTime>(type: "timestamp without time zone", nullable: false, defaultValueSql: "now()"),
                    ip_address = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false),
                    device_info = table.Column<string>(type: "text", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("login_logs_pkey", x => x.log_id);
                    table.ForeignKey(
                        name: "login_logs_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "notifications",
                columns: table => new
                {
                    notification_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    message = table.Column<string>(type: "text", nullable: false),
                    created_at = table.Column<DateTime>(type: "timestamp without time zone", nullable: false, defaultValueSql: "now()"),
                    is_read = table.Column<bool>(type: "boolean", nullable: false, defaultValue: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("notifications_pkey", x => x.notification_id);
                    table.ForeignKey(
                        name: "notifications_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "payment_templates",
                columns: table => new
                {
                    template_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    name = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    target_iban = table.Column<string>(type: "character varying(34)", maxLength: 34, nullable: false),
                    default_amount = table.Column<decimal>(type: "numeric(18,2)", precision: 18, scale: 2, nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("payment_templates_pkey", x => x.template_id);
                    table.ForeignKey(
                        name: "payment_templates_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "user_credentials",
                columns: table => new
                {
                    customer_id = table.Column<int>(type: "integer", nullable: false),
                    username = table.Column<string>(type: "character varying(100)", maxLength: 100, nullable: false),
                    password_hash = table.Column<string>(type: "text", nullable: false),
                    updated_at = table.Column<DateTime>(type: "timestamp without time zone", nullable: false, defaultValueSql: "now()")
                },
                constraints: table =>
                {
                    table.PrimaryKey("user_credentials_pkey", x => x.customer_id);
                    table.ForeignKey(
                        name: "user_credentials_customer_id_fkey",
                        column: x => x.customer_id,
                        principalTable: "customers",
                        principalColumn: "customer_id");
                });

            migrationBuilder.CreateTable(
                name: "cards",
                columns: table => new
                {
                    card_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    account_id = table.Column<int>(type: "integer", nullable: false),
                    card_number = table.Column<string>(type: "character varying(20)", maxLength: 20, nullable: false),
                    cvv = table.Column<string>(type: "character varying(4)", maxLength: 4, nullable: false),
                    expiry_date = table.Column<DateOnly>(type: "date", nullable: true),
                    status = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false, defaultValueSql: "'active'::character varying")
                },
                constraints: table =>
                {
                    table.PrimaryKey("cards_pkey", x => x.card_id);
                    table.ForeignKey(
                        name: "cards_account_id_fkey",
                        column: x => x.account_id,
                        principalTable: "accounts",
                        principalColumn: "account_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "transactions",
                columns: table => new
                {
                    transaction_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    source_account = table.Column<int>(type: "integer", nullable: false),
                    target_account = table.Column<int>(type: "integer", nullable: false),
                    amount = table.Column<decimal>(type: "numeric(18,2)", precision: 18, scale: 2, nullable: false),
                    currency = table.Column<string>(type: "character varying(10)", maxLength: 10, nullable: false),
                    created_at = table.Column<DateTime>(type: "timestamp without time zone", nullable: false, defaultValueSql: "now()"),
                    status = table.Column<string>(type: "character varying(50)", maxLength: 50, nullable: false, defaultValueSql: "'completed'::character varying")
                },
                constraints: table =>
                {
                    table.PrimaryKey("transactions_pkey", x => x.transaction_id);
                    table.ForeignKey(
                        name: "transactions_source_account_fkey",
                        column: x => x.source_account,
                        principalTable: "accounts",
                        principalColumn: "account_id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "transactions_target_account_fkey",
                        column: x => x.target_account,
                        principalTable: "accounts",
                        principalColumn: "account_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "loan_payments",
                columns: table => new
                {
                    payment_id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    loan_id = table.Column<int>(type: "integer", nullable: false),
                    amount = table.Column<decimal>(type: "numeric(18,2)", precision: 18, scale: 2, nullable: false),
                    payment_date = table.Column<DateOnly>(type: "date", nullable: false),
                    is_paid = table.Column<bool>(type: "boolean", nullable: false, defaultValue: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("loan_payments_pkey", x => x.payment_id);
                    table.ForeignKey(
                        name: "loan_payments_loan_id_fkey",
                        column: x => x.loan_id,
                        principalTable: "loans",
                        principalColumn: "loan_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "transaction_category_map",
                columns: table => new
                {
                    transaction_id = table.Column<int>(type: "integer", nullable: false),
                    category_id = table.Column<int>(type: "integer", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("transaction_category_map_pkey", x => new { x.transaction_id, x.category_id });
                    table.ForeignKey(
                        name: "transaction_category_map_category_id_fkey",
                        column: x => x.category_id,
                        principalTable: "transaction_categories",
                        principalColumn: "category_id");
                    table.ForeignKey(
                        name: "transaction_category_map_transaction_id_fkey",
                        column: x => x.transaction_id,
                        principalTable: "transactions",
                        principalColumn: "transaction_id");
                });

            migrationBuilder.CreateTable(
                name: "transaction_fees",
                columns: table => new
                {
                    id = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    transaction_id = table.Column<int>(type: "integer", nullable: false),
                    fee_id = table.Column<int>(type: "integer", nullable: false),
                    amount = table.Column<decimal>(type: "numeric(18,2)", precision: 18, scale: 2, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("transaction_fees_pkey", x => x.id);
                    table.ForeignKey(
                        name: "transaction_fees_fee_id_fkey",
                        column: x => x.fee_id,
                        principalTable: "fee_types",
                        principalColumn: "fee_id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "transaction_fees_transaction_id_fkey",
                        column: x => x.transaction_id,
                        principalTable: "transactions",
                        principalColumn: "transaction_id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "account_types_name_key",
                table: "account_types",
                column: "name",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "accounts_iban_key",
                table: "accounts",
                column: "iban",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_accounts_customer_id",
                table: "accounts",
                column: "customer_id");

            migrationBuilder.CreateIndex(
                name: "IX_accounts_type_id",
                table: "accounts",
                column: "type_id");

            migrationBuilder.CreateIndex(
                name: "IX_atms_branch_id",
                table: "atms",
                column: "branch_id");

            migrationBuilder.CreateIndex(
                name: "cards_card_number_key",
                table: "cards",
                column: "card_number",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_cards_account_id",
                table: "cards",
                column: "account_id");

            migrationBuilder.CreateIndex(
                name: "IX_customer_addresses_customer_id",
                table: "customer_addresses",
                column: "customer_id");

            migrationBuilder.CreateIndex(
                name: "customers_email_key",
                table: "customers",
                column: "email",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_deposits_customer_id",
                table: "deposits",
                column: "customer_id");

            migrationBuilder.CreateIndex(
                name: "IX_loan_payments_loan_id",
                table: "loan_payments",
                column: "loan_id");

            migrationBuilder.CreateIndex(
                name: "IX_loans_customer_id",
                table: "loans",
                column: "customer_id");

            migrationBuilder.CreateIndex(
                name: "IX_login_logs_customer_id",
                table: "login_logs",
                column: "customer_id");

            migrationBuilder.CreateIndex(
                name: "IX_notifications_customer_id",
                table: "notifications",
                column: "customer_id");

            migrationBuilder.CreateIndex(
                name: "IX_payment_templates_customer_id",
                table: "payment_templates",
                column: "customer_id");

            migrationBuilder.CreateIndex(
                name: "transaction_categories_name_key",
                table: "transaction_categories",
                column: "name",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_transaction_category_map_category_id",
                table: "transaction_category_map",
                column: "category_id");

            migrationBuilder.CreateIndex(
                name: "IX_transaction_fees_fee_id",
                table: "transaction_fees",
                column: "fee_id");

            migrationBuilder.CreateIndex(
                name: "IX_transaction_fees_transaction_id",
                table: "transaction_fees",
                column: "transaction_id");

            migrationBuilder.CreateIndex(
                name: "IX_transactions_source_account",
                table: "transactions",
                column: "source_account");

            migrationBuilder.CreateIndex(
                name: "IX_transactions_target_account",
                table: "transactions",
                column: "target_account");

            migrationBuilder.CreateIndex(
                name: "user_credentials_username_key",
                table: "user_credentials",
                column: "username",
                unique: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "atms");

            migrationBuilder.DropTable(
                name: "cards");

            migrationBuilder.DropTable(
                name: "customer_addresses");

            migrationBuilder.DropTable(
                name: "deposits");

            migrationBuilder.DropTable(
                name: "exchange_rates");

            migrationBuilder.DropTable(
                name: "loan_payments");

            migrationBuilder.DropTable(
                name: "login_logs");

            migrationBuilder.DropTable(
                name: "notifications");

            migrationBuilder.DropTable(
                name: "payment_templates");

            migrationBuilder.DropTable(
                name: "transaction_category_map");

            migrationBuilder.DropTable(
                name: "transaction_fees");

            migrationBuilder.DropTable(
                name: "user_credentials");

            migrationBuilder.DropTable(
                name: "branches");

            migrationBuilder.DropTable(
                name: "loans");

            migrationBuilder.DropTable(
                name: "transaction_categories");

            migrationBuilder.DropTable(
                name: "fee_types");

            migrationBuilder.DropTable(
                name: "transactions");

            migrationBuilder.DropTable(
                name: "accounts");

            migrationBuilder.DropTable(
                name: "customers");

            migrationBuilder.DropTable(
                name: "account_types");
        }
    }
}
