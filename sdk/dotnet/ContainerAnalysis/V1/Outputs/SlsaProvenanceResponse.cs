// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    [OutputType]
    public sealed class SlsaProvenanceResponse
    {
        /// <summary>
        /// required
        /// </summary>
        public readonly Outputs.SlsaBuilderResponse Builder;
        /// <summary>
        /// The collection of artifacts that influenced the build including sources, dependencies, build tools, base images, and so on. This is considered to be incomplete unless metadata.completeness.materials is true. Unset or null is equivalent to empty.
        /// </summary>
        public readonly ImmutableArray<Outputs.MaterialResponse> Materials;
        public readonly Outputs.SlsaMetadataResponse Metadata;
        /// <summary>
        /// Identifies the configuration used for the build. When combined with materials, this SHOULD fully describe the build, such that re-running this recipe results in bit-for-bit identical output (if the build is reproducible). required
        /// </summary>
        public readonly Outputs.SlsaRecipeResponse Recipe;

        [OutputConstructor]
        private SlsaProvenanceResponse(
            Outputs.SlsaBuilderResponse builder,

            ImmutableArray<Outputs.MaterialResponse> materials,

            Outputs.SlsaMetadataResponse metadata,

            Outputs.SlsaRecipeResponse recipe)
        {
            Builder = builder;
            Materials = materials;
            Metadata = metadata;
            Recipe = recipe;
        }
    }
}
