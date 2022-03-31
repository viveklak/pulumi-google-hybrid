// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Inputs
{

    /// <summary>
    /// Instance represents the status of an instance of a Job.
    /// </summary>
    public sealed class InstanceStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Represents time when the instance was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. +optional
        /// </summary>
        [Input("completionTime")]
        public Input<string>? CompletionTime { get; set; }

        /// <summary>
        /// Optional. The number of times this instance exited with code &gt; 0; +optional
        /// </summary>
        [Input("failed")]
        public Input<int>? Failed { get; set; }

        /// <summary>
        /// Index of the instance, unique per Job, and beginning at 0.
        /// </summary>
        [Input("index", required: true)]
        public Input<int> Index { get; set; } = null!;

        /// <summary>
        /// Optional. Result of the last attempt of this instance. +optional
        /// </summary>
        [Input("lastAttemptResult")]
        public Input<Inputs.InstanceAttemptResultArgs>? LastAttemptResult { get; set; }

        /// <summary>
        /// Optional. Last exit code seen for this instance. +optional
        /// </summary>
        [Input("lastExitCode")]
        public Input<int>? LastExitCode { get; set; }

        /// <summary>
        /// Optional. The number of times this instance was restarted. Instances are restarted according the restartPolicy configured in the Job template. +optional
        /// </summary>
        [Input("restarted")]
        public Input<int>? Restarted { get; set; }

        /// <summary>
        /// Optional. Represents time when the instance was created by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. +optional
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Optional. The number of times this instance exited with code == 0. +optional
        /// </summary>
        [Input("succeeded")]
        public Input<int>? Succeeded { get; set; }

        public InstanceStatusArgs()
        {
        }
    }
}
