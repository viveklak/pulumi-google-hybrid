// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// Defines the network configuration for the pool.
    /// </summary>
    public sealed class NetworkConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Option to configure network egress for the workers.
        /// </summary>
        [Input("egressOption")]
        public Input<Pulumi.GoogleNative.CloudBuild.V1.NetworkConfigEgressOption>? EgressOption { get; set; }

        /// <summary>
        /// Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to `WorkerPool.project_id` on the service producer network. Must be in the format `projects/{project}/global/networks/{network}`, where `{project}` is a project number, such as `12345`, and `{network}` is the name of a VPC network in the project. See [Understanding network configuration options](https://cloud.google.com/build/docs/private-pools/set-up-private-pool-environment)
        /// </summary>
        [Input("peeredNetwork", required: true)]
        public Input<string> PeeredNetwork { get; set; } = null!;

        public NetworkConfigArgs()
        {
        }
    }
}
