// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SecretManager.V1.Outputs
{

    /// <summary>
    /// Represents a Replica for this Secret.
    /// </summary>
    [OutputType]
    public sealed class ReplicaResponse
    {
        /// <summary>
        /// Optional. The customer-managed encryption configuration of the User-Managed Replica. If no configuration is provided, Google-managed default encryption is used. Updates to the Secret encryption configuration only apply to SecretVersions added afterwards. They do not apply retroactively to existing SecretVersions.
        /// </summary>
        public readonly Outputs.CustomerManagedEncryptionResponse CustomerManagedEncryption;
        /// <summary>
        /// The canonical IDs of the location to replicate data. For example: `"us-east1"`.
        /// </summary>
        public readonly string Location;

        [OutputConstructor]
        private ReplicaResponse(
            Outputs.CustomerManagedEncryptionResponse customerManagedEncryption,

            string location)
        {
            CustomerManagedEncryption = customerManagedEncryption;
            Location = location;
        }
    }
}
