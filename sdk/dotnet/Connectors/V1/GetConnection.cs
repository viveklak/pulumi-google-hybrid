// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Connectors.V1
{
    public static class GetConnection
    {
        /// <summary>
        /// Gets details of a single Connection.
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("google-native:connectors/v1:getConnection", args ?? new GetConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Connection.
        /// </summary>
        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("google-native:connectors/v1:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionArgs : Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public string ConnectionId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        public GetConnectionArgs()
        {
        }
    }

    public sealed class GetConnectionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("connectionId", required: true)]
        public Input<string> ConnectionId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetConnectionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        /// <summary>
        /// Optional. Configuration for establishing the connection's authentication with an external system.
        /// </summary>
        public readonly Outputs.AuthConfigResponse AuthConfig;
        /// <summary>
        /// Optional. Configuration for configuring the connection with an external system.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigVariableResponse> ConfigVariables;
        /// <summary>
        /// Connector version on which the connection is created. The format is: projects/*/locations/global/providers/*/connectors/*/versions/*
        /// </summary>
        public readonly string ConnectorVersion;
        /// <summary>
        /// Created time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Description of the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Outbound domains/hosts needs to be allowlisted.
        /// </summary>
        public readonly ImmutableArray<string> EgressBackends;
        /// <summary>
        /// GCR location where the envoy image is stored. formatted like: gcr.io/{bucketName}/{imageName}
        /// </summary>
        public readonly string EnvoyImageLocation;
        /// <summary>
        /// GCR location where the runtime image is stored. formatted like: gcr.io/{bucketName}/{imageName}
        /// </summary>
        public readonly string ImageLocation;
        /// <summary>
        /// Optional. Inactive indicates the connection is active to use or not.
        /// </summary>
        public readonly bool Inactive;
        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Optional. Configuration that indicates whether or not the Connection can be edited.
        /// </summary>
        public readonly Outputs.LockConfigResponse LockConfig;
        /// <summary>
        /// Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. Service account needed for runtime plane to access GCP resources.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g. "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
        /// </summary>
        public readonly string ServiceDirectory;
        /// <summary>
        /// Current status of the connection.
        /// </summary>
        public readonly Outputs.ConnectionStatusResponse Status;
        /// <summary>
        /// Updated time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetConnectionResult(
            Outputs.AuthConfigResponse authConfig,

            ImmutableArray<Outputs.ConfigVariableResponse> configVariables,

            string connectorVersion,

            string createTime,

            string description,

            ImmutableArray<string> egressBackends,

            string envoyImageLocation,

            string imageLocation,

            bool inactive,

            ImmutableDictionary<string, string> labels,

            Outputs.LockConfigResponse lockConfig,

            string name,

            string serviceAccount,

            string serviceDirectory,

            Outputs.ConnectionStatusResponse status,

            string updateTime)
        {
            AuthConfig = authConfig;
            ConfigVariables = configVariables;
            ConnectorVersion = connectorVersion;
            CreateTime = createTime;
            Description = description;
            EgressBackends = egressBackends;
            EnvoyImageLocation = envoyImageLocation;
            ImageLocation = imageLocation;
            Inactive = inactive;
            Labels = labels;
            LockConfig = lockConfig;
            Name = name;
            ServiceAccount = serviceAccount;
            ServiceDirectory = serviceDirectory;
            Status = status;
            UpdateTime = updateTime;
        }
    }
}
