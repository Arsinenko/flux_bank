using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class CustomerAddress
{
    public int AddressId { get; set; }

    public required int CustomerId { get; set; }

    public required string Country { get; set; }

    public required string City { get; set; }

    public required string Street { get; set; }

    public required string ZipCode { get; set; }

    public required bool IsPrimary { get; set; }

    public required Customer Customer { get; set; }
}
