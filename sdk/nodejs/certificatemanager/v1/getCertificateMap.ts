// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single CertificateMap.
 */
export function getCertificateMap(args: GetCertificateMapArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateMapResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:certificatemanager/v1:getCertificateMap", {
        "certificateMapId": args.certificateMapId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetCertificateMapArgs {
    certificateMapId: string;
    location: string;
    project?: string;
}

export interface GetCertificateMapResult {
    /**
     * The creation timestamp of a Certificate Map.
     */
    readonly createTime: string;
    /**
     * One or more paragraphs of text description of a certificate map.
     */
    readonly description: string;
    /**
     * A list of GCLB targets which use this Certificate Map.
     */
    readonly gclbTargets: outputs.certificatemanager.v1.GclbTargetResponse[];
    /**
     * Set of labels associated with a Certificate Map.
     */
    readonly labels: {[key: string]: string};
    /**
     * A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match pattern `projects/*&#47;locations/*&#47;certificateMaps/*`.
     */
    readonly name: string;
    /**
     * The update timestamp of a Certificate Map.
     */
    readonly updateTime: string;
}

export function getCertificateMapOutput(args: GetCertificateMapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateMapResult> {
    return pulumi.output(args).apply(a => getCertificateMap(a, opts))
}

export interface GetCertificateMapOutputArgs {
    certificateMapId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
