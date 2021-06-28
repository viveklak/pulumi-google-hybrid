// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class FixedOrPercentResponse
    {
        /// <summary>
        /// Absolute value of VM instances calculated based on the specific mode.
        /// 
        ///  
        /// - If the value is fixed, then the calculated value is equal to the fixed value. 
        /// - If the value is a percent, then the calculated value is percent/100 * targetSize. For example, the calculated value of a 80% of a managed instance group with 150 instances would be (80/100 * 150) = 120 VM instances. If there is a remainder, the number is rounded up.
        /// </summary>
        public readonly int Calculated;
        /// <summary>
        /// Specifies a fixed number of VM instances. This must be a positive integer.
        /// </summary>
        public readonly int Fixed;
        /// <summary>
        /// Specifies a percentage of instances between 0 to 100%, inclusive. For example, specify 80 for 80%.
        /// </summary>
        public readonly int Percent;

        [OutputConstructor]
        private FixedOrPercentResponse(
            int calculated,

            int @fixed,

            int percent)
        {
            Calculated = calculated;
            Fixed = @fixed;
            Percent = percent;
        }
    }
}
