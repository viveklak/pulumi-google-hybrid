// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a managed instance group using the information that you specify in the request. After the group is created, instances in the group are created using the specified instance template. This operation is marked as DONE when the group is created even if the instances in the group have not yet been created. You must separately verify the status of the individual instances with the listmanagedinstances method. A regional managed instance group can contain up to 2000 instances.
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
    public static readonly __pulumiType = 'google-native:compute/alpha:RegionInstanceGroupManager';

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
     * Specifies the instances configs overrides that should be applied for all instances in the MIG.
     */
    public readonly allInstancesConfig!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerAllInstancesConfigResponse>;
    /**
     * The autohealing policy for this managed instance group. You can specify only one value.
     */
    public readonly autoHealingPolicies!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerAutoHealingPolicyResponse[]>;
    /**
     * The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
     */
    public readonly baseInstanceName!: pulumi.Output<string>;
    /**
     * The creation timestamp for this managed instance group in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
     */
    public /*out*/ readonly currentActions!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerActionsSummaryResponse>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
     */
    public readonly distributionPolicy!: pulumi.Output<outputs.compute.alpha.DistributionPolicyResponse>;
    /**
     * The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
     */
    public readonly failoverAction!: pulumi.Output<string>;
    /**
     * Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The URL of the Instance Group resource.
     */
    public /*out*/ readonly instanceGroup!: pulumi.Output<string>;
    /**
     * Instance lifecycle policy for this Instance Group Manager.
     */
    public readonly instanceLifecyclePolicy!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerInstanceLifecyclePolicyResponse>;
    /**
     * The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
     */
    public readonly instanceTemplate!: pulumi.Output<string>;
    /**
     * The resource type, which is always compute#instanceGroupManager for managed instance groups.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Named ports configured for the Instance Groups complementary to this Instance Group Manager.
     */
    public readonly namedPorts!: pulumi.Output<outputs.compute.alpha.NamedPortResponse[]>;
    /**
     * The URL of the region where the managed instance group resides (for regional resources).
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URL for this managed instance group. The server defines this URL.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Standby policy for stopped and suspended instances.
     */
    public readonly standbyPolicy!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerStandbyPolicyResponse>;
    /**
     * Stateful configuration for this Instanced Group Manager
     */
    public readonly statefulPolicy!: pulumi.Output<outputs.compute.alpha.StatefulPolicyResponse>;
    /**
     * The status of this managed instance group.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerStatusResponse>;
    /**
     * The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
     */
    public readonly targetPools!: pulumi.Output<string[]>;
    /**
     * The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
     */
    public readonly targetSize!: pulumi.Output<number>;
    /**
     * The target number of stopped instances for this managed instance group. This number changes when you: - Stop instance using the stopInstances method or start instances using the startInstances method. - Manually change the targetStoppedSize using the update method. 
     */
    public readonly targetStoppedSize!: pulumi.Output<number>;
    /**
     * The target number of suspended instances for this managed instance group. This number changes when you: - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. - Manually change the targetSuspendedSize using the update method. 
     */
    public readonly targetSuspendedSize!: pulumi.Output<number>;
    /**
     * The update policy for this managed instance group.
     */
    public readonly updatePolicy!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerUpdatePolicyResponse>;
    /**
     * Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
     */
    public readonly versions!: pulumi.Output<outputs.compute.alpha.InstanceGroupManagerVersionResponse[]>;
    /**
     * The URL of a zone where the managed instance group is located (for zonal resources).
     */
    public /*out*/ readonly zone!: pulumi.Output<string>;

    /**
     * Create a RegionInstanceGroupManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionInstanceGroupManagerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["allInstancesConfig"] = args ? args.allInstancesConfig : undefined;
            resourceInputs["autoHealingPolicies"] = args ? args.autoHealingPolicies : undefined;
            resourceInputs["baseInstanceName"] = args ? args.baseInstanceName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["distributionPolicy"] = args ? args.distributionPolicy : undefined;
            resourceInputs["failoverAction"] = args ? args.failoverAction : undefined;
            resourceInputs["instanceLifecyclePolicy"] = args ? args.instanceLifecyclePolicy : undefined;
            resourceInputs["instanceTemplate"] = args ? args.instanceTemplate : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namedPorts"] = args ? args.namedPorts : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["standbyPolicy"] = args ? args.standbyPolicy : undefined;
            resourceInputs["statefulPolicy"] = args ? args.statefulPolicy : undefined;
            resourceInputs["targetPools"] = args ? args.targetPools : undefined;
            resourceInputs["targetSize"] = args ? args.targetSize : undefined;
            resourceInputs["targetStoppedSize"] = args ? args.targetStoppedSize : undefined;
            resourceInputs["targetSuspendedSize"] = args ? args.targetSuspendedSize : undefined;
            resourceInputs["updatePolicy"] = args ? args.updatePolicy : undefined;
            resourceInputs["versions"] = args ? args.versions : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["currentActions"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["instanceGroup"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        } else {
            resourceInputs["allInstancesConfig"] = undefined /*out*/;
            resourceInputs["autoHealingPolicies"] = undefined /*out*/;
            resourceInputs["baseInstanceName"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["currentActions"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["distributionPolicy"] = undefined /*out*/;
            resourceInputs["failoverAction"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["instanceGroup"] = undefined /*out*/;
            resourceInputs["instanceLifecyclePolicy"] = undefined /*out*/;
            resourceInputs["instanceTemplate"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namedPorts"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["standbyPolicy"] = undefined /*out*/;
            resourceInputs["statefulPolicy"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["targetPools"] = undefined /*out*/;
            resourceInputs["targetSize"] = undefined /*out*/;
            resourceInputs["targetStoppedSize"] = undefined /*out*/;
            resourceInputs["targetSuspendedSize"] = undefined /*out*/;
            resourceInputs["updatePolicy"] = undefined /*out*/;
            resourceInputs["versions"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionInstanceGroupManager.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionInstanceGroupManager resource.
 */
export interface RegionInstanceGroupManagerArgs {
    /**
     * Specifies the instances configs overrides that should be applied for all instances in the MIG.
     */
    allInstancesConfig?: pulumi.Input<inputs.compute.alpha.InstanceGroupManagerAllInstancesConfigArgs>;
    /**
     * The autohealing policy for this managed instance group. You can specify only one value.
     */
    autoHealingPolicies?: pulumi.Input<pulumi.Input<inputs.compute.alpha.InstanceGroupManagerAutoHealingPolicyArgs>[]>;
    /**
     * The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
     */
    baseInstanceName?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
     */
    distributionPolicy?: pulumi.Input<inputs.compute.alpha.DistributionPolicyArgs>;
    /**
     * The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
     */
    failoverAction?: pulumi.Input<enums.compute.alpha.RegionInstanceGroupManagerFailoverAction>;
    /**
     * Instance lifecycle policy for this Instance Group Manager.
     */
    instanceLifecyclePolicy?: pulumi.Input<inputs.compute.alpha.InstanceGroupManagerInstanceLifecyclePolicyArgs>;
    /**
     * The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
     */
    instanceTemplate?: pulumi.Input<string>;
    /**
     * The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
     */
    name?: pulumi.Input<string>;
    /**
     * Named ports configured for the Instance Groups complementary to this Instance Group Manager.
     */
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.alpha.NamedPortArgs>[]>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Standby policy for stopped and suspended instances.
     */
    standbyPolicy?: pulumi.Input<inputs.compute.alpha.InstanceGroupManagerStandbyPolicyArgs>;
    /**
     * Stateful configuration for this Instanced Group Manager
     */
    statefulPolicy?: pulumi.Input<inputs.compute.alpha.StatefulPolicyArgs>;
    /**
     * The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
     */
    targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
     */
    targetSize?: pulumi.Input<number>;
    /**
     * The target number of stopped instances for this managed instance group. This number changes when you: - Stop instance using the stopInstances method or start instances using the startInstances method. - Manually change the targetStoppedSize using the update method. 
     */
    targetStoppedSize?: pulumi.Input<number>;
    /**
     * The target number of suspended instances for this managed instance group. This number changes when you: - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. - Manually change the targetSuspendedSize using the update method. 
     */
    targetSuspendedSize?: pulumi.Input<number>;
    /**
     * The update policy for this managed instance group.
     */
    updatePolicy?: pulumi.Input<inputs.compute.alpha.InstanceGroupManagerUpdatePolicyArgs>;
    /**
     * Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
     */
    versions?: pulumi.Input<pulumi.Input<inputs.compute.alpha.InstanceGroupManagerVersionArgs>[]>;
}
