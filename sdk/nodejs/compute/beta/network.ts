// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a network in the specified project using the data included in the request.
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * Must be set to create a VPC network. If not set, a legacy network is created.
     *
     * When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode.
     *
     * An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges.
     *
     * For custom mode VPC networks, you can add subnets using the subnetworks insert method.
     */
    public readonly autoCreateSubnetworks!: pulumi.Output<boolean>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The gateway address for default routing out of the network, selected by GCP.
     */
    public /*out*/ readonly gatewayIPv4!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#network for networks.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes.
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of network peerings for the resource.
     */
    public /*out*/ readonly peerings!: pulumi.Output<outputs.compute.beta.NetworkPeeringResponse[]>;
    /**
     * The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
     */
    public readonly routingConfig!: pulumi.Output<outputs.compute.beta.NetworkRoutingConfigResponse>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined fully-qualified URLs for all subnetworks in this VPC network.
     */
    public /*out*/ readonly subnetworks!: pulumi.Output<string[]>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["autoCreateSubnetworks"] = args ? args.autoCreateSubnetworks : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["routingConfig"] = args ? args.routingConfig : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["gatewayIPv4"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["peerings"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["subnetworks"] = undefined /*out*/;
        } else {
            inputs["autoCreateSubnetworks"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["gatewayIPv4"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["mtu"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["peerings"] = undefined /*out*/;
            inputs["routingConfig"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["subnetworks"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Network.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * Must be set to create a VPC network. If not set, a legacy network is created.
     *
     * When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode.
     *
     * An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges.
     *
     * For custom mode VPC networks, you can add subnets using the subnetworks insert method.
     */
    autoCreateSubnetworks?: pulumi.Input<boolean>;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1460 and the maximum value is 1500 bytes.
     */
    mtu?: pulumi.Input<number>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
     */
    routingConfig?: pulumi.Input<inputs.compute.beta.NetworkRoutingConfigArgs>;
}
