// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Rocket.Surgery.OnePasswordNativeUnofficial
{
    [OnePasswordNativeUnofficialResourceType("one-password-native-unofficial:index:DatabaseItem")]
    public partial class DatabaseItem : Pulumi.CustomResource
    {
        [Output("alias")]
        public Output<string?> Alias { get; private set; } = null!;

        [Output("attachments")]
        public Output<ImmutableDictionary<string, Outputs.OutputAttachment>> Attachments { get; private set; } = null!;

        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        [Output("connectionOptions")]
        public Output<string?> ConnectionOptions { get; private set; } = null!;

        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

        [Output("fields")]
        public Output<ImmutableDictionary<string, Outputs.OutputField>> Fields { get; private set; } = null!;

        /// <summary>
        /// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        /// </summary>
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("port")]
        public Output<string?> Port { get; private set; } = null!;

        [Output("references")]
        public Output<ImmutableArray<Outputs.OutputReference>> References { get; private set; } = null!;

        [Output("sections")]
        public Output<ImmutableDictionary<string, Outputs.OutputSection>> Sections { get; private set; } = null!;

        [Output("server")]
        public Output<string?> Server { get; private set; } = null!;

        [Output("sid")]
        public Output<string?> Sid { get; private set; } = null!;

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

        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("urls")]
        public Output<ImmutableArray<Outputs.OutputUrl>> Urls { get; private set; } = null!;

        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        [Output("vault")]
        public Output<ImmutableDictionary<string, string>> Vault { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseItem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseItem(string name, DatabaseItemArgs? args = null, CustomResourceOptions? options = null)
            : base("one-password-native-unofficial:index:DatabaseItem", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseItem(string name, Input<string> id, DatabaseItemState? state = null, CustomResourceOptions? options = null)
            : base("one-password-native-unofficial:index:DatabaseItem", name, state, MakeResourceOptions(options, id))
        {
        }

        private static DatabaseItemArgs? MakeArgs(DatabaseItemArgs? args)
        {
            args ??= new DatabaseItemArgs();
            args.Category = "Database";
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
                    "password",
                    "sections",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseItem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseItem Get(string name, Input<string> id, DatabaseItemState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseItem(name, id, state, options);
        }
    }

    public sealed class DatabaseItemArgs : Pulumi.ResourceArgs
    {
        [Input("alias")]
        public Input<string>? Alias { get; set; }

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

        [Input("connectionOptions")]
        public Input<string>? ConnectionOptions { get; set; }

        [Input("database")]
        public Input<string>? Database { get; set; }

        [Input("fields")]
        private InputMap<Union<Inputs.FieldArgs, string>>? _fields;
        public InputMap<Union<Inputs.FieldArgs, string>> Fields
        {
            get => _fields ?? (_fields = new InputMap<Union<Inputs.FieldArgs, string>>());
            set => _fields = value;
        }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("references")]
        private InputList<string>? _references;
        public InputList<string> References
        {
            get => _references ?? (_references = new InputList<string>());
            set => _references = value;
        }

        [Input("sections")]
        private InputMap<Inputs.SectionArgs>? _sections;
        public InputMap<Inputs.SectionArgs> Sections
        {
            get => _sections ?? (_sections = new InputMap<Inputs.SectionArgs>());
            set => _sections = value;
        }

        [Input("server")]
        public Input<string>? Server { get; set; }

        [Input("sid")]
        public Input<string>? Sid { get; set; }

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

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("urls")]
        private InputList<Union<Inputs.UrlArgs, string>>? _urls;
        public InputList<Union<Inputs.UrlArgs, string>> Urls
        {
            get => _urls ?? (_urls = new InputList<Union<Inputs.UrlArgs, string>>());
            set => _urls = value;
        }

        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault")]
        public Input<string>? Vault { get; set; }

        public DatabaseItemArgs()
        {
        }
    }

    public sealed class DatabaseItemState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the vault the item is in.
        /// </summary>
        [Input("vault", required: true)]
        public Input<string> Vault { get; set; } = null!;

        public DatabaseItemState()
        {
        }
    }
}
