// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class BqmlTrainingRunResponse
    {
        /// <summary>
        /// [Output-only, Beta] List of each iteration results.
        /// </summary>
        public readonly ImmutableArray<Outputs.BqmlIterationResultResponse> IterationResults;
        /// <summary>
        /// [Output-only, Beta] Training run start time in milliseconds since the epoch.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// [Output-only, Beta] Different state applicable for a training run. IN PROGRESS: Training run is in progress. FAILED: Training run ended due to a non-retryable failure. SUCCEEDED: Training run successfully completed. CANCELLED: Training run cancelled by the user.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// [Output-only, Beta] Training options used by this training run. These options are mutable for subsequent training runs. Default values are explicitly stored for options not specified in the input query of the first training run. For subsequent training runs, any option not explicitly specified in the input query will be copied from the previous training run.
        /// </summary>
        public readonly Outputs.BqmlTrainingRunTrainingOptionsResponse TrainingOptions;

        [OutputConstructor]
        private BqmlTrainingRunResponse(
            ImmutableArray<Outputs.BqmlIterationResultResponse> iterationResults,

            string startTime,

            string state,

            Outputs.BqmlTrainingRunTrainingOptionsResponse trainingOptions)
        {
            IterationResults = iterationResults;
            StartTime = startTime;
            State = state;
            TrainingOptions = trainingOptions;
        }
    }
}