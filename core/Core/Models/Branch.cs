using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Branch
{
    public int BranchId { get; set; }

    public required string Name { get; set; }

    public required string City { get; set; }

    public required string Address { get; set; }

    public required string Phone { get; set; }

    public virtual ICollection<Atm> Atms { get; set; } = new List<Atm>();
}
