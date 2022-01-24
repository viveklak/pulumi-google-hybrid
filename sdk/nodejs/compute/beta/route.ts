// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a Route resource in the specified project using the data included in the request.
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * AS path.
     */
    public /*out*/ readonly asPaths!: pulumi.Output<outputs.compute.beta.RouteAsPathResponse[]>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
     */
    public readonly destRange!: pulumi.Output<string>;
    /**
     * Type of this resource. Always compute#routes for Route resources.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Fully-qualified URL of the network that this route applies to.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL: projects/ project/global/gateways/default-internet-gateway
     */
    public readonly nextHopGateway!: pulumi.Output<string>;
    /**
     * The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs: - 10.128.0.56 - https://www.googleapis.com/compute/v1/projects/project/regions/region /forwardingRules/forwardingRule - regions/region/forwardingRules/forwardingRule 
     */
    public readonly nextHopIlb!: pulumi.Output<string>;
    /**
     * The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
     */
    public readonly nextHopInstance!: pulumi.Output<string>;
    /**
     * The URL to an InterconnectAttachment which is the next hop for the route. This field will only be populated for the dynamic routes generated by Cloud Router with a linked interconnectAttachment.
     */
    public /*out*/ readonly nextHopInterconnectAttachment!: pulumi.Output<string>;
    /**
     * The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
     */
    public readonly nextHopIp!: pulumi.Output<string>;
    /**
     * The URL of the local network if it should handle matching packets.
     */
    public readonly nextHopNetwork!: pulumi.Output<string>;
    /**
     * The network peering name that should handle matching packets, which should conform to RFC1035.
     */
    public /*out*/ readonly nextHopPeering!: pulumi.Output<string>;
    /**
     * The URL to a VpnTunnel that should handle matching packets.
     */
    public readonly nextHopVpnTunnel!: pulumi.Output<string>;
    /**
     * The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The type of this route, which can be one of the following values: - 'TRANSIT' for a transit route that this router learned from another Cloud Router and will readvertise to one of its BGP peers - 'SUBNET' for a route from a subnet of the VPC - 'BGP' for a route learned from a BGP peer of this router - 'STATIC' for a static route
     */
    public /*out*/ readonly routeType!: pulumi.Output<string>;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * A list of instance tags to which this route applies.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * If potential misconfigurations are detected for this route, this field will be populated with warning messages.
     */
    public /*out*/ readonly warnings!: pulumi.Output<outputs.compute.beta.RouteWarningsItemResponse[]>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouteArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destRange"] = args ? args.destRange : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["nextHopGateway"] = args ? args.nextHopGateway : undefined;
            resourceInputs["nextHopIlb"] = args ? args.nextHopIlb : undefined;
            resourceInputs["nextHopInstance"] = args ? args.nextHopInstance : undefined;
            resourceInputs["nextHopIp"] = args ? args.nextHopIp : undefined;
            resourceInputs["nextHopNetwork"] = args ? args.nextHopNetwork : undefined;
            resourceInputs["nextHopVpnTunnel"] = args ? args.nextHopVpnTunnel : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["asPaths"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["nextHopInterconnectAttachment"] = undefined /*out*/;
            resourceInputs["nextHopPeering"] = undefined /*out*/;
            resourceInputs["routeType"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["warnings"] = undefined /*out*/;
        } else {
            resourceInputs["asPaths"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["destRange"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["nextHopGateway"] = undefined /*out*/;
            resourceInputs["nextHopIlb"] = undefined /*out*/;
            resourceInputs["nextHopInstance"] = undefined /*out*/;
            resourceInputs["nextHopInterconnectAttachment"] = undefined /*out*/;
            resourceInputs["nextHopIp"] = undefined /*out*/;
            resourceInputs["nextHopNetwork"] = undefined /*out*/;
            resourceInputs["nextHopPeering"] = undefined /*out*/;
            resourceInputs["nextHopVpnTunnel"] = undefined /*out*/;
            resourceInputs["priority"] = undefined /*out*/;
            resourceInputs["routeType"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["warnings"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Route.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
     */
    destRange?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    name?: pulumi.Input<string>;
    /**
     * Fully-qualified URL of the network that this route applies to.
     */
    network?: pulumi.Input<string>;
    /**
     * The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL: projects/ project/global/gateways/default-internet-gateway
     */
    nextHopGateway?: pulumi.Input<string>;
    /**
     * The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs: - 10.128.0.56 - https://www.googleapis.com/compute/v1/projects/project/regions/region /forwardingRules/forwardingRule - regions/region/forwardingRules/forwardingRule 
     */
    nextHopIlb?: pulumi.Input<string>;
    /**
     * The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
     */
    nextHopInstance?: pulumi.Input<string>;
    /**
     * The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
     */
    nextHopIp?: pulumi.Input<string>;
    /**
     * The URL of the local network if it should handle matching packets.
     */
    nextHopNetwork?: pulumi.Input<string>;
    /**
     * The URL to a VpnTunnel that should handle matching packets.
     */
    nextHopVpnTunnel?: pulumi.Input<string>;
    /**
     * The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
     */
    priority?: pulumi.Input<number>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * A list of instance tags to which this route applies.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
