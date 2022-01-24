// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.IAM.V1
{
    public static class GetKey
    {
        /// <summary>
        /// Gets a ServiceAccountKey.
        /// </summary>
        public static Task<GetKeyResult> InvokeAsync(GetKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeyResult>("google-native:iam/v1:getKey", args ?? new GetKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a ServiceAccountKey.
        /// </summary>
        public static Output<GetKeyResult> Invoke(GetKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetKeyResult>("google-native:iam/v1:getKey", args ?? new GetKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeyArgs : Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public string KeyId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("publicKeyType")]
        public string? PublicKeyType { get; set; }

        [Input("serviceAccountId", required: true)]
        public string ServiceAccountId { get; set; } = null!;

        public GetKeyArgs()
        {
        }
    }

    public sealed class GetKeyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("publicKeyType")]
        public Input<string>? PublicKeyType { get; set; }

        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        public GetKeyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeyResult
    {
        /// <summary>
        /// The key status.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// Specifies the algorithm (and possibly key size) for the key.
        /// </summary>
        public readonly string KeyAlgorithm;
        /// <summary>
        /// The key origin.
        /// </summary>
        public readonly string KeyOrigin;
        /// <summary>
        /// The key type.
        /// </summary>
        public readonly string KeyType;
        /// <summary>
        /// The resource name of the service account key in the following format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{key}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The private key data. Only provided in `CreateServiceAccountKey` responses. Make sure to keep the private key data secure because it allows for the assertion of the service account identity. When base64 decoded, the private key data can be used to authenticate with Google API client libraries and with gcloud auth activate-service-account.
        /// </summary>
        public readonly string PrivateKeyData;
        /// <summary>
        /// The output format for the private key. Only provided in `CreateServiceAccountKey` responses, not in `GetServiceAccountKey` or `ListServiceAccountKey` responses. Google never exposes system-managed private keys, and never retains user-managed private keys.
        /// </summary>
        public readonly string PrivateKeyType;
        /// <summary>
        /// The public key data. Only provided in `GetServiceAccountKey` responses.
        /// </summary>
        public readonly string PublicKeyData;
        /// <summary>
        /// The key can be used after this timestamp.
        /// </summary>
        public readonly string ValidAfterTime;
        /// <summary>
        /// The key can be used before this timestamp. For system-managed key pairs, this timestamp is the end time for the private key signing operation. The public key could still be used for verification for a few hours after this time.
        /// </summary>
        public readonly string ValidBeforeTime;

        [OutputConstructor]
        private GetKeyResult(
            bool disabled,

            string keyAlgorithm,

            string keyOrigin,

            string keyType,

            string name,

            string privateKeyData,

            string privateKeyType,

            string publicKeyData,

            string validAfterTime,

            string validBeforeTime)
        {
            Disabled = disabled;
            KeyAlgorithm = keyAlgorithm;
            KeyOrigin = keyOrigin;
            KeyType = keyType;
            Name = name;
            PrivateKeyData = privateKeyData;
            PrivateKeyType = privateKeyType;
            PublicKeyData = publicKeyData;
            ValidAfterTime = validAfterTime;
            ValidBeforeTime = validBeforeTime;
        }
    }
}
