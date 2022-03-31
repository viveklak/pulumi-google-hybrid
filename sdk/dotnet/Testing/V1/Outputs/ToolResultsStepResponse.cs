// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Outputs
{

    /// <summary>
    /// Represents a tool results step resource. This has the results of a TestExecution.
    /// </summary>
    [OutputType]
    public sealed class ToolResultsStepResponse
    {
        /// <summary>
        /// A tool results execution ID.
        /// </summary>
        public readonly string ExecutionId;
        /// <summary>
        /// A tool results history ID.
        /// </summary>
        public readonly string HistoryId;
        /// <summary>
        /// The cloud project that owns the tool results step.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// A tool results step ID.
        /// </summary>
        public readonly string StepId;

        [OutputConstructor]
        private ToolResultsStepResponse(
            string executionId,

            string historyId,

            string project,

            string stepId)
        {
            ExecutionId = executionId;
            HistoryId = historyId;
            Project = project;
            StepId = stepId;
        }
    }
}
