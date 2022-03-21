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
    /// A quasi-identifier column has a custom_tag, used to know which column in the data corresponds to which column in the statistical model.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2QuasiIdFieldResponse
    {
        /// <summary>
        /// A auxiliary field.
        /// </summary>
        public readonly string CustomTag;
        /// <summary>
        /// Identifies the column.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2FieldIdResponse Field;

        [OutputConstructor]
        private GooglePrivacyDlpV2QuasiIdFieldResponse(
            string customTag,

            Outputs.GooglePrivacyDlpV2FieldIdResponse field)
        {
            CustomTag = customTag;
            Field = field;
        }
    }
}
