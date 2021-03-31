// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Alpha
{
    /// <summary>
    /// Create a NotificationEndpoint in the specified project in the given region using the parameters that are included in the request.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:compute/alpha:RegionNotificationEndpoint")]
    public partial class RegionNotificationEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a RegionNotificationEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionNotificationEndpoint(string name, RegionNotificationEndpointArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:compute/alpha:RegionNotificationEndpoint", name, args ?? new RegionNotificationEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionNotificationEndpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:compute/alpha:RegionNotificationEndpoint", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RegionNotificationEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionNotificationEndpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionNotificationEndpoint(name, id, options);
        }
    }

    public sealed class RegionNotificationEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
        /// </summary>
        [Input("grpcSettings")]
        public Input<Inputs.NotificationEndpointGrpcSettingsArgs>? GrpcSettings { get; set; }

        /// <summary>
        /// [Output Only] A unique identifier for this resource type. The server generates this identifier.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#notificationEndpoint for notification endpoints.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationEndpoint", required: true)]
        public Input<string> NotificationEndpoint { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the notification endpoint resides. This field applies only to the regional resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public RegionNotificationEndpointArgs()
        {
        }
    }
}
