// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new resource policy.
 */
export class ResourcePolicy extends pulumi.CustomResource {
    /**
     * Get an existing ResourcePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ResourcePolicy {
        return new ResourcePolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:ResourcePolicy';

    /**
     * Returns true if the given object is an instance of ResourcePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourcePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourcePolicy.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    /**
     * Resource policy for instances for placement configuration.
     */
    public readonly groupPlacementPolicy!: pulumi.Output<outputs.compute.alpha.ResourcePolicyGroupPlacementPolicyResponse>;
    /**
     * Resource policy for scheduling instance operations.
     */
    public readonly instanceSchedulePolicy!: pulumi.Output<outputs.compute.alpha.ResourcePolicyInstanceSchedulePolicyResponse>;
    /**
     * Type of the resource. Always compute#resource_policies for resource policies.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The system status of the resource policy.
     */
    public /*out*/ readonly resourceStatus!: pulumi.Output<outputs.compute.alpha.ResourcePolicyResourceStatusResponse>;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Resource policy for persistent disks for creating snapshots.
     */
    public readonly snapshotSchedulePolicy!: pulumi.Output<outputs.compute.alpha.ResourcePolicySnapshotSchedulePolicyResponse>;
    /**
     * The status of resource policy creation.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Resource policy applicable to VMs for infrastructure maintenance.
     */
    public readonly vmMaintenancePolicy!: pulumi.Output<outputs.compute.alpha.ResourcePolicyVmMaintenancePolicyResponse>;

    /**
     * Create a ResourcePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourcePolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["groupPlacementPolicy"] = args ? args.groupPlacementPolicy : undefined;
            inputs["instanceSchedulePolicy"] = args ? args.instanceSchedulePolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["snapshotSchedulePolicy"] = args ? args.snapshotSchedulePolicy : undefined;
            inputs["vmMaintenancePolicy"] = args ? args.vmMaintenancePolicy : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["resourceStatus"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["groupPlacementPolicy"] = undefined /*out*/;
            inputs["instanceSchedulePolicy"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["resourceStatus"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["snapshotSchedulePolicy"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["vmMaintenancePolicy"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ResourcePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ResourcePolicy resource.
 */
export interface ResourcePolicyArgs {
    description?: pulumi.Input<string>;
    /**
     * Resource policy for instances for placement configuration.
     */
    groupPlacementPolicy?: pulumi.Input<inputs.compute.alpha.ResourcePolicyGroupPlacementPolicyArgs>;
    /**
     * Resource policy for scheduling instance operations.
     */
    instanceSchedulePolicy?: pulumi.Input<inputs.compute.alpha.ResourcePolicyInstanceSchedulePolicyArgs>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    region: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * Resource policy for persistent disks for creating snapshots.
     */
    snapshotSchedulePolicy?: pulumi.Input<inputs.compute.alpha.ResourcePolicySnapshotSchedulePolicyArgs>;
    /**
     * Resource policy applicable to VMs for infrastructure maintenance.
     */
    vmMaintenancePolicy?: pulumi.Input<inputs.compute.alpha.ResourcePolicyVmMaintenancePolicyArgs>;
}
