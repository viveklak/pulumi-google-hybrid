// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified firewall policy.
func LookupFirewallPolicy(ctx *pulumi.Context, args *LookupFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyResult, error) {
	var rv LookupFirewallPolicyResult
	err := ctx.Invoke("google-native:compute/v1:getFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyArgs struct {
	FirewallPolicy string `pulumi:"firewallPolicy"`
}

type LookupFirewallPolicyResult struct {
	// A list of associations that belong to this firewall policy.
	Associations []FirewallPolicyAssociationResponse `pulumi:"associations"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Deprecated: Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName string `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the firewall policy.
	Fingerprint string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
	Kind string `pulumi:"kind"`
	// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name string `pulumi:"name"`
	// The parent of the firewall policy.
	Parent string `pulumi:"parent"`
	// URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount int `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
	Rules []FirewallPolicyRuleResponse `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName string `pulumi:"shortName"`
}

func LookupFirewallPolicyOutput(ctx *pulumi.Context, args LookupFirewallPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicyResult, error) {
			args := v.(LookupFirewallPolicyArgs)
			r, err := LookupFirewallPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupFirewallPolicyResultOutput)
}

type LookupFirewallPolicyOutputArgs struct {
	FirewallPolicy pulumi.StringInput `pulumi:"firewallPolicy"`
}

func (LookupFirewallPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyArgs)(nil)).Elem()
}

type LookupFirewallPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyResult)(nil)).Elem()
}

func (o LookupFirewallPolicyResultOutput) ToLookupFirewallPolicyResultOutput() LookupFirewallPolicyResultOutput {
	return o
}

func (o LookupFirewallPolicyResultOutput) ToLookupFirewallPolicyResultOutputWithContext(ctx context.Context) LookupFirewallPolicyResultOutput {
	return o
}

// A list of associations that belong to this firewall policy.
func (o LookupFirewallPolicyResultOutput) Associations() FirewallPolicyAssociationResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) []FirewallPolicyAssociationResponse { return v.Associations }).(FirewallPolicyAssociationResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupFirewallPolicyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupFirewallPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
//
// Deprecated: Deprecated, please use short name instead. User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupFirewallPolicyResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the firewall policy.
func (o LookupFirewallPolicyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// [Output only] Type of the resource. Always compute#firewallPolicyfor firewall policies
func (o LookupFirewallPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
func (o LookupFirewallPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parent of the firewall policy.
func (o LookupFirewallPolicyResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Parent }).(pulumi.StringOutput)
}

// URL of the region where the regional firewall policy resides. This field is not applicable to global firewall policies. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupFirewallPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
func (o LookupFirewallPolicyResultOutput) RuleTupleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) int { return v.RuleTupleCount }).(pulumi.IntOutput)
}

// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a firewall policy, a default rule with action "allow" will be added.
func (o LookupFirewallPolicyResultOutput) Rules() FirewallPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) []FirewallPolicyRuleResponse { return v.Rules }).(FirewallPolicyRuleResponseArrayOutput)
}

// Server-defined URL for the resource.
func (o LookupFirewallPolicyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupFirewallPolicyResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// User-provided name of the Organization firewall plicy. The name should be unique in the organization in which the firewall policy is created. This name must be set on creation and cannot be changed. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupFirewallPolicyResultOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyResult) string { return v.ShortName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicyResultOutput{})
}
