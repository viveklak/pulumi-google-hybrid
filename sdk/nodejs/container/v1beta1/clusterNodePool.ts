// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a node pool for a cluster.
 */
export class ClusterNodePool extends pulumi.CustomResource {
    /**
     * Get an existing ClusterNodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClusterNodePool {
        return new ClusterNodePool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:container/v1beta1:ClusterNodePool';

    /**
     * Returns true if the given object is an instance of ClusterNodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterNodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterNodePool.__pulumiType;
    }

    /**
     * Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
     */
    public readonly autoscaling!: pulumi.Output<outputs.container.v1beta1.NodePoolAutoscalingResponse>;
    /**
     * Which conditions caused the current node pool state.
     */
    public readonly conditions!: pulumi.Output<outputs.container.v1beta1.StatusConditionResponse[]>;
    /**
     * The node configuration of the pool.
     */
    public readonly config!: pulumi.Output<outputs.container.v1beta1.NodeConfigResponse>;
    /**
     * The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
     */
    public readonly initialNodeCount!: pulumi.Output<number>;
    /**
     * [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
     */
    public readonly instanceGroupUrls!: pulumi.Output<string[]>;
    /**
     * The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
     */
    public readonly locations!: pulumi.Output<string[]>;
    /**
     * NodeManagement configuration for this NodePool.
     */
    public readonly management!: pulumi.Output<outputs.container.v1beta1.NodeManagementResponse>;
    /**
     * The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
     */
    public readonly maxPodsConstraint!: pulumi.Output<outputs.container.v1beta1.MaxPodsConstraintResponse>;
    /**
     * The name of the node pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
     */
    public readonly networkConfig!: pulumi.Output<outputs.container.v1beta1.NodeNetworkConfigResponse>;
    /**
     * [Output only] The pod CIDR block size per node in this node pool.
     */
    public readonly podIpv4CidrSize!: pulumi.Output<number>;
    /**
     * [Output only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * [Output only] The status of the nodes in this pool instance.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Upgrade settings control disruption and speed of the upgrade.
     */
    public readonly upgradeSettings!: pulumi.Output<outputs.container.v1beta1.UpgradeSettingsResponse>;
    /**
     * The version of the Kubernetes of this node.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a ClusterNodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterNodePoolArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clustersId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clustersId'");
            }
            if ((!args || args.locationsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locationsId'");
            }
            if ((!args || args.nodePoolsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodePoolsId'");
            }
            if ((!args || args.projectsId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectsId'");
            }
            inputs["autoscaling"] = args ? args.autoscaling : undefined;
            inputs["clustersId"] = args ? args.clustersId : undefined;
            inputs["conditions"] = args ? args.conditions : undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["initialNodeCount"] = args ? args.initialNodeCount : undefined;
            inputs["instanceGroupUrls"] = args ? args.instanceGroupUrls : undefined;
            inputs["locations"] = args ? args.locations : undefined;
            inputs["locationsId"] = args ? args.locationsId : undefined;
            inputs["management"] = args ? args.management : undefined;
            inputs["maxPodsConstraint"] = args ? args.maxPodsConstraint : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkConfig"] = args ? args.networkConfig : undefined;
            inputs["nodePoolsId"] = args ? args.nodePoolsId : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["podIpv4CidrSize"] = args ? args.podIpv4CidrSize : undefined;
            inputs["projectsId"] = args ? args.projectsId : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["upgradeSettings"] = args ? args.upgradeSettings : undefined;
            inputs["version"] = args ? args.version : undefined;
        } else {
            inputs["autoscaling"] = undefined /*out*/;
            inputs["conditions"] = undefined /*out*/;
            inputs["config"] = undefined /*out*/;
            inputs["initialNodeCount"] = undefined /*out*/;
            inputs["instanceGroupUrls"] = undefined /*out*/;
            inputs["locations"] = undefined /*out*/;
            inputs["management"] = undefined /*out*/;
            inputs["maxPodsConstraint"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkConfig"] = undefined /*out*/;
            inputs["podIpv4CidrSize"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["upgradeSettings"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClusterNodePool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClusterNodePool resource.
 */
export interface ClusterNodePoolArgs {
    /**
     * Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
     */
    readonly autoscaling?: pulumi.Input<inputs.container.v1beta1.NodePoolAutoscalingArgs>;
    readonly clustersId: pulumi.Input<string>;
    /**
     * Which conditions caused the current node pool state.
     */
    readonly conditions?: pulumi.Input<pulumi.Input<inputs.container.v1beta1.StatusConditionArgs>[]>;
    /**
     * The node configuration of the pool.
     */
    readonly config?: pulumi.Input<inputs.container.v1beta1.NodeConfigArgs>;
    /**
     * The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
     */
    readonly initialNodeCount?: pulumi.Input<number>;
    /**
     * [Output only] The resource URLs of the [managed instance groups](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances) associated with this node pool.
     */
    readonly instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
     */
    readonly locations?: pulumi.Input<pulumi.Input<string>[]>;
    readonly locationsId: pulumi.Input<string>;
    /**
     * NodeManagement configuration for this NodePool.
     */
    readonly management?: pulumi.Input<inputs.container.v1beta1.NodeManagementArgs>;
    /**
     * The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
     */
    readonly maxPodsConstraint?: pulumi.Input<inputs.container.v1beta1.MaxPodsConstraintArgs>;
    /**
     * The name of the node pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
     */
    readonly networkConfig?: pulumi.Input<inputs.container.v1beta1.NodeNetworkConfigArgs>;
    readonly nodePoolsId: pulumi.Input<string>;
    /**
     * The parent (project, location, cluster id) where the node pool will be created. Specified in the format `projects/*&#47;locations/*&#47;clusters/*`.
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * [Output only] The pod CIDR block size per node in this node pool.
     */
    readonly podIpv4CidrSize?: pulumi.Input<number>;
    readonly projectsId: pulumi.Input<string>;
    /**
     * [Output only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * [Output only] The status of the nodes in this pool instance.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Upgrade settings control disruption and speed of the upgrade.
     */
    readonly upgradeSettings?: pulumi.Input<inputs.container.v1beta1.UpgradeSettingsArgs>;
    /**
     * The version of the Kubernetes of this node.
     */
    readonly version?: pulumi.Input<string>;
}
