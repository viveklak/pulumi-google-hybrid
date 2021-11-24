// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a managed instance group using the information that you specify in the request. After the group is created, instances in the group are created using the specified instance template. This operation is marked as DONE when the group is created even if the instances in the group have not yet been created. You must separately verify the status of the individual instances with the listmanagedinstances method. A managed instance group can have up to 1000 VM instances per group. Please contact Cloud Support if you need an increase in this limit.
 */
export class InstanceGroupManager extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGroupManager resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InstanceGroupManager {
        return new InstanceGroupManager(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:InstanceGroupManager';

    /**
     * Returns true if the given object is an instance of InstanceGroupManager.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceGroupManager {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceGroupManager.__pulumiType;
    }

    /**
     * The autohealing policy for this managed instance group. You can specify only one value.
     */
    public readonly autoHealingPolicies!: pulumi.Output<outputs.compute.beta.InstanceGroupManagerAutoHealingPolicyResponse[]>;
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
    public /*out*/ readonly currentActions!: pulumi.Output<outputs.compute.beta.InstanceGroupManagerActionsSummaryResponse>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
     */
    public readonly distributionPolicy!: pulumi.Output<outputs.compute.beta.DistributionPolicyResponse>;
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
    public readonly namedPorts!: pulumi.Output<outputs.compute.beta.NamedPortResponse[]>;
    /**
     * The URL of the region where the managed instance group resides (for regional resources).
     */
    public /*out*/ readonly region!: pulumi.Output<string>;
    /**
     * The URL for this managed instance group. The server defines this URL.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Stateful configuration for this Instanced Group Manager
     */
    public readonly statefulPolicy!: pulumi.Output<outputs.compute.beta.StatefulPolicyResponse>;
    /**
     * The status of this managed instance group.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.compute.beta.InstanceGroupManagerStatusResponse>;
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
    public readonly updatePolicy!: pulumi.Output<outputs.compute.beta.InstanceGroupManagerUpdatePolicyResponse>;
    /**
     * Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
     */
    public readonly versions!: pulumi.Output<outputs.compute.beta.InstanceGroupManagerVersionResponse[]>;
    /**
     * The URL of a zone where the managed instance group is located (for zonal resources).
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceGroupManager resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceGroupManagerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["autoHealingPolicies"] = args ? args.autoHealingPolicies : undefined;
            resourceInputs["baseInstanceName"] = args ? args.baseInstanceName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["distributionPolicy"] = args ? args.distributionPolicy : undefined;
            resourceInputs["failoverAction"] = args ? args.failoverAction : undefined;
            resourceInputs["instanceTemplate"] = args ? args.instanceTemplate : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namedPorts"] = args ? args.namedPorts : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["statefulPolicy"] = args ? args.statefulPolicy : undefined;
            resourceInputs["targetPools"] = args ? args.targetPools : undefined;
            resourceInputs["targetSize"] = args ? args.targetSize : undefined;
            resourceInputs["updatePolicy"] = args ? args.updatePolicy : undefined;
            resourceInputs["versions"] = args ? args.versions : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["currentActions"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["instanceGroup"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["autoHealingPolicies"] = undefined /*out*/;
            resourceInputs["baseInstanceName"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["currentActions"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["distributionPolicy"] = undefined /*out*/;
            resourceInputs["failoverAction"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["instanceGroup"] = undefined /*out*/;
            resourceInputs["instanceTemplate"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namedPorts"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["statefulPolicy"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["targetPools"] = undefined /*out*/;
            resourceInputs["targetSize"] = undefined /*out*/;
            resourceInputs["updatePolicy"] = undefined /*out*/;
            resourceInputs["versions"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InstanceGroupManager.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a InstanceGroupManager resource.
 */
export interface InstanceGroupManagerArgs {
    /**
     * The autohealing policy for this managed instance group. You can specify only one value.
     */
    autoHealingPolicies?: pulumi.Input<pulumi.Input<inputs.compute.beta.InstanceGroupManagerAutoHealingPolicyArgs>[]>;
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
    distributionPolicy?: pulumi.Input<inputs.compute.beta.DistributionPolicyArgs>;
    /**
     * The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
     */
    failoverAction?: pulumi.Input<enums.compute.beta.InstanceGroupManagerFailoverAction>;
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
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.beta.NamedPortArgs>[]>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Stateful configuration for this Instanced Group Manager
     */
    statefulPolicy?: pulumi.Input<inputs.compute.beta.StatefulPolicyArgs>;
    /**
     * The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
     */
    targetPools?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
     */
    targetSize?: pulumi.Input<number>;
    /**
     * The update policy for this managed instance group.
     */
    updatePolicy?: pulumi.Input<inputs.compute.beta.InstanceGroupManagerUpdatePolicyArgs>;
    /**
     * Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
     */
    versions?: pulumi.Input<pulumi.Input<inputs.compute.beta.InstanceGroupManagerVersionArgs>[]>;
    zone?: pulumi.Input<string>;
}
