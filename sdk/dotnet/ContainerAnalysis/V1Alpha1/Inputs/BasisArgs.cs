// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// Basis describes the base image portion (Note) of the DockerImage relationship. Linked occurrences are derived from this or an equivalent image via: FROM Or an equivalent reference, e.g. a tag of the resource_url.
    /// </summary>
    public sealed class BasisArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The fingerprint of the base image.
        /// </summary>
        [Input("fingerprint")]
        public Input<Inputs.FingerprintArgs>? Fingerprint { get; set; }

        /// <summary>
        /// The resource_url for the resource representing the basis of associated occurrence images.
        /// </summary>
        [Input("resourceUrl")]
        public Input<string>? ResourceUrl { get; set; }

        public BasisArgs()
        {
        }
    }
}
