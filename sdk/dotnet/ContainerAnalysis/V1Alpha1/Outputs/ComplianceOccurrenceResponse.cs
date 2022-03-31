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
    /// An indication that the compliance checks in the associated ComplianceNote were not satisfied for particular resources or a specified reason.
    /// </summary>
    [OutputType]
    public sealed class ComplianceOccurrenceResponse
    {
        /// <summary>
        /// The reason for non compliance of these files.
        /// </summary>
        public readonly string NonComplianceReason;
        /// <summary>
        /// A list of files which are violating compliance checks.
        /// </summary>
        public readonly ImmutableArray<Outputs.NonCompliantFileResponse> NonCompliantFiles;

        [OutputConstructor]
        private ComplianceOccurrenceResponse(
            string nonComplianceReason,

            ImmutableArray<Outputs.NonCompliantFileResponse> nonCompliantFiles)
        {
            NonComplianceReason = nonComplianceReason;
            NonCompliantFiles = nonCompliantFiles;
        }
    }
}
