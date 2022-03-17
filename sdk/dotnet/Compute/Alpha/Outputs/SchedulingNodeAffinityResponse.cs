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
    /// Node Affinity: the configuration of desired nodes onto which this Instance could be scheduled.
    /// </summary>
    [OutputType]
    public sealed class SchedulingNodeAffinityResponse
    {
        /// <summary>
        /// Corresponds to the label key of Node resource.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Defines the operation of node selection. Valid operators are IN for affinity and NOT_IN for anti-affinity.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Corresponds to the label values of Node resource.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private SchedulingNodeAffinityResponse(
            string key,

            string @operator,

            ImmutableArray<string> values)
        {
            Key = key;
            Operator = @operator;
            Values = values;
        }
    }
}
