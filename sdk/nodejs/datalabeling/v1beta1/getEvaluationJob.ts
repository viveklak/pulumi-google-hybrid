// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets an evaluation job by resource name.
 */
export function getEvaluationJob(args: GetEvaluationJobArgs, opts?: pulumi.InvokeOptions): Promise<GetEvaluationJobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:datalabeling/v1beta1:getEvaluationJob", {
        "evaluationJobId": args.evaluationJobId,
        "project": args.project,
    }, opts);
}

export interface GetEvaluationJobArgs {
    evaluationJobId: string;
    project: string;
}

export interface GetEvaluationJobResult {
    /**
     * Name of the AnnotationSpecSet describing all the labels that your machine learning model outputs. You must create this resource before you create an evaluation job and provide its name in the following format: "projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}"
     */
    readonly annotationSpecSet: string;
    /**
     * Every time the evaluation job runs and an error occurs, the failed attempt is appended to this array.
     */
    readonly attempts: outputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1AttemptResponse[];
    /**
     * Timestamp of when this evaluation job was created.
     */
    readonly createTime: string;
    /**
     * Description of the job. The description can be up to 25,000 characters long.
     */
    readonly description: string;
    /**
     * Configuration details for the evaluation job.
     */
    readonly evaluationJobConfig: outputs.datalabeling.v1beta1.GoogleCloudDatalabelingV1beta1EvaluationJobConfigResponse;
    /**
     * Whether you want Data Labeling Service to provide ground truth labels for prediction input. If you want the service to assign human labelers to annotate your data, set this to `true`. If you want to provide your own ground truth labels in the evaluation job's BigQuery table, set this to `false`.
     */
    readonly labelMissingGroundTruth: boolean;
    /**
     * The [AI Platform Prediction model version](/ml-engine/docs/prediction-overview) to be evaluated. Prediction input and output is sampled from this model version. When creating an evaluation job, specify the model version in the following format: "projects/{project_id}/models/{model_name}/versions/{version_name}" There can only be one evaluation job per model version.
     */
    readonly modelVersion: string;
    /**
     * After you create a job, Data Labeling Service assigns a name to the job with the following format: "projects/{project_id}/evaluationJobs/ {evaluation_job_id}"
     */
    readonly name: string;
    /**
     * Describes the interval at which the job runs. This interval must be at least 1 day, and it is rounded to the nearest day. For example, if you specify a 50-hour interval, the job runs every 2 days. You can provide the schedule in [crontab format](/scheduler/docs/configuring/cron-job-schedules) or in an [English-like format](/appengine/docs/standard/python/config/cronref#schedule_format). Regardless of what you specify, the job will run at 10:00 AM UTC. Only the interval from this schedule is used, not the specific time of day.
     */
    readonly schedule: string;
    /**
     * Describes the current state of the job.
     */
    readonly state: string;
}
