// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// Message defining a field of a BigQuery table.
    /// </summary>
    public sealed class GooglePrivacyDlpV2BigQueryFieldArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Designated field in the BigQuery table.
        /// </summary>
        [Input("field")]
        public Input<Inputs.GooglePrivacyDlpV2FieldIdArgs>? Field { get; set; }

        /// <summary>
        /// Source table of the field.
        /// </summary>
        [Input("table")]
        public Input<Inputs.GooglePrivacyDlpV2BigQueryTableArgs>? Table { get; set; }

        public GooglePrivacyDlpV2BigQueryFieldArgs()
        {
        }
    }
}
