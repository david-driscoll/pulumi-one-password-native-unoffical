// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class CreditCard extends pulumi.CustomResource {
    /**
     * Get an existing CreditCard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CreditCard {
        return new CreditCard(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'onepassword:index:CreditCard';

    /**
     * Returns true if the given object is an instance of CreditCard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CreditCard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CreditCard.__pulumiType;
    }

    public readonly additionalDetails!: pulumi.Output<outputs.creditCard.AdditionalDetails | undefined>;
    public readonly cardholderName!: pulumi.Output<string | undefined>;
    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly contactInformation!: pulumi.Output<outputs.creditCard.ContactInformation | undefined>;
    public readonly expiryDate!: pulumi.Output<string | undefined>;
    public readonly fields!: pulumi.Output<outputs.GetField[] | undefined>;
    public /*out*/ readonly id!: pulumi.Output<string>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public readonly number!: pulumi.Output<string | undefined>;
    public readonly sections!: pulumi.Output<outputs.GetSection[] | undefined>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The title of the item.
     */
    public readonly title!: pulumi.Output<string>;
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;
    public readonly validFrom!: pulumi.Output<string | undefined>;
    /**
     * The UUID of the vault the item is in.
     */
    public readonly vault!: pulumi.Output<string>;
    public readonly verificationNumber!: pulumi.Output<string | undefined>;

    /**
     * Create a CreditCard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CreditCardArgs, opts?: pulumi.CustomResourceOptions) {
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
            resourceInputs["additionalDetails"] = args ? args.additionalDetails : undefined;
            resourceInputs["cardholderName"] = args ? args.cardholderName : undefined;
            resourceInputs["category"] = (args ? args.category : undefined) ?? "Item";
            resourceInputs["contactInformation"] = args ? args.contactInformation : undefined;
            resourceInputs["expiryDate"] = args ? args.expiryDate : undefined;
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["number"] = args ? args.number : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["validFrom"] = args ? args.validFrom : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["verificationNumber"] = args ? args.verificationNumber : undefined;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        } else {
            resourceInputs["additionalDetails"] = undefined /*out*/;
            resourceInputs["cardholderName"] = undefined /*out*/;
            resourceInputs["category"] = undefined /*out*/;
            resourceInputs["contactInformation"] = undefined /*out*/;
            resourceInputs["expiryDate"] = undefined /*out*/;
            resourceInputs["fields"] = undefined /*out*/;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["notes"] = undefined /*out*/;
            resourceInputs["number"] = undefined /*out*/;
            resourceInputs["sections"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
            resourceInputs["validFrom"] = undefined /*out*/;
            resourceInputs["vault"] = undefined /*out*/;
            resourceInputs["verificationNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CreditCard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CreditCard resource.
 */
export interface CreditCardArgs {
    additionalDetails?: pulumi.Input<inputs.creditCard.AdditionalDetailsArgs>;
    cardholderName?: pulumi.Input<string>;
    category: pulumi.Input<enums.Category | string>;
    contactInformation?: pulumi.Input<inputs.creditCard.ContactInformationArgs>;
    expiryDate?: pulumi.Input<string>;
    fields?: pulumi.Input<pulumi.Input<inputs.FieldArgs>[]>;
    notes?: pulumi.Input<string>;
    number?: pulumi.Input<string>;
    sections?: pulumi.Input<pulumi.Input<inputs.SectionArgs>[]>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    validFrom?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
    verificationNumber?: pulumi.Input<string>;
}
