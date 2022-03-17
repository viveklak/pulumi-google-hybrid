// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// Settings for revision-level scaling settings.
    /// </summary>
    public sealed class GoogleCloudRunV2RevisionScalingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of serving instances that this resource should have.
        /// </summary>
        [Input("maxInstanceCount")]
        public Input<int>? MaxInstanceCount { get; set; }

        /// <summary>
        /// Minimum number of serving instances that this resource should have.
        /// </summary>
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        public GoogleCloudRunV2RevisionScalingArgs()
        {
        }
    }
}
