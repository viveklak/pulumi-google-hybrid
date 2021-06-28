// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified NotificationEndpoint resource in the given region.
 */
export function getRegionNotificationEndpoint(args: GetRegionNotificationEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionNotificationEndpointResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/beta:getRegionNotificationEndpoint", {
        "notificationEndpoint": args.notificationEndpoint,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetRegionNotificationEndpointArgs {
    notificationEndpoint: string;
    project: string;
    region: string;
}

export interface GetRegionNotificationEndpointResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
     */
    readonly grpcSettings: outputs.compute.beta.NotificationEndpointGrpcSettingsResponse;
    /**
     * Type of the resource. Always compute#notificationEndpoint for notification endpoints.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * URL of the region where the notification endpoint resides. This field applies only to the regional resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
}
