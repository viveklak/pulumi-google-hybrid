// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Outputs
{

    /// <summary>
    /// Represent a user-facing Error.
    /// </summary>
    [OutputType]
    public sealed class ErrorResponse
    {
        /// <summary>
        /// Additional information about the error.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Details;
        /// <summary>
        /// The time when the error occurred.
        /// </summary>
        public readonly string ErrorTime;
        /// <summary>
        /// A unique identifier for this specific error, allowing it to be traced throughout the system in logs and API responses.
        /// </summary>
        public readonly string ErrorUuid;
        /// <summary>
        /// A message containing more information about the error that occurred.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// A title that explains the reason for the error.
        /// </summary>
        public readonly string Reason;

        [OutputConstructor]
        private ErrorResponse(
            ImmutableDictionary<string, string> details,

            string errorTime,

            string errorUuid,

            string message,

            string reason)
        {
            Details = details;
            ErrorTime = errorTime;
            ErrorUuid = errorUuid;
            Message = message;
            Reason = reason;
        }
    }
}
