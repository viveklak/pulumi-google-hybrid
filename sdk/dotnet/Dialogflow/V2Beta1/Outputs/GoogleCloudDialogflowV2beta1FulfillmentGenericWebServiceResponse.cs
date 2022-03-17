// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    /// <summary>
    /// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers. More information could be found at: https://cloud.google.com/dialogflow/docs/fulfillment-configure.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1FulfillmentGenericWebServiceResponse
    {
        /// <summary>
        /// Optional. Indicates if generic web service is created through Cloud Functions integration. Defaults to false. is_cloud_function is deprecated. Cloud functions can be configured by its uri as a regular web service now.
        /// </summary>
        public readonly bool IsCloudFunction;
        /// <summary>
        /// The password for HTTP Basic authentication.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The HTTP request headers to send together with fulfillment requests.
        /// </summary>
        public readonly ImmutableDictionary<string, string> RequestHeaders;
        /// <summary>
        /// The fulfillment URI for receiving POST requests. It must use https protocol.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// The user name for HTTP Basic authentication.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1FulfillmentGenericWebServiceResponse(
            bool isCloudFunction,

            string password,

            ImmutableDictionary<string, string> requestHeaders,

            string uri,

            string username)
        {
            IsCloudFunction = isCloudFunction;
            Password = password;
            RequestHeaders = requestHeaders;
            Uri = uri;
            Username = username;
        }
    }
}
