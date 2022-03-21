// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1
{
    public static class GetRoute
    {
        /// <summary>
        /// Use this method to get details about a route.
        /// </summary>
        public static Task<GetRouteResult> InvokeAsync(GetRouteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteResult>("google-native:datastream/v1alpha1:getRoute", args ?? new GetRouteArgs(), options.WithDefaults());

        /// <summary>
        /// Use this method to get details about a route.
        /// </summary>
        public static Output<GetRouteResult> Invoke(GetRouteInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRouteResult>("google-native:datastream/v1alpha1:getRoute", args ?? new GetRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("privateConnectionId", required: true)]
        public string PrivateConnectionId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("routeId", required: true)]
        public string RouteId { get; set; } = null!;

        public GetRouteArgs()
        {
        }
    }

    public sealed class GetRouteInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("privateConnectionId", required: true)]
        public Input<string> PrivateConnectionId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("routeId", required: true)]
        public Input<string> RouteId { get; set; } = null!;

        public GetRouteInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRouteResult
    {
        /// <summary>
        /// The create time of the resource.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Destination address for connection
        /// </summary>
        public readonly string DestinationAddress;
        /// <summary>
        /// Destination port for connection
        /// </summary>
        public readonly int DestinationPort;
        /// <summary>
        /// Display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The resource's name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The update time of the resource.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetRouteResult(
            string createTime,

            string destinationAddress,

            int destinationPort,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            DestinationAddress = destinationAddress;
            DestinationPort = destinationPort;
            DisplayName = displayName;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}
