// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetRegionNotificationEndpoint
    {
        /// <summary>
        /// Returns the specified NotificationEndpoint resource in the given region.
        /// </summary>
        public static Task<GetRegionNotificationEndpointResult> InvokeAsync(GetRegionNotificationEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionNotificationEndpointResult>("google-native:compute/v1:getRegionNotificationEndpoint", args ?? new GetRegionNotificationEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified NotificationEndpoint resource in the given region.
        /// </summary>
        public static Output<GetRegionNotificationEndpointResult> Invoke(GetRegionNotificationEndpointInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRegionNotificationEndpointResult>("google-native:compute/v1:getRegionNotificationEndpoint", args ?? new GetRegionNotificationEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionNotificationEndpointArgs : Pulumi.InvokeArgs
    {
        [Input("notificationEndpoint", required: true)]
        public string NotificationEndpoint { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        public GetRegionNotificationEndpointArgs()
        {
        }
    }

    public sealed class GetRegionNotificationEndpointInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("notificationEndpoint", required: true)]
        public Input<string> NotificationEndpoint { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public GetRegionNotificationEndpointInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegionNotificationEndpointResult
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
        /// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
        /// </summary>
        public readonly Outputs.NotificationEndpointGrpcSettingsResponse GrpcSettings;
        /// <summary>
        /// Type of the resource. Always compute#notificationEndpoint for notification endpoints.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// URL of the region where the notification endpoint resides. This field applies only to the regional resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;

        [OutputConstructor]
        private GetRegionNotificationEndpointResult(
            string creationTimestamp,

            string description,

            Outputs.NotificationEndpointGrpcSettingsResponse grpcSettings,

            string kind,

            string name,

            string region,

            string selfLink)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            GrpcSettings = grpcSettings;
            Kind = kind;
            Name = name;
            Region = region;
            SelfLink = selfLink;
        }
    }
}
