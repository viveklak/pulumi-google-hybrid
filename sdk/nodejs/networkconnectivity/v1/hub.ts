// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new hub in the specified project.
 * Auto-naming is currently not supported for this resource.
 */
export class Hub extends pulumi.CustomResource {
    /**
     * Get an existing Hub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Hub {
        return new Hub(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networkconnectivity/v1:Hub';

    /**
     * Returns true if the given object is an instance of Hub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Hub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Hub.__pulumiType;
    }

    /**
     * The time the hub was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * An optional description of the hub.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The VPC network associated with this hub's spokes. All of the VPN tunnels, VLAN attachments, and router appliance instances referenced by this hub's spokes must belong to this VPC network. This field is read-only. Network Connectivity Center automatically populates it based on the set of spokes attached to the hub.
     */
    public readonly routingVpcs!: pulumi.Output<outputs.networkconnectivity.v1.RoutingVPCResponse[]>;
    /**
     * The current lifecycle state of this hub.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted and another with the same name is created, the new hub is assigned a different unique_id.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    /**
     * The time the hub was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Hub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HubArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.hubId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hubId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hubId"] = args ? args.hubId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["routingVpcs"] = args ? args.routingVpcs : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["routingVpcs"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Hub.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Hub resource.
 */
export interface HubArgs {
    /**
     * An optional description of the hub.
     */
    description?: pulumi.Input<string>;
    hubId: pulumi.Input<string>;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The VPC network associated with this hub's spokes. All of the VPN tunnels, VLAN attachments, and router appliance instances referenced by this hub's spokes must belong to this VPC network. This field is read-only. Network Connectivity Center automatically populates it based on the set of spokes attached to the hub.
     */
    routingVpcs?: pulumi.Input<pulumi.Input<inputs.networkconnectivity.v1.RoutingVPCArgs>[]>;
}
