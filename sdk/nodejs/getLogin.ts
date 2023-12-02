// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getLogin(args: GetLoginArgs, opts?: pulumi.InvokeOptions): Promise<GetLoginResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("one-password-native-unoffical:index:GetLogin", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetLoginArgs {
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

export interface GetLoginResult {
    readonly attachments: {[key: string]: outputs.OutAttachment};
    readonly category: enums.Category | string;
    readonly fields: {[key: string]: outputs.OutField};
    readonly notes?: string;
    readonly password?: string;
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
    readonly username?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly uuid: string;
    /**
     * The UUID of the vault the item is in.
     */
    readonly vault: string;
}

export function getLoginOutput(args: GetLoginOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoginResult> {
    return pulumi.output(args).apply(a => getLogin(a, opts))
}

export interface GetLoginOutputArgs {
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
