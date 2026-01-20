using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class ExchangeRateService(IExchangeRateRepository exchangeRateRepository, IMapper mapper)
    : Core.ExchangeRateService.ExchangeRateServiceBase
{
    public override async Task<GetAllExchangeRatesResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        var exchangeRates = await exchangeRateRepository.GetAllAsync(request.PageN, request.PageSize, request.OrderBy, request.IsDesc ?? false);

        return new GetAllExchangeRatesResponse
        {
            ExchangeRates = { mapper.Map<IEnumerable<ExchangeRateModel>>(exchangeRates) }
        };
    }

    public override async Task<ExchangeRateModel> Add(AddExchangeRateRequest request, ServerCallContext context)
    {
        var exchangeRate = mapper.Map<ExchangeRate>(request);

        await exchangeRateRepository.AddAsync(exchangeRate);

        return mapper.Map<ExchangeRateModel>(exchangeRate);
    }

    public override async Task<ExchangeRateModel> GetById(GetExchangeRateByIdRequest request, ServerCallContext context)
    {
        var exchangeRate = await exchangeRateRepository.GetByIdAsync(request.RateId);

        if (exchangeRate == null)
            throw new NotFoundException("ExchangeRate not found");

        return mapper.Map<ExchangeRateModel>(exchangeRate);
    }

    public override async Task<Empty> Update(UpdateExchangeRateRequest request, ServerCallContext context)
    {
        var exchangeRate = await exchangeRateRepository.GetByIdAsync(request.RateId);

        if (exchangeRate == null)
            throw new NotFoundException("ExchangeRate not found");

        mapper.Map(request, exchangeRate);

        await exchangeRateRepository.UpdateAsync(exchangeRate);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteExchangeRateRequest request, ServerCallContext context)
    {
        await exchangeRateRepository.DeleteAsync(request.RateId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteExchangeRateBulkRequest request, ServerCallContext context)
    {
        var ids = request.ExchangeRates.Select(e => e.RateId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No exchange to delete!");
        }
        var exchangeRates = await exchangeRateRepository.GetByIdsAsync(ids);
        var foundExchangeRates = exchangeRates.Where(er => er != null).ToList();
        if (foundExchangeRates.Count != ids.Count)
        {
            throw new ValidationException("Some exchange rates not found");
        }
        await exchangeRateRepository.DeleteRangeAsync(foundExchangeRates!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateExchangeRateBulkRequest request, ServerCallContext context)
    {
        var exchangeRates = request.ExchangeRates.Select(mapper.Map<ExchangeRate>).ToList();
        if (!exchangeRates.Any())
        {
            throw new ValidationException("No exchange to update!");
        }
        await exchangeRateRepository.UpdateRangeAsync(exchangeRates);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddExchangeRateBulkRequest request, ServerCallContext context)
    {
        var exchangeRates = request.ExchangeRates.Select(mapper.Map<ExchangeRate>).ToList();
        await exchangeRateRepository.AddRangeAsync(exchangeRates);
        return new Empty();
    }

    public override async Task<GetAllExchangeRatesResponse> GetByBaseCurrency(GetExchangeRateByBaseCurrencyRequest request, ServerCallContext context)
    {
        var rates = await exchangeRateRepository.FindAsync(e => e.BaseCurrency == request.BaseCurrency);
        return new GetAllExchangeRatesResponse()
        {
            ExchangeRates = { mapper.Map<ExchangeRateModel>(rates) }
        };
    }

    public override async Task<GetAllExchangeRatesResponse> GetByIds(GetExchangeRateByIdsRequest request, ServerCallContext context)
    {
        var rates = await exchangeRateRepository.GetByIdsAsync(request.RateIds);
        return new GetAllExchangeRatesResponse()
        {
            ExchangeRates = { mapper.Map<IEnumerable<ExchangeRateModel>>(rates) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await exchangeRateRepository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }

}
