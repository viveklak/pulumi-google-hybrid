// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the specified regional HealthCheckService resource.
 */
export function getRegionHealthCheckService(args: GetRegionHealthCheckServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionHealthCheckServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:compute/beta:getRegionHealthCheckService", {
        "healthCheckService": args.healthCheckService,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetRegionHealthCheckServiceArgs {
    healthCheckService: string;
    project?: string;
    region: string;
}

export interface GetRegionHealthCheckServiceResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a HealthCheckService. An up-to-date fingerprint must be provided in order to patch/update the HealthCheckService; Otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the HealthCheckService.
     */
    readonly fingerprint: string;
    /**
     * A list of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT or portSpecification=USE_FIXED_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks must belong to the same region as zones of NEGs.
     */
    readonly healthChecks: string[];
    /**
     * Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified. - NO_AGGREGATION. An EndpointHealth message is returned for each pair in the health check service. - AND. If any health check of an endpoint reports UNHEALTHY, then UNHEALTHY is the HealthState of the endpoint. If all health checks report HEALTHY, the HealthState of the endpoint is HEALTHY. .
     */
    readonly healthStatusAggregationPolicy: string;
    /**
     * [Output only] Type of the resource. Always compute#healthCheckServicefor health check services.
     */
    readonly kind: string;
    /**
     * Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * A list of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
     */
    readonly networkEndpointGroups: string[];
    /**
     * A list of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
     */
    readonly notificationEndpoints: string[];
    /**
     * URL of the region where the health check service resides. This field is not applicable to global health check services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
}

export function getRegionHealthCheckServiceOutput(args: GetRegionHealthCheckServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionHealthCheckServiceResult> {
    return pulumi.output(args).apply(a => getRegionHealthCheckService(a, opts))
}

export interface GetRegionHealthCheckServiceOutputArgs {
    healthCheckService: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
