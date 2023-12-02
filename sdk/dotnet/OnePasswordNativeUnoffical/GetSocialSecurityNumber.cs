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
    public static class GetSocialSecurityNumber
    {
        public static Task<GetSocialSecurityNumberResult> InvokeAsync(GetSocialSecurityNumberArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSocialSecurityNumberResult>("one-password-native-unoffical:index:GetSocialSecurityNumber", args ?? new GetSocialSecurityNumberArgs(), options.WithDefaults());

        public static Output<GetSocialSecurityNumberResult> Invoke(GetSocialSecurityNumberInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSocialSecurityNumberResult>("one-password-native-unoffical:index:GetSocialSecurityNumber", args ?? new GetSocialSecurityNumberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSocialSecurityNumberArgs : Pulumi.InvokeArgs
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

        public GetSocialSecurityNumberArgs()
        {
        }
    }

    public sealed class GetSocialSecurityNumberInvokeArgs : Pulumi.InvokeArgs
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

        public GetSocialSecurityNumberInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSocialSecurityNumberResult
    {
        public readonly ImmutableDictionary<string, Outputs.OutAttachment> Attachments;
        public readonly string Category;
        public readonly ImmutableDictionary<string, Outputs.OutField> Fields;
        public readonly string? Name;
        public readonly string? Notes;
        public readonly string? Number;
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
        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        public readonly string Vault;

        [OutputConstructor]
        private GetSocialSecurityNumberResult(
            ImmutableDictionary<string, Outputs.OutAttachment> attachments,

            string category,

            ImmutableDictionary<string, Outputs.OutField> fields,

            string? name,

            string? notes,

            string? number,

            ImmutableDictionary<string, Outputs.OutField> references,

            ImmutableDictionary<string, Outputs.OutSection> sections,

            ImmutableArray<string> tags,

            string title,

            string uuid,

            string vault)
        {
            Attachments = attachments;
            Category = category;
            Fields = fields;
            Name = name;
            Notes = notes;
            Number = number;
            References = references;
            Sections = sections;
            Tags = tags;
            Title = title;
            Uuid = uuid;
            Vault = vault;
        }
    }
}
