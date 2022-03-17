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
    /// A policy that defines the replication and encryption configuration of data.
    /// </summary>
    [OutputType]
    public sealed class ReplicationResponse
    {
        /// <summary>
        /// The Secret will automatically be replicated without any restrictions.
        /// </summary>
        public readonly Outputs.AutomaticResponse Automatic;
        /// <summary>
        /// The Secret will only be replicated into the locations specified.
        /// </summary>
        public readonly Outputs.UserManagedResponse UserManaged;

        [OutputConstructor]
        private ReplicationResponse(
            Outputs.AutomaticResponse automatic,

            Outputs.UserManagedResponse userManaged)
        {
            Automatic = automatic;
            UserManaged = userManaged;
        }
    }
}
