// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1.Inputs
{

    /// <summary>
    /// Output configuration for asset feed destination.
    /// </summary>
    public sealed class FeedOutputConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination on Pub/Sub.
        /// </summary>
        [Input("pubsubDestination")]
        public Input<Inputs.PubsubDestinationArgs>? PubsubDestination { get; set; }

        public FeedOutputConfigArgs()
        {
        }
    }
}
