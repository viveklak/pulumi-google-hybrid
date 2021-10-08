// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1
{
    public static class GetIssueModel
    {
        /// <summary>
        /// Gets an issue model.
        /// </summary>
        public static Task<GetIssueModelResult> InvokeAsync(GetIssueModelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIssueModelResult>("google-native:contactcenterinsights/v1:getIssueModel", args ?? new GetIssueModelArgs(), options.WithVersion());
    }


    public sealed class GetIssueModelArgs : Pulumi.InvokeArgs
    {
        [Input("issueModelId", required: true)]
        public string IssueModelId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetIssueModelArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIssueModelResult
    {
        /// <summary>
        /// The time at which this issue model was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The representative name for the issue model.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Configs for the input data that used to create the issue model.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigResponse InputDataConfig;
        /// <summary>
        /// Immutable. The resource name of the issue model. Format: projects/{project}/locations/{location}/issueModels/{issue_model}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the model.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Immutable. The issue model's label statistics on its training data.
        /// </summary>
        public readonly Outputs.GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponse TrainingStats;
        /// <summary>
        /// The most recent time at which the issue model was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetIssueModelResult(
            string createTime,

            string displayName,

            Outputs.GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigResponse inputDataConfig,

            string name,

            string state,

            Outputs.GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponse trainingStats,

            string updateTime)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            InputDataConfig = inputDataConfig;
            Name = name;
            State = state;
            TrainingStats = trainingStats;
            UpdateTime = updateTime;
        }
    }
}