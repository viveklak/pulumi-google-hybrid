// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1Alpha1.Inputs
{

    /// <summary>
    /// Not supported by Cloud Run EnvFromSource represents the source of a set of ConfigMaps
    /// </summary>
    public sealed class EnvFromSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) The ConfigMap to select from
        /// </summary>
        [Input("configMapRef")]
        public Input<Inputs.ConfigMapEnvSourceArgs>? ConfigMapRef { get; set; }

        /// <summary>
        /// (Optional) An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// (Optional) The Secret to select from
        /// </summary>
        [Input("secretRef")]
        public Input<Inputs.SecretEnvSourceArgs>? SecretRef { get; set; }

        public EnvFromSourceArgs()
        {
        }
    }
}
