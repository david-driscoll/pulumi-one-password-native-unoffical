// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getRewardProgram(args: GetRewardProgramArgs, opts?: pulumi.InvokeOptions): Promise<GetRewardProgramResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("onepassword:index:GetRewardProgram", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetRewardProgramArgs {
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

export interface GetRewardProgramResult {
    readonly category?: enums.Category | string;
    readonly companyName?: string;
    readonly fields?: {[key: string]: outputs.GetField};
    readonly id?: string;
    readonly memberId?: string;
    readonly memberName?: string;
    readonly moreInformation?: outputs.rewardProgram.MoreInformationSection;
    readonly notes?: string;
    readonly pin?: string;
    readonly sections?: {[key: string]: outputs.GetSection};
    /**
     * An array of strings of the tags assigned to the item.
     */
    readonly tags?: string[];
    /**
     * The title of the item.
     */
    readonly title?: string;
    /**
     * The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
     */
    readonly uuid?: string;
    /**
     * The UUID of the vault the item is in.
     */
    readonly vault?: string;
}

export function getRewardProgramOutput(args: GetRewardProgramOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRewardProgramResult> {
    return pulumi.output(args).apply(a => getRewardProgram(a, opts))
}

export interface GetRewardProgramOutputArgs {
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
