// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class DatabaseItem extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseItemState, opts?: pulumi.CustomResourceOptions): DatabaseItem {
        return new DatabaseItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'one-password-native-unoffical:index:DatabaseItem';

    /**
     * Returns true if the given object is an instance of DatabaseItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseItem.__pulumiType;
    }

    public readonly alias!: pulumi.Output<string | undefined>;
    public /*out*/ readonly attachments!: pulumi.Output<{[key: string]: outputs.OutAttachment}>;
    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly connectionOptions!: pulumi.Output<string | undefined>;
    public readonly database!: pulumi.Output<string | undefined>;
    public readonly fields!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<string | undefined>;
    public /*out*/ readonly references!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly sections!: pulumi.Output<{[key: string]: outputs.OutSection}>;
    public readonly server!: pulumi.Output<string | undefined>;
    public readonly sid!: pulumi.Output<string | undefined>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The title of the item.
     */
    public readonly title!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly username!: pulumi.Output<string | undefined>;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;
    /**
     * The UUID of the vault the item is in.
     */
    public readonly vault!: pulumi.Output<string>;

    /**
     * Create a DatabaseItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseItemArgs | DatabaseItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseItemState | undefined;
            resourceInputs["vault"] = state ? state.vault : undefined;
        } else {
            const args = argsOrState as DatabaseItemArgs | undefined;
            if ((!args || args.vault === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vault'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["category"] = "Database";
            resourceInputs["connectionOptions"] = args ? args.connectionOptions : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["inputAttachments"] = args ? args.inputAttachments : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["sid"] = args ? args.sid : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["attachments"] = undefined /*out*/;
            resourceInputs["references"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["attachments", "fields", "password", "references", "sections"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DatabaseItem.__pulumiType, name, resourceInputs, opts);
    }

    getAttachment(args: DatabaseItem.GetAttachmentArgs): pulumi.Output<DatabaseItem.GetAttachmentResult> {
        return pulumi.runtime.call("one-password-native-unoffical:index:DatabaseItem/attachment", {
            "__self__": this,
            "name": args.name,
        }, this);
    }
}

export interface DatabaseItemState {
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseItem resource.
 */
export interface DatabaseItemArgs {
    alias?: pulumi.Input<string>;
    /**
     * The category of the vault the item is in.
     */
    category?: pulumi.Input<"Database">;
    connectionOptions?: pulumi.Input<string>;
    database?: pulumi.Input<string>;
    fields?: pulumi.Input<{[key: string]: pulumi.Input<inputs.FieldArgs>}>;
    inputAttachments?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>}>;
    notes?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    port?: pulumi.Input<string>;
    sections?: pulumi.Input<{[key: string]: pulumi.Input<inputs.SectionArgs>}>;
    server?: pulumi.Input<string>;
    sid?: pulumi.Input<string>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

export namespace DatabaseItem {
    /**
     * The set of arguments for the DatabaseItem.getAttachment method.
     */
    export interface GetAttachmentArgs {
        /**
         * The name or uuid of the attachment to get
         */
        name: pulumi.Input<string>;
    }

    /**
     * The results of the DatabaseItem.getAttachment method.
     */
    export interface GetAttachmentResult {
        /**
         * the value of the attachment
         */
        readonly value: string;
    }

}
