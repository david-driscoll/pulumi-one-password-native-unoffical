// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAttachment(args: GetAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetAttachmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("one-password-native:index:GetAttachment", {
        "reference": args.reference,
    }, opts);
}

export interface GetAttachmentArgs {
    /**
     * The 1Password secret reference path to the item.  eg: op://vault/item/[section]/file 
     */
    reference: string;
}

/**
 * The resolved reference value
 */
export interface GetAttachmentResult {
    readonly value?: string;
}

export function getAttachmentOutput(args: GetAttachmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAttachmentResult> {
    return pulumi.output(args).apply(a => getAttachment(a, opts))
}

export interface GetAttachmentOutputArgs {
    /**
     * The 1Password secret reference path to the item.  eg: op://vault/item/[section]/file 
     */
    reference: pulumi.Input<string>;
}
