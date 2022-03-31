// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1.Inputs
{

    /// <summary>
    /// Specification of certificate provider. Defines the mechanism to obtain the certificate and private key for peer to peer authentication.
    /// </summary>
    public sealed class GoogleCloudNetworksecurityV1beta1CertificateProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate provider instance specification that will be passed to the data plane, which will be used to load necessary credential information.
        /// </summary>
        [Input("certificateProviderInstance")]
        public Input<Inputs.CertificateProviderInstanceArgs>? CertificateProviderInstance { get; set; }

        /// <summary>
        /// gRPC specific configuration to access the gRPC server to obtain the cert and private key.
        /// </summary>
        [Input("grpcEndpoint")]
        public Input<Inputs.GoogleCloudNetworksecurityV1beta1GrpcEndpointArgs>? GrpcEndpoint { get; set; }

        public GoogleCloudNetworksecurityV1beta1CertificateProviderArgs()
        {
        }
    }
}
