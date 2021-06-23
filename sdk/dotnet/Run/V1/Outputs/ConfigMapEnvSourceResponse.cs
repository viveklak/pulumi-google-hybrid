// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Outputs
{

    [OutputType]
    public sealed class ConfigMapEnvSourceResponse
    {
        /// <summary>
        /// This field should not be used directly as it is meant to be inlined directly into the message. Use the "name" field instead.
        /// </summary>
        public readonly Outputs.LocalObjectReferenceResponse LocalObjectReference;
        /// <summary>
        /// The ConfigMap to select from.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (Optional) Specify whether the ConfigMap must be defined
        /// </summary>
        public readonly bool Optional;

        [OutputConstructor]
        private ConfigMapEnvSourceResponse(
            Outputs.LocalObjectReferenceResponse localObjectReference,

            string name,

            bool optional)
        {
            LocalObjectReference = localObjectReference;
            Name = name;
            Optional = optional;
        }
    }
}
