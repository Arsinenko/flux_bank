using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class CustomerAddress
{
    public int AddressId { get; set; }

    public int? CustomerId { get; set; }

    public string? Country { get; set; }

    public string? City { get; set; }

    public string? Street { get; set; }

    public string? ZipCode { get; set; }

    public bool? IsPrimary { get; set; }

    public virtual Customer? Customer { get; set; }
}
