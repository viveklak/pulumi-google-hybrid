// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2Beta.Outputs
{

    /// <summary>
    /// Informational messages about the state of the Cloud Function or Operation.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudFunctionsV2betaStateMessageResponse
    {
        /// <summary>
        /// The message.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Severity of the state message.
        /// </summary>
        public readonly string Severity;
        /// <summary>
        /// One-word CamelCase type of the state message.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GoogleCloudFunctionsV2betaStateMessageResponse(
            string message,

            string severity,

            string type)
        {
            Message = message;
            Severity = severity;
            Type = type;
        }
    }
}
