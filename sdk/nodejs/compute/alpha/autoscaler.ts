// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an autoscaler in the specified project using the data included in the request.
 */
export class Autoscaler extends pulumi.CustomResource {
    /**
     * Get an existing Autoscaler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Autoscaler {
        return new Autoscaler(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:Autoscaler';

    /**
     * Returns true if the given object is an instance of Autoscaler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Autoscaler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Autoscaler.__pulumiType;
    }

    /**
     * The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
     *
     * If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
     */
    public readonly autoscalingPolicy!: pulumi.Output<outputs.compute.alpha.AutoscalingPolicyResponse>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    public readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * [Output Only] Type of the resource. Always compute#autoscaler for autoscalers.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * [Output Only] Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
     */
    public readonly recommendedSize!: pulumi.Output<number>;
    /**
     * [Output Only] URL of the region where the instance group resides (for autoscalers living in regional scope).
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * [Output Only] Status information of existing scaling schedules.
     */
    public readonly scalingScheduleStatus!: pulumi.Output<{[key: string]: string}>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * [Output Only] Server-defined URL for this resource with the resource id.
     */
    public readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * [Output Only] The status of the autoscaler configuration. Current set of possible values:  
     * - PENDING: Autoscaler backend hasn't read new/updated configuration. 
     * - DELETING: Configuration is being deleted. 
     * - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. 
     * - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * [Output Only] Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
     */
    public readonly statusDetails!: pulumi.Output<outputs.compute.alpha.AutoscalerStatusDetailsResponse[]>;
    /**
     * URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
     */
    public readonly target!: pulumi.Output<string>;
    /**
     * [Output Only] URL of the zone where the instance group resides (for autoscalers living in zonal scope).
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Autoscaler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoscalerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["creationTimestamp"] = args ? args.creationTimestamp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["recommendedSize"] = args ? args.recommendedSize : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["scalingScheduleStatus"] = args ? args.scalingScheduleStatus : undefined;
            inputs["selfLink"] = args ? args.selfLink : undefined;
            inputs["selfLinkWithId"] = args ? args.selfLinkWithId : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["statusDetails"] = args ? args.statusDetails : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        } else {
            inputs["autoscalingPolicy"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["recommendedSize"] = undefined /*out*/;
            inputs["region"] = undefined /*out*/;
            inputs["scalingScheduleStatus"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusDetails"] = undefined /*out*/;
            inputs["target"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Autoscaler.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Autoscaler resource.
 */
export interface AutoscalerArgs {
    /**
     * The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization.
     *
     * If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%.
     */
    readonly autoscalingPolicy?: pulumi.Input<inputs.compute.alpha.AutoscalingPolicyArgs>;
    /**
     * [Output Only] Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * [Output Only] The unique identifier for the resource. This identifier is defined by the server.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * [Output Only] Type of the resource. Always compute#autoscaler for autoscalers.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * [Output Only] Target recommended MIG size (number of instances) computed by autoscaler. Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON. This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction.
     */
    readonly recommendedSize?: pulumi.Input<number>;
    /**
     * [Output Only] URL of the region where the instance group resides (for autoscalers living in regional scope).
     */
    readonly region?: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
    /**
     * [Output Only] Status information of existing scaling schedules.
     */
    readonly scalingScheduleStatus?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * [Output Only] Server-defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * [Output Only] Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId?: pulumi.Input<string>;
    /**
     * [Output Only] The status of the autoscaler configuration. Current set of possible values:  
     * - PENDING: Autoscaler backend hasn't read new/updated configuration. 
     * - DELETING: Configuration is being deleted. 
     * - ACTIVE: Configuration is acknowledged to be effective. Some warnings might be present in the statusDetails field. 
     * - ERROR: Configuration has errors. Actionable for users. Details are present in the statusDetails field.  New values might be added in the future.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * [Output Only] Human-readable details about the current state of the autoscaler. Read the documentation for Commonly returned status messages for examples of status messages you might encounter.
     */
    readonly statusDetails?: pulumi.Input<pulumi.Input<inputs.compute.alpha.AutoscalerStatusDetailsArgs>[]>;
    /**
     * URL of the managed instance group that this autoscaler will scale. This field is required when creating an autoscaler.
     */
    readonly target?: pulumi.Input<string>;
    /**
     * [Output Only] URL of the zone where the instance group resides (for autoscalers living in zonal scope).
     */
    readonly zone: pulumi.Input<string>;
}
