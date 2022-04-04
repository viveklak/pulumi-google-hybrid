// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single GrpcRoute.
 */
export function getGrpcRoute(args: GetGrpcRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetGrpcRouteResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-hybrid:networkservices/v1beta1:getGrpcRoute", {
        "grpcRouteId": args.grpcRouteId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetGrpcRouteArgs {
    grpcRouteId: string;
    location: string;
    project?: string;
}

export interface GetGrpcRouteResult {
    /**
     * The timestamp when the resource was created.
     */
    readonly createTime: string;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    readonly description: string;
    /**
     * Optional. Gateways defines a list of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*&#47;locations/global/gateways/`
     */
    readonly gateways: string[];
    /**
     * Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. "foo.example.com") or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateway must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service").
     */
    readonly hostnames: string[];
    /**
     * Optional. Set of label tags associated with the GrpcRoute resource.
     */
    readonly labels: {[key: string]: string};
    /**
     * Optional. Meshes defines a list of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*&#47;locations/global/meshes/`
     */
    readonly meshes: string[];
    /**
     * Name of the GrpcRoute resource. It matches pattern `projects/*&#47;locations/global/grpcRoutes/`
     */
    readonly name: string;
    /**
     * A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied.
     */
    readonly rules: outputs.networkservices.v1beta1.GrpcRouteRouteRuleResponse[];
    /**
     * Server-defined URL of this resource
     */
    readonly selfLink: string;
    /**
     * The timestamp when the resource was updated.
     */
    readonly updateTime: string;
}

export function getGrpcRouteOutput(args: GetGrpcRouteOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGrpcRouteResult> {
    return pulumi.output(args).apply(a => getGrpcRoute(a, opts))
}

export interface GetGrpcRouteOutputArgs {
    grpcRouteId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}