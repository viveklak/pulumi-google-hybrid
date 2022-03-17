// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// matchRule criteria for request header matches.
    /// </summary>
    [OutputType]
    public sealed class HttpHeaderMatchResponse
    {
        /// <summary>
        /// The value should exactly match contents of exactMatch. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        public readonly string ExactMatch;
        /// <summary>
        /// The name of the HTTP header to match. For matching against the HTTP request's authority, use a headerMatch with the header name ":authority". For matching a request's method, use the headerName ":method". When the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true, only non-binary user-specified custom metadata and the `content-type` header are supported. The following transport-level headers cannot be used in header matching rules: `:authority`, `:method`, `:path`, `:scheme`, `user-agent`, `accept-encoding`, `content-encoding`, `grpc-accept-encoding`, `grpc-encoding`, `grpc-previous-rpc-attempts`, `grpc-tags-bin`, `grpc-timeout` and `grpc-trace-bin`.
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// If set to false, the headerMatch is considered a match if the preceding match criteria are met. If set to true, the headerMatch is considered a match if the preceding match criteria are NOT met. The default setting is false. 
        /// </summary>
        public readonly bool InvertMatch;
        /// <summary>
        /// The value of the header must start with the contents of prefixMatch. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        public readonly string PrefixMatch;
        /// <summary>
        /// A header with the contents of headerName must exist. The match takes place whether or not the request's header has a value. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        public readonly bool PresentMatch;
        /// <summary>
        /// The header value must be an integer and its value must be in the range specified in rangeMatch. If the header does not contain an integer, number or is empty, the match fails. For example for a range [-5, 0] - -3 will match. - 0 will not match. - 0.25 will not match. - -3someString will not match. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set. rangeMatch is not supported for load balancers that have loadBalancingScheme set to EXTERNAL.
        /// </summary>
        public readonly Outputs.Int64RangeMatchResponse RangeMatch;
        /// <summary>
        /// The value of the header must match the regular expression specified in regexMatch. For more information about regular expression syntax, see Syntax. For matching against a port specified in the HTTP request, use a headerMatch with headerName set to PORT and a regular expression that satisfies the RFC2616 Host header's port specifier. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set. regexMatch only applies to load balancers that have loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// </summary>
        public readonly string RegexMatch;
        /// <summary>
        /// The value of the header must end with the contents of suffixMatch. Only one of exactMatch, prefixMatch, suffixMatch, regexMatch, presentMatch or rangeMatch must be set.
        /// </summary>
        public readonly string SuffixMatch;

        [OutputConstructor]
        private HttpHeaderMatchResponse(
            string exactMatch,

            string headerName,

            bool invertMatch,

            string prefixMatch,

            bool presentMatch,

            Outputs.Int64RangeMatchResponse rangeMatch,

            string regexMatch,

            string suffixMatch)
        {
            ExactMatch = exactMatch;
            HeaderName = headerName;
            InvertMatch = invertMatch;
            PrefixMatch = prefixMatch;
            PresentMatch = presentMatch;
            RangeMatch = rangeMatch;
            RegexMatch = regexMatch;
            SuffixMatch = suffixMatch;
        }
    }
}
