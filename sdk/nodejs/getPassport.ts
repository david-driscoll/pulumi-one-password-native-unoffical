// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getPassport(args: GetPassportArgs, opts?: pulumi.InvokeOptions): Promise<GetPassportResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("onepassword:index:GetPassport", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetPassportArgs {
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

export interface GetPassportResult {
    readonly attachments: {[key: string]: outputs.OutField};
    readonly category: enums.Category | string;
    readonly dateOfBirth?: string;
    readonly expiryDate?: string;
    readonly fields: {[key: string]: outputs.OutField};
    readonly fullName?: string;
    readonly gender?: string;
    readonly issuedOn?: string;
    readonly issuingAuthority?: string;
    readonly issuingCountry?: string;
    readonly nationality?: string;
    readonly notes?: string;
    readonly number?: string;
    readonly placeOfBirth?: string;
    readonly references: {[key: string]: outputs.OutField};
    readonly sections: {[key: string]: outputs.OutSection};
    /**
     * An array of strings of the tags assigned to the item.
     */
    readonly tags: string[];
    /**
     * The title of the item.
     */
    readonly title: string;
    readonly type?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly uuid: string;
    /**
     * The UUID of the vault the item is in.
     */
    readonly vault: string;
}

export function getPassportOutput(args: GetPassportOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPassportResult> {
    return pulumi.output(args).apply(a => getPassport(a, opts))
}

export interface GetPassportOutputArgs {
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
