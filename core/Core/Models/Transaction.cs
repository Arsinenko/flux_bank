using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Transaction
{
    public int TransactionId { get; set; }

    public int? SourceAccount { get; set; }

    public int? TargetAccount { get; set; }

    public decimal Amount { get; set; }

    public string Currency { get; set; } = null!;

    public DateTime? CreatedAt { get; set; }

    public string? Status { get; set; }

    public virtual Account? SourceAccountNavigation { get; set; }

    public virtual Account? TargetAccountNavigation { get; set; }

    public virtual ICollection<TransactionFee> TransactionFees { get; set; } = new List<TransactionFee>();

    public virtual ICollection<TransactionCategory> Categories { get; set; } = new List<TransactionCategory>();
}
