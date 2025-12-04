using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class ExchangeRateService(IExchangeRateRepository exchangeRateRepository, IMapper mapper)
    : Core.ExchangeRateService.ExchangeRateServiceBase
{
    public override async Task<GetAllExchangeRatesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var exchangeRates = await exchangeRateRepository.GetAllAsync(request.PageN, request.PageSize);

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
            throw new RpcException(new Status(StatusCode.NotFound, "ExchangeRate not found"));

        return mapper.Map<ExchangeRateModel>(exchangeRate);
    }

    public override async Task<Empty> Update(UpdateExchangeRateRequest request, ServerCallContext context)
    {
        var exchangeRate = await exchangeRateRepository.GetByIdAsync(request.RateId);

        if (exchangeRate == null)
            throw new RpcException(new Status(StatusCode.NotFound, "ExchangeRate not found"));

        mapper.Map(request, exchangeRate);

        await exchangeRateRepository.UpdateAsync(exchangeRate);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteExchangeRateRequest request, ServerCallContext context)
    {
        await exchangeRateRepository.DeleteAsync(request.RateId);
        return new Empty();
    }
}
