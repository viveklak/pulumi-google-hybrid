// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1.Inputs
{

    /// <summary>
    /// Config for image classification human labeling task.
    /// </summary>
    public sealed class GoogleCloudDatalabelingV1beta1ImageClassificationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. If allow_multi_label is true, contributors are able to choose multiple labels for one image.
        /// </summary>
        [Input("allowMultiLabel")]
        public Input<bool>? AllowMultiLabel { get; set; }

        /// <summary>
        /// Annotation spec set resource name.
        /// </summary>
        [Input("annotationSpecSet", required: true)]
        public Input<string> AnnotationSpecSet { get; set; } = null!;

        /// <summary>
        /// Optional. The type of how to aggregate answers.
        /// </summary>
        [Input("answerAggregationType")]
        public Input<Pulumi.GoogleNative.DataLabeling.V1Beta1.GoogleCloudDatalabelingV1beta1ImageClassificationConfigAnswerAggregationType>? AnswerAggregationType { get; set; }

        public GoogleCloudDatalabelingV1beta1ImageClassificationConfigArgs()
        {
        }
    }
}
