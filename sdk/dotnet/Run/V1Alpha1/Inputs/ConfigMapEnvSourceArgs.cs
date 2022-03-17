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
    /// Not supported by Cloud Run ConfigMapEnvSource selects a ConfigMap to populate the environment variables with. The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
    /// </summary>
    public sealed class ConfigMapEnvSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This field should not be used directly as it is meant to be inlined directly into the message. Use the "name" field instead.
        /// </summary>
        [Input("localObjectReference")]
        public Input<Inputs.LocalObjectReferenceArgs>? LocalObjectReference { get; set; }

        /// <summary>
        /// The ConfigMap to select from.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional) Specify whether the ConfigMap must be defined
        /// </summary>
        [Input("optional")]
        public Input<bool>? Optional { get; set; }

        public ConfigMapEnvSourceArgs()
        {
        }
    }
}
