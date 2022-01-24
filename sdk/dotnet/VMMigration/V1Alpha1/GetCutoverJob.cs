// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1
{
    public static class GetCutoverJob
    {
        /// <summary>
        /// Gets details of a single CutoverJob.
        /// </summary>
        public static Task<GetCutoverJobResult> InvokeAsync(GetCutoverJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCutoverJobResult>("google-native:vmmigration/v1alpha1:getCutoverJob", args ?? new GetCutoverJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single CutoverJob.
        /// </summary>
        public static Output<GetCutoverJobResult> Invoke(GetCutoverJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCutoverJobResult>("google-native:vmmigration/v1alpha1:getCutoverJob", args ?? new GetCutoverJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCutoverJobArgs : Pulumi.InvokeArgs
    {
        [Input("cutoverJobId", required: true)]
        public string CutoverJobId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public string MigratingVmId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetCutoverJobArgs()
        {
        }
    }

    public sealed class GetCutoverJobInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("cutoverJobId", required: true)]
        public Input<string> CutoverJobId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public Input<string> MigratingVmId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetCutoverJobInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCutoverJobResult
    {
        /// <summary>
        /// Details of the target VM in Compute Engine.
        /// </summary>
        public readonly Outputs.ComputeEngineTargetDetailsResponse ComputeEngineTargetDetails;
        /// <summary>
        /// The time the cutover job was created (as an API call, not when it was actually created in the target).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Provides details for the errors that led to the Cutover Job's state.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The name of the cutover job.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current progress in percentage of the cutover job.
        /// </summary>
        public readonly int Progress;
        /// <summary>
        /// The current progress in percentage of the cutover job.
        /// </summary>
        public readonly int ProgressPercent;
        /// <summary>
        /// State of the cutover job.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// A message providing possible extra details about the current state.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// The time the state was last updated.
        /// </summary>
        public readonly string StateTime;

        [OutputConstructor]
        private GetCutoverJobResult(
            Outputs.ComputeEngineTargetDetailsResponse computeEngineTargetDetails,

            string createTime,

            Outputs.StatusResponse error,

            string name,

            int progress,

            int progressPercent,

            string state,

            string stateMessage,

            string stateTime)
        {
            ComputeEngineTargetDetails = computeEngineTargetDetails;
            CreateTime = createTime;
            Error = error;
            Name = name;
            Progress = progress;
            ProgressPercent = progressPercent;
            State = state;
            StateMessage = stateMessage;
            StateTime = stateTime;
        }
    }
}
