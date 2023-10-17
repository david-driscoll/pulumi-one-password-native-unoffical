// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    public static class GetAPICredential
    {
        public static Task<GetAPICredentialResult> InvokeAsync(GetAPICredentialArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAPICredentialResult>("onepassword:index:GetAPICredential", args ?? new GetAPICredentialArgs(), options.WithDefaults());

        public static Output<GetAPICredentialResult> Invoke(GetAPICredentialInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAPICredentialResult>("onepassword:index:GetAPICredential", args ?? new GetAPICredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAPICredentialArgs : Pulumi.InvokeArgs
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

        public GetAPICredentialArgs()
        {
        }
    }

    public sealed class GetAPICredentialInvokeArgs : Pulumi.InvokeArgs
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

        public GetAPICredentialInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAPICredentialResult
    {
        public readonly ImmutableDictionary<string, Outputs.OutField> Attachments;
        public readonly string Category;
        public readonly string? Credential;
        public readonly string? Expires;
        public readonly ImmutableDictionary<string, Outputs.OutField> Fields;
        public readonly string? Filename;
        public readonly string? Hostname;
        public readonly string? Notes;
        public readonly ImmutableDictionary<string, Outputs.OutField> References;
        public readonly ImmutableDictionary<string, Outputs.OutSection> Sections;
        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The title of the item.
        /// </summary>
        public readonly string Title;
        public readonly string? Type;
        public readonly string? Username;
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        public readonly string Uuid;
        public readonly string? ValidFrom;
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        public readonly string Vault;

        [OutputConstructor]
        private GetAPICredentialResult(
            ImmutableDictionary<string, Outputs.OutField> attachments,

            string category,

            string? credential,

            string? expires,

            ImmutableDictionary<string, Outputs.OutField> fields,

            string? filename,

            string? hostname,

            string? notes,

            ImmutableDictionary<string, Outputs.OutField> references,

            ImmutableDictionary<string, Outputs.OutSection> sections,

            ImmutableArray<string> tags,

            string title,

            string? type,

            string? username,

            string uuid,

            string? validFrom,

            string vault)
        {
            Attachments = attachments;
            Category = category;
            Credential = credential;
            Expires = expires;
            Fields = fields;
            Filename = filename;
            Hostname = hostname;
            Notes = notes;
            References = references;
            Sections = sections;
            Tags = tags;
            Title = title;
            Type = type;
            Username = username;
            Uuid = uuid;
            ValidFrom = validFrom;
            Vault = vault;
        }
    }
}
