// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the details of a node.
 */
export function getNode(args: GetNodeArgs, opts?: pulumi.InvokeOptions): Promise<GetNodeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:tpu/v1alpha1:getNode", {
        "location": args.location,
        "nodeId": args.nodeId,
        "project": args.project,
    }, opts);
}

export interface GetNodeArgs {
    location: string;
    nodeId: string;
    project: string;
}

export interface GetNodeResult {
    /**
     * The type of hardware accelerators associated with this node.
     */
    readonly acceleratorType: string;
    /**
     * The API version that created this Node.
     */
    readonly apiVersion: string;
    /**
     * The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
     */
    readonly cidrBlock: string;
    /**
     * The time when the node was created.
     */
    readonly createTime: string;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    readonly description: string;
    /**
     * The health status of the TPU node.
     */
    readonly health: string;
    /**
     * If this field is populated, it contains a description of why the TPU Node is unhealthy.
     */
    readonly healthDescription: string;
    /**
     * Resource labels to represent user-provided metadata.
     */
    readonly labels: {[key: string]: string};
    /**
     * Immutable. The name of the TPU
     */
    readonly name: string;
    /**
     * The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
     */
    readonly network: string;
    /**
     * The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
     */
    readonly networkEndpoints: outputs.tpu.v1alpha1.NetworkEndpointResponse[];
    /**
     * The scheduling options for this node.
     */
    readonly schedulingConfig: outputs.tpu.v1alpha1.SchedulingConfigResponse;
    /**
     * The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
     */
    readonly serviceAccount: string;
    /**
     * The current state for the TPU Node.
     */
    readonly state: string;
    /**
     * The Symptoms that have occurred to the TPU Node.
     */
    readonly symptoms: outputs.tpu.v1alpha1.SymptomResponse[];
    /**
     * The version of Tensorflow running in the Node.
     */
    readonly tensorflowVersion: string;
    /**
     * Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
     */
    readonly useServiceNetworking: boolean;
}
