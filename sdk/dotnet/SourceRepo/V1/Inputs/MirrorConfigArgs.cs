// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SourceRepo.V1.Inputs
{

    /// <summary>
    /// Configuration to automatically mirror a repository from another hosting service, for example GitHub or Bitbucket.
    /// </summary>
    public sealed class MirrorConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the SSH deploy key at the other hosting service. Removing this key from the other service would deauthorize Google Cloud Source Repositories from mirroring.
        /// </summary>
        [Input("deployKeyId")]
        public Input<string>? DeployKeyId { get; set; }

        /// <summary>
        /// URL of the main repository at the other hosting service.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// ID of the webhook listening to updates to trigger mirroring. Removing this webhook from the other hosting service will stop Google Cloud Source Repositories from receiving notifications, and thereby disabling mirroring.
        /// </summary>
        [Input("webhookId")]
        public Input<string>? WebhookId { get; set; }

        public MirrorConfigArgs()
        {
        }
    }
}
