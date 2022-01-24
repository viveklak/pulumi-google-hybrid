// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new Membership. **This is currently only supported for GKE clusters on Google Cloud**. To register other clusters, follow the instructions at https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster.
 * Auto-naming is currently not supported for this resource.
 */
export class Membership extends pulumi.CustomResource {
    /**
     * Get an existing Membership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Membership {
        return new Membership(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gkehub/v1alpha2:Membership';

    /**
     * Returns true if the given object is an instance of Membership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Membership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Membership.__pulumiType;
    }

    /**
     * Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
     */
    public readonly authority!: pulumi.Output<outputs.gkehub.v1alpha2.AuthorityResponse>;
    /**
     * When the Membership was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * When the Membership was deleted.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*` This field is present for legacy purposes.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * Optional. Endpoint information to reach this member.
     */
    public readonly endpoint!: pulumi.Output<outputs.gkehub.v1alpha2.MembershipEndpointResponse>;
    /**
     * Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
     */
    public readonly externalId!: pulumi.Output<string>;
    /**
     * Optional. The infrastructure type this Membership is running on.
     */
    public readonly infrastructureType!: pulumi.Output<string>;
    /**
     * Optional. GCP labels for this membership.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.
     */
    public /*out*/ readonly lastConnectionTime!: pulumi.Output<string>;
    /**
     * The full, unique name of this Membership resource in the format `projects/*&#47;locations/*&#47;memberships/{membership_id}`, set during creation. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * State of the Membership resource.
     */
    public /*out*/ readonly state!: pulumi.Output<outputs.gkehub.v1alpha2.MembershipStateResponse>;
    /**
     * Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;
    /**
     * When the Membership was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Membership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MembershipArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.membershipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'membershipId'");
            }
            resourceInputs["authority"] = args ? args.authority : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["externalId"] = args ? args.externalId : undefined;
            resourceInputs["infrastructureType"] = args ? args.infrastructureType : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["membershipId"] = args ? args.membershipId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["lastConnectionTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["authority"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["infrastructureType"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["lastConnectionTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uniqueId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Membership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Membership resource.
 */
export interface MembershipArgs {
    /**
     * Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
     */
    authority?: pulumi.Input<inputs.gkehub.v1alpha2.AuthorityArgs>;
    /**
     * Optional. Endpoint information to reach this member.
     */
    endpoint?: pulumi.Input<inputs.gkehub.v1alpha2.MembershipEndpointArgs>;
    /**
     * Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
     */
    externalId?: pulumi.Input<string>;
    /**
     * Optional. The infrastructure type this Membership is running on.
     */
    infrastructureType?: pulumi.Input<enums.gkehub.v1alpha2.MembershipInfrastructureType>;
    /**
     * Optional. GCP labels for this membership.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    membershipId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
