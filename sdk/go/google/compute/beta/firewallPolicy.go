// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new policy in the specified project using the data included in the request.
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// A list of associations that belong to this firewall policy.
	Associations FirewallPolicyAssociationResponseArrayOutput `pulumi:"associations"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the firewall policy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind pulumi.StringOutput `pulumi:"kind"`
	// [Output Only] Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// [Output Only] The parent of the firewall policy.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// [Output Only] Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntOutput `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules FirewallPolicyRuleResponseArrayOutput `pulumi:"rules"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		args = &FirewallPolicyArgs{}
	}

	var resource FirewallPolicy
	err := ctx.RegisterResource("google-native:compute/beta:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("google-native:compute/beta:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
	// A list of associations that belong to this firewall policy.
	Associations []FirewallPolicyAssociationResponse `pulumi:"associations"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName *string `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the firewall policy.
	Fingerprint *string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind *string `pulumi:"kind"`
	// [Output Only] Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name *string `pulumi:"name"`
	// [Output Only] The parent of the firewall policy.
	Parent *string `pulumi:"parent"`
	// [Output Only] Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount *int `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules []FirewallPolicyRuleResponse `pulumi:"rules"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName *string `pulumi:"shortName"`
}

type FirewallPolicyState struct {
	// A list of associations that belong to this firewall policy.
	Associations FirewallPolicyAssociationResponseArrayInput
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Depreacted, please use short name instead. User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringPtrInput
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the firewall policy.
	Fingerprint pulumi.StringPtrInput
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind pulumi.StringPtrInput
	// [Output Only] Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name pulumi.StringPtrInput
	// [Output Only] The parent of the firewall policy.
	Parent pulumi.StringPtrInput
	// [Output Only] Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntPtrInput
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules FirewallPolicyRuleResponseArrayInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringPtrInput
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
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
	Fingerprint *string `pulumi:"fingerprint"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind *string `pulumi:"kind"`
	// [Output Only] Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name *string `pulumi:"name"`
	// [Output Only] The parent of the firewall policy.
	Parent    *string `pulumi:"parent"`
	ParentId  *string `pulumi:"parentId"`
	RequestId *string `pulumi:"requestId"`
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

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
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
	Fingerprint pulumi.StringPtrInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind pulumi.StringPtrInput
	// [Output Only] Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name pulumi.StringPtrInput
	// [Output Only] The parent of the firewall policy.
	Parent    pulumi.StringPtrInput
	ParentId  pulumi.StringPtrInput
	RequestId pulumi.StringPtrInput
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

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

type FirewallPolicyOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
}
