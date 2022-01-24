// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.ServiceDirectory.V1Beta1
{
    public static class GetEndpoint
    {
        /// <summary>
        /// Gets an endpoint.
        /// </summary>
        public static Task<GetEndpointResult> InvokeAsync(GetEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEndpointResult>("google-native:servicedirectory/v1beta1:getEndpoint", args ?? new GetEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an endpoint.
        /// </summary>
        public static Output<GetEndpointResult> Invoke(GetEndpointInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEndpointResult>("google-native:servicedirectory/v1beta1:getEndpoint", args ?? new GetEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEndpointArgs : Pulumi.InvokeArgs
    {
        [Input("endpointId", required: true)]
        public string EndpointId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("namespaceId", required: true)]
        public string NamespaceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceId", required: true)]
        public string ServiceId { get; set; } = null!;

        public GetEndpointArgs()
        {
        }
    }

    public sealed class GetEndpointInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public GetEndpointInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEndpointResult
    {
        /// <summary>
        /// Optional. An IPv4 or IPv6 address. Service Directory rejects bad addresses like: * `8.8.8` * `8.8.8.8:53` * `test:bad:address` * `[::1]` * `[::1]:8080` Limited to 45 characters.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The timestamp when the endpoint was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Metadata for the endpoint. This data can be consumed by service clients. Restrictions: * The entire metadata dictionary may contain up to 512 characters, spread accoss all key-value pairs. Metadata that goes beyond this limit are rejected * Valid metadata keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/). Metadata that fails to meet these requirements are rejected Note: This field is equivalent to the `annotations` field in the v1 API. They have the same syntax and read/write to the same location in Service Directory.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// Immutable. The resource name for the endpoint in the format `projects/*/locations/*/namespaces/*/services/*/endpoints/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The Google Compute Engine network (VPC) of the endpoint in the format `projects//locations/global/networks/*`. The project must be specified by project number (project id is rejected). Incorrectly formatted networks are rejected, but no other validation is performed on this field (ex. network or project existence, reachability, or permissions).
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Optional. Service Directory rejects values outside of `[0, 65535]`.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The timestamp when the endpoint was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetEndpointResult(
            string address,

            string createTime,

            ImmutableDictionary<string, string> metadata,

            string name,

            string network,

            int port,

            string updateTime)
        {
            Address = address;
            CreateTime = createTime;
            Metadata = metadata;
            Name = name;
            Network = network;
            Port = port;
            UpdateTime = updateTime;
        }
    }
}
