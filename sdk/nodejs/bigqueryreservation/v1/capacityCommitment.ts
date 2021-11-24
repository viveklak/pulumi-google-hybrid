// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new capacity commitment resource.
 * Auto-naming is currently not supported for this resource.
 */
export class CapacityCommitment extends pulumi.CustomResource {
    /**
     * Get an existing CapacityCommitment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CapacityCommitment {
        return new CapacityCommitment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigqueryreservation/v1:CapacityCommitment';

    /**
     * Returns true if the given object is an instance of CapacityCommitment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CapacityCommitment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CapacityCommitment.__pulumiType;
    }

    /**
     * The end of the current commitment period. It is applicable only for ACTIVE capacity commitments.
     */
    public /*out*/ readonly commitmentEndTime!: pulumi.Output<string>;
    /**
     * The start of the current commitment period. It is applicable only for ACTIVE capacity commitments.
     */
    public /*out*/ readonly commitmentStartTime!: pulumi.Output<string>;
    /**
     * For FAILED commitment plan, provides the reason of failure.
     */
    public /*out*/ readonly failureStatus!: pulumi.Output<outputs.bigqueryreservation.v1.StatusResponse>;
    /**
     * The resource name of the capacity commitment, e.g., `projects/myproject/locations/US/capacityCommitments/123`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Capacity commitment commitment plan.
     */
    public readonly plan!: pulumi.Output<string>;
    /**
     * The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
     */
    public readonly renewalPlan!: pulumi.Output<string>;
    /**
     * Number of slots in this commitment.
     */
    public readonly slotCount!: pulumi.Output<string>;
    /**
     * State of the commitment.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a CapacityCommitment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CapacityCommitmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["capacityCommitmentId"] = args ? args.capacityCommitmentId : undefined;
            resourceInputs["enforceSingleAdminProjectPerOrg"] = args ? args.enforceSingleAdminProjectPerOrg : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["renewalPlan"] = args ? args.renewalPlan : undefined;
            resourceInputs["slotCount"] = args ? args.slotCount : undefined;
            resourceInputs["commitmentEndTime"] = undefined /*out*/;
            resourceInputs["commitmentStartTime"] = undefined /*out*/;
            resourceInputs["failureStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["commitmentEndTime"] = undefined /*out*/;
            resourceInputs["commitmentStartTime"] = undefined /*out*/;
            resourceInputs["failureStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["plan"] = undefined /*out*/;
            resourceInputs["renewalPlan"] = undefined /*out*/;
            resourceInputs["slotCount"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CapacityCommitment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CapacityCommitment resource.
 */
export interface CapacityCommitmentArgs {
    capacityCommitmentId?: pulumi.Input<string>;
    enforceSingleAdminProjectPerOrg?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Capacity commitment commitment plan.
     */
    plan?: pulumi.Input<enums.bigqueryreservation.v1.CapacityCommitmentPlan>;
    project?: pulumi.Input<string>;
    /**
     * The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
     */
    renewalPlan?: pulumi.Input<enums.bigqueryreservation.v1.CapacityCommitmentRenewalPlan>;
    /**
     * Number of slots in this commitment.
     */
    slotCount?: pulumi.Input<string>;
}
