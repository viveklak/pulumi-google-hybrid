// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class NodeGroupAutoscalingPolicyResponse
    {
        /// <summary>
        /// The maximum number of nodes that the group should have. Must be set if autoscaling is enabled. Maximum value allowed is 100.
        /// </summary>
        public readonly int MaxNodes;
        /// <summary>
        /// The minimum number of nodes that the group should have.
        /// </summary>
        public readonly int MinNodes;
        /// <summary>
        /// The autoscaling mode. Set to one of: ON, OFF, or ONLY_SCALE_OUT. For more information, see Autoscaler modes.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private NodeGroupAutoscalingPolicyResponse(
            int maxNodes,

            int minNodes,

            string mode)
        {
            MaxNodes = maxNodes;
            MinNodes = minNodes;
            Mode = mode;
        }
    }
}
