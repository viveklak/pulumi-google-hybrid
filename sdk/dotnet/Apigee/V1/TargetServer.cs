// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates a TargetServer in the specified environment.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:TargetServer")]
    public partial class TargetServer : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. A human-readable description of this TargetServer.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Required. The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Optional. Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        /// </summary>
        [Output("isEnabled")]
        public Output<bool> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// Required. The resource id of this target server. Values must match the regular expression 
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Immutable. The protocol used by this TargetServer.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Optional. Specifies TLS configuration info for this TargetServer. The JSON name is `sSLInfo` for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        /// </summary>
        [Output("sSLInfo")]
        public Output<Outputs.GoogleCloudApigeeV1TlsInfoResponse> SSLInfo { get; private set; } = null!;


        /// <summary>
        /// Create a TargetServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetServer(string name, TargetServerArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:TargetServer", name, args ?? new TargetServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetServer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:TargetServer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing TargetServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetServer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TargetServer(name, id, options);
        }
    }

    public sealed class TargetServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. A human-readable description of this TargetServer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// Required. The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Optional. Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// Required. The resource id of this target server. Values must match the regular expression 
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        /// <summary>
        /// Required. The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Immutable. The protocol used by this TargetServer.
        /// </summary>
        [Input("protocol")]
        public Input<Pulumi.GoogleNative.Apigee.V1.TargetServerProtocol>? Protocol { get; set; }

        /// <summary>
        /// Optional. Specifies TLS configuration info for this TargetServer. The JSON name is `sSLInfo` for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        /// </summary>
        [Input("sSLInfo")]
        public Input<Inputs.GoogleCloudApigeeV1TlsInfoArgs>? SSLInfo { get; set; }

        public TargetServerArgs()
        {
        }
    }
}
