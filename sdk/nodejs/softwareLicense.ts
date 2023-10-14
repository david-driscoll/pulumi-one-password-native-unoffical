// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class SoftwareLicense extends pulumi.CustomResource {
    /**
     * Get an existing SoftwareLicense resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SoftwareLicense {
        return new SoftwareLicense(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'onepassword:index:SoftwareLicense';

    /**
     * Returns true if the given object is an instance of SoftwareLicense.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SoftwareLicense {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SoftwareLicense.__pulumiType;
    }

    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly customer!: pulumi.Output<outputs.softwareLicense.Customer | undefined>;
    public readonly fields!: pulumi.Output<outputs.GetField[] | undefined>;
    public /*out*/ readonly id!: pulumi.Output<string>;
    public readonly licenseKey!: pulumi.Output<string | undefined>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public readonly order!: pulumi.Output<outputs.softwareLicense.Order | undefined>;
    public readonly publisher!: pulumi.Output<outputs.softwareLicense.Publisher | undefined>;
    public readonly sections!: pulumi.Output<outputs.GetSection[] | undefined>;
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
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a SoftwareLicense resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SoftwareLicenseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.category === undefined) && !opts.urn) {
                throw new Error("Missing required property 'category'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            if ((!args || args.vault === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vault'");
            }
            resourceInputs["category"] = (args ? args.category : undefined) ?? "Item";
            resourceInputs["customer"] = args ? args.customer : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["licenseKey"] = args ? args.licenseKey : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["publisher"] = args ? args.publisher : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        } else {
            resourceInputs["category"] = undefined /*out*/;
            resourceInputs["customer"] = undefined /*out*/;
            resourceInputs["fields"] = undefined /*out*/;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["licenseKey"] = undefined /*out*/;
            resourceInputs["notes"] = undefined /*out*/;
            resourceInputs["order"] = undefined /*out*/;
            resourceInputs["publisher"] = undefined /*out*/;
            resourceInputs["sections"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
            resourceInputs["vault"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SoftwareLicense.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SoftwareLicense resource.
 */
export interface SoftwareLicenseArgs {
    category: pulumi.Input<enums.Category | string>;
    customer?: pulumi.Input<inputs.softwareLicense.CustomerArgs>;
    fields?: pulumi.Input<pulumi.Input<inputs.FieldArgs>[]>;
    licenseKey?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    order?: pulumi.Input<inputs.softwareLicense.OrderArgs>;
    publisher?: pulumi.Input<inputs.softwareLicense.PublisherArgs>;
    sections?: pulumi.Input<pulumi.Input<inputs.SectionArgs>[]>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}
