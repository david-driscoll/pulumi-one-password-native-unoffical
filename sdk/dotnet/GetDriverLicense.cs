// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    public static class GetDriverLicense
    {
        public static Task<GetDriverLicenseResult> InvokeAsync(GetDriverLicenseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDriverLicenseResult>("onepassword:index:GetDriverLicense", args ?? new GetDriverLicenseArgs(), options.WithDefaults());

        public static Output<GetDriverLicenseResult> Invoke(GetDriverLicenseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDriverLicenseResult>("onepassword:index:GetDriverLicense", args ?? new GetDriverLicenseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDriverLicenseArgs : Pulumi.InvokeArgs
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

        public GetDriverLicenseArgs()
        {
        }
    }

    public sealed class GetDriverLicenseInvokeArgs : Pulumi.InvokeArgs
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

        public GetDriverLicenseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDriverLicenseResult
    {
        public readonly string? Address;
        public readonly ImmutableDictionary<string, Outputs.OutField> Attachments;
        public readonly string Category;
        public readonly string? ConditionsRestrictions;
        public readonly string? Country;
        public readonly string? DateOfBirth;
        public readonly string? ExpiryDate;
        public readonly ImmutableDictionary<string, Outputs.OutField> Fields;
        public readonly string? FullName;
        public readonly string? Gender;
        public readonly string? Height;
        public readonly string? LicenseClass;
        public readonly string? Notes;
        public readonly string? Number;
        public readonly ImmutableDictionary<string, Outputs.OutField> References;
        public readonly ImmutableDictionary<string, Outputs.OutSection> Sections;
        public readonly string? State;
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
        private GetDriverLicenseResult(
            string? address,

            ImmutableDictionary<string, Outputs.OutField> attachments,

            string category,

            string? conditionsRestrictions,

            string? country,

            string? dateOfBirth,

            string? expiryDate,

            ImmutableDictionary<string, Outputs.OutField> fields,

            string? fullName,

            string? gender,

            string? height,

            string? licenseClass,

            string? notes,

            string? number,

            ImmutableDictionary<string, Outputs.OutField> references,

            ImmutableDictionary<string, Outputs.OutSection> sections,

            string? state,

            ImmutableArray<string> tags,

            string title,

            string uuid,

            string vault)
        {
            Address = address;
            Attachments = attachments;
            Category = category;
            ConditionsRestrictions = conditionsRestrictions;
            Country = country;
            DateOfBirth = dateOfBirth;
            ExpiryDate = expiryDate;
            Fields = fields;
            FullName = fullName;
            Gender = gender;
            Height = height;
            LicenseClass = licenseClass;
            Notes = notes;
            Number = number;
            References = references;
            Sections = sections;
            State = state;
            Tags = tags;
            Title = title;
            Uuid = uuid;
            Vault = vault;
        }
    }
}
