// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IDS.V1
{
    /// <summary>
    /// Creates a new Endpoint in a given project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:ids/v1:Endpoint")]
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The create time timestamp.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// User-provided description of the endpoint
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The fully qualified URL of the endpoint's ILB Forwarding Rule.
        /// </summary>
        [Output("endpointForwardingRule")]
        public Output<string> EndpointForwardingRule { get; private set; } = null!;

        /// <summary>
        /// The IP address of the IDS Endpoint's ILB.
        /// </summary>
        [Output("endpointIp")]
        public Output<string> EndpointIp { get; private set; } = null!;

        /// <summary>
        /// The labels of the endpoint.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The fully qualified URL of the network to which the IDS Endpoint is attached.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Lowest threat severity that this endpoint will alert on.
        /// </summary>
        [Output("severity")]
        public Output<string> Severity { get; private set; } = null!;

        /// <summary>
        /// Current state of the endpoint.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Whether the endpoint should report traffic logs in addition to threat logs.
        /// </summary>
        [Output("trafficLogs")]
        public Output<bool> TrafficLogs { get; private set; } = null!;

        /// <summary>
        /// The update time timestamp.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("google-native:ids/v1:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:ids/v1:Endpoint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User-provided description of the endpoint
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. The endpoint identifier. This will be part of the endpoint's resource name. This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, or hyphens, and cannot end with a hyphen. Values that do not match this pattern will trigger an INVALID_ARGUMENT error.
        /// </summary>
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels of the endpoint.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The fully qualified URL of the network to which the IDS Endpoint is attached.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Lowest threat severity that this endpoint will alert on.
        /// </summary>
        [Input("severity", required: true)]
        public Input<Pulumi.GoogleNative.IDS.V1.EndpointSeverity> Severity { get; set; } = null!;

        /// <summary>
        /// Whether the endpoint should report traffic logs in addition to threat logs.
        /// </summary>
        [Input("trafficLogs")]
        public Input<bool>? TrafficLogs { get; set; }

        public EndpointArgs()
        {
        }
    }
}
