// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class CustomerEncryptionKeyResponse
    {
        /// <summary>
        /// The name of the encryption key that is stored in Google Cloud KMS.
        /// </summary>
        public readonly string KmsKeyName;
        /// <summary>
        /// The service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used.
        /// </summary>
        public readonly string KmsKeyServiceAccount;
        /// <summary>
        /// Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.
        /// </summary>
        public readonly string RawKey;
        /// <summary>
        /// Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit customer-supplied encryption key to either encrypt or decrypt this resource. The key must meet the following requirements before you can provide it to Compute Engine: 1. The key is wrapped using a RSA public key certificate provided by Google. 2. After being wrapped, the key must be encoded in RFC 4648 base64 encoding. Gets the RSA public key certificate provided by Google at: https://cloud-certs.storage.googleapis.com/google-cloud-csek-ingress.pem 
        /// </summary>
        public readonly string RsaEncryptedKey;
        /// <summary>
        /// [Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
        /// </summary>
        public readonly string Sha256;

        [OutputConstructor]
        private CustomerEncryptionKeyResponse(
            string kmsKeyName,

            string kmsKeyServiceAccount,

            string rawKey,

            string rsaEncryptedKey,

            string sha256)
        {
            KmsKeyName = kmsKeyName;
            KmsKeyServiceAccount = kmsKeyServiceAccount;
            RawKey = rawKey;
            RsaEncryptedKey = rsaEncryptedKey;
            Sha256 = sha256;
        }
    }
}
