// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Adds a user provided trial to a study.
 */
export class Trial extends pulumi.CustomResource {
    /**
     * Get an existing Trial resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Trial {
        return new Trial(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:ml/v1:Trial';

    /**
     * Returns true if the given object is an instance of Trial.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trial {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trial.__pulumiType;
    }

    /**
     * The identifier of the client that originally requested this trial.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * Time at which the trial's status changed to COMPLETED.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The final measurement containing the objective value.
     */
    public readonly finalMeasurement!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__MeasurementResponse>;
    /**
     * A human readable string describing why the trial is infeasible. This should only be set if trial_infeasible is true.
     */
    public /*out*/ readonly infeasibleReason!: pulumi.Output<string>;
    /**
     * A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
     */
    public readonly measurements!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__MeasurementResponse[]>;
    /**
     * Name of the trial assigned by the service.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parameters of the trial.
     */
    public readonly parameters!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1_Trial_ParameterResponse[]>;
    /**
     * Time at which the trial was started.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * The detailed state of a trial.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * If true, the parameters in this trial are not attempted again.
     */
    public /*out*/ readonly trialInfeasible!: pulumi.Output<boolean>;

    /**
     * Create a Trial resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrialArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.studyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'studyId'");
            }
            inputs["finalMeasurement"] = args ? args.finalMeasurement : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["measurements"] = args ? args.measurements : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["studyId"] = args ? args.studyId : undefined;
            inputs["clientId"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["infeasibleReason"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["trialInfeasible"] = undefined /*out*/;
        } else {
            inputs["clientId"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["finalMeasurement"] = undefined /*out*/;
            inputs["infeasibleReason"] = undefined /*out*/;
            inputs["measurements"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["trialInfeasible"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Trial.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Trial resource.
 */
export interface TrialArgs {
    /**
     * The final measurement containing the objective value.
     */
    finalMeasurement?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__MeasurementArgs>;
    location: pulumi.Input<string>;
    /**
     * A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
     */
    measurements?: pulumi.Input<pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__MeasurementArgs>[]>;
    /**
     * The parameters of the trial.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.ml.v1.GoogleCloudMlV1_Trial_ParameterArgs>[]>;
    project: pulumi.Input<string>;
    /**
     * The detailed state of a trial.
     */
    state?: pulumi.Input<enums.ml.v1.TrialState>;
    studyId: pulumi.Input<string>;
}