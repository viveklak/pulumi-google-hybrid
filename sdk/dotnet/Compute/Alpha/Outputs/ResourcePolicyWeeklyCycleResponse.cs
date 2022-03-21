// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Time window specified for weekly operations.
    /// </summary>
    [OutputType]
    public sealed class ResourcePolicyWeeklyCycleResponse
    {
        /// <summary>
        /// Up to 7 intervals/windows, one for each day of the week.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourcePolicyWeeklyCycleDayOfWeekResponse> DayOfWeeks;

        [OutputConstructor]
        private ResourcePolicyWeeklyCycleResponse(ImmutableArray<Outputs.ResourcePolicyWeeklyCycleDayOfWeekResponse> dayOfWeeks)
        {
            DayOfWeeks = dayOfWeeks;
        }
    }
}
