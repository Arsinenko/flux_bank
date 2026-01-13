using Core.Interfaces;
using Core.Models.Dtos;

namespace Core.Services;

public interface IStatsService
{
    Task<BankStatsDto> GetStatsAsync();
}

public class StatsService : IStatsService
{
    private readonly ICustomerRepository _customerRepository;
    private readonly IAccountRepository _accountRepository;
    private readonly ICardRepository _cardRepository;
    private readonly ITransactionRepository _transactionRepository;
    private readonly IAtmRepository _atmRepository;
    private readonly ILoanRepository _loanRepository;
    private readonly IDepositRepository _depositRepository;
    private readonly ICacheService _cacheService;

    private const string StatsCacheKey = "BankStats";

    public StatsService(
        ICustomerRepository customerRepository,
        IAccountRepository accountRepository,
        ICardRepository cardRepository,
        ITransactionRepository transactionRepository,
        IAtmRepository atmRepository,
        ILoanRepository loanRepository,
        IDepositRepository depositRepository,
        ICacheService cacheService)
    {
        _customerRepository = customerRepository;
        _accountRepository = accountRepository;
        _cardRepository = cardRepository;
        _transactionRepository = transactionRepository;
        _atmRepository = atmRepository;
        _loanRepository = loanRepository;
        _depositRepository = depositRepository;
        _cacheService = cacheService;
    }

    public async Task<BankStatsDto> GetStatsAsync()
    {
        var cachedStats = _cacheService.Get<BankStatsDto>(StatsCacheKey);
        if (cachedStats != null)
        {
            return cachedStats;
        }

        var stats = new BankStatsDto
        {
            CustomerCount = await _customerRepository.GetCountAsync(),
            ActiveAccountCount = await _accountRepository.GetCountByStatusAsync(true),
            ActiveCardCount = await _cardRepository.GetCountByStatusAsync("active"),
            TotalBalance = await _accountRepository.GetTotalBalanceAsync() ?? 0,
            TotalTransactionSum = await _transactionRepository.GetTotalAmountAsync(),
            ActiveAtmCount = await _atmRepository.GetCountByStatusAsync("active"),
            InactiveAtmCount = await _atmRepository.GetCountByStatusAsync("inactive"),
            ActiveLoanCount = await _loanRepository.GetCountByStatus("active"),
            ActiveDepositCount = await _depositRepository.GetCountByStatus("active"),
        };

        _cacheService.Set(StatsCacheKey, stats, TimeSpan.FromMinutes(5));
        return stats;
    }
}
