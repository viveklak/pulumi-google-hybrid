// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Instance in a given location.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:memcache/v1:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
     */
    public readonly authorizedNetwork!: pulumi.Output<string>;
    /**
     * The time the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Endpoint for the Discovery API.
     */
    public /*out*/ readonly discoveryEndpoint!: pulumi.Output<string>;
    /**
     * User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * List of messages that describe the current state of the Memcached instance.
     */
    public readonly instanceMessages!: pulumi.Output<outputs.memcache.v1.InstanceMessageResponse[]>;
    /**
     * Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The full version of memcached server running on this instance. System automatically determines the full memcached version for an instance based on the input MemcacheVersion. The full version format will be "memcached-1.5.16".
     */
    public /*out*/ readonly memcacheFullVersion!: pulumi.Output<string>;
    /**
     * List of Memcached nodes. Refer to Node message for more details.
     */
    public /*out*/ readonly memcacheNodes!: pulumi.Output<outputs.memcache.v1.NodeResponse[]>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
     */
    public readonly memcacheVersion!: pulumi.Output<string>;
    /**
     * Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for Memcached nodes.
     */
    public readonly nodeConfig!: pulumi.Output<outputs.memcache.v1.NodeConfigResponse>;
    /**
     * Number of nodes in the Memcached instance.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * User defined parameters to apply to the memcached process on each node.
     */
    public readonly parameters!: pulumi.Output<outputs.memcache.v1.MemcacheParametersResponse>;
    /**
     * The state of this Memcached instance.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The time the instance was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
     */
    public readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.nodeConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeConfig'");
            }
            if ((!args || args.nodeCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeCount'");
            }
            resourceInputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceMessages"] = args ? args.instanceMessages : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["memcacheVersion"] = args ? args.memcacheVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["discoveryEndpoint"] = undefined /*out*/;
            resourceInputs["memcacheFullVersion"] = undefined /*out*/;
            resourceInputs["memcacheNodes"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["authorizedNetwork"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["discoveryEndpoint"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["instanceMessages"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["memcacheFullVersion"] = undefined /*out*/;
            resourceInputs["memcacheNodes"] = undefined /*out*/;
            resourceInputs["memcacheVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nodeConfig"] = undefined /*out*/;
            resourceInputs["nodeCount"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["zones"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
     */
    authorizedNetwork?: pulumi.Input<string>;
    /**
     * User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
     */
    displayName?: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    /**
     * List of messages that describe the current state of the Memcached instance.
     */
    instanceMessages?: pulumi.Input<pulumi.Input<inputs.memcache.v1.InstanceMessageArgs>[]>;
    /**
     * Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
     */
    memcacheVersion?: pulumi.Input<enums.memcache.v1.InstanceMemcacheVersion>;
    /**
     * Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration for Memcached nodes.
     */
    nodeConfig: pulumi.Input<inputs.memcache.v1.NodeConfigArgs>;
    /**
     * Number of nodes in the Memcached instance.
     */
    nodeCount: pulumi.Input<number>;
    /**
     * User defined parameters to apply to the memcached process on each node.
     */
    parameters?: pulumi.Input<inputs.memcache.v1.MemcacheParametersArgs>;
    project?: pulumi.Input<string>;
    /**
     * Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
