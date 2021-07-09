// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new canary evaluation for an organization.
 */
export class CanaryEvaluation extends pulumi.CustomResource {
    /**
     * Get an existing CanaryEvaluation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CanaryEvaluation {
        return new CanaryEvaluation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:CanaryEvaluation';

    /**
     * Returns true if the given object is an instance of CanaryEvaluation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CanaryEvaluation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CanaryEvaluation.__pulumiType;
    }

    /**
     * The stable version that is serving requests.
     */
    public readonly control!: pulumi.Output<string>;
    /**
     * Create time of the canary evaluation.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * End time for the evaluation's analysis.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * Labels used to filter the metrics used for a canary evaluation.
     */
    public readonly metricLabels!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1CanaryEvaluationMetricLabelsResponse>;
    /**
     * Name of the canary evalution.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Start time for the canary evaluation's analysis.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * The current state of the canary evaluation.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The newer version that is serving requests.
     */
    public readonly treatment!: pulumi.Output<string>;
    /**
     * The resulting verdict of the canary evaluations: NONE, PASS, or FAIL.
     */
    public /*out*/ readonly verdict!: pulumi.Output<string>;

    /**
     * Create a CanaryEvaluation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CanaryEvaluationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.control === undefined) && !opts.urn) {
                throw new Error("Missing required property 'control'");
            }
            if ((!args || args.endTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endTime'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.metricLabels === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricLabels'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.startTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startTime'");
            }
            if ((!args || args.treatment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'treatment'");
            }
            inputs["control"] = args ? args.control : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["metricLabels"] = args ? args.metricLabels : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["treatment"] = args ? args.treatment : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["verdict"] = undefined /*out*/;
        } else {
            inputs["control"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["metricLabels"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["treatment"] = undefined /*out*/;
            inputs["verdict"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CanaryEvaluation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CanaryEvaluation resource.
 */
export interface CanaryEvaluationArgs {
    /**
     * The stable version that is serving requests.
     */
    control: pulumi.Input<string>;
    /**
     * End time for the evaluation's analysis.
     */
    endTime: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    /**
     * Labels used to filter the metrics used for a canary evaluation.
     */
    metricLabels: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1CanaryEvaluationMetricLabelsArgs>;
    organizationId: pulumi.Input<string>;
    /**
     * Start time for the canary evaluation's analysis.
     */
    startTime: pulumi.Input<string>;
    /**
     * The newer version that is serving requests.
     */
    treatment: pulumi.Input<string>;
}
