// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the specified Route resource. Gets a list of available routes by making a list() request.
 */
export function getRoute(args: GetRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/v1:getRoute", {
        "project": args.project,
        "route": args.route,
    }, opts);
}

export interface GetRouteArgs {
    project: string;
    route: string;
}

export interface GetRouteResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    readonly description: string;
    /**
     * The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
     */
    readonly destRange: string;
    /**
     * Type of this resource. Always compute#routes for Route resources.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    readonly name: string;
    /**
     * Fully-qualified URL of the network that this route applies to.
     */
    readonly network: string;
    /**
     * The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
     */
    readonly nextHopGateway: string;
    /**
     * The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:  
     * - 10.128.0.56 
     * - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule 
     * - regions/region/forwardingRules/forwardingRule
     */
    readonly nextHopIlb: string;
    /**
     * The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
     * https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
     */
    readonly nextHopInstance: string;
    /**
     * The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
     */
    readonly nextHopIp: string;
    /**
     * The URL of the local network if it should handle matching packets.
     */
    readonly nextHopNetwork: string;
    /**
     * The network peering name that should handle matching packets, which should conform to RFC1035.
     */
    readonly nextHopPeering: string;
    /**
     * The URL to a VpnTunnel that should handle matching packets.
     */
    readonly nextHopVpnTunnel: string;
    /**
     * The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
     */
    readonly priority: number;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    readonly selfLink: string;
    /**
     * A list of instance tags to which this route applies.
     */
    readonly tags: string[];
    /**
     * If potential misconfigurations are detected for this route, this field will be populated with warning messages.
     */
    readonly warnings: outputs.compute.v1.RouteWarningsItemResponse[];
}
