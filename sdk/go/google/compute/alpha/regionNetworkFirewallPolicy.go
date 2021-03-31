// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new network firewall policy in the specified project and region.
type RegionNetworkFirewallPolicy struct {
	pulumi.CustomResourceState
}

// NewRegionNetworkFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewRegionNetworkFirewallPolicy(ctx *pulumi.Context,
	name string, args *RegionNetworkFirewallPolicyArgs, opts ...pulumi.ResourceOption) (*RegionNetworkFirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource RegionNetworkFirewallPolicy
	err := ctx.RegisterResource("google-cloud:compute/alpha:RegionNetworkFirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionNetworkFirewallPolicy gets an existing RegionNetworkFirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionNetworkFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionNetworkFirewallPolicyState, opts ...pulumi.ResourceOption) (*RegionNetworkFirewallPolicy, error) {
	var resource RegionNetworkFirewallPolicy
	err := ctx.ReadResource("google-cloud:compute/alpha:RegionNetworkFirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionNetworkFirewallPolicy resources.
type regionNetworkFirewallPolicyState struct {
}

type RegionNetworkFirewallPolicyState struct {
}

func (RegionNetworkFirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkFirewallPolicyState)(nil)).Elem()
}

type regionNetworkFirewallPolicyArgs struct {
	// A list of associations that belong to this firewall policy.
	Associations []FirewallPolicyAssociation `pulumi:"associations"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName *string `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the firewall policy.
	Fingerprint    *string `pulumi:"fingerprint"`
	FirewallPolicy string  `pulumi:"firewallPolicy"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind *string `pulumi:"kind"`
	// [Output Only] Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name *string `pulumi:"name"`
	// [Output Only] The parent of the firewall policy.
	Parent  *string `pulumi:"parent"`
	Project string  `pulumi:"project"`
	// [Output Only] URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// [Output Only] Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount *int `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules []FirewallPolicyRule `pulumi:"rules"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName *string `pulumi:"shortName"`
}

// The set of arguments for constructing a RegionNetworkFirewallPolicy resource.
type RegionNetworkFirewallPolicyArgs struct {
	// A list of associations that belong to this firewall policy.
	Associations FirewallPolicyAssociationArrayInput
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringPtrInput
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the firewall policy.
	Fingerprint    pulumi.StringPtrInput
	FirewallPolicy pulumi.StringInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind pulumi.StringPtrInput
	// [Output Only] Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name pulumi.StringPtrInput
	// [Output Only] The parent of the firewall policy.
	Parent  pulumi.StringPtrInput
	Project pulumi.StringInput
	// [Output Only] URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringInput
	// [Output Only] Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntPtrInput
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules FirewallPolicyRuleArrayInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringPtrInput
}

func (RegionNetworkFirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkFirewallPolicyArgs)(nil)).Elem()
}

type RegionNetworkFirewallPolicyInput interface {
	pulumi.Input

	ToRegionNetworkFirewallPolicyOutput() RegionNetworkFirewallPolicyOutput
	ToRegionNetworkFirewallPolicyOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyOutput
}

func (*RegionNetworkFirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionNetworkFirewallPolicy)(nil))
}

func (i *RegionNetworkFirewallPolicy) ToRegionNetworkFirewallPolicyOutput() RegionNetworkFirewallPolicyOutput {
	return i.ToRegionNetworkFirewallPolicyOutputWithContext(context.Background())
}

func (i *RegionNetworkFirewallPolicy) ToRegionNetworkFirewallPolicyOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkFirewallPolicyOutput)
}

type RegionNetworkFirewallPolicyOutput struct {
	*pulumi.OutputState
}

func (RegionNetworkFirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionNetworkFirewallPolicy)(nil))
}

func (o RegionNetworkFirewallPolicyOutput) ToRegionNetworkFirewallPolicyOutput() RegionNetworkFirewallPolicyOutput {
	return o
}

func (o RegionNetworkFirewallPolicyOutput) ToRegionNetworkFirewallPolicyOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegionNetworkFirewallPolicyOutput{})
}
