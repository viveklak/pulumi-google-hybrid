// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha.Outputs
{

    /// <summary>
    /// The credential used by Deployment Manager and TypeProvider. Only one of the options is permitted.
    /// </summary>
    [OutputType]
    public sealed class CredentialResponse
    {
        /// <summary>
        /// Basic Auth Credential, only used by TypeProvider.
        /// </summary>
        public readonly Outputs.BasicAuthResponse BasicAuth;
        /// <summary>
        /// Service Account Credential, only used by Deployment.
        /// </summary>
        public readonly Outputs.ServiceAccountResponse ServiceAccount;
        /// <summary>
        /// Specify to use the project default credential, only supported by Deployment.
        /// </summary>
        public readonly bool UseProjectDefault;

        [OutputConstructor]
        private CredentialResponse(
            Outputs.BasicAuthResponse basicAuth,

            Outputs.ServiceAccountResponse serviceAccount,

            bool useProjectDefault)
        {
            BasicAuth = basicAuth;
            ServiceAccount = serviceAccount;
            UseProjectDefault = useProjectDefault;
        }
    }
}
