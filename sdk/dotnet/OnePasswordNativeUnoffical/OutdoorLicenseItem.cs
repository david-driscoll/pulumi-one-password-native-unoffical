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
    [OnePasswordNativeUnofficalResourceType("one-password-native-unoffical:index:OutdoorLicenseItem")]
    public partial class OutdoorLicenseItem : Pulumi.CustomResource
    {
        [Output("approvedWildlife")]
        public Output<string?> ApprovedWildlife { get; private set; } = null!;

        [Output("attachments")]
        public Output<ImmutableDictionary<string, Outputs.OutAttachment>> Attachments { get; private set; } = null!;

        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        [Output("country")]
        public Output<string?> Country { get; private set; } = null!;

        [Output("expires")]
        public Output<string?> Expires { get; private set; } = null!;

        [Output("fields")]
        public Output<ImmutableDictionary<string, Outputs.OutField>> Fields { get; private set; } = null!;

        [Output("fullName")]
        public Output<string?> FullName { get; private set; } = null!;

        [Output("maximumQuota")]
        public Output<string?> MaximumQuota { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("references")]
        public Output<ImmutableDictionary<string, Outputs.OutField>> References { get; private set; } = null!;

        [Output("sections")]
        public Output<ImmutableDictionary<string, Outputs.OutSection>> Sections { get; private set; } = null!;

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

        [Output("validFrom")]
        public Output<string?> ValidFrom { get; private set; } = null!;

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Output("vault")]
        public Output<string> Vault { get; private set; } = null!;


        /// <summary>
        /// Create a OutdoorLicenseItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OutdoorLicenseItem(string name, OutdoorLicenseItemArgs args, CustomResourceOptions? options = null)
            : base("one-password-native-unoffical:index:OutdoorLicenseItem", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private OutdoorLicenseItem(string name, Input<string> id, OutdoorLicenseItemState? state = null, CustomResourceOptions? options = null)
            : base("one-password-native-unoffical:index:OutdoorLicenseItem", name, state, MakeResourceOptions(options, id))
        {
        }

        private static OutdoorLicenseItemArgs MakeArgs(OutdoorLicenseItemArgs args)
        {
            args ??= new OutdoorLicenseItemArgs();
            args.Category = "Outdoor License";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/david-driscoll",
                AdditionalSecretOutputs =
                {
                    "attachments",
                    "fields",
                    "references",
                    "sections",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OutdoorLicenseItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OutdoorLicenseItem Get(string name, Input<string> id, OutdoorLicenseItemState? state = null, CustomResourceOptions? options = null)
        {
            return new OutdoorLicenseItem(name, id, state, options);
        }

        public Pulumi.Output<OutdoorLicenseItemGetAttachmentResult> GetAttachment(OutdoorLicenseItemGetAttachmentArgs args)
            => Pulumi.Deployment.Instance.Call<OutdoorLicenseItemGetAttachmentResult>("one-password-native-unoffical:index:OutdoorLicenseItem/attachment", args ?? new OutdoorLicenseItemGetAttachmentArgs(), this);
    }

    public sealed class OutdoorLicenseItemArgs : Pulumi.ResourceArgs
    {
        [Input("approvedWildlife")]
        public Input<string>? ApprovedWildlife { get; set; }

        /// <summary>
        /// The category of the vault the item is in.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("expires")]
        public Input<string>? Expires { get; set; }

        [Input("fields")]
        private InputMap<Inputs.FieldArgs>? _fields;
        public InputMap<Inputs.FieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputMap<Inputs.FieldArgs>());
            set => _fields = value;
        }

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        [Input("inputAttachments")]
        private InputMap<AssetOrArchive>? _inputAttachments;
        public InputMap<AssetOrArchive> InputAttachments
        {
            get => _inputAttachments ?? (_inputAttachments = new InputMap<AssetOrArchive>());
            set => _inputAttachments = value;
        }

        [Input("maximumQuota")]
        public Input<string>? MaximumQuota { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("sections")]
        private InputMap<Inputs.SectionArgs>? _sections;
        public InputMap<Inputs.SectionArgs> Sections
        {
            get => _sections ?? (_sections = new InputMap<Inputs.SectionArgs>());
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
        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("validFrom")]
        public Input<string>? ValidFrom { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public OutdoorLicenseItemArgs()
        {
        }
    }

    public sealed class OutdoorLicenseItemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public OutdoorLicenseItemState()
        {
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="OutdoorLicenseItem.GetAttachment"/> method.
    /// </summary>
    public sealed class OutdoorLicenseItemGetAttachmentArgs : Pulumi.CallArgs
    {
        /// <summary>
        /// The name or uuid of the attachment to get
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public OutdoorLicenseItemGetAttachmentArgs()
        {
        }
    }

    /// <summary>
    /// The results of the <see cref="OutdoorLicenseItem.GetAttachment"/> method.
    /// </summary>
    [OutputType]
    public sealed class OutdoorLicenseItemGetAttachmentResult
    {
        /// <summary>
        /// the value of the attachment
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private OutdoorLicenseItemGetAttachmentResult(string value)
        {
            Value = value;
        }
    }
}
