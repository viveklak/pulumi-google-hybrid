// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1.Outputs
{

    /// <summary>
    /// Details of the final state "forward" and associated resource.
    /// </summary>
    [OutputType]
    public sealed class ForwardInfoResponse
    {
        /// <summary>
        /// URI of the resource that the packet is forwarded to.
        /// </summary>
        public readonly string ResourceUri;
        /// <summary>
        /// Target type where this packet is forwarded to.
        /// </summary>
        public readonly string Target;

        [OutputConstructor]
        private ForwardInfoResponse(
            string resourceUri,

            string target)
        {
            ResourceUri = resourceUri;
            Target = target;
        }
    }
}
