using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Account
{
    public int AccountId { get; set; }

    public int? CustomerId { get; set; }

    public int? TypeId { get; set; }

    public string Iban { get; set; } = null!;

    public decimal? Balance { get; set; }

    public DateTime? CreatedAt { get; set; }

    public bool? IsActive { get; set; }

    public virtual ICollection<Card> Cards { get; set; } = new List<Card>();

    public virtual Customer? Customer { get; set; }

    public virtual ICollection<Transaction> TransactionSourceAccountNavigations { get; set; } = new List<Transaction>();

    public virtual ICollection<Transaction> TransactionTargetAccountNavigations { get; set; } = new List<Transaction>();

    public virtual AccountType? Type { get; set; }
}
