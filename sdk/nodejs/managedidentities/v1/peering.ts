// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a Peering for Managed AD instance.
 * Auto-naming is currently not supported for this resource.
 */
export class Peering extends pulumi.CustomResource {
    /**
     * Get an existing Peering resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Peering {
        return new Peering(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:managedidentities/v1:Peering';

    /**
     * Returns true if the given object is an instance of Peering.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Peering {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Peering.__pulumiType;
    }

    /**
     * The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
     */
    public readonly authorizedNetwork!: pulumi.Output<string>;
    /**
     * The time the instance was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
     */
    public readonly domainResource!: pulumi.Output<string>;
    /**
     * Optional. Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Unique name of the peering in this scope including projects and location using the form: `projects/{project_id}/locations/global/peerings/{peering_id}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current state of this Peering.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current status of this peering, if available.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * Last update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Peering resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PeeringArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.authorizedNetwork === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizedNetwork'");
            }
            if ((!args || args.domainResource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainResource'");
            }
            if ((!args || args.peeringId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peeringId'");
            }
            resourceInputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            resourceInputs["domainResource"] = args ? args.domainResource : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["peeringId"] = args ? args.peeringId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["authorizedNetwork"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["domainResource"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Peering.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Peering resource.
 */
export interface PeeringArgs {
    /**
     * The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
     */
    authorizedNetwork: pulumi.Input<string>;
    /**
     * Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form: `projects/{project_id}/locations/global/domains/{domain_name}`
     */
    domainResource: pulumi.Input<string>;
    /**
     * Optional. Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    peeringId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
