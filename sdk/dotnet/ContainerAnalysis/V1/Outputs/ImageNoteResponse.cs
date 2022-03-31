// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// Basis describes the base image portion (Note) of the DockerImage relationship. Linked occurrences are derived from this or an equivalent image via: FROM Or an equivalent reference, e.g., a tag of the resource_url.
    /// </summary>
    [OutputType]
    public sealed class ImageNoteResponse
    {
        /// <summary>
        /// Immutable. The fingerprint of the base image.
        /// </summary>
        public readonly Outputs.FingerprintResponse Fingerprint;
        /// <summary>
        /// Immutable. The resource_url for the resource representing the basis of associated occurrence images.
        /// </summary>
        public readonly string ResourceUrl;

        [OutputConstructor]
        private ImageNoteResponse(
            Outputs.FingerprintResponse fingerprint,

            string resourceUrl)
        {
            Fingerprint = fingerprint;
            ResourceUrl = resourceUrl;
        }
    }
}
