// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CertificateManager.V1.Outputs
{

    /// <summary>
    /// Information about issues with provisioning a Managed Certificate.
    /// </summary>
    [OutputType]
    public sealed class ProvisioningIssueResponse
    {
        /// <summary>
        /// Human readable explanation about the issue. Provided to help address the configuration issues. Not guaranteed to be stable. For programmatic access use Reason enum.
        /// </summary>
        public readonly string Details;
        /// <summary>
        /// Reason for provisioning failures.
        /// </summary>
        public readonly string Reason;

        [OutputConstructor]
        private ProvisioningIssueResponse(
            string details,

            string reason)
        {
            Details = details;
            Reason = reason;
        }
    }
}
