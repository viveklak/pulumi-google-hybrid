// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new policy in the specified project using the data included in the request.
type SecurityPolicy struct {
	pulumi.CustomResourceState

	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigResponseOutput `pulumi:"adaptiveProtectionConfig"`
	AdvancedOptionsConfig    SecurityPolicyAdvancedOptionsConfigResponseOutput    `pulumi:"advancedOptionsConfig"`
	// A list of associations that belong to this policy.
	Associations     SecurityPolicyAssociationResponseArrayOutput `pulumi:"associations"`
	CloudArmorConfig SecurityPolicyCloudArmorConfigResponseOutput `pulumi:"cloudArmorConfig"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp    pulumi.StringOutput                              `pulumi:"creationTimestamp"`
	DdosProtectionConfig SecurityPolicyDdosProtectionConfigResponseOutput `pulumi:"ddosProtectionConfig"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of the security policy.
	Parent                 pulumi.StringOutput                                `pulumi:"parent"`
	RecaptchaOptionsConfig SecurityPolicyRecaptchaOptionsConfigResponseOutput `pulumi:"recaptchaOptionsConfig"`
	// URL of the region where the regional security policy resides. This field is not applicable to global security policies.
	Region pulumi.StringOutput `pulumi:"region"`
	// Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntOutput `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules SecurityPolicyRuleResponseArrayOutput `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecurityPolicy(ctx *pulumi.Context,
	name string, args *SecurityPolicyArgs, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	if args == nil {
		args = &SecurityPolicyArgs{}
	}

	var resource SecurityPolicy
	err := ctx.RegisterResource("google-native:compute/alpha:SecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityPolicy gets an existing SecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPolicyState, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	var resource SecurityPolicy
	err := ctx.ReadResource("google-native:compute/alpha:SecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityPolicy resources.
type securityPolicyState struct {
}

type SecurityPolicyState struct {
}

func (SecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyState)(nil)).Elem()
}

type securityPolicyArgs struct {
	AdaptiveProtectionConfig *SecurityPolicyAdaptiveProtectionConfig `pulumi:"adaptiveProtectionConfig"`
	AdvancedOptionsConfig    *SecurityPolicyAdvancedOptionsConfig    `pulumi:"advancedOptionsConfig"`
	// A list of associations that belong to this policy.
	Associations         []SecurityPolicyAssociation         `pulumi:"associations"`
	CloudArmorConfig     *SecurityPolicyCloudArmorConfig     `pulumi:"cloudArmorConfig"`
	DdosProtectionConfig *SecurityPolicyDdosProtectionConfig `pulumi:"ddosProtectionConfig"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName *string `pulumi:"displayName"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name                   *string                               `pulumi:"name"`
	Project                *string                               `pulumi:"project"`
	RecaptchaOptionsConfig *SecurityPolicyRecaptchaOptionsConfig `pulumi:"recaptchaOptionsConfig"`
	RequestId              *string                               `pulumi:"requestId"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules []SecurityPolicyRule `pulumi:"rules"`
	// The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
	Type         *SecurityPolicyType `pulumi:"type"`
	ValidateOnly *string             `pulumi:"validateOnly"`
}

// The set of arguments for constructing a SecurityPolicy resource.
type SecurityPolicyArgs struct {
	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigPtrInput
	AdvancedOptionsConfig    SecurityPolicyAdvancedOptionsConfigPtrInput
	// A list of associations that belong to this policy.
	Associations         SecurityPolicyAssociationArrayInput
	CloudArmorConfig     SecurityPolicyCloudArmorConfigPtrInput
	DdosProtectionConfig SecurityPolicyDdosProtectionConfigPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName pulumi.StringPtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name                   pulumi.StringPtrInput
	Project                pulumi.StringPtrInput
	RecaptchaOptionsConfig SecurityPolicyRecaptchaOptionsConfigPtrInput
	RequestId              pulumi.StringPtrInput
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules SecurityPolicyRuleArrayInput
	// The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
	Type         SecurityPolicyTypePtrInput
	ValidateOnly pulumi.StringPtrInput
}

func (SecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyArgs)(nil)).Elem()
}

type SecurityPolicyInput interface {
	pulumi.Input

	ToSecurityPolicyOutput() SecurityPolicyOutput
	ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput
}

func (*SecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (i *SecurityPolicy) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return i.ToSecurityPolicyOutputWithContext(context.Background())
}

func (i *SecurityPolicy) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyOutput)
}

type SecurityPolicyOutput struct{ *pulumi.OutputState }

func (SecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPolicy)(nil)).Elem()
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return o
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityPolicyInput)(nil)).Elem(), &SecurityPolicy{})
	pulumi.RegisterOutputType(SecurityPolicyOutput{})
}
