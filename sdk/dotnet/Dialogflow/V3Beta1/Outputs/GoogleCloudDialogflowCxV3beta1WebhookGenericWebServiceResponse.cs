// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    /// <summary>
    /// Represents configuration for a generic web service.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1WebhookGenericWebServiceResponse
    {
        /// <summary>
        /// Optional. Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification. This overrides the default SSL trust store. If this is empty or unspecified, Dialogflow will use Google's default trust store to verify certificates. N.B. Make sure the HTTPS server certificates are signed with "subject alt name". For instance a certificate can be self-signed using the following command, ``` openssl x509 -req -days 200 -in example.com.csr \ -signkey example.com.key \ -out example.com.crt \ -extfile &lt;(printf "\nsubjectAltName='DNS:www.example.com'") ```
        /// </summary>
        public readonly ImmutableArray<string> AllowedCaCerts;
        /// <summary>
        /// The password for HTTP Basic authentication.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The HTTP request headers to send together with webhook requests.
        /// </summary>
        public readonly ImmutableDictionary<string, string> RequestHeaders;
        /// <summary>
        /// The webhook URI for receiving POST requests. It must use https protocol.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// The user name for HTTP Basic authentication.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1WebhookGenericWebServiceResponse(
            ImmutableArray<string> allowedCaCerts,

            string password,

            ImmutableDictionary<string, string> requestHeaders,

            string uri,

            string username)
        {
            AllowedCaCerts = allowedCaCerts;
            Password = password;
            RequestHeaders = requestHeaders;
            Uri = uri;
            Username = username;
        }
    }
}
