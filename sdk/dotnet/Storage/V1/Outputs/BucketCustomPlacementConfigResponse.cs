// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Outputs
{

    /// <summary>
    /// The bucket's custom placement configuration for Custom Dual Regions.
    /// </summary>
    [OutputType]
    public sealed class BucketCustomPlacementConfigResponse
    {
        /// <summary>
        /// The list of regional locations in which data is placed.
        /// </summary>
        public readonly ImmutableArray<string> DataLocations;

        [OutputConstructor]
        private BucketCustomPlacementConfigResponse(ImmutableArray<string> dataLocations)
        {
            DataLocations = dataLocations;
        }
    }
}
