// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a cluster within an instance. Note that exactly one of Cluster.serve_nodes and Cluster.cluster_config.cluster_autoscaling_config can be set. If serve_nodes is set to non-zero, then the cluster is manually scaled. If cluster_config.cluster_autoscaling_config is non-empty, then autoscaling is enabled.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigtableadmin/v2:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Configuration for this cluster.
     */
    public readonly clusterConfig!: pulumi.Output<outputs.bigtableadmin.v2.ClusterConfigResponse>;
    /**
     * Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
     */
    public readonly defaultStorageType!: pulumi.Output<string>;
    /**
     * Immutable. The encryption configuration for CMEK-protected clusters.
     */
    public readonly encryptionConfig!: pulumi.Output<outputs.bigtableadmin.v2.EncryptionConfigResponse>;
    /**
     * Immutable. The location where this cluster's nodes and storage reside. For best performance, clients should be located as close as possible to this cluster. Currently only zones are supported, so values should be of the form `projects/{project}/locations/{zone}`.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The unique name of the cluster. Values are of the form `projects/{project}/instances/{instance}/clusters/a-z*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of nodes allocated to this cluster. More nodes enable higher throughput and more consistent performance.
     */
    public readonly serveNodes!: pulumi.Output<number>;
    /**
     * The current state of the cluster.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["clusterConfig"] = args ? args.clusterConfig : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["defaultStorageType"] = args ? args.defaultStorageType : undefined;
            resourceInputs["encryptionConfig"] = args ? args.encryptionConfig : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serveNodes"] = args ? args.serveNodes : undefined;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["clusterConfig"] = undefined /*out*/;
            resourceInputs["defaultStorageType"] = undefined /*out*/;
            resourceInputs["encryptionConfig"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serveNodes"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Configuration for this cluster.
     */
    clusterConfig?: pulumi.Input<inputs.bigtableadmin.v2.ClusterConfigArgs>;
    /**
     * Required. The ID to be used when referring to the new cluster within its instance, e.g., just `mycluster` rather than `projects/myproject/instances/myinstance/clusters/mycluster`.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
     */
    defaultStorageType?: pulumi.Input<enums.bigtableadmin.v2.ClusterDefaultStorageType>;
    /**
     * Immutable. The encryption configuration for CMEK-protected clusters.
     */
    encryptionConfig?: pulumi.Input<inputs.bigtableadmin.v2.EncryptionConfigArgs>;
    instanceId: pulumi.Input<string>;
    /**
     * Immutable. The location where this cluster's nodes and storage reside. For best performance, clients should be located as close as possible to this cluster. Currently only zones are supported, so values should be of the form `projects/{project}/locations/{zone}`.
     */
    location?: pulumi.Input<string>;
    /**
     * The unique name of the cluster. Values are of the form `projects/{project}/instances/{instance}/clusters/a-z*`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The number of nodes allocated to this cluster. More nodes enable higher throughput and more consistent performance.
     */
    serveNodes?: pulumi.Input<number>;
}
