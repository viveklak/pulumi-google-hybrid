// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V1.Outputs
{

    /// <summary>
    /// Configuration for a secret environment variable. It has the information necessary to fetch the secret value from secret manager and expose it as an environment variable.
    /// </summary>
    [OutputType]
    public sealed class SecretEnvVarResponse
    {
        /// <summary>
        /// Name of the environment variable.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Name of the secret in secret manager (not the full resource name).
        /// </summary>
        public readonly string Secret;
        /// <summary>
        /// Version of the secret (version number or the string 'latest'). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new instances start.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private SecretEnvVarResponse(
            string key,

            string project,

            string secret,

            string version)
        {
            Key = key;
            Project = project;
            Secret = secret;
            Version = version;
        }
    }
}
