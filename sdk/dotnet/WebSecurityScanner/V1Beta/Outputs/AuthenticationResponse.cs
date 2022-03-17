// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1Beta.Outputs
{

    /// <summary>
    /// Scan authentication configuration.
    /// </summary>
    [OutputType]
    public sealed class AuthenticationResponse
    {
        /// <summary>
        /// Authentication using a custom account.
        /// </summary>
        public readonly Outputs.CustomAccountResponse CustomAccount;
        /// <summary>
        /// Authentication using a Google account.
        /// </summary>
        public readonly Outputs.GoogleAccountResponse GoogleAccount;
        /// <summary>
        /// Authentication using Identity-Aware-Proxy (IAP).
        /// </summary>
        public readonly Outputs.IapCredentialResponse IapCredential;

        [OutputConstructor]
        private AuthenticationResponse(
            Outputs.CustomAccountResponse customAccount,

            Outputs.GoogleAccountResponse googleAccount,

            Outputs.IapCredentialResponse iapCredential)
        {
            CustomAccount = customAccount;
            GoogleAccount = googleAccount;
            IapCredential = iapCredential;
        }
    }
}
