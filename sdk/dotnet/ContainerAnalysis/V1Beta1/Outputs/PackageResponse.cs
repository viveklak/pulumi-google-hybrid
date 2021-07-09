// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    [OutputType]
    public sealed class PackageResponse
    {
        /// <summary>
        /// The various channels by which a package is distributed.
        /// </summary>
        public readonly ImmutableArray<Outputs.DistributionResponse> Distribution;
        /// <summary>
        /// Immutable. The name of the package.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private PackageResponse(
            ImmutableArray<Outputs.DistributionResponse> distribution,

            string name)
        {
            Distribution = distribution;
            Name = name;
        }
    }
}
