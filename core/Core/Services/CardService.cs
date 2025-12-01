using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class CardService(ICardRepository repository) : Core.CardService.CardServiceBase
{
    public override async Task<CardModel> Add(AddCardRequest request, ServerCallContext context)
    {
        var card = new Card
        {
            AccountId = request.AccountId,
            CardNumber = request.CardNumber,
            Cvv = request.Cvv,
            Status = request.Status
        };        
        await repository.AddAsync(card);
        return new CardModel
        {
            CardId = card.CardId,
            AccountId = card.AccountId,
            CardNumber = card.CardNumber,
            Cvv = card.Cvv,
            Status = card.Status
        };

    }

    public override async Task<GetAllCardsResponse> GetAll(Empty request, ServerCallContext context)
    {
        var cards = await repository.GetAllAsync();
        var response = new GetAllCardsResponse();
        foreach (var card in cards)
        {
            response.Cards.Add(new CardModel
            {
                CardId = card.CardId,
                AccountId = card.AccountId,
                CardNumber = card.CardNumber,
                Cvv = card.Cvv,
                Status = card.Status
            });
        }
        return response;
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

        card.CardNumber = request.CardNumber;
        card.Cvv = request.Cvv;
        card.Status = request.Status;

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

        return new CardModel
        {
            CardId = card.CardId,
            AccountId = card.AccountId,
            CardNumber = card.CardNumber,
            Cvv = card.Cvv,
            Status = card.Status
        };
    }
}