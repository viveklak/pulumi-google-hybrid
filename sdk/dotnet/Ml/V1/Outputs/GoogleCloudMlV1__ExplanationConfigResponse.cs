// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Outputs
{

    /// <summary>
    /// Message holding configuration options for explaining model predictions. There are three feature attribution methods supported for TensorFlow models: integrated gradients, sampled Shapley, and XRAI. [Learn more about feature attributions.](/ai-platform/prediction/docs/ai-explanations/overview)
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudMlV1__ExplanationConfigResponse
    {
        /// <summary>
        /// Attributes credit by computing the Aumann-Shapley value taking advantage of the model's fully differentiable structure. Refer to this paper for more details: https://arxiv.org/abs/1703.01365
        /// </summary>
        public readonly Outputs.GoogleCloudMlV1__IntegratedGradientsAttributionResponse IntegratedGradientsAttribution;
        /// <summary>
        /// An attribution method that approximates Shapley values for features that contribute to the label being predicted. A sampling strategy is used to approximate the value rather than considering all subsets of features.
        /// </summary>
        public readonly Outputs.GoogleCloudMlV1__SampledShapleyAttributionResponse SampledShapleyAttribution;
        /// <summary>
        /// Attributes credit by computing the XRAI taking advantage of the model's fully differentiable structure. Refer to this paper for more details: https://arxiv.org/abs/1906.02825 Currently only implemented for models with natural image inputs.
        /// </summary>
        public readonly Outputs.GoogleCloudMlV1__XraiAttributionResponse XraiAttribution;

        [OutputConstructor]
        private GoogleCloudMlV1__ExplanationConfigResponse(
            Outputs.GoogleCloudMlV1__IntegratedGradientsAttributionResponse integratedGradientsAttribution,

            Outputs.GoogleCloudMlV1__SampledShapleyAttributionResponse sampledShapleyAttribution,

            Outputs.GoogleCloudMlV1__XraiAttributionResponse xraiAttribution)
        {
            IntegratedGradientsAttribution = integratedGradientsAttribution;
            SampledShapleyAttribution = sampledShapleyAttribution;
            XraiAttribution = xraiAttribution;
        }
    }
}
