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
    /// A collection of conditions.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2ConditionsResponse
    {
        /// <summary>
        /// A collection of conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2ConditionResponse> Conditions;

        [OutputConstructor]
        private GooglePrivacyDlpV2ConditionsResponse(ImmutableArray<Outputs.GooglePrivacyDlpV2ConditionResponse> conditions)
        {
            Conditions = conditions;
        }
    }
}
