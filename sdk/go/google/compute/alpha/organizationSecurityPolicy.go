// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new policy in the specified project using the data included in the request.
type OrganizationSecurityPolicy struct {
	pulumi.CustomResourceState

	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigResponseOutput `pulumi:"adaptiveProtectionConfig"`
	// A list of associations that belong to this policy.
	Associations     SecurityPolicyAssociationResponseArrayOutput `pulumi:"associations"`
	CloudArmorConfig SecurityPolicyCloudArmorConfigResponseOutput `pulumi:"cloudArmorConfig"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// [Output Only] The parent of the security policy.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// [Output Only] Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntOutput `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules SecurityPolicyRuleResponseArrayOutput `pulumi:"rules"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// The type indicates the intended use of the security policy. CLOUD_ARMOR policies apply to backend services. FIREWALL policies apply to organizations.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOrganizationSecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSecurityPolicy(ctx *pulumi.Context,
	name string, args *OrganizationSecurityPolicyArgs, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicy, error) {
	if args == nil {
		args = &OrganizationSecurityPolicyArgs{}
	}

	var resource OrganizationSecurityPolicy
	err := ctx.RegisterResource("google-native:compute/alpha:OrganizationSecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSecurityPolicy gets an existing OrganizationSecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSecurityPolicyState, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicy, error) {
	var resource OrganizationSecurityPolicy
	err := ctx.ReadResource("google-native:compute/alpha:OrganizationSecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSecurityPolicy resources.
type organizationSecurityPolicyState struct {
	AdaptiveProtectionConfig *SecurityPolicyAdaptiveProtectionConfigResponse `pulumi:"adaptiveProtectionConfig"`
	// A list of associations that belong to this policy.
	Associations     []SecurityPolicyAssociationResponse     `pulumi:"associations"`
	CloudArmorConfig *SecurityPolicyCloudArmorConfigResponse `pulumi:"cloudArmorConfig"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName *string `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	Fingerprint *string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind *string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// [Output Only] The parent of the security policy.
	Parent *string `pulumi:"parent"`
	// [Output Only] Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount *int `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules []SecurityPolicyRuleResponse `pulumi:"rules"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// The type indicates the intended use of the security policy. CLOUD_ARMOR policies apply to backend services. FIREWALL policies apply to organizations.
	Type *string `pulumi:"type"`
}

type OrganizationSecurityPolicyState struct {
	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigResponsePtrInput
	// A list of associations that belong to this policy.
	Associations     SecurityPolicyAssociationResponseArrayInput
	CloudArmorConfig SecurityPolicyCloudArmorConfigResponsePtrInput
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringPtrInput
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	Fingerprint pulumi.StringPtrInput
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind pulumi.StringPtrInput
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint pulumi.StringPtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// [Output Only] The parent of the security policy.
	Parent pulumi.StringPtrInput
	// [Output Only] Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntPtrInput
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules SecurityPolicyRuleResponseArrayInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// The type indicates the intended use of the security policy. CLOUD_ARMOR policies apply to backend services. FIREWALL policies apply to organizations.
	Type pulumi.StringPtrInput
}

func (OrganizationSecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyState)(nil)).Elem()
}

type organizationSecurityPolicyArgs struct {
	AdaptiveProtectionConfig *SecurityPolicyAdaptiveProtectionConfig `pulumi:"adaptiveProtectionConfig"`
	// A list of associations that belong to this policy.
	Associations     []SecurityPolicyAssociation     `pulumi:"associations"`
	CloudArmorConfig *SecurityPolicyCloudArmorConfig `pulumi:"cloudArmorConfig"`
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName *string `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	Fingerprint *string `pulumi:"fingerprint"`
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id *string `pulumi:"id"`
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind *string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// [Output Only] The parent of the security policy.
	Parent    *string `pulumi:"parent"`
	ParentId  *string `pulumi:"parentId"`
	RequestId *string `pulumi:"requestId"`
	// [Output Only] Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount *int `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules []SecurityPolicyRule `pulumi:"rules"`
	// [Output Only] Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// The type indicates the intended use of the security policy. CLOUD_ARMOR policies apply to backend services. FIREWALL policies apply to organizations.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a OrganizationSecurityPolicy resource.
type OrganizationSecurityPolicyArgs struct {
	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigPtrInput
	// A list of associations that belong to this policy.
	Associations     SecurityPolicyAssociationArrayInput
	CloudArmorConfig SecurityPolicyCloudArmorConfigPtrInput
	// [Output Only] Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringPtrInput
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	Fingerprint pulumi.StringPtrInput
	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	Id pulumi.StringPtrInput
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind pulumi.StringPtrInput
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels.
	//
	// To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint pulumi.StringPtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// [Output Only] The parent of the security policy.
	Parent    pulumi.StringPtrInput
	ParentId  pulumi.StringPtrInput
	RequestId pulumi.StringPtrInput
	// [Output Only] Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntPtrInput
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules SecurityPolicyRuleArrayInput
	// [Output Only] Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// [Output Only] Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// The type indicates the intended use of the security policy. CLOUD_ARMOR policies apply to backend services. FIREWALL policies apply to organizations.
	Type pulumi.StringPtrInput
}

func (OrganizationSecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyArgs)(nil)).Elem()
}

type OrganizationSecurityPolicyInput interface {
	pulumi.Input

	ToOrganizationSecurityPolicyOutput() OrganizationSecurityPolicyOutput
	ToOrganizationSecurityPolicyOutputWithContext(ctx context.Context) OrganizationSecurityPolicyOutput
}

func (*OrganizationSecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSecurityPolicy)(nil))
}

func (i *OrganizationSecurityPolicy) ToOrganizationSecurityPolicyOutput() OrganizationSecurityPolicyOutput {
	return i.ToOrganizationSecurityPolicyOutputWithContext(context.Background())
}

func (i *OrganizationSecurityPolicy) ToOrganizationSecurityPolicyOutputWithContext(ctx context.Context) OrganizationSecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSecurityPolicyOutput)
}

type OrganizationSecurityPolicyOutput struct {
	*pulumi.OutputState
}

func (OrganizationSecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSecurityPolicy)(nil))
}

func (o OrganizationSecurityPolicyOutput) ToOrganizationSecurityPolicyOutput() OrganizationSecurityPolicyOutput {
	return o
}

func (o OrganizationSecurityPolicyOutput) ToOrganizationSecurityPolicyOutputWithContext(ctx context.Context) OrganizationSecurityPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationSecurityPolicyOutput{})
}
