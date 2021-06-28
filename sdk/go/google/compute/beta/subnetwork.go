// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a subnetwork in the specified project using the data included in the request.
type Subnetwork struct {
	pulumi.CustomResourceState

	// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length.
	//
	// Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature.
	//
	// The default value is false and applies to all existing subnetworks and automatically created subnetworks.
	//
	// This field cannot be set to true at resource creation time.
	AllowSubnetCidrRoutesOverlap pulumi.BoolOutput `pulumi:"allowSubnetCidrRoutesOverlap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	EnableFlowLogs pulumi.BoolOutput `pulumi:"enableFlowLogs"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringOutput `pulumi:"gatewayAddress"`
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// The range of internal IPv6 addresses that are owned by this subnetwork.
	Ipv6CidrRange pulumi.StringOutput `pulumi:"ipv6CidrRange"`
	// Type of the resource. Always compute#subnetwork for Subnetwork resources.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
	LogConfig SubnetworkLogConfigResponseOutput `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
	Network pulumi.StringOutput `pulumi:"network"`
	// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
	PrivateIpGoogleAccess pulumi.BoolOutput `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority.
	//
	// This field can be both set at resource creation time and updated using patch.
	PrivateIpv6GoogleAccess pulumi.StringOutput `pulumi:"privateIpv6GoogleAccess"`
	// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
	Role pulumi.StringOutput `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
	SecondaryIpRanges SubnetworkSecondaryRangeResponseArrayOutput `pulumi:"secondaryIpRanges"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY CREATING: Subnetwork is provisioning DELETING: Subnetwork is being deleted UPDATING: Subnetwork is being updated
	State pulumi.StringOutput `pulumi:"state"`
}

// NewSubnetwork registers a new resource with the given unique name, arguments, and options.
func NewSubnetwork(ctx *pulumi.Context,
	name string, args *SubnetworkArgs, opts ...pulumi.ResourceOption) (*Subnetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource Subnetwork
	err := ctx.RegisterResource("google-native:compute/beta:Subnetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetwork gets an existing Subnetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetworkState, opts ...pulumi.ResourceOption) (*Subnetwork, error) {
	var resource Subnetwork
	err := ctx.ReadResource("google-native:compute/beta:Subnetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnetwork resources.
type subnetworkState struct {
	// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length.
	//
	// Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature.
	//
	// The default value is false and applies to all existing subnetworks and automatically created subnetworks.
	//
	// This field cannot be set to true at resource creation time.
	AllowSubnetCidrRoutesOverlap *bool `pulumi:"allowSubnetCidrRoutesOverlap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
	Description *string `pulumi:"description"`
	// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	EnableFlowLogs *bool `pulumi:"enableFlowLogs"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
	Fingerprint *string `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress *string `pulumi:"gatewayAddress"`
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// The range of internal IPv6 addresses that are owned by this subnetwork.
	Ipv6CidrRange *string `pulumi:"ipv6CidrRange"`
	// Type of the resource. Always compute#subnetwork for Subnetwork resources.
	Kind *string `pulumi:"kind"`
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
	LogConfig *SubnetworkLogConfigResponse `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
	Network *string `pulumi:"network"`
	// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority.
	//
	// This field can be both set at resource creation time and updated using patch.
	PrivateIpv6GoogleAccess *string `pulumi:"privateIpv6GoogleAccess"`
	// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	Purpose *string `pulumi:"purpose"`
	// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
	Region *string `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
	Role *string `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
	SecondaryIpRanges []SubnetworkSecondaryRangeResponse `pulumi:"secondaryIpRanges"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY CREATING: Subnetwork is provisioning DELETING: Subnetwork is being deleted UPDATING: Subnetwork is being updated
	State *string `pulumi:"state"`
}

type SubnetworkState struct {
	// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length.
	//
	// Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature.
	//
	// The default value is false and applies to all existing subnetworks and automatically created subnetworks.
	//
	// This field cannot be set to true at resource creation time.
	AllowSubnetCidrRoutesOverlap pulumi.BoolPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
	Description pulumi.StringPtrInput
	// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	EnableFlowLogs pulumi.BoolPtrInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a Subnetwork. An up-to-date fingerprint must be provided in order to update the Subnetwork, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a Subnetwork.
	Fingerprint pulumi.StringPtrInput
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringPtrInput
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
	IpCidrRange pulumi.StringPtrInput
	// The range of internal IPv6 addresses that are owned by this subnetwork.
	Ipv6CidrRange pulumi.StringPtrInput
	// Type of the resource. Always compute#subnetwork for Subnetwork resources.
	Kind pulumi.StringPtrInput
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
	LogConfig SubnetworkLogConfigResponsePtrInput
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
	Network pulumi.StringPtrInput
	// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority.
	//
	// This field can be both set at resource creation time and updated using patch.
	PrivateIpv6GoogleAccess pulumi.StringPtrInput
	// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	Purpose pulumi.StringPtrInput
	// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
	Region pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
	Role pulumi.StringPtrInput
	// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
	SecondaryIpRanges SubnetworkSecondaryRangeResponseArrayInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// The state of the subnetwork, which can be one of the following values: READY: Subnetwork is created and ready to use DRAINING: only applicable to subnetworks that have the purpose set to INTERNAL_HTTPS_LOAD_BALANCER and indicates that connections to the load balancer are being drained. A subnetwork that is draining cannot be used or modified until it reaches a status of READY CREATING: Subnetwork is provisioning DELETING: Subnetwork is being deleted UPDATING: Subnetwork is being updated
	State pulumi.StringPtrInput
}

func (SubnetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkState)(nil)).Elem()
}

type subnetworkArgs struct {
	// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length.
	//
	// Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature.
	//
	// The default value is false and applies to all existing subnetworks and automatically created subnetworks.
	//
	// This field cannot be set to true at resource creation time.
	AllowSubnetCidrRoutesOverlap *bool `pulumi:"allowSubnetCidrRoutesOverlap"`
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
	Description *string `pulumi:"description"`
	// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	EnableFlowLogs *bool `pulumi:"enableFlowLogs"`
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
	LogConfig *SubnetworkLogConfig `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
	Network *string `pulumi:"network"`
	// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority.
	//
	// This field can be both set at resource creation time and updated using patch.
	PrivateIpv6GoogleAccess *string `pulumi:"privateIpv6GoogleAccess"`
	Project                 string  `pulumi:"project"`
	// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	Purpose *string `pulumi:"purpose"`
	// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
	// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
	Role *string `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
	SecondaryIpRanges []SubnetworkSecondaryRange `pulumi:"secondaryIpRanges"`
}

// The set of arguments for constructing a Subnetwork resource.
type SubnetworkArgs struct {
	// Whether this subnetwork can conflict with static routes. Setting this to true allows this subnetwork's primary and secondary ranges to conflict with routes that have already been configured on the corresponding network. Static routes will take precedence over the subnetwork route if the route prefix length is at least as large as the subnetwork prefix length.
	//
	// Also, packets destined to IPs within subnetwork may contain private/sensitive data and are prevented from leaving the virtual network. Setting this field to true will disable this feature.
	//
	// The default value is false and applies to all existing subnetworks and automatically created subnetworks.
	//
	// This field cannot be set to true at resource creation time.
	AllowSubnetCidrRoutesOverlap pulumi.BoolPtrInput
	// An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.
	Description pulumi.StringPtrInput
	// Whether to enable flow logging for this subnetwork. If this field is not explicitly set, it will not appear in get listings. If not set the default behavior is to disable flow logging. This field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	EnableFlowLogs pulumi.BoolPtrInput
	// The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 100.64.0.0/10. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported. This field is set at resource creation time. The range can be any range listed in the Valid ranges list. The range can be expanded after creation using expandIpCidrRange.
	IpCidrRange pulumi.StringPtrInput
	// This field denotes the VPC flow logging options for this subnetwork. If logging is enabled, logs are exported to Cloud Logging.
	LogConfig SubnetworkLogConfigPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network to which this subnetwork belongs, provided by the client when initially creating the subnetwork. Only networks that are in the distributed mode can have subnetworks. This field can be set only at resource creation time.
	Network pulumi.StringPtrInput
	// Whether the VMs in this subnet can access Google services without assigned external IP addresses. This field can be both set at resource creation time and updated using setPrivateIpGoogleAccess.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The private IPv6 google access type for the VMs in this subnet. This is an expanded field of enablePrivateV6Access. If both fields are set, privateIpv6GoogleAccess will take priority.
	//
	// This field can be both set at resource creation time and updated using patch.
	PrivateIpv6GoogleAccess *SubnetworkPrivateIpv6GoogleAccess
	Project                 pulumi.StringInput
	// The purpose of the resource. This field can be either PRIVATE_RFC_1918 or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is reserved for Internal HTTP(S) Load Balancing. If unspecified, the purpose defaults to PRIVATE_RFC_1918. The enableFlowLogs field isn't supported with the purpose field set to INTERNAL_HTTPS_LOAD_BALANCER.
	Purpose *SubnetworkPurpose
	// URL of the region where the Subnetwork resides. This field can be set only at resource creation time.
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE or BACKUP. An ACTIVE subnetwork is one that is currently being used for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that is ready to be promoted to ACTIVE or is currently draining. This field can be updated with a patch request.
	Role *SubnetworkRole
	// An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field can be updated with a patch request.
	SecondaryIpRanges SubnetworkSecondaryRangeArrayInput
}

func (SubnetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkArgs)(nil)).Elem()
}

type SubnetworkInput interface {
	pulumi.Input

	ToSubnetworkOutput() SubnetworkOutput
	ToSubnetworkOutputWithContext(ctx context.Context) SubnetworkOutput
}

func (*Subnetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnetwork)(nil))
}

func (i *Subnetwork) ToSubnetworkOutput() SubnetworkOutput {
	return i.ToSubnetworkOutputWithContext(context.Background())
}

func (i *Subnetwork) ToSubnetworkOutputWithContext(ctx context.Context) SubnetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkOutput)
}

type SubnetworkOutput struct {
	*pulumi.OutputState
}

func (SubnetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnetwork)(nil))
}

func (o SubnetworkOutput) ToSubnetworkOutput() SubnetworkOutput {
	return o
}

func (o SubnetworkOutput) ToSubnetworkOutputWithContext(ctx context.Context) SubnetworkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetworkOutput{})
}
