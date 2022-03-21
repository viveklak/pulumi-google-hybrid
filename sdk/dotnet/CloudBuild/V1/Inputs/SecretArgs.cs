// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// Pairs a set of secret environment variables containing encrypted values with the Cloud KMS key to use to decrypt the value. Note: Use `kmsKeyName` with `available_secrets` instead of using `kmsKeyName` with `secret`. For instructions see: https://cloud.google.com/cloud-build/docs/securing-builds/use-encrypted-credentials.
    /// </summary>
    public sealed class SecretArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud KMS key name to use to decrypt these envs.
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        [Input("secretEnv")]
        private InputMap<string>? _secretEnv;

        /// <summary>
        /// Map of environment variable name to its encrypted value. Secret environment variables must be unique across all of a build's secrets, and must be used by at least one build step. Values can be at most 64 KB in size. There can be at most 100 secret values across all of a build's secrets.
        /// </summary>
        public InputMap<string> SecretEnv
        {
            get => _secretEnv ?? (_secretEnv = new InputMap<string>());
            set => _secretEnv = value;
        }

        public SecretArgs()
        {
        }
    }
}
