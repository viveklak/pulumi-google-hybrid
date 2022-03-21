// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the specified HttpsHealthCheck resource. Gets a list of available HTTPS health checks by making a list() request.
 */
export function getHttpsHealthCheck(args: GetHttpsHealthCheckArgs, opts?: pulumi.InvokeOptions): Promise<GetHttpsHealthCheckResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/beta:getHttpsHealthCheck", {
        "httpsHealthCheck": args.httpsHealthCheck,
        "project": args.project,
    }, opts);
}

export interface GetHttpsHealthCheckArgs {
    httpsHealthCheck: string;
    project?: string;
}

export interface GetHttpsHealthCheckResult {
    /**
     * How often (in seconds) to send a health check. The default value is 5 seconds.
     */
    readonly checkIntervalSec: number;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
     */
    readonly healthyThreshold: number;
    /**
     * The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
     */
    readonly host: string;
    /**
     * Type of the resource.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The TCP port number for the HTTPS health check request. The default value is 443.
     */
    readonly port: number;
    /**
     * The request path of the HTTPS health check request. The default value is "/".
     */
    readonly requestPath: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have a greater value than checkIntervalSec.
     */
    readonly timeoutSec: number;
    /**
     * A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
     */
    readonly unhealthyThreshold: number;
}

export function getHttpsHealthCheckOutput(args: GetHttpsHealthCheckOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHttpsHealthCheckResult> {
    return pulumi.output(args).apply(a => getHttpsHealthCheck(a, opts))
}

export interface GetHttpsHealthCheckOutputArgs {
    httpsHealthCheck: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
