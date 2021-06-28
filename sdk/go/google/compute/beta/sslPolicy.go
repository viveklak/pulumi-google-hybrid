// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified SSL policy resource. Gets a list of available SSL policies by making a list() request.
type SslPolicy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// A list of features enabled when the selected profile is CUSTOM. The
	// - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
	CustomFeatures pulumi.StringArrayOutput `pulumi:"customFeatures"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayOutput `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
	MinTlsVersion pulumi.StringOutput `pulumi:"minTlsVersion"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
	Profile pulumi.StringOutput `pulumi:"profile"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
	Warnings SslPolicyWarningsItemResponseArrayOutput `pulumi:"warnings"`
}

// NewSslPolicy registers a new resource with the given unique name, arguments, and options.
func NewSslPolicy(ctx *pulumi.Context,
	name string, args *SslPolicyArgs, opts ...pulumi.ResourceOption) (*SslPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource SslPolicy
	err := ctx.RegisterResource("google-native:compute/beta:SslPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslPolicy gets an existing SslPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslPolicyState, opts ...pulumi.ResourceOption) (*SslPolicy, error) {
	var resource SslPolicy
	err := ctx.ReadResource("google-native:compute/beta:SslPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslPolicy resources.
type sslPolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// A list of features enabled when the selected profile is CUSTOM. The
	// - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures []string `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
	Fingerprint *string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
	Kind *string `pulumi:"kind"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
	Profile *string `pulumi:"profile"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
	Warnings []SslPolicyWarningsItemResponse `pulumi:"warnings"`
}

type SslPolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// A list of features enabled when the selected profile is CUSTOM. The
	// - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The list of features enabled in the SSL policy.
	EnabledFeatures pulumi.StringArrayInput
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
	Fingerprint pulumi.StringPtrInput
	// [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
	Kind pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
	MinTlsVersion pulumi.StringPtrInput
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
	Profile pulumi.StringPtrInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
	Warnings SslPolicyWarningsItemResponseArrayInput
}

func (SslPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslPolicyState)(nil)).Elem()
}

type sslPolicyArgs struct {
	// A list of features enabled when the selected profile is CUSTOM. The
	// - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
	MinTlsVersion *string `pulumi:"minTlsVersion"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
	Profile   *string `pulumi:"profile"`
	Project   string  `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a SslPolicy resource.
type SslPolicyArgs struct {
	// A list of features enabled when the selected profile is CUSTOM. The
	// - method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
	CustomFeatures pulumi.StringArrayInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
	MinTlsVersion *SslPolicyMinTlsVersion
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
	Profile   *SslPolicyProfile
	Project   pulumi.StringInput
	RequestId pulumi.StringPtrInput
}

func (SslPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslPolicyArgs)(nil)).Elem()
}

type SslPolicyInput interface {
	pulumi.Input

	ToSslPolicyOutput() SslPolicyOutput
	ToSslPolicyOutputWithContext(ctx context.Context) SslPolicyOutput
}

func (*SslPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SslPolicy)(nil))
}

func (i *SslPolicy) ToSslPolicyOutput() SslPolicyOutput {
	return i.ToSslPolicyOutputWithContext(context.Background())
}

func (i *SslPolicy) ToSslPolicyOutputWithContext(ctx context.Context) SslPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslPolicyOutput)
}

type SslPolicyOutput struct {
	*pulumi.OutputState
}

func (SslPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslPolicy)(nil))
}

func (o SslPolicyOutput) ToSslPolicyOutput() SslPolicyOutput {
	return o
}

func (o SslPolicyOutput) ToSslPolicyOutputWithContext(ctx context.Context) SslPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SslPolicyOutput{})
}
