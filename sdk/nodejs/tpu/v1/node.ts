// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a node.
 * Auto-naming is currently not supported for this resource.
 */
export class Node extends pulumi.CustomResource {
    /**
     * Get an existing Node resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Node {
        return new Node(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:tpu/v1:Node';

    /**
     * Returns true if the given object is an instance of Node.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Node {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Node.__pulumiType;
    }

    /**
     * The type of hardware accelerators associated with this node.
     */
    public readonly acceleratorType!: pulumi.Output<string>;
    /**
     * The API version that created this Node.
     */
    public /*out*/ readonly apiVersion!: pulumi.Output<string>;
    /**
     * The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The time when the node was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The health status of the TPU node.
     */
    public readonly health!: pulumi.Output<string>;
    /**
     * If this field is populated, it contains a description of why the TPU Node is unhealthy.
     */
    public /*out*/ readonly healthDescription!: pulumi.Output<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Immutable. The name of the TPU
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
     */
    public /*out*/ readonly networkEndpoints!: pulumi.Output<outputs.tpu.v1.NetworkEndpointResponse[]>;
    /**
     * The scheduling options for this node.
     */
    public readonly schedulingConfig!: pulumi.Output<outputs.tpu.v1.SchedulingConfigResponse>;
    /**
     * The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
     */
    public /*out*/ readonly serviceAccount!: pulumi.Output<string>;
    /**
     * The current state for the TPU Node.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The Symptoms that have occurred to the TPU Node.
     */
    public /*out*/ readonly symptoms!: pulumi.Output<outputs.tpu.v1.SymptomResponse[]>;
    /**
     * The version of Tensorflow running in the Node.
     */
    public readonly tensorflowVersion!: pulumi.Output<string>;
    /**
     * Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
     */
    public readonly useServiceNetworking!: pulumi.Output<boolean>;

    /**
     * Create a Node resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.acceleratorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceleratorType'");
            }
            if ((!args || args.tensorflowVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tensorflowVersion'");
            }
            resourceInputs["acceleratorType"] = args ? args.acceleratorType : undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["health"] = args ? args.health : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["nodeId"] = args ? args.nodeId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["schedulingConfig"] = args ? args.schedulingConfig : undefined;
            resourceInputs["tensorflowVersion"] = args ? args.tensorflowVersion : undefined;
            resourceInputs["useServiceNetworking"] = args ? args.useServiceNetworking : undefined;
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["healthDescription"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkEndpoints"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["symptoms"] = undefined /*out*/;
        } else {
            resourceInputs["acceleratorType"] = undefined /*out*/;
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["cidrBlock"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["healthDescription"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["networkEndpoints"] = undefined /*out*/;
            resourceInputs["schedulingConfig"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["symptoms"] = undefined /*out*/;
            resourceInputs["tensorflowVersion"] = undefined /*out*/;
            resourceInputs["useServiceNetworking"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Node.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Node resource.
 */
export interface NodeArgs {
    /**
     * The type of hardware accelerators associated with this node.
     */
    acceleratorType: pulumi.Input<string>;
    /**
     * The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * The health status of the TPU node.
     */
    health?: pulumi.Input<enums.tpu.v1.NodeHealth>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
     */
    network?: pulumi.Input<string>;
    nodeId?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The scheduling options for this node.
     */
    schedulingConfig?: pulumi.Input<inputs.tpu.v1.SchedulingConfigArgs>;
    /**
     * The version of Tensorflow running in the Node.
     */
    tensorflowVersion: pulumi.Input<string>;
    /**
     * Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
     */
    useServiceNetworking?: pulumi.Input<boolean>;
}
