// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetExternalVpnGateway
    {
        /// <summary>
        /// Returns the specified externalVpnGateway. Get a list of available externalVpnGateways by making a list() request.
        /// </summary>
        public static Task<GetExternalVpnGatewayResult> InvokeAsync(GetExternalVpnGatewayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExternalVpnGatewayResult>("google-native:compute/alpha:getExternalVpnGateway", args ?? new GetExternalVpnGatewayArgs(), options.WithVersion());
    }


    public sealed class GetExternalVpnGatewayArgs : Pulumi.InvokeArgs
    {
        [Input("externalVpnGateway", required: true)]
        public string ExternalVpnGateway { get; set; } = null!;

        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        public GetExternalVpnGatewayArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetExternalVpnGatewayResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// List of interfaces for this external VPN gateway.
        /// </summary>
        public readonly ImmutableArray<Outputs.ExternalVpnGatewayInterfaceResponse> Interfaces;
        /// <summary>
        /// Type of the resource. Always compute#externalVpnGateway for externalVpnGateways.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this ExternalVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve an ExternalVpnGateway.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Indicates the user-supplied redundancy type of this external VPN gateway.
        /// </summary>
        public readonly string RedundancyType;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;

        [OutputConstructor]
        private GetExternalVpnGatewayResult(
            string creationTimestamp,

            string description,

            ImmutableArray<Outputs.ExternalVpnGatewayInterfaceResponse> interfaces,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            string name,

            string redundancyType,

            string selfLink)
        {
            CreationTimestamp = creationTimestamp;
            Description = description;
            Interfaces = interfaces;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            Name = name;
            RedundancyType = redundancyType;
            SelfLink = selfLink;
        }
    }
}
