using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class AccountType
{
    public int TypeId { get; set; }

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    public virtual ICollection<Account> Accounts { get; set; } = new List<Account>();
}
