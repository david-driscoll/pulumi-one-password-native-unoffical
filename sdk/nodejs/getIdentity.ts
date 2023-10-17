// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getIdentity(args: GetIdentityArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("onepassword:index:GetIdentity", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetIdentityArgs {
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

export interface GetIdentityResult {
    readonly address?: outputs.identity.AddressSection;
    readonly attachments: {[key: string]: outputs.OutField};
    readonly category: enums.Category | string;
    readonly fields: {[key: string]: outputs.OutField};
    readonly identification?: outputs.identity.IdentificationSection;
    readonly internetDetails?: outputs.identity.InternetDetailsSection;
    readonly notes?: string;
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
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly uuid: string;
    /**
     * The UUID of the vault the item is in.
     */
    readonly vault: string;
}

export function getIdentityOutput(args: GetIdentityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdentityResult> {
    return pulumi.output(args).apply(a => getIdentity(a, opts))
}

export interface GetIdentityOutputArgs {
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
