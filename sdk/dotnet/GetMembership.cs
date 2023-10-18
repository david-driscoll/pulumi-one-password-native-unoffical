// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OnePasswordNative
{
    public static class GetMembership
    {
        public static Task<GetMembershipResult> InvokeAsync(GetMembershipArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMembershipResult>("one-password-native:index:GetMembership", args ?? new GetMembershipArgs(), options.WithDefaults());

        public static Output<GetMembershipResult> Invoke(GetMembershipInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMembershipResult>("one-password-native:index:GetMembership", args ?? new GetMembershipInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMembershipArgs : Pulumi.InvokeArgs
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

        public GetMembershipArgs()
        {
        }
    }

    public sealed class GetMembershipInvokeArgs : Pulumi.InvokeArgs
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

        public GetMembershipInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMembershipResult
    {
        public readonly ImmutableDictionary<string, Outputs.OutField> Attachments;
        public readonly string Category;
        public readonly string? ExpiryDate;
        public readonly ImmutableDictionary<string, Outputs.OutField> Fields;
        public readonly string? Group;
        public readonly string? MemberId;
        public readonly string? MemberName;
        public readonly string? MemberSince;
        public readonly string? Notes;
        public readonly string? Pin;
        public readonly ImmutableDictionary<string, Outputs.OutField> References;
        public readonly ImmutableDictionary<string, Outputs.OutSection> Sections;
        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly string? Telephone;
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
        public readonly string? Website;

        [OutputConstructor]
        private GetMembershipResult(
            ImmutableDictionary<string, Outputs.OutField> attachments,

            string category,

            string? expiryDate,

            ImmutableDictionary<string, Outputs.OutField> fields,

            string? group,

            string? memberId,

            string? memberName,

            string? memberSince,

            string? notes,

            string? pin,

            ImmutableDictionary<string, Outputs.OutField> references,

            ImmutableDictionary<string, Outputs.OutSection> sections,

            ImmutableArray<string> tags,

            string? telephone,

            string title,

            string uuid,

            string vault,

            string? website)
        {
            Attachments = attachments;
            Category = category;
            ExpiryDate = expiryDate;
            Fields = fields;
            Group = group;
            MemberId = memberId;
            MemberName = memberName;
            MemberSince = memberSince;
            Notes = notes;
            Pin = pin;
            References = references;
            Sections = sections;
            Tags = tags;
            Telephone = telephone;
            Title = title;
            Uuid = uuid;
            Vault = vault;
            Website = website;
        }
    }
}
