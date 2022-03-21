// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// A note that indicates a type of analysis a provider would perform. This note exists in a provider's project. A `Discovery` occurrence is created in a consumer's project at the start of analysis.
    /// </summary>
    [OutputType]
    public sealed class DiscoveryResponse
    {
        /// <summary>
        /// Immutable. The kind of analysis that is handled by this discovery.
        /// </summary>
        public readonly string AnalysisKind;

        [OutputConstructor]
        private DiscoveryResponse(string analysisKind)
        {
            AnalysisKind = analysisKind;
        }
    }
}
