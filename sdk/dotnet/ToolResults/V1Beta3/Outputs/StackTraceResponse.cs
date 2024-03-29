// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Outputs
{

    /// <summary>
    /// A stacktrace.
    /// </summary>
    [OutputType]
    public sealed class StackTraceResponse
    {
        /// <summary>
        /// The stack trace message. Required
        /// </summary>
        public readonly string Exception;

        [OutputConstructor]
        private StackTraceResponse(string exception)
        {
            Exception = exception;
        }
    }
}
