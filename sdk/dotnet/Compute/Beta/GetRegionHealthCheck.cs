// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetRegionHealthCheck
    {
        /// <summary>
        /// Returns the specified HealthCheck resource. Gets a list of available health checks by making a list() request.
        /// </summary>
        public static Task<GetRegionHealthCheckResult> InvokeAsync(GetRegionHealthCheckArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionHealthCheckResult>("google-native:compute/beta:getRegionHealthCheck", args ?? new GetRegionHealthCheckArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified HealthCheck resource. Gets a list of available health checks by making a list() request.
        /// </summary>
        public static Output<GetRegionHealthCheckResult> Invoke(GetRegionHealthCheckInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionHealthCheckResult>("google-native:compute/beta:getRegionHealthCheck", args ?? new GetRegionHealthCheckInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionHealthCheckArgs : Pulumi.InvokeArgs
    {
        [Input("healthCheck", required: true)]
        public string HealthCheck { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetRegionHealthCheckArgs()
        {
        }
    }

    public sealed class GetRegionHealthCheckInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("healthCheck", required: true)]
        public Input<string> HealthCheck { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetRegionHealthCheckInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionHealthCheckResult
    {
        /// <summary>
        /// How often (in seconds) to send a health check. The default value is 5 seconds.
        /// </summary>
        public readonly int CheckIntervalSec;
        /// <summary>
        /// Creation timestamp in 3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        public readonly Outputs.GRPCHealthCheckResponse GrpcHealthCheck;
        /// <summary>
        /// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
        /// </summary>
        public readonly int HealthyThreshold;
        public readonly Outputs.HTTP2HealthCheckResponse Http2HealthCheck;
        public readonly Outputs.HTTPHealthCheckResponse HttpHealthCheck;
        public readonly Outputs.HTTPSHealthCheckResponse HttpsHealthCheck;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Configure logging on this health check.
        /// </summary>
        public readonly Outputs.HealthCheckLogConfigResponse LogConfig;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Region where the health check resides. Not applicable to global health checks.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        public readonly Outputs.SSLHealthCheckResponse SslHealthCheck;
        public readonly Outputs.TCPHealthCheckResponse TcpHealthCheck;
        /// <summary>
        /// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
        /// </summary>
        public readonly int TimeoutSec;
        /// <summary>
        /// Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
        /// </summary>
        public readonly int UnhealthyThreshold;

        [OutputConstructor]
        private GetRegionHealthCheckResult(
            int checkIntervalSec,

            string creationTimestamp,

            string description,

            Outputs.GRPCHealthCheckResponse grpcHealthCheck,

            int healthyThreshold,

            Outputs.HTTP2HealthCheckResponse http2HealthCheck,

            Outputs.HTTPHealthCheckResponse httpHealthCheck,

            Outputs.HTTPSHealthCheckResponse httpsHealthCheck,

            string kind,

            Outputs.HealthCheckLogConfigResponse logConfig,

            string name,

            string region,

            string selfLink,

            Outputs.SSLHealthCheckResponse sslHealthCheck,

            Outputs.TCPHealthCheckResponse tcpHealthCheck,

            int timeoutSec,

            string type,

            int unhealthyThreshold)
        {
            CheckIntervalSec = checkIntervalSec;
            CreationTimestamp = creationTimestamp;
            Description = description;
            GrpcHealthCheck = grpcHealthCheck;
            HealthyThreshold = healthyThreshold;
            Http2HealthCheck = http2HealthCheck;
            HttpHealthCheck = httpHealthCheck;
            HttpsHealthCheck = httpsHealthCheck;
            Kind = kind;
            LogConfig = logConfig;
            Name = name;
            Region = region;
            SelfLink = selfLink;
            SslHealthCheck = sslHealthCheck;
            TcpHealthCheck = tcpHealthCheck;
            TimeoutSec = timeoutSec;
            Type = type;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
