// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetRegionHealthCheckService
    {
        /// <summary>
        /// Returns the specified regional HealthCheckService resource.
        /// </summary>
        public static Task<GetRegionHealthCheckServiceResult> InvokeAsync(GetRegionHealthCheckServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionHealthCheckServiceResult>("google-native:compute/beta:getRegionHealthCheckService", args ?? new GetRegionHealthCheckServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified regional HealthCheckService resource.
        /// </summary>
        public static Output<GetRegionHealthCheckServiceResult> Invoke(GetRegionHealthCheckServiceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionHealthCheckServiceResult>("google-native:compute/beta:getRegionHealthCheckService", args ?? new GetRegionHealthCheckServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionHealthCheckServiceArgs : Pulumi.InvokeArgs
    {
        [Input("healthCheckService", required: true)]
        public string HealthCheckService { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetRegionHealthCheckServiceArgs()
        {
        }
    }

    public sealed class GetRegionHealthCheckServiceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("healthCheckService", required: true)]
        public Input<string> HealthCheckService { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetRegionHealthCheckServiceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionHealthCheckServiceResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a HealthCheckService. An up-to-date fingerprint must be provided in order to patch/update the HealthCheckService; Otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the HealthCheckService.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// A list of URLs to the HealthCheck resources. Must have at least one HealthCheck, and not more than 10. HealthCheck resources must have portSpecification=USE_SERVING_PORT or portSpecification=USE_FIXED_PORT. For regional HealthCheckService, the HealthCheck must be regional and in the same region. For global HealthCheckService, HealthCheck must be global. Mix of regional and global HealthChecks is not supported. Multiple regional HealthChecks must belong to the same region. Regional HealthChecks must belong to the same region as zones of NEGs.
        /// </summary>
        public readonly ImmutableArray<string> HealthChecks;
        /// <summary>
        /// Optional. Policy for how the results from multiple health checks for the same endpoint are aggregated. Defaults to NO_AGGREGATION if unspecified. - NO_AGGREGATION. An EndpointHealth message is returned for each pair in the health check service. - AND. If any health check of an endpoint reports UNHEALTHY, then UNHEALTHY is the HealthState of the endpoint. If all health checks report HEALTHY, the HealthState of the endpoint is HEALTHY. .
        /// </summary>
        public readonly string HealthStatusAggregationPolicy;
        /// <summary>
        /// This field is deprecated. Use health_status_aggregation_policy instead. Policy for how the results from multiple health checks for the same endpoint are aggregated. - NO_AGGREGATION. An EndpointHealth message is returned for each backend in the health check service. - AND. If any backend's health check reports UNHEALTHY, then UNHEALTHY is the HealthState of the entire health check service. If all backend's are healthy, the HealthState of the health check service is HEALTHY. .
        /// </summary>
        public readonly string HealthStatusAggregationStrategy;
        /// <summary>
        /// [Output only] Type of the resource. Always compute#healthCheckServicefor health check services.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of URLs to the NetworkEndpointGroup resources. Must not have more than 100. For regional HealthCheckService, NEGs must be in zones in the region of the HealthCheckService.
        /// </summary>
        public readonly ImmutableArray<string> NetworkEndpointGroups;
        /// <summary>
        /// A list of URLs to the NotificationEndpoint resources. Must not have more than 10. A list of endpoints for receiving notifications of change in health status. For regional HealthCheckService, NotificationEndpoint must be regional and in the same region. For global HealthCheckService, NotificationEndpoint must be global.
        /// </summary>
        public readonly ImmutableArray<string> NotificationEndpoints;
        /// <summary>
        /// URL of the region where the health check service resides. This field is not applicable to global health check services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;

        [OutputConstructor]
        private GetRegionHealthCheckServiceResult(
            string creationTimestamp,

            string description,

            string fingerprint,

            ImmutableArray<string> healthChecks,

            string healthStatusAggregationPolicy,

            string healthStatusAggregationStrategy,

            string kind,

            string name,

            ImmutableArray<string> networkEndpointGroups,

            ImmutableArray<string> notificationEndpoints,

            string region,

            string selfLink)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            HealthChecks = healthChecks;
            HealthStatusAggregationPolicy = healthStatusAggregationPolicy;
            HealthStatusAggregationStrategy = healthStatusAggregationStrategy;
            Kind = kind;
            Name = name;
            NetworkEndpointGroups = networkEndpointGroups;
            NotificationEndpoints = notificationEndpoints;
            Region = region;
            SelfLink = selfLink;
        }
    }
}
