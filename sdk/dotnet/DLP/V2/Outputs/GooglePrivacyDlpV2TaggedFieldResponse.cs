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
    /// A column with a semantic tag attached.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2TaggedFieldResponse
    {
        /// <summary>
        /// A column can be tagged with a custom tag. In this case, the user must indicate an auxiliary table that contains statistical information on the possible values of this column (below).
        /// </summary>
        public readonly string CustomTag;
        /// <summary>
        /// Identifies the column.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2FieldIdResponse Field;
        /// <summary>
        /// If no semantic tag is indicated, we infer the statistical model from the distribution of values in the input data
        /// </summary>
        public readonly Outputs.GoogleProtobufEmptyResponse Inferred;
        /// <summary>
        /// A column can be tagged with a InfoType to use the relevant public dataset as a statistical model of population, if available. We currently support US ZIP codes, region codes, ages and genders. To programmatically obtain the list of supported InfoTypes, use ListInfoTypes with the supported_by=RISK_ANALYSIS filter.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2InfoTypeResponse InfoType;

        [OutputConstructor]
        private GooglePrivacyDlpV2TaggedFieldResponse(
            string customTag,

            Outputs.GooglePrivacyDlpV2FieldIdResponse field,

            Outputs.GoogleProtobufEmptyResponse inferred,

            Outputs.GooglePrivacyDlpV2InfoTypeResponse infoType)
        {
            CustomTag = customTag;
            Field = field;
            Inferred = inferred;
            InfoType = infoType;
        }
    }
}
