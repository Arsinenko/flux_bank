using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class TransactionCategory
{
    public required int CategoryId { get; set; }

    public string Name { get; set; } = null!;

    public virtual ICollection<Transaction> Transactions { get; set; } = new List<Transaction>();
}
