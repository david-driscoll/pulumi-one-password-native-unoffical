// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class SocialSecurityNumberItem extends pulumi.CustomResource {
    /**
     * Get an existing SocialSecurityNumberItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SocialSecurityNumberItemState, opts?: pulumi.CustomResourceOptions): SocialSecurityNumberItem {
        return new SocialSecurityNumberItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'onepassword:index:SocialSecurityNumberItem';

    /**
     * Returns true if the given object is an instance of SocialSecurityNumberItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SocialSecurityNumberItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SocialSecurityNumberItem.__pulumiType;
    }

    public readonly attachments!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly fields!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public readonly number!: pulumi.Output<string | undefined>;
    public /*out*/ readonly references!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly sections!: pulumi.Output<{[key: string]: outputs.OutSection}>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The title of the item.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;
    /**
     * The UUID of the vault the item is in.
     */
    public readonly vault!: pulumi.Output<string>;

    /**
     * Create a SocialSecurityNumberItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SocialSecurityNumberItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SocialSecurityNumberItemArgs | SocialSecurityNumberItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SocialSecurityNumberItemState | undefined;
            resourceInputs["vault"] = state ? state.vault : undefined;
        } else {
            const args = argsOrState as SocialSecurityNumberItemArgs | undefined;
            if ((!args || args.vault === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vault'");
            }
            resourceInputs["attachments"] = args ? args.attachments : undefined;
            resourceInputs["category"] = "Social Security Number";
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["number"] = args?.number ? pulumi.secret(args.number) : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["references"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["attachments", "fields", "number", "references", "sections"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SocialSecurityNumberItem.__pulumiType, name, resourceInputs, opts);
    }

    getAttachment(args: SocialSecurityNumberItem.GetAttachmentArgs): pulumi.Output<SocialSecurityNumberItem.GetAttachmentResult> {
        return pulumi.runtime.call("onepassword:index:SocialSecurityNumberItem/attachment", {
            "__self__": this,
            "name": args.name,
        }, this);
    }
}

export interface SocialSecurityNumberItemState {
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SocialSecurityNumberItem resource.
 */
export interface SocialSecurityNumberItemArgs {
    attachments?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>}>;
    /**
     * The category of the vault the item is in.
     */
    category?: pulumi.Input<"Social Security Number">;
    fields?: pulumi.Input<{[key: string]: pulumi.Input<inputs.FieldArgs>}>;
    name?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    number?: pulumi.Input<string>;
    sections?: pulumi.Input<{[key: string]: pulumi.Input<inputs.SectionArgs>}>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

export namespace SocialSecurityNumberItem {
    /**
     * The set of arguments for the SocialSecurityNumberItem.getAttachment method.
     */
    export interface GetAttachmentArgs {
        /**
         * The name or uuid of the attachment to get
         */
        name: pulumi.Input<string>;
    }

    /**
     * The results of the SocialSecurityNumberItem.getAttachment method.
     */
    export interface GetAttachmentResult {
        /**
         * the value of the attachment
         */
        readonly value: string;
    }

}
