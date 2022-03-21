// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// NotificationConfig is the configuration of notifications.
    /// </summary>
    public sealed class NotificationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Notification config for Pub/Sub.
        /// </summary>
        [Input("pubsub")]
        public Input<Inputs.PubSubArgs>? Pubsub { get; set; }

        public NotificationConfigArgs()
        {
        }
    }
}
