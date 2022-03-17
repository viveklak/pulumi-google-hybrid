// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2Beta.Outputs
{

    /// <summary>
    /// Service Account used as a credential.
    /// </summary>
    [OutputType]
    public sealed class ServiceAccountResponse
    {
        /// <summary>
        /// The IAM service account email address like test@myproject.iam.gserviceaccount.com
        /// </summary>
        public readonly string Email;

        [OutputConstructor]
        private ServiceAccountResponse(string email)
        {
            Email = email;
        }
    }
}
