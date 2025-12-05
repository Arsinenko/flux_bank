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

    public override async Task<GetAllCardsResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        var cards = await repository.GetAllAsync(request.PageN, request.PageSize);
        return new GetAllCardsResponse
        {
            Cards = { mapper.Map<IEnumerable<CardModel>>(cards) }
        };
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

    public override async Task<Empty> DeleteBulk(DeleteCardBulkRequest request, ServerCallContext context)
    {
        var ids = request.Cards.Select(c => c.CardId).ToList();
        if (ids.Count == 0)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No cards to delete"));
        }
        var cards = await repository.GetByIdsAsync(ids);
        var foundCards = cards.Where(c => c != null).ToList();
        if (foundCards.Count != ids.Count)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Some cards not found"));
        }
        await repository.DeleteRangeAsync(foundCards!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateCardBulkRequest request, ServerCallContext context)
    {
        var cards = request.Cards.Select(mapper.Map<Card>).ToList();
        if (!cards.Any())
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No cards to update"));
        }
        await repository.UpdateRangeAsync(cards);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddCardBulkRequest request, ServerCallContext context)
    {
        var cards = request.Cards.Select(mapper.Map<Card>).ToList();
        try
        {
            await repository.AddRangeAsync(cards);
            return new Empty();
        }
        catch (Exception e)
        {
            throw new RpcException(new Status(StatusCode.Internal, e.Message));
        }
    }
}
     