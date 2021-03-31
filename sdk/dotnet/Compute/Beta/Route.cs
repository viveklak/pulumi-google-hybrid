// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Compute.Beta
{
    /// <summary>
    /// Creates a Route resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:compute/beta:Route")]
    public partial class Route : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:compute/beta:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:compute/beta:Route", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Route(name, id, options);
        }
    }

    public sealed class RouteArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
        /// </summary>
        [Input("destRange")]
        public Input<string>? DestRange { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// [Output Only] Type of this resource. Always compute#routes for Route resources.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Fully-qualified URL of the network that this route applies to.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL:  projects/project/global/gateways/default-internet-gateway
        /// </summary>
        [Input("nextHopGateway")]
        public Input<string>? NextHopGateway { get; set; }

        /// <summary>
        /// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs:  
        /// - 10.128.0.56 
        /// - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule 
        /// - regions/region/forwardingRules/forwardingRule
        /// </summary>
        [Input("nextHopIlb")]
        public Input<string>? NextHopIlb { get; set; }

        /// <summary>
        /// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:
        /// https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
        /// </summary>
        [Input("nextHopInstance")]
        public Input<string>? NextHopInstance { get; set; }

        /// <summary>
        /// [Output Only] The URL to an InterconnectAttachment which is the next hop for the route. This field will only be populated for the dynamic routes generated by Cloud Router with a linked interconnectAttachment.
        /// </summary>
        [Input("nextHopInterconnectAttachment")]
        public Input<string>? NextHopInterconnectAttachment { get; set; }

        /// <summary>
        /// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
        /// </summary>
        [Input("nextHopIp")]
        public Input<string>? NextHopIp { get; set; }

        /// <summary>
        /// The URL of the local network if it should handle matching packets.
        /// </summary>
        [Input("nextHopNetwork")]
        public Input<string>? NextHopNetwork { get; set; }

        /// <summary>
        /// [Output Only] The network peering name that should handle matching packets, which should conform to RFC1035.
        /// </summary>
        [Input("nextHopPeering")]
        public Input<string>? NextHopPeering { get; set; }

        /// <summary>
        /// The URL to a VpnTunnel that should handle matching packets.
        /// </summary>
        [Input("nextHopVpnTunnel")]
        public Input<string>? NextHopVpnTunnel { get; set; }

        /// <summary>
        /// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("route", required: true)]
        public Input<string> Route { get; set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined fully-qualified URL for this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of instance tags to which this route applies.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("warnings")]
        private InputList<ImmutableDictionary<string, string>>? _warnings;

        /// <summary>
        /// [Output Only] If potential misconfigurations are detected for this route, this field will be populated with warning messages.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> Warnings
        {
            get => _warnings ?? (_warnings = new InputList<ImmutableDictionary<string, string>>());
            set => _warnings = value;
        }

        public RouteArgs()
        {
        }
    }
}
