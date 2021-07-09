// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a training or a batch prediction job.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:ml/v1:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * When the job was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * When the job processing was completed.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The details of a failure or a cancellation.
     */
    public /*out*/ readonly errorMessage!: pulumi.Output<string>;
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a job from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform job updates in order to avoid race conditions: An `etag` is returned in the response to `GetJob`, and systems are expected to put that etag in the request to `UpdateJob` to ensure that their change will be applied to the same version of the job.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The user-specified id of the job.
     */
    public readonly jobId!: pulumi.Output<string>;
    /**
     * Optional. One or more labels that you can add, to organize your jobs. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Input parameters to create a prediction job.
     */
    public readonly predictionInput!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__PredictionInputResponse>;
    /**
     * The current prediction job result.
     */
    public readonly predictionOutput!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__PredictionOutputResponse>;
    /**
     * When the job processing was started.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * The detailed state of a job.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Input parameters to create a training job.
     */
    public readonly trainingInput!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__TrainingInputResponse>;
    /**
     * The current training job result.
     */
    public readonly trainingOutput!: pulumi.Output<outputs.ml.v1.GoogleCloudMlV1__TrainingOutputResponse>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.jobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobId'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["jobId"] = args ? args.jobId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["predictionInput"] = args ? args.predictionInput : undefined;
            inputs["predictionOutput"] = args ? args.predictionOutput : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["trainingInput"] = args ? args.trainingInput : undefined;
            inputs["trainingOutput"] = args ? args.trainingOutput : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["errorMessage"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        } else {
            inputs["createTime"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["errorMessage"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["jobId"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["predictionInput"] = undefined /*out*/;
            inputs["predictionOutput"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["trainingInput"] = undefined /*out*/;
            inputs["trainingOutput"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a job from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform job updates in order to avoid race conditions: An `etag` is returned in the response to `GetJob`, and systems are expected to put that etag in the request to `UpdateJob` to ensure that their change will be applied to the same version of the job.
     */
    etag?: pulumi.Input<string>;
    /**
     * The user-specified id of the job.
     */
    jobId: pulumi.Input<string>;
    /**
     * Optional. One or more labels that you can add, to organize your jobs. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Input parameters to create a prediction job.
     */
    predictionInput?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__PredictionInputArgs>;
    /**
     * The current prediction job result.
     */
    predictionOutput?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__PredictionOutputArgs>;
    project: pulumi.Input<string>;
    /**
     * Input parameters to create a training job.
     */
    trainingInput?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__TrainingInputArgs>;
    /**
     * The current training job result.
     */
    trainingOutput?: pulumi.Input<inputs.ml.v1.GoogleCloudMlV1__TrainingOutputArgs>;
}
