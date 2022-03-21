// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// Layer holds metadata specific to a layer of a Docker image.
    /// </summary>
    public sealed class LayerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The recovered arguments to the Dockerfile directive.
        /// </summary>
        [Input("arguments")]
        public Input<string>? Arguments { get; set; }

        /// <summary>
        /// The recovered Dockerfile directive used to construct this layer.
        /// </summary>
        [Input("directive", required: true)]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.LayerDirective> Directive { get; set; } = null!;

        public LayerArgs()
        {
        }
    }
}
