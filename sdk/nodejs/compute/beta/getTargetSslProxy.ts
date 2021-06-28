// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified TargetSslProxy resource. Gets a list of available target SSL proxies by making a list() request.
 */
export function getTargetSslProxy(args: GetTargetSslProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetTargetSslProxyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/beta:getTargetSslProxy", {
        "project": args.project,
        "targetSslProxy": args.targetSslProxy,
    }, opts);
}

export interface GetTargetSslProxyArgs {
    project: string;
    targetSslProxy: string;
}

export interface GetTargetSslProxyResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Type of the resource. Always compute#targetSslProxy for target SSL proxies.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
     */
    readonly proxyHeader: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * URL to the BackendService resource.
     */
    readonly service: string;
    /**
     * URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates.
     */
    readonly sslCertificates: string[];
    /**
     * URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
     */
    readonly sslPolicy: string;
}
