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
    /// A note that indicates a type of analysis a provider would perform. This note exists in a provider's project. A `Discovery` occurrence is created in a consumer's project at the start of analysis. The occurrence's operation will indicate the status of the analysis. Absence of an occurrence linked to this note for a resource indicates that analysis hasn't started.
    /// </summary>
    [OutputType]
    public sealed class DiscoveryResponse
    {
        /// <summary>
        /// The kind of analysis that is handled by this discovery.
        /// </summary>
        public readonly string AnalysisKind;

        [OutputConstructor]
        private DiscoveryResponse(string analysisKind)
        {
            AnalysisKind = analysisKind;
        }
    }
}
