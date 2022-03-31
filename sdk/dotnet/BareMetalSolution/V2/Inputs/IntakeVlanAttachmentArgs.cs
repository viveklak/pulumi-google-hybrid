// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Inputs
{

    /// <summary>
    /// A GCP vlan attachment.
    /// </summary>
    public sealed class IntakeVlanAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the VLAN attachment.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Attachment pairing key.
        /// </summary>
        [Input("pairingKey")]
        public Input<string>? PairingKey { get; set; }

        public IntakeVlanAttachmentArgs()
        {
        }
    }
}
