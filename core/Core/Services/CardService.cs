using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class CardService(ICardRepository repository, IMapper mapper) : Core.CardService.CardServiceBase
{
    public override async Task<CardModel> Add(AddCardRequest request, ServerCallContext context)
    {
       var card = mapper.Map<Card>(request);
       await repository.AddAsync(card);
       return mapper.Map<CardModel>(card);
    }

    public override Task<GetAllCardsResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        return base.GetAll(request, context);
    }

    public override async Task<Empty> Delete(DeleteCardRequest request, ServerCallContext context)
    {
        var card = await repository.GetByIdAsync(request.CardId);
        if (card == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Card not found"));
        }
        await repository.DeleteAsync(card.CardId);
        return new Empty();
    }

    public override async Task<Empty> Update(UpdateCardRequest request, ServerCallContext context)
    {
        var card = await repository.GetByIdAsync(request.CardId);
        if (card == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Card not found"));
        }

        mapper.Map(request, card);

        await repository.UpdateAsync(card);

        return new Empty();
    }

    public override async Task<CardModel> GetById(GetCardByIdRequest request, ServerCallContext context)
    {
        var card = await repository.GetByIdAsync(request.CardId);
        if (card == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Card not found"));
        }

        return mapper.Map<CardModel>(card);
    }
}
     