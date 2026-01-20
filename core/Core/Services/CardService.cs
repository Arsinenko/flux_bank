using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class CardService(ICardRepository repository, IMapper mapper, ICacheService cacheService, IStatsService statsService) : Core.CardService.CardServiceBase
{
    public override async Task<CardModel> Add(AddCardRequest request, ServerCallContext context)
    {
        var card = mapper.Map<Card>(request);
        await repository.AddAsync(card);
        cacheService.Remove("BankStats");
        return mapper.Map<CardModel>(card);
    }

    public override async Task<GetAllCardsResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        var cards = await repository.GetAllAsync(request.PageN, request.PageSize, request.OrderBy, request.IsDesc ?? false);
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
            throw new NotFoundException("Card not found");
        }
        await repository.DeleteAsync(card.CardId);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> Update(UpdateCardRequest request, ServerCallContext context)
    {
        var card = await repository.GetByIdAsync(request.CardId);
        if (card == null)
        {
            throw new NotFoundException("Card not found");
        }

        mapper.Map(request, card);

        await repository.UpdateAsync(card);
        cacheService.Remove("BankStats");

        return new Empty();
    }

    public override async Task<CardModel> GetById(GetCardByIdRequest request, ServerCallContext context)
    {
        var card = await repository.GetByIdAsync(request.CardId);
        if (card == null)
        {
            throw new NotFoundException("Card not found");
        }

        return mapper.Map<CardModel>(card);
    }

    public override async Task<Empty> DeleteBulk(DeleteCardBulkRequest request, ServerCallContext context)
    {
        var ids = request.Cards.Select(c => c.CardId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No cards to delete");
        }
        var cards = await repository.GetByIdsAsync(ids);
        var foundCards = cards.Where(c => c != null).ToList();
        if (foundCards.Count != ids.Count)
        {
            throw new ValidationException("Some cards not found");
        }
        await repository.DeleteRangeAsync(foundCards!);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateCardBulkRequest request, ServerCallContext context)
    {
        var cards = request.Cards.Select(mapper.Map<Card>).ToList();
        if (!cards.Any())
        {
            throw new ValidationException("No cards to update");
        }
        await repository.UpdateRangeAsync(cards);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddCardBulkRequest request, ServerCallContext context)
    {
        var cards = request.Cards.Select(mapper.Map<Card>).ToList();
        await repository.AddRangeAsync(cards);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<GetAllCardsResponse> GetByAccount(GetCardsByAccountRequest request, ServerCallContext context)
    {
        var cards = await repository.FindAsync(c => c.AccountId == request.AccountId);
        return new GetAllCardsResponse()
        {
            Cards = { mapper.Map<IEnumerable<CardModel>>(cards) }
        };
    }

    public override async Task<GetAllCardsResponse> GetByIds(GetCardByIdsRequest request, ServerCallContext context)
    {
        var cards = await repository.GetByIdsAsync(request.CardIds);
        return new GetAllCardsResponse()
        {
            Cards = { mapper.Map<IEnumerable<CardModel>>(cards) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await repository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountByStatus(GetCardCountByStatus request, ServerCallContext context)
    {
        if (request.Status == "active")
        {
            var stats = await statsService.GetStatsAsync();
            return new CountResponse()
            {
                Count = stats.ActiveCardCount
            };
        }
        var count = await repository.GetCountAsync(c => c.Status == request.Status);
        return new CountResponse()
        {
            Count = count
        };
    }
}
