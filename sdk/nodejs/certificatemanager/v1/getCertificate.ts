// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Certificate.
 */
export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:certificatemanager/v1:getCertificate", {
        "certificateId": args.certificateId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetCertificateArgs {
    certificateId: string;
    location: string;
    project?: string;
}

export interface GetCertificateResult {
    /**
     * The creation timestamp of a Certificate.
     */
    readonly createTime: string;
    /**
     * One or more paragraphs of text description of a certificate.
     */
    readonly description: string;
    /**
     * The expiry timestamp of a Certificate.
     */
    readonly expireTime: string;
    /**
     * Set of labels associated with a Certificate.
     */
    readonly labels: {[key: string]: string};
    /**
     * If set, contains configuration and state of a managed certificate.
     */
    readonly managed: outputs.certificatemanager.v1.ManagedCertificateResponse;
    /**
     * A user-defined name of the certificate. Certificate names must be unique globally and match pattern `projects/*&#47;locations/*&#47;certificates/*`.
     */
    readonly name: string;
    /**
     * The PEM-encoded certificate chain.
     */
    readonly pemCertificate: string;
    /**
     * The list of Subject Alternative Names of dnsName type defined in the certificate (see RFC 5280 4.2.1.6)
     */
    readonly sanDnsnames: string[];
    /**
     * Immutable. The scope of the certificate.
     */
    readonly scope: string;
    /**
     * If set, defines data of a self-managed certificate.
     */
    readonly selfManaged: outputs.certificatemanager.v1.SelfManagedCertificateResponse;
    /**
     * The last update timestamp of a Certificate.
     */
    readonly updateTime: string;
}

export function getCertificateOutput(args: GetCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateResult> {
    return pulumi.output(args).apply(a => getCertificate(a, opts))
}

export interface GetCertificateOutputArgs {
    certificateId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
