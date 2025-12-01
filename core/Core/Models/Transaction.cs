using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Transaction
{
    public int TransactionId { get; set; }

    public required int SourceAccount { get; set; }

    public required int TargetAccount { get; set; }

    public required decimal Amount { get; set; }

    public required string Currency { get; set; }

    public required DateTime CreatedAt { get; set; }

    public required string Status { get; set; }

    public required Account SourceAccountNavigation { get; set; }

    public required Account TargetAccountNavigation { get; set; }

    public virtual ICollection<TransactionFee> TransactionFees { get; set; } = new List<TransactionFee>();

    public virtual ICollection<TransactionCategory> Categories { get; set; } = new List<TransactionCategory>();
}
