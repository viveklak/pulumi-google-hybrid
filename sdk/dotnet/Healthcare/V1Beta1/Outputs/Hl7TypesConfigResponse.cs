// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Outputs
{

    /// <summary>
    /// Root config for HL7v2 datatype definitions for a specific HL7v2 version.
    /// </summary>
    [OutputType]
    public sealed class Hl7TypesConfigResponse
    {
        /// <summary>
        /// The HL7v2 type definitions.
        /// </summary>
        public readonly ImmutableArray<Outputs.TypeResponse> Type;
        /// <summary>
        /// The version selectors that this config applies to. A message must match ALL version sources to apply.
        /// </summary>
        public readonly ImmutableArray<Outputs.VersionSourceResponse> Version;

        [OutputConstructor]
        private Hl7TypesConfigResponse(
            ImmutableArray<Outputs.TypeResponse> type,

            ImmutableArray<Outputs.VersionSourceResponse> version)
        {
            Type = type;
            Version = version;
        }
    }
}
