// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets an attestor. Returns NOT_FOUND if the attestor does not exist.
 */
export function getAttestor(args: GetAttestorArgs, opts?: pulumi.InvokeOptions): Promise<GetAttestorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:binaryauthorization/v1beta1:getAttestor", {
        "attestorId": args.attestorId,
        "project": args.project,
    }, opts);
}

export interface GetAttestorArgs {
    attestorId: string;
    project: string;
}

export interface GetAttestorResult {
    /**
     * Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
     */
    readonly description: string;
    /**
     * The resource name, in the format: `projects/*&#47;attestors/*`. This field may not be updated.
     */
    readonly name: string;
    /**
     * Time when the attestor was last updated.
     */
    readonly updateTime: string;
    /**
     * A Drydock ATTESTATION_AUTHORITY Note, created by the user.
     */
    readonly userOwnedDrydockNote: outputs.binaryauthorization.v1beta1.UserOwnedDrydockNoteResponse;
}
