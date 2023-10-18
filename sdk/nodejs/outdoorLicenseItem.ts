// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class OutdoorLicenseItem extends pulumi.CustomResource {
    /**
     * Get an existing OutdoorLicenseItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OutdoorLicenseItemState, opts?: pulumi.CustomResourceOptions): OutdoorLicenseItem {
        return new OutdoorLicenseItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'onepassword:index:OutdoorLicenseItem';

    /**
     * Returns true if the given object is an instance of OutdoorLicenseItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OutdoorLicenseItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OutdoorLicenseItem.__pulumiType;
    }

    public readonly approvedWildlife!: pulumi.Output<string | undefined>;
    public readonly attachments!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly country!: pulumi.Output<string | undefined>;
    public readonly expires!: pulumi.Output<string | undefined>;
    public readonly fields!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly fullName!: pulumi.Output<string | undefined>;
    public readonly maximumQuota!: pulumi.Output<string | undefined>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly references!: pulumi.Output<{[key: string]: outputs.OutField}>;
    public readonly sections!: pulumi.Output<{[key: string]: outputs.OutSection}>;
    public readonly state!: pulumi.Output<string | undefined>;
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
    public readonly validFrom!: pulumi.Output<string | undefined>;
    /**
     * The UUID of the vault the item is in.
     */
    public readonly vault!: pulumi.Output<string>;

    /**
     * Create a OutdoorLicenseItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OutdoorLicenseItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OutdoorLicenseItemArgs | OutdoorLicenseItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OutdoorLicenseItemState | undefined;
            resourceInputs["vault"] = state ? state.vault : undefined;
        } else {
            const args = argsOrState as OutdoorLicenseItemArgs | undefined;
            if ((!args || args.vault === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vault'");
            }
            resourceInputs["approvedWildlife"] = args ? args.approvedWildlife : undefined;
            resourceInputs["attachments"] = args ? args.attachments : undefined;
            resourceInputs["category"] = "Outdoor License";
            resourceInputs["country"] = args ? args.country : undefined;
            resourceInputs["expires"] = args ? args.expires : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["fullName"] = args ? args.fullName : undefined;
            resourceInputs["maximumQuota"] = args ? args.maximumQuota : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["validFrom"] = args ? args.validFrom : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["references"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["attachments", "fields", "references", "sections"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(OutdoorLicenseItem.__pulumiType, name, resourceInputs, opts);
    }

    getAttachment(args: OutdoorLicenseItem.GetAttachmentArgs): pulumi.Output<OutdoorLicenseItem.GetAttachmentResult> {
        return pulumi.runtime.call("onepassword:index:OutdoorLicenseItem/attachment", {
            "__self__": this,
            "name": args.name,
        }, this);
    }
}

export interface OutdoorLicenseItemState {
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OutdoorLicenseItem resource.
 */
export interface OutdoorLicenseItemArgs {
    approvedWildlife?: pulumi.Input<string>;
    attachments?: pulumi.Input<{[key: string]: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>}>;
    /**
     * The category of the vault the item is in.
     */
    category?: pulumi.Input<"Outdoor License">;
    country?: pulumi.Input<string>;
    expires?: pulumi.Input<string>;
    fields?: pulumi.Input<{[key: string]: pulumi.Input<inputs.FieldArgs>}>;
    fullName?: pulumi.Input<string>;
    maximumQuota?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    sections?: pulumi.Input<{[key: string]: pulumi.Input<inputs.SectionArgs>}>;
    state?: pulumi.Input<string>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    validFrom?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}

export namespace OutdoorLicenseItem {
    /**
     * The set of arguments for the OutdoorLicenseItem.getAttachment method.
     */
    export interface GetAttachmentArgs {
        /**
         * The name or uuid of the attachment to get
         */
        name: pulumi.Input<string>;
    }

    /**
     * The results of the OutdoorLicenseItem.getAttachment method.
     */
    export interface GetAttachmentResult {
        /**
         * the value of the attachment
         */
        readonly value: string;
    }

}
