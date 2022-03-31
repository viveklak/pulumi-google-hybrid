// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Bypass the cache when the specified request headers are present, e.g. Pragma or Authorization headers. Values are case insensitive. The presence of such a header overrides the cache_mode setting.
    /// </summary>
    public sealed class BackendServiceCdnPolicyBypassCacheOnRequestHeaderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The header field name to match on when bypassing cache. Values are case-insensitive.
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        public BackendServiceCdnPolicyBypassCacheOnRequestHeaderArgs()
        {
        }
    }
}
