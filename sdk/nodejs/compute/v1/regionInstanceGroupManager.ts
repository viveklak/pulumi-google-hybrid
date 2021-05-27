// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a managed instance group using the information that you specify in the request. After the group is created, instances in the group are created using the specified instance template. This operation is marked as DONE when the group is created even if the instances in the group have not yet been created. You must separately verify the status of the individual instances with the listmanagedinstances method.
 *
 * A regional managed instance group can contain up to 2000 instances.
 */
export class RegionInstanceGroupManager extends pulumi.CustomResource {
    /**
     * Get an existing RegionInstanceGroupManager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionInstanceGroupManager {
        return new RegionInstanceGroupManager(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/v1:RegionInstanceGroupManager';

    /**
     * Returns true if the given object is an instance of RegionInstanceGroupManager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionInstanceGroupManager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionInstanceGroupManager.__pulumiType;
    }

    /**
     * The autohealing policy for this managed instance group. You can specify only one value.
     */
    public readonly autoHealingPolicies!: pulumi.Output<outputs.compute.v1.InstanceGroupManagerAutoHealingPolicyResponse[]>;
    /**
     * The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
     */
    public readonly baseInstanceName!: pulumi.Output<string>;
    /**
     * [Output Only] The creation timestamp for this managed instance group in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * [Output Only] The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
     */
    public readonly currentActions!: pulumi.Output<outputs.compute.v1.InstanceGroupManagerActionsSummaryResponse>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
     */
    public readonly distributionPolicy!: pulumi.Output<outputs.compute.v1.DistributionPolicyResponse>;
    /**
     * Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
     */
    public readonly fingerprint!: pulumi.Output<string>;
    /**
     * [Output Only] The URL of the Instance Group resource.
     */
    public readonly instanceGroup!: pulumi.Output<string>;
    /**
     * The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
     */
    public readonly instanceTemplate!: pulumi.Output<string>;
    /**
     * [Output Only] The resource type, which is always compute#instanceGroupManager for managed instance groups.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Named ports configured for the Instance Groups complementary to this Instance Group Manager.
     */
    public readonly namedPorts!: pulumi.Output<outputs.compute.v1.NamedPortResponse[]>;
    /**
     * [Output Only] The URL of the region where the managed instance group resides (for regional resources).
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output Only] The URL for this managed instance group. The server defines this URL.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * Stateful configuration for this Instanced Group Manager
     */
    public readonly statefulPolicy!: pulumi.Output<outputs.compute.v1.StatefulPolicyResponse>;
    /**
     * [Output Only] The status of this managed instance group.
     */
    public readonly status!: pulumi.Output<outputs.compute.v1.InstanceGroupManagerStatusResponse>;
    /**
     * The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
     */
    public readonly targetPools!: pulumi.Output<string[]>;
    /**
     * The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
     */
    public readonly targetSize!: pulumi.Output<number>;
    /**
     * The update policy for this managed instance group.
     */
    public readonly updatePolicy!: pulumi.Output<outputs.compute.v1.InstanceGroupManagerUpdatePolicyResponse>;
    /**
     * Specifies the instance templates used by this managed instance group to create instances.
     *
     * Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
     */
    public readonly versions!: pulumi.Output<outputs.compute.v1.InstanceGroupManagerVersionResponse[]>;
    /**
     * [Output Only] The URL of a zone where the managed instance group is located (for zonal resources).
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a RegionInstanceGroupManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionInstanceGroupManagerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["autoHealingPolicies"] = args ? args.autoHealingPolicies : undefined;
            inputs["baseInstanceName"] = args ? args.baseInstanceName : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["currentActions"] = args ? args.currentActions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["distributionPolicy"] = args ? args.distributionPolicy : undefined;
            inputs["fingerprint"] = args ? args.fingerprint : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["instanceGroup"] = args ? args.instanceGroup : undefined;
            inputs["instanceTemplate"] = args ? args.instanceTemplate : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namedPorts"] = args ? args.namedPorts : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["statefulPolicy"] = args ? args.statefulPolicy : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["targetPools"] = args ? args.targetPools : undefined;
            inputs["targetSize"] = args ? args.targetSize : undefined;
            inputs["updatePolicy"] = args ? args.updatePolicy : undefined;
            inputs["versions"] = args ? args.versions : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        } else {
            inputs["autoHealingPolicies"] = undefined /*out*/;
            inputs["baseInstanceName"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["currentActions"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["distributionPolicy"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["instanceGroup"] = undefined /*out*/;
            inputs["instanceTemplate"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["namedPorts"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["statefulPolicy"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["targetPools"] = undefined /*out*/;
            inputs["targetSize"] = undefined /*out*/;
            inputs["updatePolicy"] = undefined /*out*/;
            inputs["versions"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionInstanceGroupManager.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionInstanceGroupManager resource.
 */
export interface RegionInstanceGroupManagerArgs {
    /**
     * The autohealing policy for this managed instance group. You can specify only one value.
     */
    readonly autoHealingPolicies?: pulumi.Input<pulumi.Input<inputs.compute.v1.InstanceGroupManagerAutoHealingPolicyArgs>[]>;
    /**
     * The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
     */
    readonly baseInstanceName?: pulumi.Input<string>;
    /**
     * [Output Only] The creation timestamp for this managed instance group in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * [Output Only] The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
     */
    readonly currentActions?: pulumi.Input<inputs.compute.v1.InstanceGroupManagerActionsSummaryArgs>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
     */
    readonly distributionPolicy?: pulumi.Input<inputs.compute.v1.DistributionPolicyArgs>;
    /**
     * Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
     */
    readonly fingerprint?: pulumi.Input<string>;
    /**
     * [Output Only] A unique identifier for this resource type. The server generates this identifier.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output Only] The URL of the Instance Group resource.
     */
    readonly instanceGroup?: pulumi.Input<string>;
    /**
     * The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
     */
    readonly instanceTemplate?: pulumi.Input<string>;
    /**
     * [Output Only] The resource type, which is always compute#instanceGroupManager for managed instance groups.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Named ports configured for the Instance Groups complementary to this Instance Group Manager.
     */
    readonly namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.v1.NamedPortArgs>[]>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] The URL of the region where the managed instance group resides (for regional resources).
     */
    readonly region: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * [Output Only] The URL for this managed instance group. The server defines this URL.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Stateful configuration for this Instanced Group Manager
     */
    readonly statefulPolicy?: pulumi.Input<inputs.compute.v1.StatefulPolicyArgs>;
    /**
     * [Output Only] The status of this managed instance group.
     */
    readonly status?: pulumi.Input<inputs.compute.v1.InstanceGroupManagerStatusArgs>;
    /**
     * The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
     */
    readonly targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
     */
    readonly targetSize?: pulumi.Input<number>;
    /**
     * The update policy for this managed instance group.
     */
    readonly updatePolicy?: pulumi.Input<inputs.compute.v1.InstanceGroupManagerUpdatePolicyArgs>;
    /**
     * Specifies the instance templates used by this managed instance group to create instances.
     *
     * Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
     */
    readonly versions?: pulumi.Input<pulumi.Input<inputs.compute.v1.InstanceGroupManagerVersionArgs>[]>;
    /**
     * [Output Only] The URL of a zone where the managed instance group is located (for zonal resources).
     */
    readonly zone?: pulumi.Input<string>;
}
