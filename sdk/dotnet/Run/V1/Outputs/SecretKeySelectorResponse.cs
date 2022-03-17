// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Outputs
{

    /// <summary>
    /// SecretKeySelector selects a key of a Secret.
    /// </summary>
    [OutputType]
    public sealed class SecretKeySelectorResponse
    {
        /// <summary>
        /// A Cloud Secret Manager secret version. Must be 'latest' for the latest version or an integer for a specific version. The key of the secret to select from. Must be a valid secret key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// This field should not be used directly as it is meant to be inlined directly into the message. Use the "name" field instead.
        /// </summary>
        public readonly Outputs.LocalObjectReferenceResponse LocalObjectReference;
        /// <summary>
        /// The name of the secret in Cloud Secret Manager. By default, the secret is assumed to be in the same project. If the secret is in another project, you must define an alias. An alias definition has the form: :projects//secrets/. If multiple alias definitions are needed, they must be separated by commas. The alias definitions must be set on the run.googleapis.com/secrets annotation. The name of the secret in the pod's namespace to select from.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (Optional) Specify whether the Secret or its key must be defined
        /// </summary>
        public readonly bool Optional;

        [OutputConstructor]
        private SecretKeySelectorResponse(
            string key,

            Outputs.LocalObjectReferenceResponse localObjectReference,

            string name,

            bool optional)
        {
            Key = key;
            LocalObjectReference = localObjectReference;
            Name = name;
            Optional = optional;
        }
    }
}
