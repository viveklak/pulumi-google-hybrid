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
    /// Reidentifiability metric. This corresponds to a risk model similar to what is called "journalist risk" in the literature, except the attack dataset is statistically modeled instead of being perfectly known. This can be done using publicly available data (like the US Census), or using a custom statistical model (indicated as one or several BigQuery tables), or by extrapolating from the distribution of values in the input dataset.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2KMapEstimationConfigResponse
    {
        /// <summary>
        /// Several auxiliary tables can be used in the analysis. Each custom_tag used to tag a quasi-identifiers column must appear in exactly one column of one auxiliary table.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2AuxiliaryTableResponse> AuxiliaryTables;
        /// <summary>
        /// Fields considered to be quasi-identifiers. No two columns can have the same tag.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2TaggedFieldResponse> QuasiIds;
        /// <summary>
        /// ISO 3166-1 alpha-2 region code to use in the statistical modeling. Set if no column is tagged with a region-specific InfoType (like US_ZIP_5) or a region code.
        /// </summary>
        public readonly string RegionCode;

        [OutputConstructor]
        private GooglePrivacyDlpV2KMapEstimationConfigResponse(
            ImmutableArray<Outputs.GooglePrivacyDlpV2AuxiliaryTableResponse> auxiliaryTables,

            ImmutableArray<Outputs.GooglePrivacyDlpV2TaggedFieldResponse> quasiIds,

            string regionCode)
        {
            AuxiliaryTables = auxiliaryTables;
            QuasiIds = quasiIds;
            RegionCode = regionCode;
        }
    }
}
