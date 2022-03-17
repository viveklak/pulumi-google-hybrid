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
    /// If set, the detailed findings will be persisted to the specified OutputStorageConfig. Only a single instance of this action can be specified. Compatible with: Inspect, Risk
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2SaveFindingsResponse
    {
        /// <summary>
        /// Location to store findings outside of DLP.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2OutputStorageConfigResponse OutputConfig;

        [OutputConstructor]
        private GooglePrivacyDlpV2SaveFindingsResponse(Outputs.GooglePrivacyDlpV2OutputStorageConfigResponse outputConfig)
        {
            OutputConfig = outputConfig;
        }
    }
}
