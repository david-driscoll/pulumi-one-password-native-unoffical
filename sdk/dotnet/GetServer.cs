// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    public static class GetServer
    {
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("onepassword:index:GetServer", args ?? new GetServerArgs(), options.WithDefaults());

        public static Output<GetServerResult> Invoke(GetServerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetServerResult>("onepassword:index:GetServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title")]
        public string? Title { get; set; }

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Input("uuid")]
        public string? Uuid { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public string Vault { get; set; } = null!;

        public GetServerArgs()
        {
        }
    }

    public sealed class GetServerInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public GetServerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerResult
    {
        public readonly Pulumi.Onepassword.Server.Outputs.AdminConsoleSection? AdminConsole;
        public readonly string? Category;
        public readonly ImmutableArray<Outputs.GetField> Fields;
        public readonly Pulumi.Onepassword.Server.Outputs.HostingProviderSection? HostingProvider;
        public readonly string? Id;
        public readonly string? Notes;
        public readonly string? Password;
        public readonly ImmutableArray<Outputs.GetSection> Sections;
        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The title of the item.
        /// </summary>
        public readonly string? Title;
        public readonly string? Url;
        public readonly string? Username;
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        public readonly string? Uuid;
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        public readonly string? Vault;

        [OutputConstructor]
        private GetServerResult(
            Pulumi.Onepassword.Server.Outputs.AdminConsoleSection? adminConsole,

            string? category,

            ImmutableArray<Outputs.GetField> fields,

            Pulumi.Onepassword.Server.Outputs.HostingProviderSection? hostingProvider,

            string? id,

            string? notes,

            string? password,

            ImmutableArray<Outputs.GetSection> sections,

            ImmutableArray<string> tags,

            string? title,

            string? url,

            string? username,

            string? uuid,

            string? vault)
        {
            AdminConsole = adminConsole;
            Category = category;
            Fields = fields;
            HostingProvider = hostingProvider;
            Id = id;
            Notes = notes;
            Password = password;
            Sections = sections;
            Tags = tags;
            Title = title;
            Url = url;
            Username = username;
            Uuid = uuid;
            Vault = vault;
        }
    }
}
