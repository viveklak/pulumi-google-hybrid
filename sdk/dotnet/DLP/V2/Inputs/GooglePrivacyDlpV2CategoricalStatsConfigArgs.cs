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
    /// Compute numerical stats over an individual column, including number of distinct values and value count distribution.
    /// </summary>
    public sealed class GooglePrivacyDlpV2CategoricalStatsConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Field to compute categorical stats on. All column types are supported except for arrays and structs. However, it may be more informative to use NumericalStats when the field type is supported, depending on the data.
        /// </summary>
        [Input("field")]
        public Input<Inputs.GooglePrivacyDlpV2FieldIdArgs>? Field { get; set; }

        public GooglePrivacyDlpV2CategoricalStatsConfigArgs()
        {
        }
    }
}
