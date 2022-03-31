// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.V2Beta.Inputs
{

    /// <summary>
    /// The credential used by Deployment Manager and TypeProvider. Only one of the options is permitted.
    /// </summary>
    public sealed class CredentialArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic Auth Credential, only used by TypeProvider.
        /// </summary>
        [Input("basicAuth")]
        public Input<Inputs.BasicAuthArgs>? BasicAuth { get; set; }

        /// <summary>
        /// Service Account Credential, only used by Deployment.
        /// </summary>
        [Input("serviceAccount")]
        public Input<Inputs.ServiceAccountArgs>? ServiceAccount { get; set; }

        /// <summary>
        /// Specify to use the project default credential, only supported by Deployment.
        /// </summary>
        [Input("useProjectDefault")]
        public Input<bool>? UseProjectDefault { get; set; }

        public CredentialArgs()
        {
        }
    }
}
