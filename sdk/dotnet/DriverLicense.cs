// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    [OnepasswordResourceType("onepassword:index:DriverLicense")]
    public partial class DriverLicense : Pulumi.CustomResource
    {
        [Output("address")]
        public Output<string?> Address { get; private set; } = null!;

        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        [Output("conditionsRestrictions")]
        public Output<string?> ConditionsRestrictions { get; private set; } = null!;

        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        [Output("dateOfBirth")]
        public Output<string?> DateOfBirth { get; private set; } = null!;

        [Output("expiryDate")]
        public Output<string?> ExpiryDate { get; private set; } = null!;

        [Output("fields")]
        public Output<ImmutableArray<Outputs.GetField>> Fields { get; private set; } = null!;

        [Output("fullName")]
        public Output<string?> FullName { get; private set; } = null!;

        [Output("gender")]
        public Output<string?> Gender { get; private set; } = null!;

        [Output("height")]
        public Output<string?> Height { get; private set; } = null!;

        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("licenseClass")]
        public Output<string?> LicenseClass { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("number")]
        public Output<string?> Number { get; private set; } = null!;

        [Output("sections")]
        public Output<ImmutableArray<Outputs.GetSection>> Sections { get; private set; } = null!;

        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The title of the item.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Output("vault")]
        public Output<string> Vault { get; private set; } = null!;


        /// <summary>
        /// Create a DriverLicense resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DriverLicense(string name, DriverLicenseArgs args, CustomResourceOptions? options = null)
            : base("onepassword:index:DriverLicense", name, args ?? new DriverLicenseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DriverLicense(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("onepassword:index:DriverLicense", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DriverLicense resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DriverLicense Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DriverLicense(name, id, options);
        }
    }

    public sealed class DriverLicenseArgs : Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("category", required: true)]
        public InputUnion<Pulumi.Onepassword.Category, string> Category { get; set; } = null!;

        [Input("conditionsRestrictions")]
        public Input<string>? ConditionsRestrictions { get; set; }

        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("dateOfBirth")]
        public Input<string>? DateOfBirth { get; set; }

        [Input("expiryDate")]
        public Input<string>? ExpiryDate { get; set; }

        [Input("fields")]
        private InputList<Inputs.FieldArgs>? _fields;
        public InputList<Inputs.FieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.FieldArgs>());
            set => _fields = value;
        }

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        [Input("gender")]
        public Input<string>? Gender { get; set; }

        [Input("height")]
        public Input<string>? Height { get; set; }

        [Input("licenseClass")]
        public Input<string>? LicenseClass { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("number")]
        public Input<string>? Number { get; set; }

        [Input("sections")]
        private InputList<Inputs.SectionArgs>? _sections;
        public InputList<Inputs.SectionArgs> Sections
        {
            get => _sections ?? (_sections = new InputList<Inputs.SectionArgs>());
            set => _sections = value;
        }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// An array of strings of the tags assigned to the item.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public DriverLicenseArgs()
        {
            Category = "Item";
        }
    }
}
