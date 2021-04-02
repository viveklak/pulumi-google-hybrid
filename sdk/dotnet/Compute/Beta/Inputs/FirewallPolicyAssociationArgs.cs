// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Beta.Inputs
{

    public sealed class FirewallPolicyAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The target that the firewall policy is attached to.
        /// </summary>
        [Input("attachmentTarget")]
        public Input<string>? AttachmentTarget { get; set; }

        /// <summary>
        /// [Output Only] Deprecated, please use short name instead. The display name of the firewall policy of the association.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// [Output Only] The firewall policy ID of the association.
        /// </summary>
        [Input("firewallPolicyId")]
        public Input<string>? FirewallPolicyId { get; set; }

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [Output Only] The short name of the firewall policy of the association.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        public FirewallPolicyAssociationArgs()
        {
        }
    }
}