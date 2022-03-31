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
    /// Represents a gRPC setting that describes one gRPC notification endpoint and the retry duration attempting to send notification to this endpoint.
    /// </summary>
    [OutputType]
    public sealed class NotificationEndpointGrpcSettingsResponse
    {
        /// <summary>
        /// Optional. If specified, this field is used to set the authority header by the sender of notifications. See https://tools.ietf.org/html/rfc7540#section-8.1.2.3
        /// </summary>
        public readonly string Authority;
        /// <summary>
        /// Endpoint to which gRPC notifications are sent. This must be a valid gRPCLB DNS name.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// Optional. If specified, this field is used to populate the "name" field in gRPC requests.
        /// </summary>
        public readonly string PayloadName;
        /// <summary>
        /// Optional. This field is used to configure how often to send a full update of all non-healthy backends. If unspecified, full updates are not sent. If specified, must be in the range between 600 seconds to 3600 seconds. Nanos are disallowed.
        /// </summary>
        public readonly Outputs.DurationResponse ResendInterval;
        /// <summary>
        /// How much time (in seconds) is spent attempting notification retries until a successful response is received. Default is 30s. Limit is 20m (1200s). Must be a positive number.
        /// </summary>
        public readonly int RetryDurationSec;

        [OutputConstructor]
        private NotificationEndpointGrpcSettingsResponse(
            string authority,

            string endpoint,

            string payloadName,

            Outputs.DurationResponse resendInterval,

            int retryDurationSec)
        {
            Authority = authority;
            Endpoint = endpoint;
            PayloadName = payloadName;
            ResendInterval = resendInterval;
            RetryDurationSec = retryDurationSec;
        }
    }
}
