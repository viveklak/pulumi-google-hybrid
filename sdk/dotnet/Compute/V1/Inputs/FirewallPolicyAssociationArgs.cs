// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    public sealed class FirewallPolicyAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target that the firewall policy is attached to.
        /// </summary>
        [Input("attachmentTarget")]
        public Input<string>? AttachmentTarget { get; set; }

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallPolicyAssociationArgs()
        {
        }
    }
}
