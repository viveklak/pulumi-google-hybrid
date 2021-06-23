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
    /// Cloud Run fully managed: not supported Cloud Run on GKE: supported EnvVarSource represents a source for the value of an EnvVar.
    /// </summary>
    public sealed class EnvVarSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Run fully managed: not supported Cloud Run on GKE: supported Selects a key of a ConfigMap. +optional
        /// </summary>
        [Input("configMapKeyRef")]
        public Input<Inputs.ConfigMapKeySelectorArgs>? ConfigMapKeyRef { get; set; }

        /// <summary>
        /// Cloud Run fully managed: supported. Selects a key (version) of a secret in Secret Manager. Cloud Run for Anthos: supported. Selects a key of a secret in the pod's namespace. +optional
        /// </summary>
        [Input("secretKeyRef")]
        public Input<Inputs.SecretKeySelectorArgs>? SecretKeyRef { get; set; }

        public EnvVarSourceArgs()
        {
        }
    }
}
