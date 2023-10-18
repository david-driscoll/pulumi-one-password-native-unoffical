// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Onepassword
{
    [OnepasswordResourceType("onepassword:index:SocialSecurityNumberItem")]
    public partial class SocialSecurityNumberItem : Pulumi.CustomResource
    {
        [Output("attachments")]
        public Output<ImmutableDictionary<string, Outputs.OutField>> Attachments { get; private set; } = null!;

        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        [Output("fields")]
        public Output<ImmutableDictionary<string, Outputs.OutField>> Fields { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("number")]
        public Output<string?> Number { get; private set; } = null!;

        [Output("references")]
        public Output<ImmutableDictionary<string, Outputs.OutField>> References { get; private set; } = null!;

        [Output("sections")]
        public Output<ImmutableDictionary<string, Outputs.OutSection>> Sections { get; private set; } = null!;

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
        /// Create a SocialSecurityNumberItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SocialSecurityNumberItem(string name, SocialSecurityNumberItemArgs args, CustomResourceOptions? options = null)
            : base("onepassword:index:SocialSecurityNumberItem", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private SocialSecurityNumberItem(string name, Input<string> id, SocialSecurityNumberItemState? state = null, CustomResourceOptions? options = null)
            : base("onepassword:index:SocialSecurityNumberItem", name, state, MakeResourceOptions(options, id))
        {
        }

        private static SocialSecurityNumberItemArgs MakeArgs(SocialSecurityNumberItemArgs args)
        {
            args ??= new SocialSecurityNumberItemArgs();
            args.Category = "Social Security Number";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "attachments",
                    "fields",
                    "number",
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
        /// Get an existing SocialSecurityNumberItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SocialSecurityNumberItem Get(string name, Input<string> id, SocialSecurityNumberItemState? state = null, CustomResourceOptions? options = null)
        {
            return new SocialSecurityNumberItem(name, id, state, options);
        }

        public Pulumi.Output<SocialSecurityNumberItemGetAttachmentResult> GetAttachment(SocialSecurityNumberItemGetAttachmentArgs args)
            => Pulumi.Deployment.Instance.Call<SocialSecurityNumberItemGetAttachmentResult>("onepassword:index:SocialSecurityNumberItem/attachment", args ?? new SocialSecurityNumberItemGetAttachmentArgs(), this);
    }

    public sealed class SocialSecurityNumberItemArgs : Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputMap<AssetOrArchive>? _attachments;
        public InputMap<AssetOrArchive> Attachments
        {
            get => _attachments ?? (_attachments = new InputMap<AssetOrArchive>());
            set => _attachments = value;
        }

        /// <summary>
        /// The category of the vault the item is in.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        [Input("fields")]
        private InputMap<Inputs.FieldArgs>? _fields;
        public InputMap<Inputs.FieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputMap<Inputs.FieldArgs>());
            set => _fields = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("number")]
        private Input<string>? _number;
        public Input<string>? Number
        {
            get => _number;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _number = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("sections")]
        private InputMap<Inputs.SectionArgs>? _sections;
        public InputMap<Inputs.SectionArgs> Sections
        {
            get => _sections ?? (_sections = new InputMap<Inputs.SectionArgs>());
            set => _sections = value;
        }

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

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public SocialSecurityNumberItemArgs()
        {
        }
    }

    public sealed class SocialSecurityNumberItemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public SocialSecurityNumberItemState()
        {
        }
    }

    /// <summary>
    /// The set of arguments for the <see cref="SocialSecurityNumberItem.GetAttachment"/> method.
    /// </summary>
    public sealed class SocialSecurityNumberItemGetAttachmentArgs : Pulumi.CallArgs
    {
        /// <summary>
        /// The name or uuid of the attachment to get
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public SocialSecurityNumberItemGetAttachmentArgs()
        {
        }
    }

    /// <summary>
    /// The results of the <see cref="SocialSecurityNumberItem.GetAttachment"/> method.
    /// </summary>
    [OutputType]
    public sealed class SocialSecurityNumberItemGetAttachmentResult
    {
        /// <summary>
        /// the value of the attachment
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private SocialSecurityNumberItemGetAttachmentResult(string value)
        {
            Value = value;
        }
    }
}
