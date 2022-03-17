// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIoT.V1.Inputs
{

    /// <summary>
    /// A server-stored device credential used for authentication.
    /// </summary>
    public sealed class DeviceCredentialArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Optional] The time at which this credential becomes invalid. This credential will be ignored for new client authentication requests after this timestamp; however, it will not be automatically deleted.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// A public key used to verify the signature of JSON Web Tokens (JWTs). When adding a new device credential, either via device creation or via modifications, this public key credential may be required to be signed by one of the registry level certificates. More specifically, if the registry contains at least one certificate, any new device credential must be signed by one of the registry certificates. As a result, when the registry contains certificates, only X.509 certificates are accepted as device credentials. However, if the registry does not contain a certificate, self-signed certificates and public keys will be accepted. New device credentials must be different from every registry-level certificate.
        /// </summary>
        [Input("publicKey")]
        public Input<Inputs.PublicKeyCredentialArgs>? PublicKey { get; set; }

        public DeviceCredentialArgs()
        {
        }
    }
}
