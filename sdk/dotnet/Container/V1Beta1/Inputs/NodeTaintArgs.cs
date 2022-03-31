// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Kubernetes taint is comprised of three fields: key, value, and effect. Effect can only be one of three types: NoSchedule, PreferNoSchedule or NoExecute. See [here](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration) for more information, including usage and the valid values.
    /// </summary>
    public sealed class NodeTaintArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Effect for taint.
        /// </summary>
        [Input("effect")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.NodeTaintEffect>? Effect { get; set; }

        /// <summary>
        /// Key for taint.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Value for taint.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NodeTaintArgs()
        {
        }
    }
}
