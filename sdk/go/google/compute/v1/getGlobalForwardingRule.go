// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified GlobalForwardingRule resource. Gets a list of available forwarding rules by making a list() request.
func LookupGlobalForwardingRule(ctx *pulumi.Context, args *LookupGlobalForwardingRuleArgs, opts ...pulumi.InvokeOption) (*LookupGlobalForwardingRuleResult, error) {
	var rv LookupGlobalForwardingRuleResult
	err := ctx.Invoke("google-native:compute/v1:getGlobalForwardingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalForwardingRuleArgs struct {
	ForwardingRule string  `pulumi:"forwardingRule"`
	Project        *string `pulumi:"project"`
}

type LookupGlobalForwardingRuleResult struct {
	// This field is used along with the backend_service field for Internal TCP/UDP Load Balancing or Network Load Balancing, or with the target field for internal and external TargetInstance. You can only use one of ports and port_range, or allPorts. The three are mutually exclusive. For TCP, UDP and SCTP traffic, packets addressed to any ports will be forwarded to the target or backendService.
	AllPorts bool `pulumi:"allPorts"`
	// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
	AllowGlobalAccess bool `pulumi:"allowGlobalAccess"`
	// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
	BackendService string `pulumi:"backendService"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
	Fingerprint string `pulumi:"fingerprint"`
	// IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule. If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address: * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region /addresses/address-name * Partial URL or by name, as in: - projects/project_id/regions/region/addresses/address-name - regions/region/addresses/address-name - global/addresses/address-name - address-name The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, see [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true. For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
	IpAddress string `pulumi:"ipAddress"`
	// The IP protocol to which this rule applies. For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP, ICMP and L3_DEFAULT. The valid IP protocols are different for different load balancing products as described in [Load balancing features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
	IpProtocol string `pulumi:"ipProtocol"`
	// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
	IpVersion string `pulumi:"ipVersion"`
	// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
	IsMirroringCollector bool `pulumi:"isMirroringCollector"`
	// Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Specifies the forwarding rule type. For more information about forwarding rules, refer to Forwarding rule concepts.
	LoadBalancingScheme string `pulumi:"loadBalancingScheme"`
	// Opaque filter criteria used by load balancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to load balancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule are not visible to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match. metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references. metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	MetadataFilters []MetadataFilterResponse `pulumi:"metadataFilters"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. For Private Service Connect forwarding rules that forward traffic to Google APIs, the forwarding rule name must be a 1-20 characters string with lowercase letters and numbers and must start with a letter.
	Name string `pulumi:"name"`
	// This field is not used for external load balancing. For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used. For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
	Network string `pulumi:"network"`
	// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD. For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM. If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
	NetworkTier string `pulumi:"networkTier"`
	// This field can be used only if: - Load balancing scheme is one of EXTERNAL, INTERNAL_SELF_MANAGED or INTERNAL_MANAGED - IPProtocol is one of TCP, UDP, or SCTP. Packets addressed to ports in the specified range will be forwarded to target or backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint ports. Some types of forwarding target have constraints on the acceptable ports. For more information, see [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications). @pattern: \\d+(?:-\\d+)?
	PortRange string `pulumi:"portRange"`
	// The ports field is only supported when the forwarding rule references a backend_service directly. Only packets addressed to the [specified list of ports](<(https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)>) are forwarded to backends. You can only use one of ports and port_range, or allPorts. The three are mutually exclusive. You can specify a list of up to five ports, which can be non-contiguous. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint ports. @pattern: \\d+(?:-\\d+)?
	Ports []string `pulumi:"ports"`
	// The PSC connection id of the PSC Forwarding Rule.
	PscConnectionId     string `pulumi:"pscConnectionId"`
	PscConnectionStatus string `pulumi:"pscConnectionStatus"`
	// URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
	ServiceDirectoryRegistrations []ForwardingRuleServiceDirectoryRegistrationResponse `pulumi:"serviceDirectoryRegistrations"`
	// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.
	ServiceLabel string `pulumi:"serviceLabel"`
	// The internal fully qualified service name for this Forwarding Rule. This field is only used for internal load balancing.
	ServiceName string `pulumi:"serviceName"`
	// This field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule, used in internal load balancing and network load balancing with IPv6. If the network specified is in auto subnet mode, this field is optional. However, a subnetwork must be specified if the network is in custom subnet mode or when creating external forwarding rule with IPv6.
	Subnetwork string `pulumi:"subnetwork"`
	Target     string `pulumi:"target"`
}

func LookupGlobalForwardingRuleOutput(ctx *pulumi.Context, args LookupGlobalForwardingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalForwardingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalForwardingRuleResult, error) {
			args := v.(LookupGlobalForwardingRuleArgs)
			r, err := LookupGlobalForwardingRule(ctx, &args, opts...)
			return *r, err
		}).(LookupGlobalForwardingRuleResultOutput)
}

type LookupGlobalForwardingRuleOutputArgs struct {
	ForwardingRule pulumi.StringInput    `pulumi:"forwardingRule"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGlobalForwardingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalForwardingRuleArgs)(nil)).Elem()
}

type LookupGlobalForwardingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalForwardingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalForwardingRuleResult)(nil)).Elem()
}

func (o LookupGlobalForwardingRuleResultOutput) ToLookupGlobalForwardingRuleResultOutput() LookupGlobalForwardingRuleResultOutput {
	return o
}

func (o LookupGlobalForwardingRuleResultOutput) ToLookupGlobalForwardingRuleResultOutputWithContext(ctx context.Context) LookupGlobalForwardingRuleResultOutput {
	return o
}

// This field is used along with the backend_service field for Internal TCP/UDP Load Balancing or Network Load Balancing, or with the target field for internal and external TargetInstance. You can only use one of ports and port_range, or allPorts. The three are mutually exclusive. For TCP, UDP and SCTP traffic, packets addressed to any ports will be forwarded to the target or backendService.
func (o LookupGlobalForwardingRuleResultOutput) AllPorts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) bool { return v.AllPorts }).(pulumi.BoolOutput)
}

// This field is used along with the backend_service field for internal load balancing or with the target field for internal TargetInstance. If the field is set to TRUE, clients can access ILB from all regions. Otherwise only allows access from clients in the same region as the internal load balancer.
func (o LookupGlobalForwardingRuleResultOutput) AllowGlobalAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) bool { return v.AllowGlobalAccess }).(pulumi.BoolOutput)
}

// Identifies the backend service to which the forwarding rule sends traffic. Required for Internal TCP/UDP Load Balancing and Network Load Balancing; must be omitted for all other load balancer types.
func (o LookupGlobalForwardingRuleResultOutput) BackendService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.BackendService }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupGlobalForwardingRuleResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupGlobalForwardingRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a ForwardingRule. Include the fingerprint in patch request to ensure that you do not overwrite changes that were applied from another concurrent request. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
func (o LookupGlobalForwardingRuleResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// IP address that this forwarding rule serves. When a client sends traffic to this IP address, the forwarding rule directs the traffic to the target that you specify in the forwarding rule. If you don't specify a reserved IP address, an ephemeral IP address is assigned. Methods for specifying an IP address: * IPv4 dotted decimal, as in `100.1.2.3` * Full URL, as in https://www.googleapis.com/compute/v1/projects/project_id/regions/region /addresses/address-name * Partial URL or by name, as in: - projects/project_id/regions/region/addresses/address-name - regions/region/addresses/address-name - global/addresses/address-name - address-name The loadBalancingScheme and the forwarding rule's target determine the type of IP address that you can use. For detailed information, see [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications). Must be set to `0.0.0.0` when the target is targetGrpcProxy that has validateForProxyless field set to true. For Private Service Connect forwarding rules that forward traffic to Google APIs, IP address must be provided.
func (o LookupGlobalForwardingRuleResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The IP protocol to which this rule applies. For protocol forwarding, valid options are TCP, UDP, ESP, AH, SCTP, ICMP and L3_DEFAULT. The valid IP protocols are different for different load balancing products as described in [Load balancing features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends).
func (o LookupGlobalForwardingRuleResultOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.IpProtocol }).(pulumi.StringOutput)
}

// The IP Version that will be used by this forwarding rule. Valid options are IPV4 or IPV6. This can only be specified for an external global forwarding rule.
func (o LookupGlobalForwardingRuleResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

// Indicates whether or not this load balancer can be used as a collector for packet mirroring. To prevent mirroring loops, instances behind this load balancer will not have their traffic mirrored even if a PacketMirroring rule applies to them. This can only be set to true for load balancers that have their loadBalancingScheme set to INTERNAL.
func (o LookupGlobalForwardingRuleResultOutput) IsMirroringCollector() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) bool { return v.IsMirroringCollector }).(pulumi.BoolOutput)
}

// Type of the resource. Always compute#forwardingRule for Forwarding Rule resources.
func (o LookupGlobalForwardingRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this resource, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a ForwardingRule.
func (o LookupGlobalForwardingRuleResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupGlobalForwardingRuleResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Specifies the forwarding rule type. For more information about forwarding rules, refer to Forwarding rule concepts.
func (o LookupGlobalForwardingRuleResultOutput) LoadBalancingScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.LoadBalancingScheme }).(pulumi.StringOutput)
}

// Opaque filter criteria used by load balancer to restrict routing configuration to a limited set of xDS compliant clients. In their xDS requests to load balancer, xDS clients present node metadata. When there is a match, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. TargetHttpProxy, UrlMap) referenced by the ForwardingRule are not visible to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. If multiple metadataFilters are specified, all of them need to be satisfied in order to be considered a match. metadataFilters specified here will be applifed before those specified in the UrlMap that this ForwardingRule references. metadataFilters only applies to Loadbalancers that have their loadBalancingScheme set to INTERNAL_SELF_MANAGED.
func (o LookupGlobalForwardingRuleResultOutput) MetadataFilters() MetadataFilterResponseArrayOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) []MetadataFilterResponse { return v.MetadataFilters }).(MetadataFilterResponseArrayOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. For Private Service Connect forwarding rules that forward traffic to Google APIs, the forwarding rule name must be a 1-20 characters string with lowercase letters and numbers and must start with a letter.
func (o LookupGlobalForwardingRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// This field is not used for external load balancing. For Internal TCP/UDP Load Balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used. For Private Service Connect forwarding rules that forward traffic to Google APIs, a network must be provided.
func (o LookupGlobalForwardingRuleResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Network }).(pulumi.StringOutput)
}

// This signifies the networking tier used for configuring this load balancer and can only take the following values: PREMIUM, STANDARD. For regional ForwardingRule, the valid values are PREMIUM and STANDARD. For GlobalForwardingRule, the valid value is PREMIUM. If this field is not specified, it is assumed to be PREMIUM. If IPAddress is specified, this value must be equal to the networkTier of the Address.
func (o LookupGlobalForwardingRuleResultOutput) NetworkTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.NetworkTier }).(pulumi.StringOutput)
}

// This field can be used only if: - Load balancing scheme is one of EXTERNAL, INTERNAL_SELF_MANAGED or INTERNAL_MANAGED - IPProtocol is one of TCP, UDP, or SCTP. Packets addressed to ports in the specified range will be forwarded to target or backend_service. You can only use one of ports, port_range, or allPorts. The three are mutually exclusive. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint ports. Some types of forwarding target have constraints on the acceptable ports. For more information, see [Port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications). @pattern: \\d+(?:-\\d+)?
func (o LookupGlobalForwardingRuleResultOutput) PortRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.PortRange }).(pulumi.StringOutput)
}

// The ports field is only supported when the forwarding rule references a backend_service directly. Only packets addressed to the [specified list of ports](<(https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)>) are forwarded to backends. You can only use one of ports and port_range, or allPorts. The three are mutually exclusive. You can specify a list of up to five ports, which can be non-contiguous. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint ports. @pattern: \\d+(?:-\\d+)?
func (o LookupGlobalForwardingRuleResultOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) []string { return v.Ports }).(pulumi.StringArrayOutput)
}

// The PSC connection id of the PSC Forwarding Rule.
func (o LookupGlobalForwardingRuleResultOutput) PscConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.PscConnectionId }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) PscConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.PscConnectionStatus }).(pulumi.StringOutput)
}

// URL of the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupGlobalForwardingRuleResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupGlobalForwardingRuleResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Service Directory resources to register this forwarding rule with. Currently, only supports a single Service Directory resource.
func (o LookupGlobalForwardingRuleResultOutput) ServiceDirectoryRegistrations() ForwardingRuleServiceDirectoryRegistrationResponseArrayOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) []ForwardingRuleServiceDirectoryRegistrationResponse {
		return v.ServiceDirectoryRegistrations
	}).(ForwardingRuleServiceDirectoryRegistrationResponseArrayOutput)
}

// An optional prefix to the service name for this Forwarding Rule. If specified, the prefix is the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. This field is only used for internal load balancing.
func (o LookupGlobalForwardingRuleResultOutput) ServiceLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.ServiceLabel }).(pulumi.StringOutput)
}

// The internal fully qualified service name for this Forwarding Rule. This field is only used for internal load balancing.
func (o LookupGlobalForwardingRuleResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// This field identifies the subnetwork that the load balanced IP should belong to for this Forwarding Rule, used in internal load balancing and network load balancing with IPv6. If the network specified is in auto subnet mode, this field is optional. However, a subnetwork must be specified if the network is in custom subnet mode or when creating external forwarding rule with IPv6.
func (o LookupGlobalForwardingRuleResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

func (o LookupGlobalForwardingRuleResultOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalForwardingRuleResult) string { return v.Target }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalForwardingRuleResultOutput{})
}
