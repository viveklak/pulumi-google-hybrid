// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a TargetHttpProxy resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:TargetHttpProxy")]
    public partial class TargetHttpProxy : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// 
        /// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
        /// 
        /// The default is false.
        /// </summary>
        [Output("proxyBind")]
        public Output<bool> ProxyBind { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
        /// </summary>
        [Output("urlMap")]
        public Output<string> UrlMap { get; private set; } = null!;


        /// <summary>
        /// Create a TargetHttpProxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetHttpProxy(string name, TargetHttpProxyArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:TargetHttpProxy", name, args ?? new TargetHttpProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetHttpProxy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:TargetHttpProxy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing TargetHttpProxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetHttpProxy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TargetHttpProxy(name, id, options);
        }
    }

    public sealed class TargetHttpProxyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// 
        /// When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them.
        /// 
        /// The default is false.
        /// </summary>
        [Input("proxyBind")]
        public Input<bool>? ProxyBind { get; set; }

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
        /// </summary>
        [Input("urlMap")]
        public Input<string>? UrlMap { get; set; }

        public TargetHttpProxyArgs()
        {
        }
    }
}
