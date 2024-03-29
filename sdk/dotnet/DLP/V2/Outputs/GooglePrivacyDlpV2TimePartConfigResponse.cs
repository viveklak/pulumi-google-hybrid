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
    /// For use with `Date`, `Timestamp`, and `TimeOfDay`, extract or preserve a portion of the value.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2TimePartConfigResponse
    {
        /// <summary>
        /// The part of the time to keep.
        /// </summary>
        public readonly string PartToExtract;

        [OutputConstructor]
        private GooglePrivacyDlpV2TimePartConfigResponse(string partToExtract)
        {
            PartToExtract = partToExtract;
        }
    }
}
