// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// A service account.
    /// </summary>
    [OutputType]
    public sealed class ServiceAccountResponse
    {
        /// <summary>
        /// Email address of the service account.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The list of scopes to be made available for this service account.
        /// </summary>
        public readonly ImmutableArray<string> Scopes;

        [OutputConstructor]
        private ServiceAccountResponse(
            string email,

            ImmutableArray<string> scopes)
        {
            Email = email;
            Scopes = scopes;
        }
    }
}
