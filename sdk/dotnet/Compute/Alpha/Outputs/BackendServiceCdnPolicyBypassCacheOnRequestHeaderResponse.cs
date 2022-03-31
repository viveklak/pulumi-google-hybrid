// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Bypass the cache when the specified request headers are present, e.g. Pragma or Authorization headers. Values are case insensitive. The presence of such a header overrides the cache_mode setting.
    /// </summary>
    [OutputType]
    public sealed class BackendServiceCdnPolicyBypassCacheOnRequestHeaderResponse
    {
        /// <summary>
        /// The header field name to match on when bypassing cache. Values are case-insensitive.
        /// </summary>
        public readonly string HeaderName;

        [OutputConstructor]
        private BackendServiceCdnPolicyBypassCacheOnRequestHeaderResponse(string headerName)
        {
            HeaderName = headerName;
        }
    }
}
