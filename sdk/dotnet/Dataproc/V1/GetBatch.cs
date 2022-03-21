// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1
{
    public static class GetBatch
    {
        /// <summary>
        /// Gets the batch workload resource representation.
        /// </summary>
        public static Task<GetBatchResult> InvokeAsync(GetBatchArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBatchResult>("google-native:dataproc/v1:getBatch", args ?? new GetBatchArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the batch workload resource representation.
        /// </summary>
        public static Output<GetBatchResult> Invoke(GetBatchInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBatchResult>("google-native:dataproc/v1:getBatch", args ?? new GetBatchInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBatchArgs : Pulumi.InvokeArgs
    {
        [Input("batchId", required: true)]
        public string BatchId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetBatchArgs()
        {
        }
    }

    public sealed class GetBatchInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("batchId", required: true)]
        public Input<string> BatchId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetBatchInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBatchResult
    {
        /// <summary>
        /// The time when the batch was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The email address of the user who created the batch.
        /// </summary>
        public readonly string Creator;
        /// <summary>
        /// Optional. Environment configuration for the batch execution.
        /// </summary>
        public readonly Outputs.EnvironmentConfigResponse EnvironmentConfig;
        /// <summary>
        /// Optional. The labels to associate with this batch. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a batch.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource name of the batch.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource name of the operation associated with this batch.
        /// </summary>
        public readonly string Operation;
        /// <summary>
        /// Optional. PySpark batch config.
        /// </summary>
        public readonly Outputs.PySparkBatchResponse PysparkBatch;
        /// <summary>
        /// Optional. Runtime configuration for the batch execution.
        /// </summary>
        public readonly Outputs.RuntimeConfigResponse RuntimeConfig;
        /// <summary>
        /// Runtime information about batch execution.
        /// </summary>
        public readonly Outputs.RuntimeInfoResponse RuntimeInfo;
        /// <summary>
        /// Optional. Spark batch config.
        /// </summary>
        public readonly Outputs.SparkBatchResponse SparkBatch;
        /// <summary>
        /// Optional. SparkR batch config.
        /// </summary>
        public readonly Outputs.SparkRBatchResponse SparkRBatch;
        /// <summary>
        /// Optional. SparkSql batch config.
        /// </summary>
        public readonly Outputs.SparkSqlBatchResponse SparkSqlBatch;
        /// <summary>
        /// The state of the batch.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Historical state information for the batch.
        /// </summary>
        public readonly ImmutableArray<Outputs.StateHistoryResponse> StateHistory;
        /// <summary>
        /// Batch state details, such as a failure description if the state is FAILED.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// The time when the batch entered a current state.
        /// </summary>
        public readonly string StateTime;
        /// <summary>
        /// A batch UUID (Unique Universal Identifier). The service generates this value when it creates the batch.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetBatchResult(
            string createTime,

            string creator,

            Outputs.EnvironmentConfigResponse environmentConfig,

            ImmutableDictionary<string, string> labels,

            string name,

            string operation,

            Outputs.PySparkBatchResponse pysparkBatch,

            Outputs.RuntimeConfigResponse runtimeConfig,

            Outputs.RuntimeInfoResponse runtimeInfo,

            Outputs.SparkBatchResponse sparkBatch,

            Outputs.SparkRBatchResponse sparkRBatch,

            Outputs.SparkSqlBatchResponse sparkSqlBatch,

            string state,

            ImmutableArray<Outputs.StateHistoryResponse> stateHistory,

            string stateMessage,

            string stateTime,

            string uuid)
        {
            CreateTime = createTime;
            Creator = creator;
            EnvironmentConfig = environmentConfig;
            Labels = labels;
            Name = name;
            Operation = operation;
            PysparkBatch = pysparkBatch;
            RuntimeConfig = runtimeConfig;
            RuntimeInfo = runtimeInfo;
            SparkBatch = sparkBatch;
            SparkRBatch = sparkRBatch;
            SparkSqlBatch = sparkSqlBatch;
            State = state;
            StateHistory = stateHistory;
            StateMessage = stateMessage;
            StateTime = stateTime;
            Uuid = uuid;
        }
    }
}
