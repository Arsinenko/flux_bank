using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class LoginLog
{
    public int LogId { get; set; }

    public required int CustomerId { get; set; }

    public required DateTime LoginTime { get; set; }

    public required string IpAddress { get; set; }

    public required string DeviceInfo { get; set; }

    public required Customer Customer { get; set; }
}
