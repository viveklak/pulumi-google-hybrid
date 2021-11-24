// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an Experiment in the specified Environment.
 */
export class Experiment extends pulumi.CustomResource {
    /**
     * Get an existing Experiment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Experiment {
        return new Experiment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3:Experiment';

    /**
     * Returns true if the given object is an instance of Experiment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Experiment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Experiment.__pulumiType;
    }

    /**
     * Creation time of this experiment.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * The definition of the experiment.
     */
    public readonly definition!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentDefinitionResponse>;
    /**
     * The human-readable description of the experiment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * End time of this experiment.
     */
    public readonly endTime!: pulumi.Output<string>;
    /**
     * Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
     */
    public readonly experimentLength!: pulumi.Output<string>;
    /**
     * Last update time of this experiment.
     */
    public readonly lastUpdateTime!: pulumi.Output<string>;
    /**
     * The name of the experiment. Format: projects//locations//agents//environments//experiments/..
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Inference result of the experiment.
     */
    public readonly result!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentResultResponse>;
    /**
     * The configuration for auto rollout. If set, there should be exactly two variants in the experiment (control variant being the default version of the flow), the traffic allocation for the non-control variant will gradually increase to 100% when conditions are met, and eventually replace the control variant to become the default version of the flow.
     */
    public readonly rolloutConfig!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3RolloutConfigResponse>;
    /**
     * The reason why rollout has failed. Should only be set when state is ROLLOUT_FAILED.
     */
    public readonly rolloutFailureReason!: pulumi.Output<string>;
    /**
     * State of the auto rollout process.
     */
    public readonly rolloutState!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3RolloutStateResponse>;
    /**
     * Start time of this experiment.
     */
    public readonly startTime!: pulumi.Output<string>;
    /**
     * The current state of the experiment. Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING. Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or RUNNING->DONE.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The history of updates to the experiment variants.
     */
    public readonly variantsHistory!: pulumi.Output<outputs.dialogflow.v3.GoogleCloudDialogflowCxV3VariantsHistoryResponse[]>;

    /**
     * Create a Experiment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExperimentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            resourceInputs["agentId"] = args ? args.agentId : undefined;
            resourceInputs["createTime"] = args ? args.createTime : undefined;
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["experimentLength"] = args ? args.experimentLength : undefined;
            resourceInputs["lastUpdateTime"] = args ? args.lastUpdateTime : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["result"] = args ? args.result : undefined;
            resourceInputs["rolloutConfig"] = args ? args.rolloutConfig : undefined;
            resourceInputs["rolloutFailureReason"] = args ? args.rolloutFailureReason : undefined;
            resourceInputs["rolloutState"] = args ? args.rolloutState : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["variantsHistory"] = args ? args.variantsHistory : undefined;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["definition"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["experimentLength"] = undefined /*out*/;
            resourceInputs["lastUpdateTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["result"] = undefined /*out*/;
            resourceInputs["rolloutConfig"] = undefined /*out*/;
            resourceInputs["rolloutFailureReason"] = undefined /*out*/;
            resourceInputs["rolloutState"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["variantsHistory"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Experiment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Experiment resource.
 */
export interface ExperimentArgs {
    agentId: pulumi.Input<string>;
    /**
     * Creation time of this experiment.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The definition of the experiment.
     */
    definition?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentDefinitionArgs>;
    /**
     * The human-readable description of the experiment.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * End time of this experiment.
     */
    endTime?: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    /**
     * Maximum number of days to run the experiment/rollout. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
     */
    experimentLength?: pulumi.Input<string>;
    /**
     * Last update time of this experiment.
     */
    lastUpdateTime?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The name of the experiment. Format: projects//locations//agents//environments//experiments/..
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Inference result of the experiment.
     */
    result?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3ExperimentResultArgs>;
    /**
     * The configuration for auto rollout. If set, there should be exactly two variants in the experiment (control variant being the default version of the flow), the traffic allocation for the non-control variant will gradually increase to 100% when conditions are met, and eventually replace the control variant to become the default version of the flow.
     */
    rolloutConfig?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3RolloutConfigArgs>;
    /**
     * The reason why rollout has failed. Should only be set when state is ROLLOUT_FAILED.
     */
    rolloutFailureReason?: pulumi.Input<string>;
    /**
     * State of the auto rollout process.
     */
    rolloutState?: pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3RolloutStateArgs>;
    /**
     * Start time of this experiment.
     */
    startTime?: pulumi.Input<string>;
    /**
     * The current state of the experiment. Transition triggered by Experiments.StartExperiment: DRAFT->RUNNING. Transition triggered by Experiments.CancelExperiment: DRAFT->DONE or RUNNING->DONE.
     */
    state?: pulumi.Input<enums.dialogflow.v3.ExperimentState>;
    /**
     * The history of updates to the experiment variants.
     */
    variantsHistory?: pulumi.Input<pulumi.Input<inputs.dialogflow.v3.GoogleCloudDialogflowCxV3VariantsHistoryArgs>[]>;
}
