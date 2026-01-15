using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class LoanService(ILoanRepository loanRepository, IMapper mapper, ICacheService cacheService, IStatsService statsService)
    : Core.LoanService.LoanServiceBase
{
    public override async Task<GetAllLoansResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        var loans = await loanRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllLoansResponse
        {
            Loans = { mapper.Map<IEnumerable<LoanModel>>(loans) }
        };
    }

    public override async Task<LoanModel> Add(AddLoanRequest request, ServerCallContext context)
    {
        var loan = mapper.Map<Loan>(request);

        await loanRepository.AddAsync(loan);
        cacheService.Remove("BankStats");

        return mapper.Map<LoanModel>(loan);
    }

    public override async Task<LoanModel> GetById(GetLoanByIdRequest request, ServerCallContext context)
    {
        var loan = await loanRepository.GetByIdAsync(request.LoanId);

        if (loan == null)
            throw new NotFoundException("Loan not found");

        return mapper.Map<LoanModel>(loan);
    }

    public override async Task<Empty> Update(UpdateLoanRequest request, ServerCallContext context)
    {
        var loan = await loanRepository.GetByIdAsync(request.LoanId);

        if (loan == null)
            throw new NotFoundException("Loan not found");

        mapper.Map(request, loan);

        await loanRepository.UpdateAsync(loan);
        cacheService.Remove("BankStats");

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteLoanRequest request, ServerCallContext context)
    {
        await loanRepository.DeleteAsync(request.LoanId);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteLoanBulkRequest request, ServerCallContext context)
    {
        var ids = request.Loans.Select(l => l.LoanId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No loans to delete");
        }

        var loans = await loanRepository.GetByIdsAsync(ids);
        var foundLoans = loans.Where(l => l != null).ToList();
        if (foundLoans.Count != ids.Count)
        {
            throw new ValidationException("Some loans not found");
        }
        await loanRepository.DeleteRangeAsync(foundLoans!);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<GetAllLoansResponse> GetByCustomer(GetLoansByCustomerRequest request, ServerCallContext context)
    {
        var loans = await loanRepository.FindAsync(l => l.CustomerId == request.CustomerId);
        return new GetAllLoansResponse()
        {
            Loans = { mapper.Map<LoanModel>(loans) }
        };
    }

    public override async Task<GetAllLoansResponse> GetByIds(GetLoanByIdsRequest request, ServerCallContext context)
    {
        var loans = await loanRepository.GetByIdsAsync(request.LoanIds);
        return new GetAllLoansResponse()
        {
            Loans = { mapper.Map<IEnumerable<LoanModel>>(loans) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await loanRepository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountByStatus(GetLoanCountByStatusRequest request, ServerCallContext context)
    {
        if (request.Status == "active")
        {
            var stats = await statsService.GetStatsAsync();
            return new CountResponse()
            {
                Count = stats.ActiveLoanCount
            };
        }
        var count = await loanRepository.GetCountAsync(l => l.Status == request.Status);
        return new CountResponse()
        {
            Count = count
        };
    }
}
