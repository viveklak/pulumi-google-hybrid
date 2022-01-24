// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Returns the job data.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:transcoder/v1beta1:getJob", {
        "jobId": args.jobId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetJobArgs {
    jobId: string;
    location: string;
    project?: string;
}

export interface GetJobResult {
    /**
     * The configuration for this job.
     */
    readonly config: outputs.transcoder.v1beta1.JobConfigResponse;
    /**
     * The time the job was created.
     */
    readonly createTime: string;
    /**
     * The time the transcoding finished.
     */
    readonly endTime: string;
    /**
     * List of failure details. This property may contain additional information about the failure when `failure_reason` is present. *Note*: This feature is not yet available.
     */
    readonly failureDetails: outputs.transcoder.v1beta1.FailureDetailResponse[];
    /**
     * A description of the reason for the failure. This property is always present when `state` is `FAILED`.
     */
    readonly failureReason: string;
    /**
     * Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
     */
    readonly inputUri: string;
    /**
     * The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
     */
    readonly name: string;
    /**
     * The origin URI. *Note*: This feature is not yet available.
     */
    readonly originUri: outputs.transcoder.v1beta1.OriginUriResponse;
    /**
     * Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
     */
    readonly outputUri: string;
    /**
     * Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
     */
    readonly priority: number;
    /**
     * Estimated fractional progress, from `0` to `1` for each step. *Note*: This feature is not yet available.
     */
    readonly progress: outputs.transcoder.v1beta1.ProgressResponse;
    /**
     * The time the transcoding started.
     */
    readonly startTime: string;
    /**
     * The current state of the job.
     */
    readonly state: string;
    /**
     * Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
     */
    readonly templateId: string;
    /**
     * Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
     */
    readonly ttlAfterCompletionDays: number;
}

export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply(a => getJob(a, opts))
}

export interface GetJobOutputArgs {
    jobId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
