// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class CustomerEncryptionKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the encryption key that is stored in Google Cloud KMS.
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        /// <summary>
        /// The service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used.
        /// </summary>
        [Input("kmsKeyServiceAccount")]
        public Input<string>? KmsKeyServiceAccount { get; set; }

        /// <summary>
        /// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
        /// </summary>
        [Input("rawKey")]
        public Input<string>? RawKey { get; set; }

        /// <summary>
        /// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource.
        /// 
        /// The key must meet the following requirements before you can provide it to Compute Engine:  
        /// - The key is wrapped using a RSA public key certificate provided by Google. 
        /// - After being wrapped, the key must be encoded in RFC 4648 base64 encoding.  Gets the RSA public key certificate provided by Google at:
        /// https://cloud-certs.storage.googleapis.com/google-cloud-csek-ingress.pem
        /// </summary>
        [Input("rsaEncryptedKey")]
        public Input<string>? RsaEncryptedKey { get; set; }

        public CustomerEncryptionKeyArgs()
        {
        }
    }
}
