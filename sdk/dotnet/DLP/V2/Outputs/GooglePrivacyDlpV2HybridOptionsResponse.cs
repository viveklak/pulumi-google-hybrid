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
    /// Configuration to control jobs where the content being inspected is outside of Google Cloud Platform.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2HybridOptionsResponse
    {
        /// <summary>
        /// A short description of where the data is coming from. Will be stored once in the job. 256 max length.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// To organize findings, these labels will be added to each finding. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`. Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`. No more than 10 labels can be associated with a given finding. Examples: * `"environment" : "production"` * `"pipeline" : "etl"`
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// These are labels that each inspection request must include within their 'finding_labels' map. Request may contain others, but any missing one of these will be rejected. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`. No more than 10 keys can be required.
        /// </summary>
        public readonly ImmutableArray<string> RequiredFindingLabelKeys;
        /// <summary>
        /// If the container is a table, additional information to make findings meaningful such as the columns that are primary keys.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2TableOptionsResponse TableOptions;

        [OutputConstructor]
        private GooglePrivacyDlpV2HybridOptionsResponse(
            string description,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> requiredFindingLabelKeys,

            Outputs.GooglePrivacyDlpV2TableOptionsResponse tableOptions)
        {
            Description = description;
            Labels = labels;
            RequiredFindingLabelKeys = requiredFindingLabelKeys;
            TableOptions = tableOptions;
        }
    }
}
