// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Informational metadata about Partner attachments from Partners to display to customers. These fields are propagated from PARTNER_PROVIDER attachments to their corresponding PARTNER attachments.
    /// </summary>
    public sealed class InterconnectAttachmentPartnerMetadataArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Plain text name of the Interconnect this attachment is connected to, as displayed in the Partner's portal. For instance "Chicago 1". This value may be validated to match approved Partner values.
        /// </summary>
        [Input("interconnectName")]
        public Input<string>? InterconnectName { get; set; }

        /// <summary>
        /// Plain text name of the Partner providing this attachment. This value may be validated to match approved Partner values.
        /// </summary>
        [Input("partnerName")]
        public Input<string>? PartnerName { get; set; }

        /// <summary>
        /// URL of the Partner's portal for this Attachment. Partners may customise this to be a deep link to the specific resource on the Partner portal. This value may be validated to match approved Partner values.
        /// </summary>
        [Input("portalUrl")]
        public Input<string>? PortalUrl { get; set; }

        public InterconnectAttachmentPartnerMetadataArgs()
        {
        }
    }
}
