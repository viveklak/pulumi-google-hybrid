// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists all of the ordered rules present in a single specified policy.
func LookupSslPolicy(ctx *pulumi.Context, args *LookupSslPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSslPolicyResult, error) {
	var rv LookupSslPolicyResult
	err := ctx.Invoke("google-native:compute/alpha:getSslPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSslPolicyArgs struct {
	Project   *string `pulumi:"project"`
	SslPolicy string  `pulumi:"sslPolicy"`
}

type LookupSslPolicyResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
	CustomFeatures []string `pulumi:"customFeatures"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// The list of features enabled in the SSL policy.
	EnabledFeatures []string `pulumi:"enabledFeatures"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
	Fingerprint string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
	Kind string `pulumi:"kind"`
	// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
	MinTlsVersion string `pulumi:"minTlsVersion"`
	// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
	Profile string `pulumi:"profile"`
	// URL of the region where the regional SSL policy resides. This field is not applicable to global SSL policies.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// Security settings for the proxy. This field is only applicable to a global backend service with the loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	TlsSettings ServerTlsSettingsResponse `pulumi:"tlsSettings"`
	// If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
	Warnings []SslPolicyWarningsItemResponse `pulumi:"warnings"`
}

func LookupSslPolicyOutput(ctx *pulumi.Context, args LookupSslPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSslPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSslPolicyResult, error) {
			args := v.(LookupSslPolicyArgs)
			r, err := LookupSslPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupSslPolicyResultOutput)
}

type LookupSslPolicyOutputArgs struct {
	Project   pulumi.StringPtrInput `pulumi:"project"`
	SslPolicy pulumi.StringInput    `pulumi:"sslPolicy"`
}

func (LookupSslPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSslPolicyArgs)(nil)).Elem()
}

type LookupSslPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSslPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSslPolicyResult)(nil)).Elem()
}

func (o LookupSslPolicyResultOutput) ToLookupSslPolicyResultOutput() LookupSslPolicyResultOutput {
	return o
}

func (o LookupSslPolicyResultOutput) ToLookupSslPolicyResultOutputWithContext(ctx context.Context) LookupSslPolicyResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupSslPolicyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
func (o LookupSslPolicyResultOutput) CustomFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) []string { return v.CustomFeatures }).(pulumi.StringArrayOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupSslPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The list of features enabled in the SSL policy.
func (o LookupSslPolicyResultOutput) EnabledFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) []string { return v.EnabledFeatures }).(pulumi.StringArrayOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
func (o LookupSslPolicyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
func (o LookupSslPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
func (o LookupSslPolicyResultOutput) MinTlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.MinTlsVersion }).(pulumi.StringOutput)
}

// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupSslPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
func (o LookupSslPolicyResultOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.Profile }).(pulumi.StringOutput)
}

// URL of the region where the regional SSL policy resides. This field is not applicable to global SSL policies.
func (o LookupSslPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupSslPolicyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupSslPolicyResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// Security settings for the proxy. This field is only applicable to a global backend service with the loadBalancingScheme set to INTERNAL_SELF_MANAGED.
func (o LookupSslPolicyResultOutput) TlsSettings() ServerTlsSettingsResponseOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) ServerTlsSettingsResponse { return v.TlsSettings }).(ServerTlsSettingsResponseOutput)
}

// If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
func (o LookupSslPolicyResultOutput) Warnings() SslPolicyWarningsItemResponseArrayOutput {
	return o.ApplyT(func(v LookupSslPolicyResult) []SslPolicyWarningsItemResponse { return v.Warnings }).(SslPolicyWarningsItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSslPolicyResultOutput{})
}
