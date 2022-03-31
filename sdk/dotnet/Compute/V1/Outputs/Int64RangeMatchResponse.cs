// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// HttpRouteRuleMatch criteria for field values that must stay within the specified integer range.
    /// </summary>
    [OutputType]
    public sealed class Int64RangeMatchResponse
    {
        /// <summary>
        /// The end of the range (exclusive) in signed long integer format.
        /// </summary>
        public readonly string RangeEnd;
        /// <summary>
        /// The start of the range (inclusive) in signed long integer format.
        /// </summary>
        public readonly string RangeStart;

        [OutputConstructor]
        private Int64RangeMatchResponse(
            string rangeEnd,

            string rangeStart)
        {
            RangeEnd = rangeEnd;
            RangeStart = rangeStart;
        }
    }
}
