// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getBankAccount(args: GetBankAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetBankAccountResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("onepassword:index:GetBankAccount", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetBankAccountArgs {
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    uuid?: string;
    /**
     * The UUID of the vault the item is in.
     */
    vault: string;
}

export interface GetBankAccountResult {
    readonly accountNumber?: string;
    readonly bankName?: string;
    readonly branchInformation?: outputs.bankAccount.BranchInformationSection;
    readonly category?: enums.Category | string;
    readonly fields?: {[key: string]: outputs.GetField};
    readonly iban?: string;
    readonly id?: string;
    readonly nameOnAccount?: string;
    readonly notes?: string;
    readonly pin?: string;
    readonly routingNumber?: string;
    readonly sections?: {[key: string]: outputs.GetSection};
    readonly swift?: string;
    /**
     * An array of strings of the tags assigned to the item.
     */
    readonly tags?: string[];
    /**
     * The title of the item.
     */
    readonly title?: string;
    readonly type?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly uuid?: string;
    /**
     * The UUID of the vault the item is in.
     */
    readonly vault?: string;
}

export function getBankAccountOutput(args: GetBankAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBankAccountResult> {
    return pulumi.output(args).apply(a => getBankAccount(a, opts))
}

export interface GetBankAccountOutputArgs {
    /**
     * The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
     */
    title?: pulumi.Input<string>;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    uuid?: pulumi.Input<string>;
    /**
     * The UUID of the vault the item is in.
     */
    vault: pulumi.Input<string>;
}
