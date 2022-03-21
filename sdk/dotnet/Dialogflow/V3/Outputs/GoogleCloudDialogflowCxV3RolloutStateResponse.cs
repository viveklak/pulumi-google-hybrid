// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// State of the auto-rollout process.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3RolloutStateResponse
    {
        /// <summary>
        /// Start time of the current step.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Display name of the current auto rollout step.
        /// </summary>
        public readonly string Step;
        /// <summary>
        /// Index of the current step in the auto rollout steps list.
        /// </summary>
        public readonly int StepIndex;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3RolloutStateResponse(
            string startTime,

            string step,

            int stepIndex)
        {
            StartTime = startTime;
            Step = step;
            StepIndex = stepIndex;
        }
    }
}
