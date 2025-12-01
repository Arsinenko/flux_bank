using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class UserCredential
{
    public required int CustomerId { get; set; }

    public required string Username { get; set; }

    public required string PasswordHash { get; set; }

    public required DateTime UpdatedAt { get; set; }

    public required Customer Customer { get; set; }
}
