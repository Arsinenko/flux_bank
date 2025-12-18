using AutoMapper;
using Core.Interfaces;
using Core.Models;
using FluentAssertions;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Moq;
using Core;
using Core.Exceptions;
using ValidationException = System.ComponentModel.DataAnnotations.ValidationException;

namespace TestProject1;

public class AccountServiceTest
{
    private readonly Mock<IAccountRepository> _accountRepositoryMock;
    private readonly Mock<IMapper> _mapperMock;
    private readonly Core.Services.AccountService _accountService;

    public AccountServiceTest()
    {
        _accountRepositoryMock = new Mock<IAccountRepository>();
        _mapperMock = new Mock<IMapper>();
        _accountService = new Core.Services.AccountService(_accountRepositoryMock.Object, _mapperMock.Object);
    }

    [Fact]
    public async Task GetAll_ShouldReturnAllAccounts()
    {
        // Arrange
        var request = new GetAllRequest { PageN = 1, PageSize = 10 };
        var accounts = new List<Account> { new Account { AccountId = 1, CustomerId = 1, TypeId = 1, Iban = "Test" } };
        var accountModels = new List<AccountModel> { new AccountModel { AccountId = 1 } };

        _accountRepositoryMock.Setup(r => r.GetAllAsync(request.PageN, request.PageSize)).ReturnsAsync(accounts);
        _mapperMock.Setup(m => m.Map<IEnumerable<AccountModel>>(accounts)).Returns(accountModels);

        // Act
        var response = await _accountService.GetAll(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Accounts.Should().BeEquivalentTo(accountModels);
    }

    [Fact]
    public async Task Add_ShouldAddAccount()
    {
        // Arrange
        var request = new AddAccountRequest { CustomerId = 1, TypeId = 1 };
        var account = new Account { AccountId = 1, CustomerId = 1, TypeId = 1, Iban = "Test" };
        var accountModel = new AccountModel { AccountId = 1 };

        _mapperMock.Setup(m => m.Map<Account>(request)).Returns(account);
        _mapperMock.Setup(m => m.Map<AccountModel>(account)).Returns(accountModel);

        // Act
        var result = await _accountService.Add(request, Mock.Of<ServerCallContext>());

        // Assert
        _accountRepositoryMock.Verify(r => r.AddAsync(account), Times.Once);
        result.Should().BeEquivalentTo(accountModel);
    }

    [Fact]
    public async Task GetById_ShouldReturnAccount_WhenAccountExists()
    {
        // Arrange
        var request = new GetAccountByIdRequest { AccountId = 1 };
        var account = new Account { AccountId = 1, CustomerId = 1, TypeId = 1, Iban = "Test" };
        var accountModel = new AccountModel { AccountId = 1 };

        _accountRepositoryMock.Setup(r => r.GetByIdAsync(request.AccountId)).ReturnsAsync(account);
        _mapperMock.Setup(m => m.Map<AccountModel>(account)).Returns(accountModel);

        // Act
        var result = await _accountService.GetById(request, Mock.Of<ServerCallContext>());

        // Assert
        result.Should().BeEquivalentTo(accountModel);
    }

    [Fact]
    public async Task GetById_ShouldThrowRpcException_WhenAccountNotFound()
    {
        // Arrange
        var request = new GetAccountByIdRequest { AccountId = 1 };
        _accountRepositoryMock.Setup(r => r.GetByIdAsync(request.AccountId)).ReturnsAsync((Account)null);

        // Act
        Func<Task> act = async () => await _accountService.GetById(request, Mock.Of<ServerCallContext>());

        // Assert
        var exception = await act.Should().ThrowAsync<NotFoundException>();
    }

    [Fact]
    public async Task Update_ShouldUpdateAccount_WhenAccountExists()
    {
        // Arrange
        var request = new UpdateAccountRequest { AccountId = 1, Iban = "Updated" };
        var account = new Account { AccountId = 1, CustomerId = 1, TypeId = 1, Iban = "Test" };

        _accountRepositoryMock.Setup(r => r.GetByIdAsync(request.AccountId)).ReturnsAsync(account);

        // Act
        var result = await _accountService.Update(request, Mock.Of<ServerCallContext>());

        // Assert
        _mapperMock.Verify(m => m.Map(request, account), Times.Once);
        _accountRepositoryMock.Verify(r => r.UpdateAsync(account), Times.Once);
        result.Should().BeOfType<Empty>();
    }

    [Fact]
    public async Task Update_ShouldThrowRpcException_WhenAccountNotFound()
    {
        // Arrange
        var request = new UpdateAccountRequest { AccountId = 1 };
        _accountRepositoryMock.Setup(r => r.GetByIdAsync(request.AccountId)).ReturnsAsync((Account)null);

        // Act
        Func<Task> act = async () => await _accountService.Update(request, Mock.Of<ServerCallContext>());

        // Assert
        var exception = await act.Should().ThrowAsync<NotFoundException>();
    }

    [Fact]
    public async Task Delete_ShouldDeleteAccount()
    {
        // Arrange
        var request = new DeleteAccountRequest { AccountId = 1 };

        // Act
        var result = await _accountService.Delete(request, Mock.Of<ServerCallContext>());

        // Assert
        _accountRepositoryMock.Verify(r => r.DeleteAsync(request.AccountId), Times.Once);
        result.Should().BeOfType<Empty>();
    }

    [Fact]
    public async Task DeleteBulk_ShouldDeleteAccounts()
    {
        // Arrange
        var request = new DeleteAccountBulkRequest();
        request.Accounts.Add(new DeleteAccountRequest { AccountId = 1 });
        var ids = new List<int> { 1 };
        var accounts = new List<Account> { new Account { AccountId = 1, CustomerId = 1, TypeId = 1, Iban = "Test" } };

        _accountRepositoryMock.Setup(r => r.GetByIdsAsync(ids)).ReturnsAsync(accounts);

        // Act
        var result = await _accountService.DeleteBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        _accountRepositoryMock.Verify(r => r.DeleteRangeAsync(It.Is<List<Account>>(list => list.Count == 1 && list[0].AccountId == 1)), Times.Once);
        result.Should().BeOfType<Empty>();
    }
    
    [Fact]
    public async Task DeleteBulk_ShouldThrowRpcException_WhenNoAccountsToDelete()
    {
        // Arrange
        var request = new DeleteAccountBulkRequest();

        // Act
        Func<Task> act = async () => await _accountService.DeleteBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        var exception = await act.Should().ThrowAsync<ValidationException>();
        exception.Which.Message.Should().Be("No accounts to delete");
    }

    [Fact]
    public async Task DeleteBulk_ShouldThrowRpcException_WhenSomeAccountsNotFound()
    {
        // Arrange
        var request = new DeleteAccountBulkRequest();
        request.Accounts.Add(new DeleteAccountRequest { AccountId = 1 });
        request.Accounts.Add(new DeleteAccountRequest { AccountId = 2 });
        var ids = new List<int> { 1, 2 };
        var accounts = new List<Account> { new Account { AccountId = 1, CustomerId = 1, TypeId = 1, Iban = "Test" } };

        _accountRepositoryMock.Setup(r => r.GetByIdsAsync(ids)).ReturnsAsync(accounts);

        // Act
        Func<Task> act = async () => await _accountService.DeleteBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        var exception = await act.Should().ThrowAsync<ValidationException>();
        exception.Which.Message.Should().Be("One or more accounts not found");
    }

    [Fact]
    public async Task UpdateBulk_ShouldUpdateAccounts()
    {
        // Arrange
        var request = new UpdateAccountBulkRequest();
        var updateAccountRequest = new UpdateAccountRequest { AccountId = 1, Iban = "Updated" };
        request.Accounts.Add(updateAccountRequest);
        var account = new Account { AccountId = 1, CustomerId = 1, TypeId = 1, Iban = "Updated" };

        _mapperMock.Setup(m => m.Map<Account>(updateAccountRequest)).Returns(account);

        // Act
        var result = await _accountService.UpdateBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        _accountRepositoryMock.Verify(r => r.UpdateRangeAsync(It.Is<List<Account>>(list => list.Count == 1 && list[0].AccountId == 1)), Times.Once);
        result.Should().BeOfType<Empty>();
    }

    [Fact]
    public async Task UpdateBulk_ShouldThrowRpcException_WhenNoAccountsToUpdate()
    {
        // Arrange
        var request = new UpdateAccountBulkRequest();

        // Act
        Func<Task> act = async () => await _accountService.UpdateBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        var exception = await act.Should().ThrowAsync<ValidationException>();
        exception.Which.Message.Should().Be("No accounts found");
    }
}