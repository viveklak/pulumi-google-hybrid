// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    [OutputType]
    public sealed class ComplianceVersionResponse
    {
        /// <summary>
        /// The CPE URI (https://cpe.mitre.org/specification/) this benchmark is applicable to.
        /// </summary>
        public readonly string CpeUri;
        /// <summary>
        /// The version of the benchmark. This is set to the version of the OS-specific CIS document the benchmark is defined in.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private ComplianceVersionResponse(
            string cpeUri,

            string version)
        {
            CpeUri = cpeUri;
            Version = version;
        }
    }
}
