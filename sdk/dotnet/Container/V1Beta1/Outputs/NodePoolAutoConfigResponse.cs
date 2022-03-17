// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// node pool configs that apply to all auto-provisioned node pools in autopilot clusters and node auto-provisioning enabled clusters
    /// </summary>
    [OutputType]
    public sealed class NodePoolAutoConfigResponse
    {
        /// <summary>
        /// The list of instance tags applied to all nodes. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during cluster creation. Each tag within the list must comply with RFC1035.
        /// </summary>
        public readonly Outputs.NetworkTagsResponse NetworkTags;

        [OutputConstructor]
        private NodePoolAutoConfigResponse(Outputs.NetworkTagsResponse networkTags)
        {
            NetworkTags = networkTags;
        }
    }
}
