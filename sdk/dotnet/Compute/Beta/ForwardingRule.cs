// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    /// <summary>
    /// Creates a ForwardingRule resource in the specified project and region using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/beta:ForwardingRule")]
    public partial class ForwardingRule : Pulumi.CustomResource
    {
        /// <summary>
        /// IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule.
        /// 
        /// If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address:
        /// 
        /// * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name * Partial URL or by name, as in:  
        /// - projects/project_id/regions/region/addresses/address-name 
        /// - regions/region/addresses/address-name 
        /// - global/addresses/address-name 
        /// - address-name  
        /// 
        /// The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
        /// 
        /// Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true.
        /// 
        /// For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
        /// </summary>
        [Output("IPAddress")]
        public Output<string> IPAddress { get; private set; } = null!;

        /// <summary>
        /// The IP protocol to which this rule applies.
        /// 
        /// For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP.
        /// 
        /// The valid IP protocols are different for different load balancing products:  
        /// - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid. 
        /// - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid.  
        /// - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid. 
        /// - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid. 
        /// - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid.
        /// </summary>
        [Output("IPProtocol")]
        public Output<string> IPProtocol { get; private set; } = null!;

        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. This field cannot be used with port or portRange fields.
        /// 
        /// When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
        /// </summary>
        [Output("allPorts")]
        public Output<bool> AllPorts { get; private set; } = null!;

        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
        /// </summary>
        [Output("allowGlobalAccess")]
        public Output<bool> AllowGlobalAccess { get; private set; } = null!;

        /// <summary>
        /// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
        /// </summary>
        [Output("backendService")]
        public Output<string> BackendService { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
        /// </summary>
        [Output("ipVersion")]
        public Output<string> IpVersion { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
        /// </summary>
        [Output("isMirroringCollector")]
        public Output<bool> IsMirroringCollector { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// Specifies the forwarding rule type.
        /// 
        ///  
        /// - EXTERNAL is used for:  
        /// - Classic Cloud VPN gateways 
        /// - Protocol forwarding to VMs from an external IP address 
        /// - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing    
        /// - INTERNAL is used for:  
        /// - Protocol forwarding to VMs from an internal IP address 
        /// - Internal TCP/UDP Load Balancing   
        /// - INTERNAL_MANAGED is used for:  
        /// - Internal HTTP(S) Load Balancing   
        /// - INTERNAL_SELF_MANAGED is used for:  
        /// - Traffic Director    
        /// 
        /// For more information about forwarding rules, refer to Forwarding rule concepts.
        /// </summary>
        [Output("loadBalancingScheme")]
        public Output<string> LoadBalancingScheme { get; private set; } = null!;

        /// <summary>
        /// Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to Loadbalancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule will not be visible to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match.
        /// metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// </summary>
        [Output("metadataFilters")]
        public Output<ImmutableArray<Outputs.MetadataFilterResponse>> MetadataFilters { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This field is not used for external load balancing.
        /// 
        /// For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
        /// 
        /// For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD.
        /// 
        /// For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM.
        /// 
        /// If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
        /// </summary>
        [Output("networkTier")]
        public Output<string> NetworkTier { get; private set; } = null!;

        /// <summary>
        /// This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP.
        /// 
        /// Packets addressed to ports in the specified range will be forwarded to target or  backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges.
        /// 
        /// Some types of forwarding target have constraints on the acceptable ports:  
        /// - TargetHttpProxy: 80, 8080 
        /// - TargetHttpsProxy: 443 
        /// - TargetGrpcProxy: no constraints 
        /// - TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
        /// - TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
        /// - TargetVpnGateway: 500, 4500
        /// </summary>
        [Output("portRange")]
        public Output<string> PortRange { get; private set; } = null!;

        /// <summary>
        /// The ports field is only supported when the forwarding rule references a backend_service directly. Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing. Only packets addressed to the specified list of ports are forwarded to backends.
        /// 
        /// You can only use one of ports and port_range, or allPorts. The three are mutually exclusive.
        /// 
        /// You can specify a list of up to five ports, which can be non-contiguous.
        /// 
        /// For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports.
        /// 
        /// For more information, see [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<string>> Ports { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The PSC connection id of the PSC Forwarding Rule.
        /// </summary>
        [Output("pscConnectionId")]
        public Output<string> PscConnectionId { get; private set; } = null!;

        /// <summary>
        /// [Output Only] URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
        /// 
        /// It is only supported for Internal TCP/UDP Load Balancing and Internal HTTP(S) Load Balancing.
        /// </summary>
        [Output("serviceDirectoryRegistrations")]
        public Output<ImmutableArray<Outputs.ForwardingRuleServiceDirectoryRegistrationResponse>> ServiceDirectoryRegistrations { get; private set; } = null!;

        /// <summary>
        /// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name.
        /// 
        /// The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// 
        /// This field is only used for internal load balancing.
        /// </summary>
        [Output("serviceLabel")]
        public Output<string> ServiceLabel { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The internal fully qualified service name for this Forwarding Rule.
        /// 
        /// This field is only used for internal load balancing.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// This field is only used for internal load balancing.
        /// 
        /// For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule.
        /// 
        /// If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
        /// </summary>
        [Output("subnetwork")]
        public Output<string> Subnetwork { get; private set; } = null!;

        [Output("target")]
        public Output<string> Target { get; private set; } = null!;


        /// <summary>
        /// Create a ForwardingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ForwardingRule(string name, ForwardingRuleArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:ForwardingRule", name, args ?? new ForwardingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ForwardingRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/beta:ForwardingRule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ForwardingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ForwardingRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ForwardingRule(name, id, options);
        }
    }

    public sealed class ForwardingRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule.
        /// 
        /// If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address:
        /// 
        /// * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name * Partial URL or by name, as in:  
        /// - projects/project_id/regions/region/addresses/address-name 
        /// - regions/region/addresses/address-name 
        /// - global/addresses/address-name 
        /// - address-name  
        /// 
        /// The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
        /// 
        /// Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true.
        /// 
        /// For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
        /// </summary>
        [Input("IPAddress")]
        public Input<string>? IPAddress { get; set; }

        /// <summary>
        /// The IP protocol to which this rule applies.
        /// 
        /// For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP.
        /// 
        /// The valid IP protocols are different for different load balancing products:  
        /// - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid. 
        /// - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid.  
        /// - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid. 
        /// - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid. 
        /// - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid.
        /// </summary>
        [Input("IPProtocol")]
        public Input<string>? IPProtocol { get; set; }

        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. This field cannot be used with port or portRange fields.
        /// 
        /// When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
        /// </summary>
        [Input("allPorts")]
        public Input<bool>? AllPorts { get; set; }

        /// <summary>
        /// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
        /// </summary>
        [Input("allowGlobalAccess")]
        public Input<bool>? AllowGlobalAccess { get; set; }

        /// <summary>
        /// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
        /// </summary>
        [Input("backendService")]
        public Input<string>? BackendService { get; set; }

        /// <summary>
        /// [Output Only] Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
        /// </summary>
        [Input("isMirroringCollector")]
        public Input<bool>? IsMirroringCollector { get; set; }

        /// <summary>
        /// [Output Only] Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
        /// 
        /// To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Specifies the forwarding rule type.
        /// 
        ///  
        /// - EXTERNAL is used for:  
        /// - Classic Cloud VPN gateways 
        /// - Protocol forwarding to VMs from an external IP address 
        /// - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing    
        /// - INTERNAL is used for:  
        /// - Protocol forwarding to VMs from an internal IP address 
        /// - Internal TCP/UDP Load Balancing   
        /// - INTERNAL_MANAGED is used for:  
        /// - Internal HTTP(S) Load Balancing   
        /// - INTERNAL_SELF_MANAGED is used for:  
        /// - Traffic Director    
        /// 
        /// For more information about forwarding rules, refer to Forwarding rule concepts.
        /// </summary>
        [Input("loadBalancingScheme")]
        public Input<string>? LoadBalancingScheme { get; set; }

        [Input("metadataFilters")]
        private InputList<Inputs.MetadataFilterArgs>? _metadataFilters;

        /// <summary>
        /// Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to Loadbalancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule will not be visible to those proxies.
        /// For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match.
        /// metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references.
        /// metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        /// </summary>
        public InputList<Inputs.MetadataFilterArgs> MetadataFilters
        {
            get => _metadataFilters ?? (_metadataFilters = new InputList<Inputs.MetadataFilterArgs>());
            set => _metadataFilters = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This field is not used for external load balancing.
        /// 
        /// For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
        /// 
        /// For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD.
        /// 
        /// For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM.
        /// 
        /// If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
        /// </summary>
        [Input("networkTier")]
        public Input<string>? NetworkTier { get; set; }

        /// <summary>
        /// This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP.
        /// 
        /// Packets addressed to ports in the specified range will be forwarded to target or  backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges.
        /// 
        /// Some types of forwarding target have constraints on the acceptable ports:  
        /// - TargetHttpProxy: 80, 8080 
        /// - TargetHttpsProxy: 443 
        /// - TargetGrpcProxy: no constraints 
        /// - TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
        /// - TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222 
        /// - TargetVpnGateway: 500, 4500
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        [Input("ports")]
        private InputList<string>? _ports;

        /// <summary>
        /// The ports field is only supported when the forwarding rule references a backend_service directly. Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing. Only packets addressed to the specified list of ports are forwarded to backends.
        /// 
        /// You can only use one of ports and port_range, or allPorts. The three are mutually exclusive.
        /// 
        /// You can specify a list of up to five ports, which can be non-contiguous.
        /// 
        /// For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports.
        /// 
        /// For more information, see [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
        /// </summary>
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// [Output Only] The PSC connection id of the PSC Forwarding Rule.
        /// </summary>
        [Input("pscConnectionId")]
        public Input<string>? PscConnectionId { get; set; }

        /// <summary>
        /// [Output Only] URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// [Output Only] Server-defined URL for the resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("serviceDirectoryRegistrations")]
        private InputList<Inputs.ForwardingRuleServiceDirectoryRegistrationArgs>? _serviceDirectoryRegistrations;

        /// <summary>
        /// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
        /// 
        /// It is only supported for Internal TCP/UDP Load Balancing and Internal HTTP(S) Load Balancing.
        /// </summary>
        public InputList<Inputs.ForwardingRuleServiceDirectoryRegistrationArgs> ServiceDirectoryRegistrations
        {
            get => _serviceDirectoryRegistrations ?? (_serviceDirectoryRegistrations = new InputList<Inputs.ForwardingRuleServiceDirectoryRegistrationArgs>());
            set => _serviceDirectoryRegistrations = value;
        }

        /// <summary>
        /// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name.
        /// 
        /// The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// 
        /// This field is only used for internal load balancing.
        /// </summary>
        [Input("serviceLabel")]
        public Input<string>? ServiceLabel { get; set; }

        /// <summary>
        /// [Output Only] The internal fully qualified service name for this Forwarding Rule.
        /// 
        /// This field is only used for internal load balancing.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// This field is only used for internal load balancing.
        /// 
        /// For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule.
        /// 
        /// If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        [Input("target")]
        public Input<string>? Target { get; set; }

        public ForwardingRuleArgs()
        {
        }
    }
}
