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
    /// Instructions regarding the table content being inspected.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2TableOptionsResponse
    {
        /// <summary>
        /// The columns that are the primary keys for table objects included in ContentItem. A copy of this cell's value will stored alongside alongside each finding so that the finding can be traced to the specific row it came from. No more than 3 may be provided.
        /// </summary>
        public readonly ImmutableArray<Outputs.GooglePrivacyDlpV2FieldIdResponse> IdentifyingFields;

        [OutputConstructor]
        private GooglePrivacyDlpV2TableOptionsResponse(ImmutableArray<Outputs.GooglePrivacyDlpV2FieldIdResponse> identifyingFields)
        {
            IdentifyingFields = identifyingFields;
        }
    }
}
