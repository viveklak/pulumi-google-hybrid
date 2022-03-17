// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1
{
    public static class GetAgentPool
    {
        /// <summary>
        /// Gets an agent pool.
        /// </summary>
        public static Task<GetAgentPoolResult> InvokeAsync(GetAgentPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAgentPoolResult>("google-native:storagetransfer/v1:getAgentPool", args ?? new GetAgentPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an agent pool.
        /// </summary>
        public static Output<GetAgentPoolResult> Invoke(GetAgentPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAgentPoolResult>("google-native:storagetransfer/v1:getAgentPool", args ?? new GetAgentPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAgentPoolArgs : Pulumi.InvokeArgs
    {
        [Input("agentPoolId", required: true)]
        public string AgentPoolId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAgentPoolArgs()
        {
        }
    }

    public sealed class GetAgentPoolInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("agentPoolId", required: true)]
        public Input<string> AgentPoolId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAgentPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAgentPoolResult
    {
        /// <summary>
        /// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
        /// </summary>
        public readonly Outputs.BandwidthLimitResponse BandwidthLimit;
        /// <summary>
        /// Specifies the client-specified AgentPool description.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Specifies a unique string that identifies the agent pool. Format: `projects/{project_id}/agentPools/{agent_pool_id}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the state of the AgentPool.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetAgentPoolResult(
            Outputs.BandwidthLimitResponse bandwidthLimit,

            string displayName,

            string name,

            string state)
        {
            BandwidthLimit = bandwidthLimit;
            DisplayName = displayName;
            Name = name;
            State = state;
        }
    }
}
