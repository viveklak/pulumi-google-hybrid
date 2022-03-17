// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIGateway.V1Beta.Inputs
{

    /// <summary>
    /// Configuration for all backends.
    /// </summary>
    public sealed class ApigatewayBackendConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Google Cloud IAM service account used to sign OIDC tokens for backends that have authentication configured (https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services.configs#backend). This may either be the Service Account's email (i.e. "{ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com") or its full resource name (i.e. "projects/{PROJECT}/accounts/{UNIQUE_ID}"). This is most often used when the backend is a GCP resource such as a Cloud Run Service or an IAP-secured service. Note that this token is always sent as an authorization header bearer token. The audience of the OIDC token is configured in the associated Service Config in the BackendRule option (https://github.com/googleapis/googleapis/blob/master/google/api/backend.proto#L125).
        /// </summary>
        [Input("googleServiceAccount")]
        public Input<string>? GoogleServiceAccount { get; set; }

        public ApigatewayBackendConfigArgs()
        {
        }
    }
}
