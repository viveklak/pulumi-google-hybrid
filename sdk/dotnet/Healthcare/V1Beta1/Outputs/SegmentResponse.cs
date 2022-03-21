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
    /// A segment in a structured format.
    /// </summary>
    [OutputType]
    public sealed class SegmentResponse
    {
        /// <summary>
        /// A mapping from the positional location to the value. The key string uses zero-based indexes separated by dots to identify Fields, components and sub-components. A bracket notation is also used to identify different instances of a repeated field. Regex for key: (\d+)(\[\d+\])?(.\d+)?(.\d+)? Examples of (key, value) pairs: * (0.1, "hemoglobin") denotes that the first component of Field 0 has the value "hemoglobin". * (1.1.2, "CBC") denotes that the second sub-component of the first component of Field 1 has the value "CBC". * (1[0].1, "HbA1c") denotes that the first component of the first Instance of Field 1, which is repeated, has the value "HbA1c".
        /// </summary>
        public readonly ImmutableDictionary<string, string> Fields;
        /// <summary>
        /// A string that indicates the type of segment. For example, EVN or PID.
        /// </summary>
        public readonly string SegmentId;
        /// <summary>
        /// Set ID for segments that can be in a set. This can be empty if it's missing or isn't applicable.
        /// </summary>
        public readonly string SetId;

        [OutputConstructor]
        private SegmentResponse(
            ImmutableDictionary<string, string> fields,

            string segmentId,

            string setId)
        {
            Fields = fields;
            SegmentId = segmentId;
            SetId = setId;
        }
    }
}
