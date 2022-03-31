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
    /// Replace each input value with a given `Value`.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2ReplaceValueConfigResponse
    {
        /// <summary>
        /// Value to replace it with.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2ValueResponse NewValue;

        [OutputConstructor]
        private GooglePrivacyDlpV2ReplaceValueConfigResponse(Outputs.GooglePrivacyDlpV2ValueResponse newValue)
        {
            NewValue = newValue;
        }
    }
}
