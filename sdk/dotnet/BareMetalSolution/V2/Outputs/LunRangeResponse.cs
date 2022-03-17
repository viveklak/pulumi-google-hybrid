// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Outputs
{

    /// <summary>
    /// A LUN(Logical Unit Number) range.
    /// </summary>
    [OutputType]
    public sealed class LunRangeResponse
    {
        /// <summary>
        /// Number of LUNs to create.
        /// </summary>
        public readonly int Quantity;
        /// <summary>
        /// The requested size of each LUN, in GB.
        /// </summary>
        public readonly int SizeGb;

        [OutputConstructor]
        private LunRangeResponse(
            int quantity,

            int sizeGb)
        {
            Quantity = quantity;
            SizeGb = sizeGb;
        }
    }
}
