// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudScheduler.V1Beta1.Outputs
{

    /// <summary>
    /// Contains information needed for generating an [OpenID Connect token](https://developers.google.com/identity/protocols/OpenIDConnect). This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.
    /// </summary>
    [OutputType]
    public sealed class OidcTokenResponse
    {
        /// <summary>
        /// Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used.
        /// </summary>
        public readonly string Audience;
        /// <summary>
        /// [Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OIDC token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.
        /// </summary>
        public readonly string ServiceAccountEmail;

        [OutputConstructor]
        private OidcTokenResponse(
            string audience,

            string serviceAccountEmail)
        {
            Audience = audience;
            ServiceAccountEmail = serviceAccountEmail;
        }
    }
}
