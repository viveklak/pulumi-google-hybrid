// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Inputs
{

    /// <summary>
    /// Message holding configuration options for explaining model predictions. There are three feature attribution methods supported for TensorFlow models: integrated gradients, sampled Shapley, and XRAI. [Learn more about feature attributions.](/ai-platform/prediction/docs/ai-explanations/overview)
    /// </summary>
    public sealed class GoogleCloudMlV1__ExplanationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attributes credit by computing the Aumann-Shapley value taking advantage of the model's fully differentiable structure. Refer to this paper for more details: https://arxiv.org/abs/1703.01365
        /// </summary>
        [Input("integratedGradientsAttribution")]
        public Input<Inputs.GoogleCloudMlV1__IntegratedGradientsAttributionArgs>? IntegratedGradientsAttribution { get; set; }

        /// <summary>
        /// An attribution method that approximates Shapley values for features that contribute to the label being predicted. A sampling strategy is used to approximate the value rather than considering all subsets of features.
        /// </summary>
        [Input("sampledShapleyAttribution")]
        public Input<Inputs.GoogleCloudMlV1__SampledShapleyAttributionArgs>? SampledShapleyAttribution { get; set; }

        /// <summary>
        /// Attributes credit by computing the XRAI taking advantage of the model's fully differentiable structure. Refer to this paper for more details: https://arxiv.org/abs/1906.02825 Currently only implemented for models with natural image inputs.
        /// </summary>
        [Input("xraiAttribution")]
        public Input<Inputs.GoogleCloudMlV1__XraiAttributionArgs>? XraiAttribution { get; set; }

        public GoogleCloudMlV1__ExplanationConfigArgs()
        {
        }
    }
}
