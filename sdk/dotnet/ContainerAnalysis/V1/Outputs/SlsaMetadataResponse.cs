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
    /// Other properties of the build.
    /// </summary>
    [OutputType]
    public sealed class SlsaMetadataResponse
    {
        /// <summary>
        /// The timestamp of when the build completed.
        /// </summary>
        public readonly string BuildFinishedOn;
        /// <summary>
        /// Identifies the particular build invocation, which can be useful for finding associated logs or other ad-hoc analysis. The value SHOULD be globally unique, per in-toto Provenance spec.
        /// </summary>
        public readonly string BuildInvocationId;
        /// <summary>
        /// The timestamp of when the build started.
        /// </summary>
        public readonly string BuildStartedOn;
        /// <summary>
        /// Indicates that the builder claims certain fields in this message to be complete.
        /// </summary>
        public readonly Outputs.SlsaCompletenessResponse Completeness;
        /// <summary>
        /// If true, the builder claims that running the recipe on materials will produce bit-for-bit identical output.
        /// </summary>
        public readonly bool Reproducible;

        [OutputConstructor]
        private SlsaMetadataResponse(
            string buildFinishedOn,

            string buildInvocationId,

            string buildStartedOn,

            Outputs.SlsaCompletenessResponse completeness,

            bool reproducible)
        {
            BuildFinishedOn = buildFinishedOn;
            BuildInvocationId = buildInvocationId;
            BuildStartedOn = buildStartedOn;
            Completeness = completeness;
            Reproducible = reproducible;
        }
    }
}
