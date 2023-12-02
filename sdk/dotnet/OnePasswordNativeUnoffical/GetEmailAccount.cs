// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Rocket.Surgery.OnePasswordNativeUnoffical
{
    public static class GetEmailAccount
    {
        public static Task<GetEmailAccountResult> InvokeAsync(GetEmailAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEmailAccountResult>("one-password-native-unoffical:index:GetEmailAccount", args ?? new GetEmailAccountArgs(), options.WithDefaults());

        public static Output<GetEmailAccountResult> Invoke(GetEmailAccountInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEmailAccountResult>("one-password-native-unoffical:index:GetEmailAccount", args ?? new GetEmailAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEmailAccountArgs : Pulumi.InvokeArgs
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

        public GetEmailAccountArgs()
        {
        }
    }

    public sealed class GetEmailAccountInvokeArgs : Pulumi.InvokeArgs
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

        public GetEmailAccountInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEmailAccountResult
    {
        public readonly ImmutableDictionary<string, Outputs.OutAttachment> Attachments;
        public readonly string? AuthMethod;
        public readonly string Category;
        public readonly Rocket.Surgery.OnePasswordNativeUnoffical.EmailAccount.Outputs.ContactInformationSection? ContactInformation;
        public readonly ImmutableDictionary<string, Outputs.OutField> Fields;
        public readonly string? Notes;
        public readonly string? Password;
        public readonly string? PortNumber;
        public readonly ImmutableDictionary<string, Outputs.OutField> References;
        public readonly ImmutableDictionary<string, Outputs.OutSection> Sections;
        public readonly string? Security;
        public readonly string? Server;
        public readonly Rocket.Surgery.OnePasswordNativeUnoffical.EmailAccount.Outputs.SmtpSection? Smtp;
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
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        public readonly string Vault;

        [OutputConstructor]
        private GetEmailAccountResult(
            ImmutableDictionary<string, Outputs.OutAttachment> attachments,

            string? authMethod,

            string category,

            Rocket.Surgery.OnePasswordNativeUnoffical.EmailAccount.Outputs.ContactInformationSection? contactInformation,

            ImmutableDictionary<string, Outputs.OutField> fields,

            string? notes,

            string? password,

            string? portNumber,

            ImmutableDictionary<string, Outputs.OutField> references,

            ImmutableDictionary<string, Outputs.OutSection> sections,

            string? security,

            string? server,

            Rocket.Surgery.OnePasswordNativeUnoffical.EmailAccount.Outputs.SmtpSection? smtp,

            ImmutableArray<string> tags,

            string title,

            string? type,

            string? username,

            string uuid,

            string vault)
        {
            Attachments = attachments;
            AuthMethod = authMethod;
            Category = category;
            ContactInformation = contactInformation;
            Fields = fields;
            Notes = notes;
            Password = password;
            PortNumber = portNumber;
            References = references;
            Sections = sections;
            Security = security;
            Server = server;
            Smtp = smtp;
            Tags = tags;
            Title = title;
            Type = type;
            Username = username;
            Uuid = uuid;
            Vault = vault;
        }
    }
}
