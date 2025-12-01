using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Branch
{
    public int BranchId { get; set; }

    public string? Name { get; set; }

    public string? City { get; set; }

    public string? Address { get; set; }

    public string? Phone { get; set; }

    public virtual ICollection<Atm> Atms { get; set; } = new List<Atm>();
}
