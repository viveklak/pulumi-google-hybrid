// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSLogin.V1Beta
{
    public static class GetSshPublicKey
    {
        /// <summary>
        /// Retrieves an SSH public key.
        /// </summary>
        public static Task<GetSshPublicKeyResult> InvokeAsync(GetSshPublicKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSshPublicKeyResult>("google-native:oslogin/v1beta:getSshPublicKey", args ?? new GetSshPublicKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves an SSH public key.
        /// </summary>
        public static Output<GetSshPublicKeyResult> Invoke(GetSshPublicKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSshPublicKeyResult>("google-native:oslogin/v1beta:getSshPublicKey", args ?? new GetSshPublicKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSshPublicKeyArgs : Pulumi.InvokeArgs
    {
        [Input("sshPublicKeyId", required: true)]
        public string SshPublicKeyId { get; set; } = null!;

        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetSshPublicKeyArgs()
        {
        }
    }

    public sealed class GetSshPublicKeyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("sshPublicKeyId", required: true)]
        public Input<string> SshPublicKeyId { get; set; } = null!;

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetSshPublicKeyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSshPublicKeyResult
    {
        /// <summary>
        /// An expiration time in microseconds since epoch.
        /// </summary>
        public readonly string ExpirationTimeUsec;
        /// <summary>
        /// The SHA-256 fingerprint of the SSH public key.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// Public key text in SSH format, defined by RFC4253 section 6.6.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The canonical resource name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetSshPublicKeyResult(
            string expirationTimeUsec,

            string fingerprint,

            string key,

            string name)
        {
            ExpirationTimeUsec = expirationTimeUsec;
            Fingerprint = fingerprint;
            Key = key;
            Name = name;
        }
    }
}
