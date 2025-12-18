using AutoMapper;
using Core;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using FluentAssertions;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Moq;

namespace TestProject1;

public class AccountTypeServiceTest
{
    private readonly Mock<IAccountTypeRepository> _accountTypeRepositoryMock;
    private readonly Mock<IMapper> _mapperMock;
    private readonly Core.Services.AccountTypeService _accountTypeService;

    public AccountTypeServiceTest()
    {
        _accountTypeRepositoryMock = new Mock<IAccountTypeRepository>();
        _mapperMock = new Mock<IMapper>();
        _accountTypeService = new Core.Services.AccountTypeService(_accountTypeRepositoryMock.Object, _mapperMock.Object);
    }

    [Fact]
    public async Task GetAll_ShouldReturnAllAccountTypes()
    {
        // Arrange
        var request = new GetAllRequest { PageN = 1, PageSize = 10 };
        var accountTypes = new List<AccountType> { new AccountType { TypeId = 1, Name = "Test" } };
        var accountTypeModels = new List<AccountTypeModel> { new AccountTypeModel { TypeId = 1 } };

        _accountTypeRepositoryMock.Setup(r => r.GetAllAsync(request.PageN, request.PageSize)).ReturnsAsync(accountTypes);
        _mapperMock.Setup(m => m.Map<IEnumerable<AccountTypeModel>>(accountTypes)).Returns(accountTypeModels);

        // Act
        var response = await _accountTypeService.GetAll(request, Mock.Of<ServerCallContext>());

        // Assert
        response.AccountTypes.Should().BeEquivalentTo(accountTypeModels);
    }
    [Fact]
    public async Task GetAll_ShouldReturnEmptyList_WhenNoAccountTypesFound()
    {
        // Arrange
        var request = new GetAllRequest { PageN = 1, PageSize = 10 };
        _accountTypeRepositoryMock.Setup(r => r.GetAllAsync(request.PageN, request.PageSize)).ReturnsAsync(new List<AccountType>());

        // Act
        var response = await _accountTypeService.GetAll(request, Mock.Of<ServerCallContext>());

        // Assert
        response.AccountTypes.Should().BeEmpty();
    }

    [Fact]
    public async Task GetById_ShouldReturnAccountType_WhenAccountTypeExists()
    {
        // Arrange
        var request = new GetAccountTypeByIdRequest { TypeId = 1 };
        var accountType = new AccountType { TypeId = 1, Name = "Test" };
        var accountTypeModel = new AccountTypeModel { TypeId = 1 };

        _accountTypeRepositoryMock.Setup(r => r.GetByIdAsync(request.TypeId)).ReturnsAsync(accountType);
        _mapperMock.Setup(m => m.Map<AccountTypeModel>(accountType)).Returns(accountTypeModel);

        // Act
        var response = await _accountTypeService.GetById(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Should().BeEquivalentTo(accountTypeModel);
    }
    [Fact]
    public async Task GetById_ShouldThrowException_WhenAccountTypeDoesNotExist()
    {
        // Arrange
        var request = new GetAccountTypeByIdRequest { TypeId = 1 };
        _accountTypeRepositoryMock.Setup(r => r.GetByIdAsync(request.TypeId)).ReturnsAsync((AccountType?)null);

        // Act & Assert
        await Assert.ThrowsAsync<NotFoundException>(() => _accountTypeService.GetById(request, Mock.Of<ServerCallContext>()));
    }
    [Fact]
    public async Task Update_ShouldUpdateAccountType_WhenAccountTypeExists()
    {
        // Arrange
        var request = new UpdateAccountTypeRequest { TypeId = 1, Name = "Test" };
        var accountType = new AccountType { TypeId = 1, Name = "Test" };
        var accountTypeModel = new AccountTypeModel { TypeId = 1 };

        _accountTypeRepositoryMock.Setup(r => r.GetByIdAsync(request.TypeId)).ReturnsAsync(accountType);
        _mapperMock.Setup(m => m.Map<AccountTypeModel>(accountType)).Returns(accountTypeModel);

        // Act
        var response = await _accountTypeService.Update(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Should().BeOfType<Empty>();
    }   
    [Fact]
    public async Task Update_ShouldThrowException_WhenAccountTypeDoesNotExist()
    {
        // Arrange
        var request = new UpdateAccountTypeRequest { TypeId = 1, Name = "Test" };
        _accountTypeRepositoryMock.Setup(r => r.GetByIdAsync(request.TypeId)).ReturnsAsync((AccountType?)null);

        // Act & Assert
        await Assert.ThrowsAsync<NotFoundException>(() => _accountTypeService.Update(request, Mock.Of<ServerCallContext>()));
    }
    [Fact]
    public async Task Add_ShouldAddAccountType_WhenRequestIsValid()
    {
        // Arrange
        var request = new AddAccountTypeRequest { Name = "Test" };
        var accountType = new AccountType { Name = "Test" };
        var accountTypeModel = new AccountTypeModel { TypeId = 1, Name = "Test" };

        _mapperMock.Setup(m => m.Map<AccountType>(request)).Returns(accountType);
        _accountTypeRepositoryMock.Setup(r => r.AddAsync(accountType)).Returns(Task.CompletedTask);
        _mapperMock.Setup(m => m.Map<AccountTypeModel>(accountType)).Returns(accountTypeModel);

        // Act
        var response = await _accountTypeService.Add(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Should().BeEquivalentTo(accountTypeModel);
    }

    [Fact]
    public async Task Delete_ShouldDeleteAccountType_WhenAccountTypeExists()
    {
        // Arrange
        var request = new DeleteAccountTypeRequest { TypeId = 1 };
        var accountType = new AccountType { TypeId = 1 };

        _accountTypeRepositoryMock.Setup(r => r.GetByIdAsync(request.TypeId)).ReturnsAsync(accountType);
        _accountTypeRepositoryMock.Setup(r => r.DeleteAsync(request.TypeId)).Returns(Task.CompletedTask);

        // Act
        var response = await _accountTypeService.Delete(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Should().BeOfType<Empty>();
    }

    [Fact]
    public async Task Delete_ShouldThrowException_WhenAccountTypeDoesNotExist()
    {
        // Arrange
        var request = new DeleteAccountTypeRequest { TypeId = 1 };
        _accountTypeRepositoryMock.Setup(r => r.GetByIdAsync(request.TypeId)).ReturnsAsync((AccountType?)null);

        // Act & Assert
        await Assert.ThrowsAsync<NotFoundException>(() => _accountTypeService.Delete(request, Mock.Of<ServerCallContext>()));
    }

    [Fact]
    public async Task DeleteBulk_ShouldDeleteAccountTypes_WhenAllExist()
    {
        // Arrange
        var request = new DeleteAccountTypeBulkRequest { AccountTypes = { new DeleteAccountTypeRequest { TypeId = 1 }, new DeleteAccountTypeRequest { TypeId = 2 } } };
        var ids = new List<int> { 1, 2 };
        var accountTypes = new List<AccountType?> { new AccountType { TypeId = 1 }, new AccountType { TypeId = 2 } };

        _accountTypeRepositoryMock.Setup(r => r.GetByIdsAsync(It.Is<IEnumerable<int>>(x => x.SequenceEqual(ids)))).ReturnsAsync(accountTypes);
        _accountTypeRepositoryMock.Setup(r => r.DeleteRangeAsync(It.IsAny<IEnumerable<AccountType>>())).Returns(Task.CompletedTask);

        // Act
        var response = await _accountTypeService.DeleteBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Should().BeOfType<Empty>();
    }

    [Fact]
    public async Task DeleteBulk_ShouldThrowException_WhenNoAccountTypesToDelete()
    {
        // Arrange
        var request = new DeleteAccountTypeBulkRequest();

        // Act & Assert
        await Assert.ThrowsAsync<ValidationException>(() => _accountTypeService.DeleteBulk(request, Mock.Of<ServerCallContext>()));
    }

    [Fact]
    public async Task DeleteBulk_ShouldThrowException_WhenSomeAccountTypesDoNotExist()
    {
        // Arrange
        var request = new DeleteAccountTypeBulkRequest { AccountTypes = { new DeleteAccountTypeRequest { TypeId = 1 }, new DeleteAccountTypeRequest { TypeId = 2 } } };
        var ids = new List<int> { 1, 2 };
        var accountTypes = new List<AccountType?> { new AccountType { TypeId = 1 } }; // Missing one

        _accountTypeRepositoryMock.Setup(r => r.GetByIdsAsync(It.Is<IEnumerable<int>>(x => x.SequenceEqual(ids)))).ReturnsAsync(accountTypes);

        // Act & Assert
        await Assert.ThrowsAsync<ValidationException>(() => _accountTypeService.DeleteBulk(request, Mock.Of<ServerCallContext>()));
    }

    [Fact]
    public async Task UpdateBulk_ShouldUpdateAccountTypes_WhenRequestIsValid()
    {
        // Arrange
        var request = new UpdateAccountTypeBulkRequest { AccountTypes = { new UpdateAccountTypeRequest { TypeId = 1 } } };
        var accountType = new AccountType { TypeId = 1 };

        _mapperMock.Setup(m => m.Map<AccountType>(It.IsAny<UpdateAccountTypeRequest>())).Returns(accountType);
        _accountTypeRepositoryMock.Setup(r => r.UpdateRangeAsync(It.IsAny<IEnumerable<AccountType>>())).Returns(Task.CompletedTask);

        // Act
        var response = await _accountTypeService.UpdateBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Should().BeOfType<Empty>();
    }

    [Fact]
    public async Task UpdateBulk_ShouldThrowException_WhenNoAccountTypesToUpdate()
    {
        // Arrange
        var request = new UpdateAccountTypeBulkRequest();

        // Act & Assert
        await Assert.ThrowsAsync<ValidationException>(() => _accountTypeService.UpdateBulk(request, Mock.Of<ServerCallContext>()));
    }

    [Fact]
    public async Task AddBulk_ShouldAddAccountTypes_WhenRequestIsValid()
    {
        // Arrange
        var request = new AddAccountTypeBulkRequest { AccountTypes = { new AddAccountTypeRequest { Name = "Test" } } };
        var accountType = new AccountType { Name = "Test" };

        _mapperMock.Setup(m => m.Map<AccountType>(It.IsAny<AddAccountTypeRequest>())).Returns(accountType);
        _accountTypeRepositoryMock.Setup(r => r.AddRangeAsync(It.IsAny<IEnumerable<AccountType>>())).Returns(Task.CompletedTask);

        // Act
        var response = await _accountTypeService.AddBulk(request, Mock.Of<ServerCallContext>());

        // Assert
        response.Should().BeOfType<Empty>();
    }

    [Fact]
    public async Task GetByIds_ShouldReturnAccountTypes_WhenTheyExist()
    {
        // Arrange
        var request = new GetAccountTypeByIdsRequest { TypeIds = { 1, 2 } };
        var accountTypes = new List<AccountType> { new AccountType { TypeId = 1 }, new AccountType { TypeId = 2 } };
        var accountTypeModels = new List<AccountTypeModel> { new AccountTypeModel { TypeId = 1 }, new AccountTypeModel { TypeId = 2 } };

        _accountTypeRepositoryMock.Setup(r => r.GetByIdsAsync(request.TypeIds)).ReturnsAsync(accountTypes);
        _mapperMock.Setup(m => m.Map<IEnumerable<AccountTypeModel>>(accountTypes)).Returns(accountTypeModels);

        // Act
        var response = await _accountTypeService.GetByIds(request, Mock.Of<ServerCallContext>());

        // Assert
        response.AccountTypes.Should().BeEquivalentTo(accountTypeModels);
    }
    
}