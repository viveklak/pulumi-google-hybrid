// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// Describes the CIS benchmark version that is applicable to a given OS and os version.
    /// </summary>
    public sealed class ComplianceVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPE URI (https://cpe.mitre.org/specification/) this benchmark is applicable to.
        /// </summary>
        [Input("cpeUri")]
        public Input<string>? CpeUri { get; set; }

        /// <summary>
        /// The version of the benchmark. This is set to the version of the OS-specific CIS document the benchmark is defined in.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ComplianceVersionArgs()
        {
        }
    }
}
