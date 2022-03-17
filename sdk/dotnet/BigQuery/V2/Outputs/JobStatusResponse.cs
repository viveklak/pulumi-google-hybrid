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
    public sealed class JobStatusResponse
    {
        /// <summary>
        /// Final error result of the job. If present, indicates that the job has completed and was unsuccessful.
        /// </summary>
        public readonly Outputs.ErrorProtoResponse ErrorResult;
        /// <summary>
        /// The first errors encountered during the running of the job. The final message includes the number of errors that caused the process to stop. Errors here do not necessarily mean that the job has completed or was unsuccessful.
        /// </summary>
        public readonly ImmutableArray<Outputs.ErrorProtoResponse> Errors;
        /// <summary>
        /// Running state of the job.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private JobStatusResponse(
            Outputs.ErrorProtoResponse errorResult,

            ImmutableArray<Outputs.ErrorProtoResponse> errors,

            string state)
        {
            ErrorResult = errorResult;
            Errors = errors;
            State = state;
        }
    }
}
