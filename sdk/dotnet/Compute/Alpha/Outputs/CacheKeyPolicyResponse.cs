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
    /// Message containing what to include in the cache key for a request for Cloud CDN.
    /// </summary>
    [OutputType]
    public sealed class CacheKeyPolicyResponse
    {
        /// <summary>
        /// If true, requests to different hosts will be cached separately.
        /// </summary>
        public readonly bool IncludeHost;
        /// <summary>
        /// Allows HTTP request headers (by name) to be used in the cache key.
        /// </summary>
        public readonly ImmutableArray<string> IncludeHttpHeaders;
        /// <summary>
        /// Allows HTTP cookies (by name) to be used in the cache key. The name=value pair will be used in the cache key Cloud CDN generates.
        /// </summary>
        public readonly ImmutableArray<string> IncludeNamedCookies;
        /// <summary>
        /// If true, http and https requests will be cached separately.
        /// </summary>
        public readonly bool IncludeProtocol;
        /// <summary>
        /// If true, include query string parameters in the cache key according to query_string_whitelist and query_string_blacklist. If neither is set, the entire query string will be included. If false, the query string will be excluded from the cache key entirely.
        /// </summary>
        public readonly bool IncludeQueryString;
        /// <summary>
        /// Names of query string parameters to exclude in cache keys. All other parameters will be included. Either specify query_string_whitelist or query_string_blacklist, not both. '&amp;' and '=' will be percent encoded and not treated as delimiters.
        /// </summary>
        public readonly ImmutableArray<string> QueryStringBlacklist;
        /// <summary>
        /// Names of query string parameters to include in cache keys. All other parameters will be excluded. Either specify query_string_whitelist or query_string_blacklist, not both. '&amp;' and '=' will be percent encoded and not treated as delimiters.
        /// </summary>
        public readonly ImmutableArray<string> QueryStringWhitelist;

        [OutputConstructor]
        private CacheKeyPolicyResponse(
            bool includeHost,

            ImmutableArray<string> includeHttpHeaders,

            ImmutableArray<string> includeNamedCookies,

            bool includeProtocol,

            bool includeQueryString,

            ImmutableArray<string> queryStringBlacklist,

            ImmutableArray<string> queryStringWhitelist)
        {
            IncludeHost = includeHost;
            IncludeHttpHeaders = includeHttpHeaders;
            IncludeNamedCookies = includeNamedCookies;
            IncludeProtocol = includeProtocol;
            IncludeQueryString = includeQueryString;
            QueryStringBlacklist = queryStringBlacklist;
            QueryStringWhitelist = queryStringWhitelist;
        }
    }
}
