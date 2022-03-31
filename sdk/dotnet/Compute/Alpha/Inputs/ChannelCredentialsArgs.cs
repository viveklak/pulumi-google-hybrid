// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// [Deprecated] gRPC channel credentials to access the SDS server. gRPC channel credentials to access the SDS server.
    /// </summary>
    public sealed class ChannelCredentialsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The call credentials to access the SDS server.
        /// </summary>
        [Input("certificates")]
        public Input<Inputs.TlsCertificatePathsArgs>? Certificates { get; set; }

        /// <summary>
        /// The channel credentials to access the SDS server. This field can be set to one of the following: CERTIFICATES: Use TLS certificates to access the SDS server. GCE_VM: Use local GCE VM credentials to access the SDS server.
        /// </summary>
        [Input("channelCredentialType")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ChannelCredentialsChannelCredentialType>? ChannelCredentialType { get; set; }

        public ChannelCredentialsArgs()
        {
        }
    }
}
