// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Retrieves a particular SSL certificate. Does not include the private key (required for usage). The private key must be saved from the response to initial creation.
 */
export function getSslCert(args: GetSslCertArgs, opts?: pulumi.InvokeOptions): Promise<GetSslCertResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:sqladmin/v1beta4:getSslCert", {
        "instance": args.instance,
        "project": args.project,
        "sha1Fingerprint": args.sha1Fingerprint,
    }, opts);
}

export interface GetSslCertArgs {
    instance: string;
    project?: string;
    sha1Fingerprint: string;
}

export interface GetSslCertResult {
    /**
     * PEM representation.
     */
    readonly cert: string;
    /**
     * Serial number, as extracted from the certificate.
     */
    readonly certSerialNumber: string;
    /**
     * User supplied name. Constrained to [a-zA-Z.-_ ]+.
     */
    readonly commonName: string;
    /**
     * The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example **2012-11-15T16:19:00.094Z**.
     */
    readonly createTime: string;
    /**
     * The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example **2012-11-15T16:19:00.094Z**.
     */
    readonly expirationTime: string;
    /**
     * Name of the database instance.
     */
    readonly instance: string;
    /**
     * This is always **sql#sslCert**.
     */
    readonly kind: string;
    /**
     * The URI of this resource.
     */
    readonly selfLink: string;
    /**
     * Sha1 Fingerprint.
     */
    readonly sha1Fingerprint: string;
}

export function getSslCertOutput(args: GetSslCertOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSslCertResult> {
    return pulumi.output(args).apply(a => getSslCert(a, opts))
}

export interface GetSslCertOutputArgs {
    instance: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sha1Fingerprint: pulumi.Input<string>;
}
