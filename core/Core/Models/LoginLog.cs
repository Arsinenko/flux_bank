using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class LoginLog
{
    public int LogId { get; set; }

    public int? CustomerId { get; set; }

    public DateTime? LoginTime { get; set; }

    public string? IpAddress { get; set; }

    public string? DeviceInfo { get; set; }

    public virtual Customer? Customer { get; set; }
}
