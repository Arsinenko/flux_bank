using AutoMapper;
using Core.Mappings;

namespace TestProject1;

public class AutoMapperConfigurationTests
{
    [Fact]
    public void AutoMapper_Configuration_IsValid()
    {
        var config = new MapperConfiguration(cfg =>
        {
            cfg.AddProfile<ProtoMappingProfile>();
        });

        config.AssertConfigurationIsValid();
    }
}