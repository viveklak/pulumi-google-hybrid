// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Transcoder.V1Beta1
{
    public static class GetJob
    {
        /// <summary>
        /// Returns the job data.
        /// </summary>
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("google-native:transcoder/v1beta1:getJob", args ?? new GetJobArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the job data.
        /// </summary>
        public static Output<GetJobResult> Invoke(GetJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetJobResult>("google-native:transcoder/v1beta1:getJob", args ?? new GetJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobArgs : Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public string JobId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetJobArgs()
        {
        }
    }

    public sealed class GetJobInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetJobInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobResult
    {
        /// <summary>
        /// The configuration for this job.
        /// </summary>
        public readonly Outputs.JobConfigResponse Config;
        /// <summary>
        /// The time the job was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time the transcoding finished.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// List of failure details. This property may contain additional information about the failure when `failure_reason` is present. *Note*: This feature is not yet available.
        /// </summary>
        public readonly ImmutableArray<Outputs.FailureDetailResponse> FailureDetails;
        /// <summary>
        /// A description of the reason for the failure. This property is always present when `state` is `FAILED`.
        /// </summary>
        public readonly string FailureReason;
        /// <summary>
        /// Input only. Specify the `input_uri` to populate empty `uri` fields in each element of `Job.config.inputs` or `JobTemplate.config.inputs` when using template. URI of the media. Input files must be at least 5 seconds in duration and stored in Cloud Storage (for example, `gs://bucket/inputs/file.mp4`).
        /// </summary>
        public readonly string InputUri;
        /// <summary>
        /// The resource name of the job. Format: `projects/{project}/locations/{location}/jobs/{job}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The origin URI. *Note*: This feature is not yet available.
        /// </summary>
        public readonly Outputs.OriginUriResponse OriginUri;
        /// <summary>
        /// Input only. Specify the `output_uri` to populate an empty `Job.config.output.uri` or `JobTemplate.config.output.uri` when using template. URI for the output file(s). For example, `gs://my-bucket/outputs/`.
        /// </summary>
        public readonly string OutputUri;
        /// <summary>
        /// Specify the priority of the job. Enter a value between 0 and 100, where 0 is the lowest priority and 100 is the highest priority. The default is 0.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Estimated fractional progress, from `0` to `1` for each step. *Note*: This feature is not yet available.
        /// </summary>
        public readonly Outputs.ProgressResponse Progress;
        /// <summary>
        /// The time the transcoding started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The current state of the job.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Input only. Specify the `template_id` to use for populating `Job.config`. The default is `preset/web-hd`. Preset Transcoder templates: - `preset/{preset_id}` - User defined JobTemplate: `{job_template_id}`
        /// </summary>
        public readonly string TemplateId;
        /// <summary>
        /// Job time to live value in days, which will be effective after job completion. Job should be deleted automatically after the given TTL. Enter a value between 1 and 90. The default is 30.
        /// </summary>
        public readonly int TtlAfterCompletionDays;

        [OutputConstructor]
        private GetJobResult(
            Outputs.JobConfigResponse config,

            string createTime,

            string endTime,

            ImmutableArray<Outputs.FailureDetailResponse> failureDetails,

            string failureReason,

            string inputUri,

            string name,

            Outputs.OriginUriResponse originUri,

            string outputUri,

            int priority,

            Outputs.ProgressResponse progress,

            string startTime,

            string state,

            string templateId,

            int ttlAfterCompletionDays)
        {
            Config = config;
            CreateTime = createTime;
            EndTime = endTime;
            FailureDetails = failureDetails;
            FailureReason = failureReason;
            InputUri = inputUri;
            Name = name;
            OriginUri = originUri;
            OutputUri = outputUri;
            Priority = priority;
            Progress = progress;
            StartTime = startTime;
            State = state;
            TemplateId = templateId;
            TtlAfterCompletionDays = ttlAfterCompletionDays;
        }
    }
}
