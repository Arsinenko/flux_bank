using AutoMapper;
using Core;
using Core.Interfaces;
using Core.Models;
using FluentAssertions;
using Grpc.Core;
using Moq;

namespace TestProject1;

public class AtmServiceTest
{
    private readonly Mock<IAtmRepository> _atmRepositoryMock;
    private readonly Mock<IMapper> _mapperMock;
    private readonly Core.Services.AtmService _atmService;

    public AtmServiceTest()
    {
        _atmRepositoryMock = new Mock<IAtmRepository>();
        _mapperMock = new Mock<IMapper>();
        _atmService = new Core.Services.AtmService(_atmRepositoryMock.Object, _mapperMock.Object);
    }

    [Fact]
    public async Task GetAll_ShouldReturnAtms()
    {
        var request = new GetAllRequest()
        {
            PageN = 1,
            PageSize = 10
        };
        var atms = new List<Atm>
        {
            new Atm()
            {
                AtmId = 1,
                Location = "test",
                Status = "status",
                BranchId = 1

            }
        };
        var atmModels = new List<AtmModel>
        {
            new AtmModel()
            {
                AtmId = 1,
                Location = "test",
                BranchId = 1,
                Status = "status"
            }
        };
        _atmRepositoryMock.Setup(r => r.GetAllAsync(request.PageN, request.PageSize)).ReturnsAsync(atms);
        _mapperMock.Setup(m => m.Map<IEnumerable<AtmModel>>(atms)).Returns(atmModels);
        var response = await _atmService.GetAll(request, Mock.Of<ServerCallContext>());

        response.Atms.Should().BeEquivalentTo(atmModels);

    }

    [Fact]
    public async Task GetAll_ShouldReturnEmptyListWhenNoAtms()
    {
        var request = new GetAllRequest()
        {
            PageN = 1,
            PageSize = 10
        };

        var atms = new List<Atm>();
        var atmModels = new List<AtmModel>();
        _atmRepositoryMock.Setup(r => r.GetAllAsync(request.PageN, request.PageSize)).ReturnsAsync(atms);
        _mapperMock.Setup(m => m.Map<IEnumerable<AtmModel>>(atms)).Returns(atmModels);

        var response = await _atmService.GetAll(request, Mock.Of<ServerCallContext>());

        response.Atms.Should().BeEmpty();
    }

    [Fact]
    public async Task Add_ShouldReturnAtm()
    {
        var request = new AddAtmRequest()
        {
            BranchId = 1,
            Location = "test",
            Status = "status",
        };
        var atm = new Atm()
        {
            AtmId = 1,
            Location = "test",
            Status = "status",
            BranchId = 1

        };

        var atmModel = new AtmModel()
        {
            AtmId = 1,
            Location = "test",
            Status = "status",
            BranchId = 1
        };
        
        _mapperMock.Setup(m => m.Map<Atm>(request)).Returns(atm);
        _mapperMock.Setup(m => m.Map<AtmModel>(atm)).Returns(atmModel);
        var result = await _atmService.Add(request, Mock.Of<ServerCallContext>());

        _atmRepositoryMock.Verify(r => r.AddAsync(atm), Times.Once);
        result.Should().BeEquivalentTo(atmModel);
    }

    [Fact]
    public async Task GetById_ShouldReturnAtm()
    {
        var atm = new Atm()
        {
            AtmId = 1,
            Location = "test",
            Status = "status",
            BranchId = 1
        };
        var atmModel = new AtmModel()
        {
            AtmId = 1,
            Location = "test",
            Status = "status",
            BranchId = 1
        };

        var request = new GetAtmByIdRequest()
        {
            AtmId = 1
        };
        _atmRepositoryMock.Setup(r => r.GetByIdAsync(request.AtmId)).ReturnsAsync(atm);
        _mapperMock.Setup(m => m.Map<AtmModel>(atm)).Returns(atmModel);

        var result = await _atmService.GetById(request, Mock.Of<ServerCallContext>());

        result.Should().BeEquivalentTo(atmModel);
    }

    [Fact]
    public async Task GetById_ShouldThrowWhenNotFound()
    {
        var request = new GetAtmByIdRequest()
        {
            AtmId = 1
        };
        _atmRepositoryMock.Setup(r => r.GetByIdAsync(request.AtmId)).ReturnsAsync((Atm)null);

        Func<Task> act = async () => await _atmService.GetById(request, Mock.Of<ServerCallContext>());

        await act.Should().ThrowAsync<Core.Exceptions.NotFoundException>().WithMessage("ATM not found");
    }

    [Fact]
    public async Task Update_ShouldThrowWhenNotFound()
    {
        var request = new UpdateAtmRequest()
        {
            AtmId = 1,
            Location = "test",
            Status = "status",
        };
        var atm = new Atm()
        {
            AtmId = 1,
            Location = "test",
            Status = "status",
            BranchId = 1
        };
        _mapperMock.Setup(m => m.Map<Atm>(request)).Returns(atm);
        _atmRepositoryMock.Setup(r => r.GetByIdAsync(request.AtmId)).ReturnsAsync((Atm)null!);
        Func<Task> act = async () => await _atmService.Update(request, Mock.Of<ServerCallContext>());

        await act.Should().ThrowAsync<Core.Exceptions.NotFoundException>().WithMessage("ATM not found");
    }

}