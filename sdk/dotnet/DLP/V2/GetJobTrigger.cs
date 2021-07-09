// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2
{
    public static class GetJobTrigger
    {
        /// <summary>
        /// Gets a job trigger. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
        /// </summary>
        public static Task<GetJobTriggerResult> InvokeAsync(GetJobTriggerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobTriggerResult>("google-native:dlp/v2:getJobTrigger", args ?? new GetJobTriggerArgs(), options.WithVersion());
    }


    public sealed class GetJobTriggerArgs : Pulumi.InvokeArgs
    {
        [Input("jobTriggerId", required: true)]
        public string JobTriggerId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetJobTriggerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobTriggerResult
    {
        /// <summary>
        /// The creation timestamp of a triggeredJob.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// User provided description (max 256 chars)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Display name (max 100 chars)
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2ErrorResponse> Errors;
        /// <summary>
        /// For inspect jobs, a snapshot of the configuration.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2InspectJobConfigResponse InspectJob;
        /// <summary>
        /// The timestamp of the last time this trigger executed.
        /// </summary>
        public readonly string LastRunTime;
        /// <summary>
        /// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A status for this trigger.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2TriggerResponse> Triggers;
        /// <summary>
        /// The last update timestamp of a triggeredJob.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetJobTriggerResult(
            string createTime,

            string description,

            string displayName,

            ImmutableArray<Outputs.GooglePrivacyDlpV2ErrorResponse> errors,

            Outputs.GooglePrivacyDlpV2InspectJobConfigResponse inspectJob,

            string lastRunTime,

            string name,

            string status,

            ImmutableArray<Outputs.GooglePrivacyDlpV2TriggerResponse> triggers,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            DisplayName = displayName;
            Errors = errors;
            InspectJob = inspectJob;
            LastRunTime = lastRunTime;
            Name = name;
            Status = status;
            Triggers = triggers;
            UpdateTime = updateTime;
        }
    }
}
