using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Atm
{
    public int AtmId { get; set; }

    public int? BranchId { get; set; }

    public required string Location { get; set; }

    public required string Status { get; set; }

    public virtual Branch? Branch { get; set; }
}
