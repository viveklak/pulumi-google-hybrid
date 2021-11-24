// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a node pool for a cluster.
 */
export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:container/v1beta1:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
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
    public /*out*/ readonly instanceGroupUrls!: pulumi.Output<string[]>;
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
    public /*out*/ readonly podIpv4CidrSize!: pulumi.Output<number>;
    /**
     * [Output only] Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * [Output only] The status of the nodes in this pool instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Upgrade settings control disruption and speed of the upgrade.
     */
    public readonly upgradeSettings!: pulumi.Output<outputs.container.v1beta1.UpgradeSettingsResponse>;
    /**
     * The version of the Kubernetes of this node.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodePoolArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["autoscaling"] = args ? args.autoscaling : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["initialNodeCount"] = args ? args.initialNodeCount : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["management"] = args ? args.management : undefined;
            resourceInputs["maxPodsConstraint"] = args ? args.maxPodsConstraint : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["upgradeSettings"] = args ? args.upgradeSettings : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["instanceGroupUrls"] = undefined /*out*/;
            resourceInputs["podIpv4CidrSize"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["autoscaling"] = undefined /*out*/;
            resourceInputs["conditions"] = undefined /*out*/;
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["initialNodeCount"] = undefined /*out*/;
            resourceInputs["instanceGroupUrls"] = undefined /*out*/;
            resourceInputs["locations"] = undefined /*out*/;
            resourceInputs["management"] = undefined /*out*/;
            resourceInputs["maxPodsConstraint"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkConfig"] = undefined /*out*/;
            resourceInputs["podIpv4CidrSize"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["upgradeSettings"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * Autoscaler configuration for this NodePool. Autoscaler is enabled only if a valid configuration is present.
     */
    autoscaling?: pulumi.Input<inputs.container.v1beta1.NodePoolAutoscalingArgs>;
    clusterId: pulumi.Input<string>;
    /**
     * Which conditions caused the current node pool state.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.container.v1beta1.StatusConditionArgs>[]>;
    /**
     * The node configuration of the pool.
     */
    config?: pulumi.Input<inputs.container.v1beta1.NodeConfigArgs>;
    /**
     * The initial node count for the pool. You must ensure that your Compute Engine [resource quota](https://cloud.google.com/compute/quotas) is sufficient for this number of instances. You must also have available firewall and routes quota.
     */
    initialNodeCount?: pulumi.Input<number>;
    location?: pulumi.Input<string>;
    /**
     * The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the NodePool's nodes should be located. If this value is unspecified during node pool creation, the [Cluster.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster.FIELDS.locations) value will be used, instead. Warning: changing node pool locations will result in nodes being added and/or removed.
     */
    locations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * NodeManagement configuration for this NodePool.
     */
    management?: pulumi.Input<inputs.container.v1beta1.NodeManagementArgs>;
    /**
     * The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
     */
    maxPodsConstraint?: pulumi.Input<inputs.container.v1beta1.MaxPodsConstraintArgs>;
    /**
     * The name of the node pool.
     */
    name?: pulumi.Input<string>;
    /**
     * Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults.
     */
    networkConfig?: pulumi.Input<inputs.container.v1beta1.NodeNetworkConfigArgs>;
    /**
     * The parent (project, location, cluster id) where the node pool will be created. Specified in the format `projects/*&#47;locations/*&#47;clusters/*`.
     */
    parent?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Upgrade settings control disruption and speed of the upgrade.
     */
    upgradeSettings?: pulumi.Input<inputs.container.v1beta1.UpgradeSettingsArgs>;
    /**
     * The version of the Kubernetes of this node.
     */
    version?: pulumi.Input<string>;
}
