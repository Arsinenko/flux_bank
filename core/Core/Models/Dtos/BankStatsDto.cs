namespace Core.Models.Dtos;

public class BankStatsDto
{
    public int CustomerCount { get; set; }
    public int ActiveAccountCount { get; set; }
    public int ActiveCardCount { get; set; }
    public decimal TotalBalance { get; set; }
    public decimal TotalTransactionSum { get; set; }
    public int ActiveAtmCount { get; set; }
    public int InactiveAtmCount { get; set; }
    public int ActiveLoanCount { get; set; }
    public int ActiveDepositCount { get; set; }
}
