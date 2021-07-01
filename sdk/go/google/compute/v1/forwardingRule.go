// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a ForwardingRule resource in the specified project and region using the data included in the request.
type ForwardingRule struct {
	pulumi.CustomResourceState

	// IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule.
	//
	// If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address:
	//
	// * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name * Partial URL or by name, as in:
	// - projects/project_id/regions/region/addresses/address-name
	// - regions/region/addresses/address-name
	// - global/addresses/address-name
	// - address-name
	//
	// The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	//
	// Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
	IPAddress pulumi.StringOutput `pulumi:"IPAddress"`
	// The IP protocol to which this rule applies.
	//
	// For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP.
	//
	// The valid IP protocols are different for different load balancing products:
	// - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid.
	// - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid.
	// - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid.
	// - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid.
	// - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid.
	IPProtocol pulumi.StringOutput `pulumi:"IPProtocol"`
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. This field cannot be used with port or portRange fields.
	//
	// When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
	AllPorts pulumi.BoolOutput `pulumi:"allPorts"`
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
	AllowGlobalAccess pulumi.BoolOutput `pulumi:"allowGlobalAccess"`
	// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
	BackendService pulumi.StringOutput `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request.
	//
	// To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector pulumi.BoolOutput `pulumi:"isMirroringCollector"`
	// Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Specifies the forwarding rule type.
	//
	// - EXTERNAL is used for:
	// - Classic Cloud VPN gateways
	// - Protocol forwarding to VMs from an external IP address
	// - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing
	// - INTERNAL is used for:
	// - Protocol forwarding to VMs from an internal IP address
	// - Internal TCP/UDP Load Balancing
	// - INTERNAL_MANAGED is used for:
	// - Internal HTTP(S) Load Balancing
	// - INTERNAL_SELF_MANAGED is used for:
	// - Traffic Director
	//
	// For more information about forwarding rules, refer to Forwarding rule concepts.
	LoadBalancingScheme pulumi.StringOutput `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to Loadbalancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule will not be visible to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match.
	// metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	MetadataFilters MetadataFilterResponseArrayOutput `pulumi:"metadataFilters"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// This field is not used for external load balancing.
	//
	// For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
	Network pulumi.StringOutput `pulumi:"network"`
	// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD.
	//
	// For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM.
	//
	// If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP.
	//
	// Packets addressed to ports in the specified range will be forwarded to target or  backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges.
	//
	// Some types of forwarding target have constraints on the acceptable ports:
	// - TargetHttpProxy: 80, 8080
	// - TargetHttpsProxy: 443
	// - TargetGrpcProxy: no constraints
	// - TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// - TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// - TargetVpnGateway: 500, 4500
	PortRange pulumi.StringOutput `pulumi:"portRange"`
	// The ports field is only supported when the forwarding rule references a backend_service directly. Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing. Only packets addressed to the specified list of ports are forwarded to backends.
	//
	// You can only use one of ports and port_range, or allPorts. The three are mutually exclusive.
	//
	// You can specify a list of up to five ports, which can be non-contiguous.
	//
	// For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports.
	//
	// For more information, see [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
	Ports pulumi.StringArrayOutput `pulumi:"ports"`
	// The PSC connection id of the PSC Forwarding Rule.
	PscConnectionId pulumi.StringOutput `pulumi:"pscConnectionId"`
	// URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
	//
	// It is only supported for Internal TCP/UDP Load Balancing and Internal HTTP(S) Load Balancing.
	ServiceDirectoryRegistrations ForwardingRuleServiceDirectoryRegistrationResponseArrayOutput `pulumi:"serviceDirectoryRegistrations"`
	// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name.
	//
	// The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// This field is only used for internal load balancing.
	ServiceLabel pulumi.StringOutput `pulumi:"serviceLabel"`
	// The internal fully qualified service name for this Forwarding Rule.
	//
	// This field is only used for internal load balancing.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// This field is only used for internal load balancing.
	//
	// For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule.
	//
	// If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	Target     pulumi.StringOutput `pulumi:"target"`
}

// NewForwardingRule registers a new resource with the given unique name, arguments, and options.
func NewForwardingRule(ctx *pulumi.Context,
	name string, args *ForwardingRuleArgs, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource ForwardingRule
	err := ctx.RegisterResource("google-native:compute/v1:ForwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForwardingRule gets an existing ForwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardingRuleState, opts ...pulumi.ResourceOption) (*ForwardingRule, error) {
	var resource ForwardingRule
	err := ctx.ReadResource("google-native:compute/v1:ForwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ForwardingRule resources.
type forwardingRuleState struct {
}

type ForwardingRuleState struct {
}

func (ForwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleState)(nil)).Elem()
}

type forwardingRuleArgs struct {
	// IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule.
	//
	// If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address:
	//
	// * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name * Partial URL or by name, as in:
	// - projects/project_id/regions/region/addresses/address-name
	// - regions/region/addresses/address-name
	// - global/addresses/address-name
	// - address-name
	//
	// The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	//
	// Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
	IPAddress *string `pulumi:"IPAddress"`
	// The IP protocol to which this rule applies.
	//
	// For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP.
	//
	// The valid IP protocols are different for different load balancing products:
	// - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid.
	// - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid.
	// - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid.
	// - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid.
	// - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid.
	IPProtocol *string `pulumi:"IPProtocol"`
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. This field cannot be used with port or portRange fields.
	//
	// When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
	AllPorts *bool `pulumi:"allPorts"`
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
	AllowGlobalAccess *bool `pulumi:"allowGlobalAccess"`
	// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
	BackendService *string `pulumi:"backendService"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
	IpVersion *string `pulumi:"ipVersion"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector *bool `pulumi:"isMirroringCollector"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Specifies the forwarding rule type.
	//
	// - EXTERNAL is used for:
	// - Classic Cloud VPN gateways
	// - Protocol forwarding to VMs from an external IP address
	// - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing
	// - INTERNAL is used for:
	// - Protocol forwarding to VMs from an internal IP address
	// - Internal TCP/UDP Load Balancing
	// - INTERNAL_MANAGED is used for:
	// - Internal HTTP(S) Load Balancing
	// - INTERNAL_SELF_MANAGED is used for:
	// - Traffic Director
	//
	// For more information about forwarding rules, refer to Forwarding rule concepts.
	LoadBalancingScheme *string `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to Loadbalancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule will not be visible to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match.
	// metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	MetadataFilters []MetadataFilter `pulumi:"metadataFilters"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// This field is not used for external load balancing.
	//
	// For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
	Network *string `pulumi:"network"`
	// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD.
	//
	// For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM.
	//
	// If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
	NetworkTier *string `pulumi:"networkTier"`
	// This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP.
	//
	// Packets addressed to ports in the specified range will be forwarded to target or  backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges.
	//
	// Some types of forwarding target have constraints on the acceptable ports:
	// - TargetHttpProxy: 80, 8080
	// - TargetHttpsProxy: 443
	// - TargetGrpcProxy: no constraints
	// - TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// - TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// - TargetVpnGateway: 500, 4500
	PortRange *string `pulumi:"portRange"`
	// The ports field is only supported when the forwarding rule references a backend_service directly. Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing. Only packets addressed to the specified list of ports are forwarded to backends.
	//
	// You can only use one of ports and port_range, or allPorts. The three are mutually exclusive.
	//
	// You can specify a list of up to five ports, which can be non-contiguous.
	//
	// For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports.
	//
	// For more information, see [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
	Ports     []string `pulumi:"ports"`
	Project   string   `pulumi:"project"`
	Region    string   `pulumi:"region"`
	RequestId *string  `pulumi:"requestId"`
	// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
	//
	// It is only supported for Internal TCP/UDP Load Balancing and Internal HTTP(S) Load Balancing.
	ServiceDirectoryRegistrations []ForwardingRuleServiceDirectoryRegistration `pulumi:"serviceDirectoryRegistrations"`
	// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name.
	//
	// The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// This field is only used for internal load balancing.
	ServiceLabel *string `pulumi:"serviceLabel"`
	// This field is only used for internal load balancing.
	//
	// For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule.
	//
	// If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
	Subnetwork *string `pulumi:"subnetwork"`
	Target     *string `pulumi:"target"`
}

// The set of arguments for constructing a ForwardingRule resource.
type ForwardingRuleArgs struct {
	// IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule.
	//
	// If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address:
	//
	// * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name * Partial URL or by name, as in:
	// - projects/project_id/regions/region/addresses/address-name
	// - regions/region/addresses/address-name
	// - global/addresses/address-name
	// - address-name
	//
	// The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).
	//
	// Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
	IPAddress pulumi.StringPtrInput
	// The IP protocol to which this rule applies.
	//
	// For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP and ICMP.
	//
	// The valid IP protocols are different for different load balancing products:
	// - Internal TCP/UDP Load Balancing: The load balancing scheme is INTERNAL, and one of TCP, UDP or ALL is valid.
	// - Traffic Director: The load balancing scheme is INTERNAL_SELF_MANAGED, and only TCP is valid.
	// - Internal HTTP(S) Load Balancing: The load balancing scheme is INTERNAL_MANAGED, and only TCP is valid.
	// - HTTP(S), SSL Proxy, and TCP Proxy Load Balancing: The load balancing scheme is EXTERNAL and only TCP is valid.
	// - Network Load Balancing: The load balancing scheme is EXTERNAL, and one of TCP or UDP is valid.
	IPProtocol *ForwardingRuleIPProtocol
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. This field cannot be used with port or portRange fields.
	//
	// When the load balancing scheme is INTERNAL and protocol is TCP/UDP, specify this field to allow packets addressed to any ports will be forwarded to the backends configured with this forwarding rule.
	AllPorts pulumi.BoolPtrInput
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
	AllowGlobalAccess pulumi.BoolPtrInput
	// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
	BackendService pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
	IpVersion *ForwardingRuleIpVersion
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector pulumi.BoolPtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Specifies the forwarding rule type.
	//
	// - EXTERNAL is used for:
	// - Classic Cloud VPN gateways
	// - Protocol forwarding to VMs from an external IP address
	// - HTTP(S), SSL Proxy, TCP Proxy, and Network Load Balancing
	// - INTERNAL is used for:
	// - Protocol forwarding to VMs from an internal IP address
	// - Internal TCP/UDP Load Balancing
	// - INTERNAL_MANAGED is used for:
	// - Internal HTTP(S) Load Balancing
	// - INTERNAL_SELF_MANAGED is used for:
	// - Traffic Director
	//
	// For more information about forwarding rules, refer to Forwarding rule concepts.
	LoadBalancingScheme *ForwardingRuleLoadBalancingScheme
	// Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to Loadbalancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule will not be visible to those proxies.
	// For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match.
	// metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references.
	// metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	MetadataFilters MetadataFilterArrayInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// This field is not used for external load balancing.
	//
	// For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used.
	//
	// For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
	Network pulumi.StringPtrInput
	// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD.
	//
	// For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM.
	//
	// If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
	NetworkTier *ForwardingRuleNetworkTier
	// This field can be used only if: * Load balancing scheme is one of EXTERNAL,  INTERNAL_SELF_MANAGED or INTERNAL_MANAGED, and * IPProtocol is one of TCP, UDP, or SCTP.
	//
	// Packets addressed to ports in the specified range will be forwarded to target or  backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges.
	//
	// Some types of forwarding target have constraints on the acceptable ports:
	// - TargetHttpProxy: 80, 8080
	// - TargetHttpsProxy: 443
	// - TargetGrpcProxy: no constraints
	// - TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// - TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222
	// - TargetVpnGateway: 500, 4500
	PortRange pulumi.StringPtrInput
	// The ports field is only supported when the forwarding rule references a backend_service directly. Supported load balancing products are Internal TCP/UDP Load Balancing and Network Load Balancing. Only packets addressed to the specified list of ports are forwarded to backends.
	//
	// You can only use one of ports and port_range, or allPorts. The three are mutually exclusive.
	//
	// You can specify a list of up to five ports, which can be non-contiguous.
	//
	// For Internal TCP/UDP Load Balancing, if you specify allPorts, you should not specify ports.
	//
	// For more information, see [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).
	Ports     pulumi.StringArrayInput
	Project   pulumi.StringInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
	//
	// It is only supported for Internal TCP/UDP Load Balancing and Internal HTTP(S) Load Balancing.
	ServiceDirectoryRegistrations ForwardingRuleServiceDirectoryRegistrationArrayInput
	// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name.
	//
	// The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// This field is only used for internal load balancing.
	ServiceLabel pulumi.StringPtrInput
	// This field is only used for internal load balancing.
	//
	// For internal load balancing, this field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule.
	//
	// If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.
	Subnetwork pulumi.StringPtrInput
	Target     pulumi.StringPtrInput
}

func (ForwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardingRuleArgs)(nil)).Elem()
}

type ForwardingRuleInput interface {
	pulumi.Input

	ToForwardingRuleOutput() ForwardingRuleOutput
	ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput
}

func (*ForwardingRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardingRule)(nil))
}

func (i *ForwardingRule) ToForwardingRuleOutput() ForwardingRuleOutput {
	return i.ToForwardingRuleOutputWithContext(context.Background())
}

func (i *ForwardingRule) ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardingRuleOutput)
}

type ForwardingRuleOutput struct {
	*pulumi.OutputState
}

func (ForwardingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardingRule)(nil))
}

func (o ForwardingRuleOutput) ToForwardingRuleOutput() ForwardingRuleOutput {
	return o
}

func (o ForwardingRuleOutput) ToForwardingRuleOutputWithContext(ctx context.Context) ForwardingRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ForwardingRuleOutput{})
}
