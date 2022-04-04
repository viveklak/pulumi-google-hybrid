// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * @deprecated google-hybrid.compute.Autoscalar has been deprecated in favor of google-hybrid.compute.Autoscaler
 */
export class Autoscalar extends pulumi.CustomResource {
    /**
     * Get an existing Autoscalar resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoscalarState, opts?: pulumi.CustomResourceOptions): Autoscalar {
        pulumi.log.warn("Autoscalar is deprecated: google-hybrid.compute.Autoscalar has been deprecated in favor of google-hybrid.compute.Autoscaler")
        return new Autoscalar(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-hybrid:compute/classic:Autoscalar';

    /**
     * Returns true if the given object is an instance of Autoscalar.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Autoscalar {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Autoscalar.__pulumiType;
    }

    /**
     * The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an
     * autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the
     * default will be to autoscale based on cpuUtilization to 0.6 or 60%.
     */
    public readonly autoscalingPolicy!: pulumi.Output<outputs.compute.AutoscalarAutoscalingPolicy.AutoscalarAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource. The name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * URL of the managed instance group that this autoscaler will scale.
     */
    public readonly target!: pulumi.Output<string>;
    /**
     * URL of the zone where the instance group resides.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Autoscalar resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated google-hybrid.compute.Autoscalar has been deprecated in favor of google-hybrid.compute.Autoscaler */
    constructor(name: string, args: AutoscalarArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated google-hybrid.compute.Autoscalar has been deprecated in favor of google-hybrid.compute.Autoscaler */
    constructor(name: string, argsOrState?: AutoscalarArgs | AutoscalarState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Autoscalar is deprecated: google-hybrid.compute.Autoscalar has been deprecated in favor of google-hybrid.compute.Autoscaler")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoscalarState | undefined;
            resourceInputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as AutoscalarArgs | undefined;
            if ((!args || args.autoscalingPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscalingPolicy'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            resourceInputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Autoscalar.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Autoscalar resources.
 */
export interface AutoscalarState {
    /**
     * The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an
     * autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the
     * default will be to autoscale based on cpuUtilization to 0.6 or 60%.
     */
    autoscalingPolicy?: pulumi.Input<inputs.compute.AutoscalarAutoscalingPolicy.AutoscalarAutoscalingPolicyArgs>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    selfLink?: pulumi.Input<string>;
    /**
     * URL of the managed instance group that this autoscaler will scale.
     */
    target?: pulumi.Input<string>;
    /**
     * URL of the zone where the instance group resides.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Autoscalar resource.
 */
export interface AutoscalarArgs {
    /**
     * The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an
     * autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the
     * default will be to autoscale based on cpuUtilization to 0.6 or 60%.
     */
    autoscalingPolicy: pulumi.Input<inputs.compute.AutoscalarAutoscalingPolicy.AutoscalarAutoscalingPolicyArgs>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * URL of the managed instance group that this autoscaler will scale.
     */
    target: pulumi.Input<string>;
    /**
     * URL of the zone where the instance group resides.
     */
    zone?: pulumi.Input<string>;
}