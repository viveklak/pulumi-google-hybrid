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
    /// Indicates that the builder claims certain fields in this message to be complete.
    /// </summary>
    [OutputType]
    public sealed class CompletenessResponse
    {
        /// <summary>
        /// If true, the builder claims that recipe.arguments is complete, meaning that all external inputs are properly captured in the recipe.
        /// </summary>
        public readonly bool Arguments;
        /// <summary>
        /// If true, the builder claims that recipe.environment is claimed to be complete.
        /// </summary>
        public readonly bool Environment;
        /// <summary>
        /// If true, the builder claims that materials are complete, usually through some controls to prevent network access. Sometimes called "hermetic".
        /// </summary>
        public readonly bool Materials;

        [OutputConstructor]
        private CompletenessResponse(
            bool arguments,

            bool environment,

            bool materials)
        {
            Arguments = arguments;
            Environment = environment;
            Materials = materials;
        }
    }
}
