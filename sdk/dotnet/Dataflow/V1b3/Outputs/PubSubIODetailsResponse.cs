// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// Metadata for a Pub/Sub connector used by the job.
    /// </summary>
    [OutputType]
    public sealed class PubSubIODetailsResponse
    {
        /// <summary>
        /// Subscription used in the connection.
        /// </summary>
        public readonly string Subscription;
        /// <summary>
        /// Topic accessed in the connection.
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private PubSubIODetailsResponse(
            string subscription,

            string topic)
        {
            Subscription = subscription;
            Topic = topic;
        }
    }
}
