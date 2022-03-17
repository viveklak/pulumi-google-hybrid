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
    /// A type of transformation that is applied over structured data such as a table.
    /// </summary>
    public sealed class GooglePrivacyDlpV2RecordTransformationsArgs : Pulumi.ResourceArgs
    {
        [Input("fieldTransformations")]
        private InputList<Inputs.GooglePrivacyDlpV2FieldTransformationArgs>? _fieldTransformations;

        /// <summary>
        /// Transform the record by applying various field transformations.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2FieldTransformationArgs> FieldTransformations
        {
            get => _fieldTransformations ?? (_fieldTransformations = new InputList<Inputs.GooglePrivacyDlpV2FieldTransformationArgs>());
            set => _fieldTransformations = value;
        }

        [Input("recordSuppressions")]
        private InputList<Inputs.GooglePrivacyDlpV2RecordSuppressionArgs>? _recordSuppressions;

        /// <summary>
        /// Configuration defining which records get suppressed entirely. Records that match any suppression rule are omitted from the output.
        /// </summary>
        public InputList<Inputs.GooglePrivacyDlpV2RecordSuppressionArgs> RecordSuppressions
        {
            get => _recordSuppressions ?? (_recordSuppressions = new InputList<Inputs.GooglePrivacyDlpV2RecordSuppressionArgs>());
            set => _recordSuppressions = value;
        }

        public GooglePrivacyDlpV2RecordTransformationsArgs()
        {
        }
    }
}
