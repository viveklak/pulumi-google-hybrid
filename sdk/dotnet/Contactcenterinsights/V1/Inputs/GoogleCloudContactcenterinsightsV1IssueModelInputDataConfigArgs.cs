// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Inputs
{

    /// <summary>
    /// Configs for the input data used to create the issue model.
    /// </summary>
    public sealed class GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A filter to reduce the conversations used for training the model to a specific subset.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Medium of conversations used in training data. This field is being deprecated. To specify the medium to be used in training a new issue model, set the `medium` field on `filter`.
        /// </summary>
        [Input("medium")]
        public Input<Pulumi.GoogleNative.Contactcenterinsights.V1.GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigMedium>? Medium { get; set; }

        public GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigArgs()
        {
        }
    }
}
