// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the requested node pool.
 */
export function getNodePool(args: GetNodePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetNodePoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:container/v1:getNodePool", {
        "clusterId": args.clusterId,
        "location": args.location,
        "nodePoolId": args.nodePoolId,
        "project": args.project,
    }, opts);
}

export interface GetNodePoolArgs {
    clusterId: string;
    location: string;
    nodePoolId: string;
    project?: string;
}

export interface GetNodePoolResult {
    /**
     * Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
     */
    readonly autoscaling: outputs.container.v1.NodePoolAutoscalingResponse;
    /**
     * Which conditions caused the current node pool state.
     */
    readonly conditions: outputs.container.v1.StatusConditionResponse[];
    /**
     * The node configuration of the pool.
     */
    readonly config: outputs.container.v1.NodeConfigResponse;
    /**
     * The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
     */
    readonly initialNodeCount: number;
    /**
     * [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
     */
    readonly instanceGroupUrls: string[];
    /**
     * The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
     */
    readonly locations: string[];
    /**
     * NodeManagement configuration for this NodePool.
     */
    readonly management: outputs.container.v1.NodeManagementResponse;
    /**
     * The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
     */
    readonly maxPodsConstraint: outputs.container.v1.MaxPodsConstraintResponse;
    /**
     * The name of the node pool.
     */
    readonly name: string;
    /**
     * Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
     */
    readonly networkConfig: outputs.container.v1.NodeNetworkConfigResponse;
    /**
     * [Output only] The pod CIDR block size per node in this node pool.
     */
    readonly podIpv4CidrSize: number;
    /**
     * [Output only] Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * [Output only] The status of the nodes in this pool instance.
     */
    readonly status: string;
    /**
     * Upgrade settings control disruption and speed of the upgrade.
     */
    readonly upgradeSettings: outputs.container.v1.UpgradeSettingsResponse;
    /**
     * The version of the Kubernetes of this node.
     */
    readonly version: string;
}

export function getNodePoolOutput(args: GetNodePoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNodePoolResult> {
    return pulumi.output(args).apply(a => getNodePool(a, opts))
}

export interface GetNodePoolOutputArgs {
    clusterId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    nodePoolId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
