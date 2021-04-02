// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Secretmanager.V1.Inputs
{

    /// <summary>
    /// A Pub/Sub topic which Secret Manager will publish to when control plane events occur on this secret.
    /// </summary>
    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The resource name of the Pub/Sub topic that will be published to, in the following format: `projects/*/topics/*`. For publication to succeed, the Secret Manager P4SA must have `pubsub.publisher` permissions on the topic.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TopicArgs()
        {
        }
    }
}