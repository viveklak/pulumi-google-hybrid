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
    public sealed class UrlMapTestResponse
    {
        /// <summary>
        /// Description of this test case.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The expected output URL evaluated by load balancer containing the scheme, host, path and query parameters. For rules that forward requests to backends, the test passes only when expectedOutputUrl matches the request forwarded by load balancer to backends. For rules with urlRewrite, the test verifies that the forwarded request matches hostRewrite and pathPrefixRewrite in the urlRewrite action. When service is specified, expectedOutputUrl`s scheme is ignored. For rules with urlRedirect, the test passes only if expectedOutputUrl matches the URL in the load balancer's redirect response. If urlRedirect specifies https_redirect, the test passes only if the scheme in expectedOutputUrl is also set to https. If urlRedirect specifies strip_query, the test passes only if expectedOutputUrl does not contain any query parameters. expectedOutputUrl is optional when service is specified.
        /// </summary>
        public readonly string ExpectedOutputUrl;
        /// <summary>
        /// For rules with urlRedirect, the test passes only if expectedRedirectResponseCode matches the HTTP status code in load balancer's redirect response. expectedRedirectResponseCode cannot be set when service is set.
        /// </summary>
        public readonly int ExpectedRedirectResponseCode;
        /// <summary>
        /// HTTP headers for this request. If headers contains a host header, then host must also match the header value.
        /// </summary>
        public readonly ImmutableArray<Outputs.UrlMapTestHeaderResponse> Headers;
        /// <summary>
        /// Host portion of the URL. If headers contains a host header, then host must also match the header value.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Path portion of the URL.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Expected BackendService or BackendBucket resource the given URL should be mapped to. service cannot be set if expectedRedirectResponseCode is set.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private UrlMapTestResponse(
            string description,

            string expectedOutputUrl,

            int expectedRedirectResponseCode,

            ImmutableArray<Outputs.UrlMapTestHeaderResponse> headers,

            string host,

            string path,

            string service)
        {
            Description = description;
            ExpectedOutputUrl = expectedOutputUrl;
            ExpectedRedirectResponseCode = expectedRedirectResponseCode;
            Headers = headers;
            Host = host;
            Path = path;
            Service = service;
        }
    }
}
