// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Outputs
{

    /// <summary>
    /// Information involved in an HTTP/HTTPS Uptime check request.
    /// </summary>
    [OutputType]
    public sealed class HttpCheckResponse
    {
        /// <summary>
        /// The authentication information. Optional when creating an HTTP check; defaults to empty.
        /// </summary>
        public readonly Outputs.BasicAuthenticationResponse AuthInfo;
        /// <summary>
        /// The request body associated with the HTTP POST request. If content_type is URL_ENCODED, the body passed in must be URL-encoded. Users can provide a Content-Length header via the headers field or the API will do so. If the request_method is GET and body is not empty, the API will return an error. The maximum byte size is 1 megabyte.Note: If client libraries aren't used (which performs the conversion automatically) base64 encode your body data since the field is of bytes type.
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// The content type header to use for the check. The following configurations result in errors: 1. Content type is specified in both the headers field and the content_type field. 2. Request method is GET and content_type is not TYPE_UNSPECIFIED 3. Request method is POST and content_type is TYPE_UNSPECIFIED. 4. Request method is POST and a "Content-Type" header is provided via headers field. The content_type field should be used instead.
        /// </summary>
        public readonly string ContentType;
        /// <summary>
        /// The list of headers to send as part of the Uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Headers;
        /// <summary>
        /// Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to true then the headers will be obscured with ******.
        /// </summary>
        public readonly bool MaskHeaders;
        /// <summary>
        /// Optional (defaults to "/"). The path to the page against which to run the check. Will be combined with the host (specified within the monitored_resource) and port to construct the full URL. If the provided path does not begin with "/", a "/" will be prepended automatically.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Optional (defaults to 80 when use_ssl is false, and 443 when use_ssl is true). The TCP port on the HTTP server against which to run the check. Will be combined with host (specified within the monitored_resource) and path to construct the full URL.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The HTTP request method to use for the check. If set to METHOD_UNSPECIFIED then request_method defaults to GET.
        /// </summary>
        public readonly string RequestMethod;
        /// <summary>
        /// If true, use HTTPS instead of HTTP to run the check.
        /// </summary>
        public readonly bool UseSsl;
        /// <summary>
        /// Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where monitored_resource is set to uptime_url. If use_ssl is false, setting validate_ssl to true has no effect.
        /// </summary>
        public readonly bool ValidateSsl;

        [OutputConstructor]
        private HttpCheckResponse(
            Outputs.BasicAuthenticationResponse authInfo,

            string body,

            string contentType,

            ImmutableDictionary<string, string> headers,

            bool maskHeaders,

            string path,

            int port,

            string requestMethod,

            bool useSsl,

            bool validateSsl)
        {
            AuthInfo = authInfo;
            Body = body;
            ContentType = contentType;
            Headers = headers;
            MaskHeaders = maskHeaders;
            Path = path;
            Port = port;
            RequestMethod = requestMethod;
            UseSsl = useSsl;
            ValidateSsl = validateSsl;
        }
    }
}
