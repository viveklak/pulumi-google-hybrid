// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1
{
    public static class GetConnectionProfile
    {
        /// <summary>
        /// Use this method to get details about a connection profile.
        /// </summary>
        public static Task<GetConnectionProfileResult> InvokeAsync(GetConnectionProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionProfileResult>("google-native:datastream/v1alpha1:getConnectionProfile", args ?? new GetConnectionProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Use this method to get details about a connection profile.
        /// </summary>
        public static Output<GetConnectionProfileResult> Invoke(GetConnectionProfileInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConnectionProfileResult>("google-native:datastream/v1alpha1:getConnectionProfile", args ?? new GetConnectionProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionProfileArgs : Pulumi.InvokeArgs
    {
        [Input("connectionProfileId", required: true)]
        public string ConnectionProfileId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetConnectionProfileArgs()
        {
        }
    }

    public sealed class GetConnectionProfileInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("connectionProfileId", required: true)]
        public Input<string> ConnectionProfileId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetConnectionProfileInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionProfileResult
    {
        /// <summary>
        /// The create time of the resource.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Forward SSH tunnel connectivity.
        /// </summary>
        public readonly Outputs.ForwardSshTunnelConnectivityResponse ForwardSshConnectivity;
        /// <summary>
        /// Cloud Storage ConnectionProfile configuration.
        /// </summary>
        public readonly Outputs.GcsProfileResponse GcsProfile;
        /// <summary>
        /// Labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// MySQL ConnectionProfile configuration.
        /// </summary>
        public readonly Outputs.MysqlProfileResponse MysqlProfile;
        /// <summary>
        /// The resource's name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// No connectivity option chosen.
        /// </summary>
        public readonly Outputs.NoConnectivitySettingsResponse NoConnectivity;
        /// <summary>
        /// Oracle ConnectionProfile configuration.
        /// </summary>
        public readonly Outputs.OracleProfileResponse OracleProfile;
        /// <summary>
        /// Private connectivity.
        /// </summary>
        public readonly Outputs.PrivateConnectivityResponse PrivateConnectivity;
        /// <summary>
        /// Static Service IP connectivity.
        /// </summary>
        public readonly Outputs.StaticServiceIpConnectivityResponse StaticServiceIpConnectivity;
        /// <summary>
        /// The update time of the resource.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetConnectionProfileResult(
            string createTime,

            string displayName,

            Outputs.ForwardSshTunnelConnectivityResponse forwardSshConnectivity,

            Outputs.GcsProfileResponse gcsProfile,

            ImmutableDictionary<string, string> labels,

            Outputs.MysqlProfileResponse mysqlProfile,

            string name,

            Outputs.NoConnectivitySettingsResponse noConnectivity,

            Outputs.OracleProfileResponse oracleProfile,

            Outputs.PrivateConnectivityResponse privateConnectivity,

            Outputs.StaticServiceIpConnectivityResponse staticServiceIpConnectivity,

            string updateTime)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            ForwardSshConnectivity = forwardSshConnectivity;
            GcsProfile = gcsProfile;
            Labels = labels;
            MysqlProfile = mysqlProfile;
            Name = name;
            NoConnectivity = noConnectivity;
            OracleProfile = oracleProfile;
            PrivateConnectivity = privateConnectivity;
            StaticServiceIpConnectivity = staticServiceIpConnectivity;
            UpdateTime = updateTime;
        }
    }
}
