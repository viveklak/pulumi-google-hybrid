// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Spoke in a given project and location.
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
    public static readonly __pulumiType = 'google-native:networkconnectivity/v1alpha1:Spoke';

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
     * The time when the Spoke was created.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * Short description of the spoke resource
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The resource URL of the hub resource that the spoke is attached to
     */
    public readonly hub!: pulumi.Output<string>;
    /**
     * User-defined labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The URIs of linked interconnect attachment resources
     */
    public readonly linkedInterconnectAttachments!: pulumi.Output<string[]>;
    /**
     * The URIs of linked Router appliance resources
     */
    public readonly linkedRouterApplianceInstances!: pulumi.Output<outputs.networkconnectivity.v1alpha1.RouterApplianceInstanceResponse[]>;
    /**
     * The URIs of linked VPN tunnel resources
     */
    public readonly linkedVpnTunnels!: pulumi.Output<string[]>;
    /**
     * Immutable. The name of a Spoke resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current lifecycle state of this Hub.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Google-generated UUID for this resource. This is unique across all Spoke resources. If a Spoke resource is deleted and another with the same name is created, it gets a different unique_id.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    /**
     * The time when the Spoke was updated.
     */
    public readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Spoke resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SpokeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["createTime"] = args ? args.createTime : undefined;
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
            resourceInputs["updateTime"] = args ? args.updateTime : undefined;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
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
     * The time when the Spoke was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Short description of the spoke resource
     */
    description?: pulumi.Input<string>;
    /**
     * The resource URL of the hub resource that the spoke is attached to
     */
    hub?: pulumi.Input<string>;
    /**
     * User-defined labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The URIs of linked interconnect attachment resources
     */
    linkedInterconnectAttachments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URIs of linked Router appliance resources
     */
    linkedRouterApplianceInstances?: pulumi.Input<pulumi.Input<inputs.networkconnectivity.v1alpha1.RouterApplianceInstanceArgs>[]>;
    /**
     * The URIs of linked VPN tunnel resources
     */
    linkedVpnTunnels?: pulumi.Input<pulumi.Input<string>[]>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. The name of a Spoke resource.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    spokeId?: pulumi.Input<string>;
    /**
     * The time when the Spoke was updated.
     */
    updateTime?: pulumi.Input<string>;
}
