// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified BackendBucket resource. Gets a list of available backend buckets by making a list() request.
 */
export function getBackendBucket(args: GetBackendBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetBackendBucketResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/beta:getBackendBucket", {
        "backendBucket": args.backendBucket,
        "project": args.project,
    }, opts);
}

export interface GetBackendBucketArgs {
    backendBucket: string;
    project?: string;
}

export interface GetBackendBucketResult {
    /**
     * Cloud Storage bucket name.
     */
    readonly bucketName: string;
    /**
     * Cloud CDN configuration for this BackendBucket.
     */
    readonly cdnPolicy: outputs.compute.beta.BackendBucketCdnPolicyResponse;
    /**
     * Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
     */
    readonly compressionMode: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * Headers that the HTTP/S load balancer should add to proxied responses.
     */
    readonly customResponseHeaders: string[];
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    readonly description: string;
    /**
     * The resource URL for the edge security policy associated with this backend bucket.
     */
    readonly edgeSecurityPolicy: string;
    /**
     * If true, enable Cloud CDN for this BackendBucket.
     */
    readonly enableCdn: boolean;
    /**
     * Type of the resource.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
}

export function getBackendBucketOutput(args: GetBackendBucketOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackendBucketResult> {
    return pulumi.output(args).apply(a => getBackendBucket(a, opts))
}

export interface GetBackendBucketOutputArgs {
    backendBucket: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
