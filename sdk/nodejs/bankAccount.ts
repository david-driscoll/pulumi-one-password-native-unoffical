// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export class BankAccount extends pulumi.CustomResource {
    /**
     * Get an existing BankAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BankAccount {
        return new BankAccount(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'onepassword:index:BankAccount';

    /**
     * Returns true if the given object is an instance of BankAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BankAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BankAccount.__pulumiType;
    }

    public readonly accountNumber!: pulumi.Output<string | undefined>;
    public readonly bankName!: pulumi.Output<string | undefined>;
    public readonly branchInformation!: pulumi.Output<outputs.bankAccount.BranchInformation | undefined>;
    public readonly category!: pulumi.Output<enums.Category | string>;
    public readonly fields!: pulumi.Output<outputs.GetField[] | undefined>;
    public readonly iban!: pulumi.Output<string | undefined>;
    public /*out*/ readonly id!: pulumi.Output<string>;
    public readonly nameOnAccount!: pulumi.Output<string | undefined>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public readonly pin!: pulumi.Output<string | undefined>;
    public readonly routingNumber!: pulumi.Output<string | undefined>;
    public readonly sections!: pulumi.Output<outputs.GetSection[] | undefined>;
    public readonly swift!: pulumi.Output<string | undefined>;
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
    /**
     * The UUID of the vault the item is in.
     */
    public readonly vault!: pulumi.Output<string>;

    /**
     * Create a BankAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BankAccountArgs, opts?: pulumi.CustomResourceOptions) {
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
            resourceInputs["accountNumber"] = args ? args.accountNumber : undefined;
            resourceInputs["bankName"] = args ? args.bankName : undefined;
            resourceInputs["branchInformation"] = args ? args.branchInformation : undefined;
            resourceInputs["category"] = (args ? args.category : undefined) ?? "Item";
            resourceInputs["fields"] = args ? args.fields : undefined;
            resourceInputs["iban"] = args ? args.iban : undefined;
            resourceInputs["nameOnAccount"] = args ? args.nameOnAccount : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["pin"] = args ? args.pin : undefined;
            resourceInputs["routingNumber"] = args ? args.routingNumber : undefined;
            resourceInputs["sections"] = args ? args.sections : undefined;
            resourceInputs["swift"] = args ? args.swift : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vault"] = args ? args.vault : undefined;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        } else {
            resourceInputs["accountNumber"] = undefined /*out*/;
            resourceInputs["bankName"] = undefined /*out*/;
            resourceInputs["branchInformation"] = undefined /*out*/;
            resourceInputs["category"] = undefined /*out*/;
            resourceInputs["fields"] = undefined /*out*/;
            resourceInputs["iban"] = undefined /*out*/;
            resourceInputs["id"] = undefined /*out*/;
            resourceInputs["nameOnAccount"] = undefined /*out*/;
            resourceInputs["notes"] = undefined /*out*/;
            resourceInputs["pin"] = undefined /*out*/;
            resourceInputs["routingNumber"] = undefined /*out*/;
            resourceInputs["sections"] = undefined /*out*/;
            resourceInputs["swift"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
            resourceInputs["vault"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BankAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BankAccount resource.
 */
export interface BankAccountArgs {
    accountNumber?: pulumi.Input<string>;
    bankName?: pulumi.Input<string>;
    branchInformation?: pulumi.Input<inputs.bankAccount.BranchInformationArgs>;
    category: pulumi.Input<enums.Category | string>;
    fields?: pulumi.Input<pulumi.Input<inputs.FieldArgs>[]>;
    iban?: pulumi.Input<string>;
    nameOnAccount?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    pin?: pulumi.Input<string>;
    routingNumber?: pulumi.Input<string>;
    sections?: pulumi.Input<pulumi.Input<inputs.SectionArgs>[]>;
    swift?: pulumi.Input<string>;
    /**
     * An array of strings of the tags assigned to the item.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}
