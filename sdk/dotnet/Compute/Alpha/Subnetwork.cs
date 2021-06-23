// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates a subnetwork in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:Subnetwork")]
    public partial class Subnetwork : Pulumi.CustomResource
    {
        /// <summary>
        /// Can only be specified if VPC flow logging for this subnetwork is enabled. Sets the aggregation interval for collecting flow logs. Increasing the interval time reduces the amount of generated flow logs for long-lasting connections. Default is an interval of 5 seconds per connection. Valid values: INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN.
        /// </summary>
        [Output("aggregationInterval")]
        public Output<string> AggregationInterval { get; private set; } = null!;

        /// <summary>
        /// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length. Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature. The default value is false and applies to all existing subnetworks and automatically created subnetworks. This field cannot be set to true at resource creation time.
        /// </summary>
        [Output("allowSubnetCidrRoutesOverlap")]
        public Output<bool> AllowSubnetCidrRoutesOverlap { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        /// </summary>
        [Output("enableFlowLogs")]
        public Output<bool> EnableFlowLogs { get; private set; } = null!;

        /// <summary>
        /// Enables Layer2 communication on the subnetwork.
        /// </summary>
        [Output("enableL2")]
        public Output<bool> EnableL2 { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The range of external IPv6 addresses that are owned by this subnetwork.
        /// </summary>
        [Output("externalIpv6Prefix")]
        public Output<string> ExternalIpv6Prefix { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// Can only be specified if VPC flow logging for this subnetwork is enabled. The value of the field must be in [0, 1]. Set the sampling rate of VPC flow logs within the subnetwork where 1.0 means all collected logs are reported and 0.0 means no logs are reported. Default is 0.5, which means half of all collected logs are reported.
        /// </summary>
        [Output("flowSampling")]
        public Output<double> FlowSampling { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The gateway address for default routes to reach destination addresses outside this subnetwork.
        /// </summary>
        [Output("gatewayAddress")]
        public Output<string> GatewayAddress { get; private set; } = null!;

        /// <summary>
        /// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
        /// </summary>
        [Output("ipCidrRange")]
        public Output<string> IpCidrRange { get; private set; } = null!;

        /// <summary>
        /// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct path.
        /// </summary>
        [Output("ipv6AccessType")]
        public Output<string> Ipv6AccessType { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The range of internal IPv6 addresses that are owned by this subnetwork.
        /// </summary>
        [Output("ipv6CidrRange")]
        public Output<string> Ipv6CidrRange { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#subnetwork for Subnetwork resources.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
        /// </summary>
        [Output("logConfig")]
        public Output<Outputs.SubnetworkLogConfigResponse> LogConfig { get; private set; } = null!;

        /// <summary>
        /// Can only be specified if VPC flow logging for this subnetwork is enabled. Configures whether metadata fields should be added to the reported VPC flow logs. Options are INCLUDE_ALL_METADATA, EXCLUDE_ALL_METADATA, and CUSTOM_METADATA. Default is EXCLUDE_ALL_METADATA.
        /// </summary>
        [Output("metadata")]
        public Output<string> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
        /// </summary>
        [Output("privateIpGoogleAccess")]
        public Output<bool> PrivateIpGoogleAccess { get; private set; } = null!;

        /// <summary>
        /// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority. This field can be both set at resource creation time and updated using patch.
        /// </summary>
        [Output("privateIpv6GoogleAccess")]
        public Output<string> PrivateIpv6GoogleAccess { get; private set; } = null!;

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        /// </summary>
        [Output("purpose")]
        public Output<string> Purpose { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URL of the reserved internal range.
        /// </summary>
        [Output("reservedInternalRange")]
        public Output<string> ReservedInternalRange { get; private set; } = null!;

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
        /// </summary>
        [Output("secondaryIpRanges")]
        public Output<ImmutableArray<Outputs.SubnetworkSecondaryRangeResponse>> SecondaryIpRanges { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will be used. This field can be both set at resource creation time and updated using patch.
        /// </summary>
        [Output("stackType")]
        public Output<string> StackType { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A repeated field indicating the VLAN IDs supported on this subnetwork. During Subnet creation, specifying vlan is valid only if enable_l2 is true. During Subnet Update, specifying vlan is allowed only for l2 enabled subnets. Restricted to only one VLAN.
        /// </summary>
        [Output("vlans")]
        public Output<ImmutableArray<int>> Vlans { get; private set; } = null!;


        /// <summary>
        /// Create a Subnetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnetwork(string name, SubnetworkArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:Subnetwork", name, args ?? new SubnetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnetwork(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:Subnetwork", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Subnetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnetwork Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Subnetwork(name, id, options);
        }
    }

    public sealed class SubnetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can only be specified if VPC flow logging for this subnetwork is enabled. Sets the aggregation interval for collecting flow logs. Increasing the interval time reduces the amount of generated flow logs for long-lasting connections. Default is an interval of 5 seconds per connection. Valid values: INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN.
        /// </summary>
        [Input("aggregationInterval")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkAggregationInterval>? AggregationInterval { get; set; }

        /// <summary>
        /// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length. Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature. The default value is false and applies to all existing subnetworks and automatically created subnetworks. This field cannot be set to true at resource creation time.
        /// </summary>
        [Input("allowSubnetCidrRoutesOverlap")]
        public Input<bool>? AllowSubnetCidrRoutesOverlap { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        /// </summary>
        [Input("enableFlowLogs")]
        public Input<bool>? EnableFlowLogs { get; set; }

        /// <summary>
        /// Enables Layer2 communication on the subnetwork.
        /// </summary>
        [Input("enableL2")]
        public Input<bool>? EnableL2 { get; set; }

        /// <summary>
        /// [Output Only] The range of external IPv6 addresses that are owned by this subnetwork.
        /// </summary>
        [Input("externalIpv6Prefix")]
        public Input<string>? ExternalIpv6Prefix { get; set; }

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// Can only be specified if VPC flow logging for this subnetwork is enabled. The value of the field must be in [0, 1]. Set the sampling rate of VPC flow logs within the subnetwork where 1.0 means all collected logs are reported and 0.0 means no logs are reported. Default is 0.5, which means half of all collected logs are reported.
        /// </summary>
        [Input("flowSampling")]
        public Input<double>? FlowSampling { get; set; }

        /// <summary>
        /// [Output Only] The gateway address for default routes to reach destination addresses outside this subnetwork.
        /// </summary>
        [Input("gatewayAddress")]
        public Input<string>? GatewayAddress { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        /// <summary>
        /// The access type of IPv6 address this subnet holds. It's immutable and can only be specified during creation or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet cannot enable direct path.
        /// </summary>
        [Input("ipv6AccessType")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkIpv6AccessType>? Ipv6AccessType { get; set; }

        /// <summary>
        /// [Output Only] The range of internal IPv6 addresses that are owned by this subnetwork.
        /// </summary>
        [Input("ipv6CidrRange")]
        public Input<string>? Ipv6CidrRange { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#subnetwork for Subnetwork resources.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
        /// </summary>
        [Input("logConfig")]
        public Input<Inputs.SubnetworkLogConfigArgs>? LogConfig { get; set; }

        /// <summary>
        /// Can only be specified if VPC flow logging for this subnetwork is enabled. Configures whether metadata fields should be added to the reported VPC flow logs. Options are INCLUDE_ALL_METADATA, EXCLUDE_ALL_METADATA, and CUSTOM_METADATA. Default is EXCLUDE_ALL_METADATA.
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkMetadata>? Metadata { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
        /// </summary>
        [Input("privateIpGoogleAccess")]
        public Input<bool>? PrivateIpGoogleAccess { get; set; }

        /// <summary>
        /// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority. This field can be both set at resource creation time and updated using patch.
        /// </summary>
        [Input("privateIpv6GoogleAccess")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkPrivateIpv6GoogleAccess>? PrivateIpv6GoogleAccess { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
        /// </summary>
        [Input("purpose")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkPurpose>? Purpose { get; set; }

        /// <summary>
        /// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The URL of the reserved internal range.
        /// </summary>
        [Input("reservedInternalRange")]
        public Input<string>? ReservedInternalRange { get; set; }

        /// <summary>
        /// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
        /// </summary>
        [Input("role")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkRole>? Role { get; set; }

        [Input("secondaryIpRanges")]
        private InputList<Inputs.SubnetworkSecondaryRangeArgs>? _secondaryIpRanges;

        /// <summary>
        /// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
        /// </summary>
        public InputList<Inputs.SubnetworkSecondaryRangeArgs> SecondaryIpRanges
        {
            get => _secondaryIpRanges ?? (_secondaryIpRanges = new InputList<Inputs.SubnetworkSecondaryRangeArgs>());
            set => _secondaryIpRanges = value;
        }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for this resource with the resource id.
        /// </summary>
        [Input("selfLinkWithId")]
        public Input<string>? SelfLinkWithId { get; set; }

        /// <summary>
        /// The stack type for this subnet to identify whether the IPv6 feature is enabled or not. If not specified IPV4_ONLY will be used. This field can be both set at resource creation time and updated using patch.
        /// </summary>
        [Input("stackType")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkStackType>? StackType { get; set; }

        /// <summary>
        /// [Output Only] The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.SubnetworkState>? State { get; set; }

        [Input("vlans")]
        private InputList<int>? _vlans;

        /// <summary>
        /// A repeated field indicating the VLAN IDs supported on this subnetwork. During Subnet creation, specifying vlan is valid only if enable_l2 is true. During Subnet Update, specifying vlan is allowed only for l2 enabled subnets. Restricted to only one VLAN.
        /// </summary>
        public InputList<int> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<int>());
            set => _vlans = value;
        }

        public SubnetworkArgs()
        {
        }
    }
}
