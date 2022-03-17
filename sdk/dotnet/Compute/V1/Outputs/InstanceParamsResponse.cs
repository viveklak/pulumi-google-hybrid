// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Additional instance params.
    /// </summary>
    [OutputType]
    public sealed class InstanceParamsResponse
    {
        /// <summary>
        /// Resource manager tags to be bound to the instance. Tag keys and values have the same definition as resource manager tags. Keys must be in the format `tagKeys/{tag_key_id}`, and values are in the format `tagValues/456`. The field is ignored (both PUT &amp; PATCH) when empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ResourceManagerTags;

        [OutputConstructor]
        private InstanceParamsResponse(ImmutableDictionary<string, string> resourceManagerTags)
        {
            ResourceManagerTags = resourceManagerTags;
        }
    }
}
