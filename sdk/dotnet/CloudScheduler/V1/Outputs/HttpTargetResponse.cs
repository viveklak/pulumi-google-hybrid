// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudScheduler.V1.Outputs
{

    [OutputType]
    public sealed class HttpTargetResponse
    {
        /// <summary>
        /// HTTP request body. A request body is allowed only if the HTTP method is POST, PUT, or PATCH. It is an error to set body on a job with an incompatible HttpMethod.
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// The user can specify HTTP request headers to send with the job's HTTP request. This map contains the header field names and values. Repeated headers are not supported, but a header value can contain commas. These headers represent a subset of the headers that will accompany the job's HTTP request. Some HTTP request headers will be ignored or replaced. A partial list of headers that will be ignored or replaced is below: - Host: This will be computed by Cloud Scheduler and derived from uri. * `Content-Length`: This will be computed by Cloud Scheduler. * `User-Agent`: This will be set to `"Google-Cloud-Scheduler"`. * `X-Google-*`: Google internal use only. * `X-AppEngine-*`: Google internal use only. The total size of headers must be less than 80KB.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Headers;
        /// <summary>
        /// Which HTTP method to use for the request.
        /// </summary>
        public readonly string HttpMethod;
        /// <summary>
        /// If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2) will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization should generally only be used when calling Google APIs hosted on *.googleapis.com.
        /// </summary>
        public readonly Outputs.OAuthTokenResponse OauthToken;
        /// <summary>
        /// If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect) token will be generated and attached as an `Authorization` header in the HTTP request. This type of authorization can be used for many scenarios, including calling Cloud Run, or endpoints where you intend to validate the token yourself.
        /// </summary>
        public readonly Outputs.OidcTokenResponse OidcToken;
        /// <summary>
        /// The full URI path that the request will be sent to. This string must begin with either "http://" or "https://". Some examples of valid values for uri are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler will encode some characters for safety and compatibility. The maximum allowed URL length is 2083 characters after encoding.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private HttpTargetResponse(
            string body,

            ImmutableDictionary<string, string> headers,

            string httpMethod,

            Outputs.OAuthTokenResponse oauthToken,

            Outputs.OidcTokenResponse oidcToken,

            string uri)
        {
            Body = body;
            Headers = headers;
            HttpMethod = httpMethod;
            OauthToken = oauthToken;
            OidcToken = oidcToken;
            Uri = uri;
        }
    }
}
