// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Configuration for the underlying infrastructure used to run workloads.
    /// </summary>
    public sealed class GoogleCloudDataplexV1EnvironmentInfrastructureSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Compute resources needed for analyze interactive workloads.
        /// </summary>
        [Input("compute")]
        public Input<Inputs.GoogleCloudDataplexV1EnvironmentInfrastructureSpecComputeResourcesArgs>? Compute { get; set; }

        /// <summary>
        /// Software Runtime Configuration for analyze interactive workloads.
        /// </summary>
        [Input("osImage", required: true)]
        public Input<Inputs.GoogleCloudDataplexV1EnvironmentInfrastructureSpecOsImageRuntimeArgs> OsImage { get; set; } = null!;

        public GoogleCloudDataplexV1EnvironmentInfrastructureSpecArgs()
        {
        }
    }
}
