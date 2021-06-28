// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns all of the details about the specified managed instance group. Gets a list of available managed instance groups by making a list() request.
 */
export function getInstanceGroupManager(args: GetInstanceGroupManagerArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceGroupManagerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/alpha:getInstanceGroupManager", {
        "instanceGroupManager": args.instanceGroupManager,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

export interface GetInstanceGroupManagerArgs {
    instanceGroupManager: string;
    project: string;
    zone: string;
}

export interface GetInstanceGroupManagerResult {
    /**
     * The autohealing policy for this managed instance group. You can specify only one value.
     */
    readonly autoHealingPolicies: outputs.compute.alpha.InstanceGroupManagerAutoHealingPolicyResponse[];
    /**
     * The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
     */
    readonly baseInstanceName: string;
    /**
     * The creation timestamp for this managed instance group in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
     */
    readonly currentActions: outputs.compute.alpha.InstanceGroupManagerActionsSummaryResponse;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
     */
    readonly distributionPolicy: outputs.compute.alpha.DistributionPolicyResponse;
    /**
     * The action to perform in case of zone failure. Only one value is supported, NO_FAILOVER. The default is NO_FAILOVER.
     */
    readonly failoverAction: string;
    /**
     * Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet.
     *
     * To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
     */
    readonly fingerprint: string;
    /**
     * The URL of the Instance Group resource.
     */
    readonly instanceGroup: string;
    /**
     * Instance lifecycle policy for this Instance Group Manager.
     */
    readonly instanceLifecyclePolicy: outputs.compute.alpha.InstanceGroupManagerInstanceLifecyclePolicyResponse;
    /**
     * The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
     */
    readonly instanceTemplate: string;
    /**
     * The resource type, which is always compute#instanceGroupManager for managed instance groups.
     */
    readonly kind: string;
    /**
     * The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
     */
    readonly name: string;
    /**
     * Named ports configured for the Instance Groups complementary to this Instance Group Manager.
     */
    readonly namedPorts: outputs.compute.alpha.NamedPortResponse[];
    /**
     * The URL of the region where the managed instance group resides (for regional resources).
     */
    readonly region: string;
    /**
     * The URL for this managed instance group. The server defines this URL.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account {projectNumber}@cloudservices.gserviceaccount.com is used.
     */
    readonly serviceAccount: string;
    /**
     * Stateful configuration for this Instanced Group Manager
     */
    readonly statefulPolicy: outputs.compute.alpha.StatefulPolicyResponse;
    /**
     * The status of this managed instance group.
     */
    readonly status: outputs.compute.alpha.InstanceGroupManagerStatusResponse;
    /**
     * The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
     */
    readonly targetPools: string[];
    /**
     * The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
     */
    readonly targetSize: number;
    /**
     * The target number of stopped instances for this managed instance group. This number changes when you:  
     * - Stop instance using the stopInstances method or start instances using the startInstances method. 
     * - Manually change the targetStoppedSize using the update method.
     */
    readonly targetStoppedSize: number;
    /**
     * The target number of suspended instances for this managed instance group. This number changes when you:  
     * - Suspend instance using the suspendInstances method or resume instances using the resumeInstances method. 
     * - Manually change the targetSuspendedSize using the update method.
     */
    readonly targetSuspendedSize: number;
    /**
     * The update policy for this managed instance group.
     */
    readonly updatePolicy: outputs.compute.alpha.InstanceGroupManagerUpdatePolicyResponse;
    /**
     * Specifies the instance templates used by this managed instance group to create instances.
     *
     * Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
     */
    readonly versions: outputs.compute.alpha.InstanceGroupManagerVersionResponse[];
    /**
     * The URL of a zone where the managed instance group is located (for zonal resources).
     */
    readonly zone: string;
}
