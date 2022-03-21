// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Outputs
{

    /// <summary>
    /// Specifies how many time series must fail a predicate to trigger a condition. If not specified, then a {count: 1} trigger is used.
    /// </summary>
    [OutputType]
    public sealed class TriggerResponse
    {
        /// <summary>
        /// The absolute number of time series that must fail the predicate for the condition to be triggered.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The percentage of time series that must fail the predicate for the condition to be triggered.
        /// </summary>
        public readonly double Percent;

        [OutputConstructor]
        private TriggerResponse(
            int count,

            double percent)
        {
            Count = count;
            Percent = percent;
        }
    }
}
