// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQueryConnection.V1Beta1.Inputs
{

    /// <summary>
    /// Credential info for the Cloud SQL.
    /// </summary>
    public sealed class CloudSqlCredentialArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password for the credential.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The username for the credential.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public CloudSqlCredentialArgs()
        {
        }
    }
}
