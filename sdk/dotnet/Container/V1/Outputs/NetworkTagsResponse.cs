// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// Collection of Compute Engine network tags that can be applied to a node's underlying VM instance.
    /// </summary>
    [OutputType]
    public sealed class NetworkTagsResponse
    {
        /// <summary>
        /// List of network tags.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private NetworkTagsResponse(ImmutableArray<string> tags)
        {
            Tags = tags;
        }
    }
}
