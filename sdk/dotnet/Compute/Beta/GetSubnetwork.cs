// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetSubnetwork
    {
        /// <summary>
        /// Returns the specified subnetwork. Gets a list of available subnetworks list() request.
        /// </summary>
        public static Task<GetSubnetworkResult> InvokeAsync(GetSubnetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetworkResult>("google-native:compute/beta:getSubnetwork", args ?? new GetSubnetworkArgs(), options.WithVersion());
    }


    public sealed class GetSubnetworkArgs : Pulumi.InvokeArgs
    {
        [Input("project", required: true)]
        public string Project { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("subnetwork", required: true)]
        public string Subnetwork { get; set; } = null!;

        public GetSubnetworkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubnetworkResult
    {
        /// <summary>
        /// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length. Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature. The default value is false and applies to all existing subnetworks and automatically created subnetworks. This field cannot be set to true at resource creation time.
        /// </summary>
        public readonly bool AllowSubnetCidrRoutesOverlap;
        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        /// </summary>
        public readonly bool EnableFlowLogs;
        /// <summary>
        /// [Output Only] The range of external IPv6 addresses that are owned by this subnetwork.
        /// </summary>
        public readonly string ExternalIpv6Prefix;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// [Output Only] The gateway address for default routes to reach destination addresses outside this subnetwork.
        /// </summary>
        public readonly string GatewayAddress;
        /// <summary>
        /// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct path.
        /// </summary>
        public readonly string Ipv6AccessType;
        /// <summary>
        /// [Output Only] The range of internal IPv6 addresses that are owned by this subnetwork.
        /// </summary>
        public readonly string Ipv6CidrRange;
        /// <summary>
        /// [Output Only] Type of the resource. Always compute#subnetwork for Subnetwork resources.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
        /// </summary>
        public readonly Outputs.SubnetworkLogConfigResponse LogConfig;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
        /// </summary>
        public readonly bool PrivateIpGoogleAccess;
        /// <summary>
        /// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority. This field can be both set at resource creation time and updated using patch.
        /// </summary>
        public readonly string PrivateIpv6GoogleAccess;
        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        /// </summary>
        public readonly string Purpose;
        /// <summary>
        /// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubnetworkSecondaryRangeResponse> SecondaryIpRanges;
        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will be used. This field can be both set at resource creation time and updated using patch.
        /// </summary>
        public readonly string StackType;
        /// <summary>
        /// [Output Only] The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetSubnetworkResult(
            bool allowSubnetCidrRoutesOverlap,

            string creationTimestamp,

            string description,

            bool enableFlowLogs,

            string externalIpv6Prefix,

            string fingerprint,

            string gatewayAddress,

            string ipCidrRange,

            string ipv6AccessType,

            string ipv6CidrRange,

            string kind,

            Outputs.SubnetworkLogConfigResponse logConfig,

            string name,

            string network,

            bool privateIpGoogleAccess,

            string privateIpv6GoogleAccess,

            string purpose,

            string region,

            string role,

            ImmutableArray<Outputs.SubnetworkSecondaryRangeResponse> secondaryIpRanges,

            string selfLink,

            string stackType,

            string state)
        {
            AllowSubnetCidrRoutesOverlap = allowSubnetCidrRoutesOverlap;
            CreationTimestamp = creationTimestamp;
            Description = description;
            EnableFlowLogs = enableFlowLogs;
            ExternalIpv6Prefix = externalIpv6Prefix;
            Fingerprint = fingerprint;
            GatewayAddress = gatewayAddress;
            IpCidrRange = ipCidrRange;
            Ipv6AccessType = ipv6AccessType;
            Ipv6CidrRange = ipv6CidrRange;
            Kind = kind;
            LogConfig = logConfig;
            Name = name;
            Network = network;
            PrivateIpGoogleAccess = privateIpGoogleAccess;
            PrivateIpv6GoogleAccess = privateIpv6GoogleAccess;
            Purpose = purpose;
            Region = region;
            Role = role;
            SecondaryIpRanges = secondaryIpRanges;
            SelfLink = selfLink;
            StackType = stackType;
            State = state;
        }
    }
}
