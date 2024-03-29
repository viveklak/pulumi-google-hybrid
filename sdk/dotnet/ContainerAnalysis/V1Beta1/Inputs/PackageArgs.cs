// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// This represents a particular package that is distributed over various channels. E.g., glibc (aka libc6) is distributed by many, at various versions.
    /// </summary>
    public sealed class PackageArgs : Pulumi.ResourceArgs
    {
        [Input("distribution")]
        private InputList<Inputs.DistributionArgs>? _distribution;

        /// <summary>
        /// The various channels by which a package is distributed.
        /// </summary>
        public InputList<Inputs.DistributionArgs> Distribution
        {
            get => _distribution ?? (_distribution = new InputList<Inputs.DistributionArgs>());
            set => _distribution = value;
        }

        /// <summary>
        /// Immutable. The name of the package.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public PackageArgs()
        {
        }
    }
}
