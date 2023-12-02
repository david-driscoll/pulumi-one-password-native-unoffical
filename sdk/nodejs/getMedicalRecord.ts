// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

export function getMedicalRecord(args: GetMedicalRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetMedicalRecordResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("one-password-native-unoffical:index:GetMedicalRecord", {
        "title": args.title,
        "uuid": args.uuid,
        "vault": args.vault,
    }, opts);
}

export interface GetMedicalRecordArgs {
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

export interface GetMedicalRecordResult {
    readonly attachments: {[key: string]: outputs.OutAttachment};
    readonly category: enums.Category | string;
    readonly date?: string;
    readonly fields: {[key: string]: outputs.OutField};
    readonly healthcareProfessional?: string;
    readonly location?: string;
    readonly medication?: outputs.medicalRecord.MedicationSection;
    readonly notes?: string;
    readonly patient?: string;
    readonly reasonForVisit?: string;
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

export function getMedicalRecordOutput(args: GetMedicalRecordOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMedicalRecordResult> {
    return pulumi.output(args).apply(a => getMedicalRecord(a, opts))
}

export interface GetMedicalRecordOutputArgs {
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
