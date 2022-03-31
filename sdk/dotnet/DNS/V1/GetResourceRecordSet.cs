// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1
{
    public static class GetResourceRecordSet
    {
        /// <summary>
        /// Fetches the representation of an existing ResourceRecordSet.
        /// </summary>
        public static Task<GetResourceRecordSetResult> InvokeAsync(GetResourceRecordSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceRecordSetResult>("google-native:dns/v1:getResourceRecordSet", args ?? new GetResourceRecordSetArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches the representation of an existing ResourceRecordSet.
        /// </summary>
        public static Output<GetResourceRecordSetResult> Invoke(GetResourceRecordSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourceRecordSetResult>("google-native:dns/v1:getResourceRecordSet", args ?? new GetResourceRecordSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceRecordSetArgs : Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public string? ClientOperationId { get; set; }

        [Input("managedZone", required: true)]
        public string ManagedZone { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetResourceRecordSetArgs()
        {
        }
    }

    public sealed class GetResourceRecordSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clientOperationId")]
        public Input<string>? ClientOperationId { get; set; }

        [Input("managedZone", required: true)]
        public Input<string> ManagedZone { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetResourceRecordSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourceRecordSetResult
    {
        public readonly string Kind;
        /// <summary>
        /// For example, www.example.com.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
        /// </summary>
        public readonly Outputs.RRSetRoutingPolicyResponse RoutingPolicy;
        /// <summary>
        /// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
        /// </summary>
        public readonly ImmutableArray<string> Rrdatas;
        /// <summary>
        /// As defined in RFC 4034 (section 3.2).
        /// </summary>
        public readonly ImmutableArray<string> SignatureRrdatas;
        /// <summary>
        /// Number of seconds that this ResourceRecordSet can be cached by resolvers.
        /// </summary>
        public readonly int Ttl;
        /// <summary>
        /// The identifier of a supported record type. See the list of Supported DNS record types.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetResourceRecordSetResult(
            string kind,

            string name,

            Outputs.RRSetRoutingPolicyResponse routingPolicy,

            ImmutableArray<string> rrdatas,

            ImmutableArray<string> signatureRrdatas,

            int ttl,

            string type)
        {
            Kind = kind;
            Name = name;
            RoutingPolicy = routingPolicy;
            Rrdatas = rrdatas;
            SignatureRrdatas = signatureRrdatas;
            Ttl = ttl;
            Type = type;
        }
    }
}
