// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// Artifact describes a build product.
    /// </summary>
    [OutputType]
    public sealed class ArtifactResponse
    {
        /// <summary>
        /// Hash or checksum value of a binary, or Docker Registry 2.0 digest of a container.
        /// </summary>
        public readonly string Checksum;
        /// <summary>
        /// Name of the artifact. This may be the path to a binary or jar file, or in the case of a container build, the name used to push the container image to Google Container Registry, as presented to `docker push`. This field is deprecated in favor of the plural `names` field; it continues to exist here to allow existing BuildProvenance serialized to json in google.devtools.containeranalysis.v1alpha1.BuildDetails.provenance_bytes to deserialize back into proto.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Related artifact names. This may be the path to a binary or jar file, or in the case of a container build, the name used to push the container image to Google Container Registry, as presented to `docker push`. Note that a single Artifact ID can have multiple names, for example if two tags are applied to one image.
        /// </summary>
        public readonly ImmutableArray<string> Names;

        [OutputConstructor]
        private ArtifactResponse(
            string checksum,

            string name,

            ImmutableArray<string> names)
        {
            Checksum = checksum;
            Name = name;
            Names = names;
        }
    }
}
