// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    public static class GetEnvironment
    {
        /// <summary>
        /// Retrieves the specified Environment.
        /// </summary>
        public static Task<GetEnvironmentResult> InvokeAsync(GetEnvironmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnvironmentResult>("google-native:dialogflow/v3:getEnvironment", args ?? new GetEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified Environment.
        /// </summary>
        public static Output<GetEnvironmentResult> Invoke(GetEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEnvironmentResult>("google-native:dialogflow/v3:getEnvironment", args ?? new GetEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnvironmentArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public string EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetEnvironmentArgs()
        {
        }
    }

    public sealed class GetEnvironmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetEnvironmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEnvironmentResult
    {
        /// <summary>
        /// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The name of the environment. Format: `projects//locations//agents//environments/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The test cases config for continuous tests of this environment.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3EnvironmentTestCasesConfigResponse TestCasesConfig;
        /// <summary>
        /// Update time of this environment.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponse> VersionConfigs;

        [OutputConstructor]
        private GetEnvironmentResult(
            string description,

            string displayName,

            string name,

            Outputs.GoogleCloudDialogflowCxV3EnvironmentTestCasesConfigResponse testCasesConfig,

            string updateTime,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponse> versionConfigs)
        {
            Description = description;
            DisplayName = displayName;
            Name = name;
            TestCasesConfig = testCasesConfig;
            UpdateTime = updateTime;
            VersionConfigs = versionConfigs;
        }
    }
}
