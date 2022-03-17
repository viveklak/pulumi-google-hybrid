// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Outputs
{

    /// <summary>
    /// Not supported by Cloud Run EnvFromSource represents the source of a set of ConfigMaps
    /// </summary>
    [OutputType]
    public sealed class EnvFromSourceResponse
    {
        /// <summary>
        /// (Optional) The ConfigMap to select from
        /// </summary>
        public readonly Outputs.ConfigMapEnvSourceResponse ConfigMapRef;
        /// <summary>
        /// (Optional) An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// (Optional) The Secret to select from
        /// </summary>
        public readonly Outputs.SecretEnvSourceResponse SecretRef;

        [OutputConstructor]
        private EnvFromSourceResponse(
            Outputs.ConfigMapEnvSourceResponse configMapRef,

            string prefix,

            Outputs.SecretEnvSourceResponse secretRef)
        {
            ConfigMapRef = configMapRef;
            Prefix = prefix;
            SecretRef = secretRef;
        }
    }
}
