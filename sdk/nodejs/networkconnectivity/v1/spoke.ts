// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a spoke in the specified project and location.
 */
export class Spoke extends pulumi.CustomResource {
    /**
     * Get an existing Spoke resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Spoke {
        return new Spoke(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networkconnectivity/v1:Spoke';

    /**
     * Returns true if the given object is an instance of Spoke.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Spoke {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Spoke.__pulumiType;
    }

    /**
     * The time the spoke was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * An optional description of the spoke.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Immutable. The name of the hub that this spoke is attached to.
     */
    public readonly hub!: pulumi.Output<string>;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * VLAN attachments that are associated with the spoke.
     */
    public readonly linkedInterconnectAttachments!: pulumi.Output<outputs.networkconnectivity.v1.LinkedInterconnectAttachmentsResponse>;
    /**
     * Router appliance instances that are associated with the spoke.
     */
    public readonly linkedRouterApplianceInstances!: pulumi.Output<outputs.networkconnectivity.v1.LinkedRouterApplianceInstancesResponse>;
    /**
     * VPN tunnels that are associated with the spoke.
     */
    public readonly linkedVpnTunnels!: pulumi.Output<outputs.networkconnectivity.v1.LinkedVpnTunnelsResponse>;
    /**
     * Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current lifecycle state of this spoke.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    /**
     * The time the spoke was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Spoke resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SpokeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.spokeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'spokeId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hub"] = args ? args.hub : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["linkedInterconnectAttachments"] = args ? args.linkedInterconnectAttachments : undefined;
            resourceInputs["linkedRouterApplianceInstances"] = args ? args.linkedRouterApplianceInstances : undefined;
            resourceInputs["linkedVpnTunnels"] = args ? args.linkedVpnTunnels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["spokeId"] = args ? args.spokeId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["hub"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["linkedInterconnectAttachments"] = undefined /*out*/;
            resourceInputs["linkedRouterApplianceInstances"] = undefined /*out*/;
            resourceInputs["linkedVpnTunnels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Spoke.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Spoke resource.
 */
export interface SpokeArgs {
    /**
     * An optional description of the spoke.
     */
    description?: pulumi.Input<string>;
    /**
     * Immutable. The name of the hub that this spoke is attached to.
     */
    hub?: pulumi.Input<string>;
    /**
     * Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * VLAN attachments that are associated with the spoke.
     */
    linkedInterconnectAttachments?: pulumi.Input<inputs.networkconnectivity.v1.LinkedInterconnectAttachmentsArgs>;
    /**
     * Router appliance instances that are associated with the spoke.
     */
    linkedRouterApplianceInstances?: pulumi.Input<inputs.networkconnectivity.v1.LinkedRouterApplianceInstancesArgs>;
    /**
     * VPN tunnels that are associated with the spoke.
     */
    linkedVpnTunnels?: pulumi.Input<inputs.networkconnectivity.v1.LinkedVpnTunnelsArgs>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    spokeId: pulumi.Input<string>;
}
