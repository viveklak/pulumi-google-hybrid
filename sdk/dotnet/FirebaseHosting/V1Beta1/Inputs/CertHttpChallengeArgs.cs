// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.FirebaseHosting.V1Beta1.Inputs
{

    /// <summary>
    /// Represents an HTTP certificate challenge.
    /// </summary>
    public sealed class CertHttpChallengeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL path on which to serve the specified token to satisfy the certificate challenge.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The token to serve at the specified URL path to satisfy the certificate challenge.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public CertHttpChallengeArgs()
        {
        }
    }
}