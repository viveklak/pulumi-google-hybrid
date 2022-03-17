// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1
{
    /// <summary>
    /// Use this method to create a connection profile in a project and location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:datastream/v1alpha1:ConnectionProfile")]
    public partial class ConnectionProfile : Pulumi.CustomResource
    {
        /// <summary>
        /// The create time of the resource.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Forward SSH tunnel connectivity.
        /// </summary>
        [Output("forwardSshConnectivity")]
        public Output<Outputs.ForwardSshTunnelConnectivityResponse> ForwardSshConnectivity { get; private set; } = null!;

        /// <summary>
        /// Cloud Storage ConnectionProfile configuration.
        /// </summary>
        [Output("gcsProfile")]
        public Output<Outputs.GcsProfileResponse> GcsProfile { get; private set; } = null!;

        /// <summary>
        /// Labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// MySQL ConnectionProfile configuration.
        /// </summary>
        [Output("mysqlProfile")]
        public Output<Outputs.MysqlProfileResponse> MysqlProfile { get; private set; } = null!;

        /// <summary>
        /// The resource's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// No connectivity option chosen.
        /// </summary>
        [Output("noConnectivity")]
        public Output<Outputs.NoConnectivitySettingsResponse> NoConnectivity { get; private set; } = null!;

        /// <summary>
        /// Oracle ConnectionProfile configuration.
        /// </summary>
        [Output("oracleProfile")]
        public Output<Outputs.OracleProfileResponse> OracleProfile { get; private set; } = null!;

        /// <summary>
        /// Private connectivity.
        /// </summary>
        [Output("privateConnectivity")]
        public Output<Outputs.PrivateConnectivityResponse> PrivateConnectivity { get; private set; } = null!;

        /// <summary>
        /// Static Service IP connectivity.
        /// </summary>
        [Output("staticServiceIpConnectivity")]
        public Output<Outputs.StaticServiceIpConnectivityResponse> StaticServiceIpConnectivity { get; private set; } = null!;

        /// <summary>
        /// The update time of the resource.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectionProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectionProfile(string name, ConnectionProfileArgs args, CustomResourceOptions? options = null)
            : base("google-native:datastream/v1alpha1:ConnectionProfile", name, args ?? new ConnectionProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectionProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datastream/v1alpha1:ConnectionProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectionProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectionProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConnectionProfile(name, id, options);
        }
    }

    public sealed class ConnectionProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The connection profile identifier.
        /// </summary>
        [Input("connectionProfileId", required: true)]
        public Input<string> ConnectionProfileId { get; set; } = null!;

        /// <summary>
        /// Display name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Forward SSH tunnel connectivity.
        /// </summary>
        [Input("forwardSshConnectivity")]
        public Input<Inputs.ForwardSshTunnelConnectivityArgs>? ForwardSshConnectivity { get; set; }

        /// <summary>
        /// Cloud Storage ConnectionProfile configuration.
        /// </summary>
        [Input("gcsProfile")]
        public Input<Inputs.GcsProfileArgs>? GcsProfile { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// MySQL ConnectionProfile configuration.
        /// </summary>
        [Input("mysqlProfile")]
        public Input<Inputs.MysqlProfileArgs>? MysqlProfile { get; set; }

        /// <summary>
        /// No connectivity option chosen.
        /// </summary>
        [Input("noConnectivity")]
        public Input<Inputs.NoConnectivitySettingsArgs>? NoConnectivity { get; set; }

        /// <summary>
        /// Oracle ConnectionProfile configuration.
        /// </summary>
        [Input("oracleProfile")]
        public Input<Inputs.OracleProfileArgs>? OracleProfile { get; set; }

        /// <summary>
        /// Private connectivity.
        /// </summary>
        [Input("privateConnectivity")]
        public Input<Inputs.PrivateConnectivityArgs>? PrivateConnectivity { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Static Service IP connectivity.
        /// </summary>
        [Input("staticServiceIpConnectivity")]
        public Input<Inputs.StaticServiceIpConnectivityArgs>? StaticServiceIpConnectivity { get; set; }

        public ConnectionProfileArgs()
        {
        }
    }
}
