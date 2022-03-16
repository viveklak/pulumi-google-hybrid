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
    /// The set of columns' values that share the same ldiversity value.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2LDiversityEquivalenceClassResponse
    {
        /// <summary>
        /// Size of the k-anonymity equivalence class.
        /// </summary>
        public readonly string EquivalenceClassSize;
        /// <summary>
        /// Number of distinct sensitive values in this equivalence class.
        /// </summary>
        public readonly string NumDistinctSensitiveValues;
        /// <summary>
        /// Quasi-identifier values defining the k-anonymity equivalence class. The order is always the same as the original request.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2ValueResponse> QuasiIdsValues;
        /// <summary>
        /// Estimated frequencies of top sensitive values.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2ValueFrequencyResponse> TopSensitiveValues;

        [OutputConstructor]
        private GooglePrivacyDlpV2LDiversityEquivalenceClassResponse(
            string equivalenceClassSize,

            string numDistinctSensitiveValues,

            ImmutableArray<Outputs.GooglePrivacyDlpV2ValueResponse> quasiIdsValues,

            ImmutableArray<Outputs.GooglePrivacyDlpV2ValueFrequencyResponse> topSensitiveValues)
        {
            EquivalenceClassSize = equivalenceClassSize;
            NumDistinctSensitiveValues = numDistinctSensitiveValues;
            QuasiIdsValues = quasiIdsValues;
            TopSensitiveValues = topSensitiveValues;
        }
    }
}