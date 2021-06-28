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
    public sealed class RouteWarningsItemDataItemResponse
    {
        /// <summary>
        /// A key that provides more detail on the warning being returned. For example, for warnings where there are no results in a list request for a particular zone, this key might be scope and the key value might be the zone name. Other examples might be a key indicating a deprecated resource and a suggested replacement, or a warning about invalid network settings (for example, if an instance attempts to perform IP forwarding but is not enabled for IP forwarding).
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// A warning data value corresponding to the key.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private RouteWarningsItemDataItemResponse(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
