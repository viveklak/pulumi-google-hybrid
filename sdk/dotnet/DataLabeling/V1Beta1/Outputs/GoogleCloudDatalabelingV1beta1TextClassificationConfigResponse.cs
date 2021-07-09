// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDatalabelingV1beta1TextClassificationConfigResponse
    {
        /// <summary>
        /// Optional. If allow_multi_label is true, contributors are able to choose multiple labels for one text segment.
        /// </summary>
        public readonly bool AllowMultiLabel;
        /// <summary>
        /// Annotation spec set resource name.
        /// </summary>
        public readonly string AnnotationSpecSet;
        /// <summary>
        /// Optional. Configs for sentiment selection. We deprecate sentiment analysis in data labeling side as it is incompatible with uCAIP.
        /// </summary>
        public readonly Outputs.GoogleCloudDatalabelingV1beta1SentimentConfigResponse SentimentConfig;

        [OutputConstructor]
        private GoogleCloudDatalabelingV1beta1TextClassificationConfigResponse(
            bool allowMultiLabel,

            string annotationSpecSet,

            Outputs.GoogleCloudDatalabelingV1beta1SentimentConfigResponse sentimentConfig)
        {
            AllowMultiLabel = allowMultiLabel;
            AnnotationSpecSet = annotationSpecSet;
            SentimentConfig = sentimentConfig;
        }
    }
}
