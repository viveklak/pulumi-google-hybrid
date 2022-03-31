// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns metadata for a given ImportJob.
 */
export function getImportJob(args: GetImportJobArgs, opts?: pulumi.InvokeOptions): Promise<GetImportJobResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:cloudkms/v1:getImportJob", {
        "importJobId": args.importJobId,
        "keyRingId": args.keyRingId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetImportJobArgs {
    importJobId: string;
    keyRingId: string;
    location: string;
    project?: string;
}

export interface GetImportJobResult {
    /**
     * Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen ImportMethod is one with a protection level of HSM.
     */
    readonly attestation: outputs.cloudkms.v1.KeyOperationAttestationResponse;
    /**
     * The time at which this ImportJob was created.
     */
    readonly createTime: string;
    /**
     * The time this ImportJob expired. Only present if state is EXPIRED.
     */
    readonly expireEventTime: string;
    /**
     * The time at which this ImportJob is scheduled for expiration and can no longer be used to import key material.
     */
    readonly expireTime: string;
    /**
     * The time this ImportJob's key material was generated.
     */
    readonly generateTime: string;
    /**
     * Immutable. The wrapping method to be used for incoming key material.
     */
    readonly importMethod: string;
    /**
     * The resource name for this ImportJob in the format `projects/*&#47;locations/*&#47;keyRings/*&#47;importJobs/*`.
     */
    readonly name: string;
    /**
     * Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
     */
    readonly protectionLevel: string;
    /**
     * The public key with which to wrap key material prior to import. Only returned if state is ACTIVE.
     */
    readonly publicKey: outputs.cloudkms.v1.WrappingPublicKeyResponse;
    /**
     * The current state of the ImportJob, indicating if it can be used.
     */
    readonly state: string;
}

export function getImportJobOutput(args: GetImportJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImportJobResult> {
    return pulumi.output(args).apply(a => getImportJob(a, opts))
}

export interface GetImportJobOutputArgs {
    importJobId: pulumi.Input<string>;
    keyRingId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
