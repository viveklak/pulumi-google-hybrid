// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// How to handle transformation errors during de-identification. A transformation error occurs when the requested transformation is incompatible with the data. For example, trying to de-identify an IP address using a `DateShift` transformation would result in a transformation error, since date info cannot be extracted from an IP address. Information about any incompatible transformations, and how they were handled, is returned in the response as part of the `TransformationOverviews`.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2TransformationErrorHandlingResponse
    {
        /// <summary>
        /// Ignore errors
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2LeaveUntransformedResponse LeaveUntransformed;
        /// <summary>
        /// Throw an error
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2ThrowErrorResponse ThrowError;

        [OutputConstructor]
        private GooglePrivacyDlpV2TransformationErrorHandlingResponse(
            Outputs.GooglePrivacyDlpV2LeaveUntransformedResponse leaveUntransformed,

            Outputs.GooglePrivacyDlpV2ThrowErrorResponse throwError)
        {
            LeaveUntransformed = leaveUntransformed;
            ThrowError = throwError;
        }
    }
}
