// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIGateway.V1Beta
{
    public static class GetGateway
    {
        /// <summary>
        /// Gets details of a single Gateway.
        /// </summary>
        public static Task<GetGatewayResult> InvokeAsync(GetGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewayResult>("google-native:apigateway/v1beta:getGateway", args ?? new GetGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Gateway.
        /// </summary>
        public static Output<GetGatewayResult> Invoke(GetGatewayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGatewayResult>("google-native:apigateway/v1beta:getGateway", args ?? new GetGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("gatewayId", required: true)]
        public string GatewayId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGatewayArgs()
        {
        }
    }

    public sealed class GetGatewayInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGatewayInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewayResult
    {
        /// <summary>
        /// Resource name of the API Config for this Gateway. Format: projects/{project}/locations/global/apis/{api}/configs/{apiConfig}
        /// </summary>
        public readonly string ApiConfig;
        /// <summary>
        /// Created time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The default API Gateway host name of the form `{gateway_id}-{hash}.{region_code}.gateway.dev`.
        /// </summary>
        public readonly string DefaultHostname;
        /// <summary>
        /// Optional. Display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Resource name of the Gateway. Format: projects/{project}/locations/{location}/gateways/{gateway}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current state of the Gateway.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Updated time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGatewayResult(
            string apiConfig,

            string createTime,

            string defaultHostname,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string name,

            string state,

            string updateTime)
        {
            ApiConfig = apiConfig;
            CreateTime = createTime;
            DefaultHostname = defaultHostname;
            DisplayName = displayName;
            Labels = labels;
            Name = name;
            State = state;
            UpdateTime = updateTime;
        }
    }
}
