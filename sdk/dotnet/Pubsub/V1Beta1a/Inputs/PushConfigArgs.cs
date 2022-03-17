// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1Beta1a.Inputs
{

    /// <summary>
    /// Configuration for a push delivery endpoint.
    /// </summary>
    public sealed class PushConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A URL locating the endpoint to which messages should be pushed. For example, a Webhook endpoint might use "https://example.com/push".
        /// </summary>
        [Input("pushEndpoint")]
        public Input<string>? PushEndpoint { get; set; }

        public PushConfigArgs()
        {
        }
    }
}
