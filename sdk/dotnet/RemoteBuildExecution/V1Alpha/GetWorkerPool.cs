// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RemoteBuildExecution.V1Alpha
{
    public static class GetWorkerPool
    {
        /// <summary>
        /// Returns the specified worker pool.
        /// </summary>
        public static Task<GetWorkerPoolResult> InvokeAsync(GetWorkerPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkerPoolResult>("google-native:remotebuildexecution/v1alpha:getWorkerPool", args ?? new GetWorkerPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified worker pool.
        /// </summary>
        public static Output<GetWorkerPoolResult> Invoke(GetWorkerPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetWorkerPoolResult>("google-native:remotebuildexecution/v1alpha:getWorkerPool", args ?? new GetWorkerPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkerPoolArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("workerpoolId", required: true)]
        public string WorkerpoolId { get; set; } = null!;

        public GetWorkerPoolArgs()
        {
        }
    }

    public sealed class GetWorkerPoolInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("workerpoolId", required: true)]
        public Input<string> WorkerpoolId { get; set; } = null!;

        public GetWorkerPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkerPoolResult
    {
        /// <summary>
        /// The autoscale policy to apply on a pool.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse Autoscale;
        /// <summary>
        /// Channel specifies the release channel of the pool.
        /// </summary>
        public readonly string Channel;
        /// <summary>
        /// WorkerPool resource name formatted as: `projects/[PROJECT_ID]/instances/[INSTANCE_ID]/workerpools/[POOL_ID]`. name should not be populated when creating a worker pool since it is provided in the `poolId` field.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the worker pool.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Specifies the properties, such as machine type and disk size, used for creating workers in a worker pool.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse WorkerConfig;
        /// <summary>
        /// The desired number of workers in the worker pool. Must be a value between 0 and 15000.
        /// </summary>
        public readonly string WorkerCount;

        [OutputConstructor]
        private GetWorkerPoolResult(
            Outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaAutoscaleResponse autoscale,

            string channel,

            string name,

            string state,

            Outputs.GoogleDevtoolsRemotebuildexecutionAdminV1alphaWorkerConfigResponse workerConfig,

            string workerCount)
        {
            Autoscale = autoscale;
            Channel = channel;
            Name = name;
            State = state;
            WorkerConfig = workerConfig;
            WorkerCount = workerCount;
        }
    }
}
