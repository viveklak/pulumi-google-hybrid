// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class SecurityPolicyAssociationResponse
    {
        /// <summary>
        /// The resource that the security policy is attached to.
        /// </summary>
        public readonly string AttachmentId;
        /// <summary>
        /// The display name of the security policy of the association.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The name for an association.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The security policy ID of the association.
        /// </summary>
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private SecurityPolicyAssociationResponse(
            string attachmentId,

            string displayName,

            string name,

            string securityPolicyId)
        {
            AttachmentId = attachmentId;
            DisplayName = displayName;
            Name = name;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
