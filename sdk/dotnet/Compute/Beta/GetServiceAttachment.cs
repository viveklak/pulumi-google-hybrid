// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetServiceAttachment
    {
        /// <summary>
        /// Returns the specified ServiceAttachment resource in the given scope.
        /// </summary>
        public static Task<GetServiceAttachmentResult> InvokeAsync(GetServiceAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceAttachmentResult>("google-native:compute/beta:getServiceAttachment", args ?? new GetServiceAttachmentArgs(), options.WithVersion());
    }


    public sealed class GetServiceAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("serviceAttachment", required: true)]
        public string ServiceAttachment { get; set; } = null!;

        public GetServiceAttachmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceAttachmentResult
    {
        /// <summary>
        /// [Output Only] An array of connections for all the consumers connected to this service attachment.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceAttachmentConnectedEndpointResponse> ConnectedEndpoints;
        /// <summary>
        /// The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
        /// </summary>
        public readonly string ConnectionPreference;
        /// <summary>
        /// Projects that are allowed to connect to this service attachment.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceAttachmentConsumerProjectLimitResponse> ConsumerAcceptLists;
        /// <summary>
        /// [Output Only] An array of forwarding rules for all the consumers connected to this service attachment.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceAttachmentConsumerForwardingRuleResponse> ConsumerForwardingRules;
        /// <summary>
        /// Projects that are not allowed to connect to this service attachment. The project can be specified using its id or number.
        /// </summary>
        public readonly ImmutableArray<string> ConsumerRejectLists;
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// If true, enable the proxy protocol which is for supplying client TCP/IP address data in TCP connections that traverse proxies on their way to destination servers.
        /// </summary>
        public readonly bool EnableProxyProtocol;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ServiceAttachment. An up-to-date fingerprint must be provided in order to patch/update the ServiceAttachment; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the ServiceAttachment.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// [Output Only] Type of the resource. Always compute#serviceAttachment for service attachments.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An array of URLs where each entry is the URL of a subnet provided by the service producer to use for NAT in this service attachment.
        /// </summary>
        public readonly ImmutableArray<string> NatSubnets;
        /// <summary>
        /// The URL of a forwarding rule with loadBalancingScheme INTERNAL* that is serving the endpoint identified by this service attachment.
        /// </summary>
        public readonly string ProducerForwardingRule;
        /// <summary>
        /// [Output Only] An 128-bit global unique ID of the PSC service attachment.
        /// </summary>
        public readonly Outputs.Uint128Response PscServiceAttachmentId;
        /// <summary>
        /// [Output Only] URL of the region where the service attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The URL of a service serving the endpoint identified by this service attachment.
        /// </summary>
        public readonly string TargetService;

        [OutputConstructor]
        private GetServiceAttachmentResult(
            ImmutableArray<Outputs.ServiceAttachmentConnectedEndpointResponse> connectedEndpoints,

            string connectionPreference,

            ImmutableArray<Outputs.ServiceAttachmentConsumerProjectLimitResponse> consumerAcceptLists,

            ImmutableArray<Outputs.ServiceAttachmentConsumerForwardingRuleResponse> consumerForwardingRules,

            ImmutableArray<string> consumerRejectLists,

            string creationTimestamp,

            string description,

            bool enableProxyProtocol,

            string fingerprint,

            string kind,

            string name,

            ImmutableArray<string> natSubnets,

            string producerForwardingRule,

            Outputs.Uint128Response pscServiceAttachmentId,

            string region,

            string selfLink,

            string targetService)
        {
            ConnectedEndpoints = connectedEndpoints;
            ConnectionPreference = connectionPreference;
            ConsumerAcceptLists = consumerAcceptLists;
            ConsumerForwardingRules = consumerForwardingRules;
            ConsumerRejectLists = consumerRejectLists;
            CreationTimestamp = creationTimestamp;
            Description = description;
            EnableProxyProtocol = enableProxyProtocol;
            Fingerprint = fingerprint;
            Kind = kind;
            Name = name;
            NatSubnets = natSubnets;
            ProducerForwardingRule = producerForwardingRule;
            PscServiceAttachmentId = pscServiceAttachmentId;
            Region = region;
            SelfLink = selfLink;
            TargetService = targetService;
        }
    }
}
