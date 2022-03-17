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
    /// Attributes credit by computing the XRAI taking advantage of the model's fully differentiable structure. Refer to this paper for more details: https://arxiv.org/abs/1906.02825 Currently only implemented for models with natural image inputs.
    /// </summary>
    public sealed class GoogleCloudMlV1__XraiAttributionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of steps for approximating the path integral. A good value to start is 50 and gradually increase until the sum to diff property is met within the desired error range.
        /// </summary>
        [Input("numIntegralSteps")]
        public Input<int>? NumIntegralSteps { get; set; }

        public GoogleCloudMlV1__XraiAttributionArgs()
        {
        }
    }
}
