using System.Collections;
using System.Collections.Immutable;
using System.Reflection;
using System.Reflection.Metadata;
using FluentAssertions;
using Microsoft.VisualStudio.TestPlatform.Utilities;
using Pulumi;
using pulumi_resource_one_password_native_unofficial;
using pulumi_resource_one_password_native_unofficial.OnePasswordCli;
using Pulumi.Experimental.Provider;
using Rocket.Surgery.OnePasswordNativeUnofficial;
using Rocket.Surgery.OnePasswordNativeUnofficial.Inputs;
using Serilog;
using Serilog.Core;
using Serilog.Events;
using TestProject.Helpers;
using Xunit.Abstractions;
using YamlDotNet.Serialization;
using FieldType = Rocket.Surgery.OnePasswordNativeUnofficial.FieldType;

namespace TestProject;

[Collection("Connect collection")]
[UsesVerify]
public class ConnectServerItemTests : IClassFixture<PulumiFixture>
{
    private readonly PulumiFixture _fixture;
    private readonly IServerFixture _serverFixture;
    private readonly Logger _logger;

    public ConnectServerItemTests(PulumiFixture fixture, ConnectServerFixture serverFixture, ITestOutputHelper output)
    {
        _fixture = fixture;
        _serverFixture = serverFixture;
        _logger = new LoggerConfiguration()
            .WriteTo.TestOutput(output, LogEventLevel.Verbose)
            .CreateLogger();
        fixture.Connect(serverFixture);
    }

    [Fact]
    public async Task Should_Throw_When_Attachments_Are_Used()
    {
        var provider = new OnePasswordProvider(_logger);

        await _serverFixture.ConfigureProvider(provider);

        var data = await _fixture.CreateRequestObject<LoginItem, LoginItemArgs>("testlogin", new()
        {
            Vault = "testing-pulumi",
            Username = "me",
            Attachments = new()
            {
                ["my-attachment"] = new StringAsset("this is an attachment"),
                // currently there is no way to have a period escaped via the cli
                // ["package.json"] = new FileAsset("./Pulumi.yaml")
            },
            // Password = "secret1234",
            Fields = new()
            {
                ["password"] = new FieldArgs()
                {
                    Value = "secret1234",
                    Type = FieldType.Concealed
                }
            },
            Sections = new()
            {
                ["mysection"] = new SectionArgs()
                {
                    Fields = new()
                    {
                        ["password2"] = new FieldArgs()
                        {
                            Value = "secret1235!",
                            Type = FieldType.Concealed
                        }
                    },
                    Attachments = new()
                    {
                        ["my-different-attachment"] = new StringAsset("this is my different attachment"),
                        // currently there is no way to have a period escaped via the cli
                        // ["package.json"] = new FileAsset("./Pulumi.yaml")
                    },
                }
            },
            Tags = new string[] { "test-tag" }
        });

        Func<Task> Action = () => provider.Create(new CreateRequest(data.Urn, data.Request, TimeSpan.MaxValue, false), CancellationToken.None);
        await Action.Should().ThrowAsync<NotSupportedException>();
    }

    [Fact]
    public async Task Should_Create_Login_Item()
    {
        var provider = new OnePasswordProvider(_logger);

        await _serverFixture.ConfigureProvider(provider);

        var data = await _fixture.CreateRequestObject<LoginItem, LoginItemArgs>("myitem", new()
        {
            Vault = "testing-pulumi",
            Username = "me",
            // Password = "secret1234",
            Fields = new()
            {
                ["password"] = new FieldArgs()
                {
                    Value = "secret1234",
                    Type = FieldType.Concealed
                }
            },
            Sections = new()
            {
                ["mysection"] = new SectionArgs()
                {
                    Fields = new()
                    {
                        ["password2"] = new FieldArgs()
                        {
                            Value = "secret1235!",
                            Type = FieldType.Concealed
                        }
                    },
                }
            },
            Tags = new string[] { "test-tag" }
        });

        var create = await provider.Create(new CreateRequest(data.Urn, data.Request, TimeSpan.MaxValue, false), CancellationToken.None);

        await Verify(create)
            .AddScrubber(z => z.Replace(create.Id!, "[server-generated]"));
    }
    

    [Fact]
    public async Task Should_Generate_Password()
    {
        var provider = new OnePasswordProvider(_logger);

        await _serverFixture.ConfigureProvider(provider);

        var data = await _fixture.CreateRequestObject<LoginItem, LoginItemArgs>("myitem", new()
        {
            Vault = "testing-pulumi",
            Username = "me",
            GeneratePassword = new PasswordRecipeArgs()
            {
                Digits = true,
                Length = 40,
                Symbols = true,
                Letters = true
            },
            Tags = new string[] { "test-tag" }
        });

        var create = await provider.Create(new CreateRequest(data.Urn, data.Request, TimeSpan.MaxValue, false), CancellationToken.None);

        await Verify(create)
            .AddScrubber(z => z.Replace(create.Id!, "[server-generated]"))
            .AddScrubber(x => x.Replace(TemplateMetadata.GetObjectStringValue(create.Properties as IReadOnlyDictionary<string, PropertyValue>, "password"), "[redacted,server-generated]"));
    }
    

    [Fact]
    public async Task Should_Generate_Random_Password()
    {
        var provider = new OnePasswordProvider(_logger);

        await _serverFixture.ConfigureProvider(provider);

        var data = await _fixture.CreateRequestObject<LoginItem, LoginItemArgs>("myitem", new()
        {
            Vault = "testing-pulumi",
            Username = "me",
            GeneratePassword = true,
            Tags = new string[] { "test-tag" }
        });

        var create = await provider.Create(new CreateRequest(data.Urn, data.Request, TimeSpan.MaxValue, false), CancellationToken.None);

        await Verify(create)
            .AddScrubber(z => z.Replace(create.Id!, "[server-generated]"))
            .AddScrubber(x => x.Replace(TemplateMetadata.GetObjectStringValue(create.Properties as IReadOnlyDictionary<string, PropertyValue>, "password"), "[redacted,server-generated]"));
    }
}
