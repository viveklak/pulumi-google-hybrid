// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified global PublicDelegatedPrefix resource.
func LookupGlobalPublicDelegatedPrefix(ctx *pulumi.Context, args *LookupGlobalPublicDelegatedPrefixArgs, opts ...pulumi.InvokeOption) (*LookupGlobalPublicDelegatedPrefixResult, error) {
	var rv LookupGlobalPublicDelegatedPrefixResult
	err := ctx.Invoke("google-native:compute/v1:getGlobalPublicDelegatedPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalPublicDelegatedPrefixArgs struct {
	Project               *string `pulumi:"project"`
	PublicDelegatedPrefix string  `pulumi:"publicDelegatedPrefix"`
}

type LookupGlobalPublicDelegatedPrefixResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
	Fingerprint string `pulumi:"fingerprint"`
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange string `pulumi:"ipCidrRange"`
	// If true, the prefix will be live migrated.
	IsLiveMigration bool `pulumi:"isLiveMigration"`
	// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix string `pulumi:"parentPrefix"`
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs []PublicDelegatedPrefixPublicDelegatedSubPrefixResponse `pulumi:"publicDelegatedSubPrefixs"`
	// URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned.
	Status string `pulumi:"status"`
}

func LookupGlobalPublicDelegatedPrefixOutput(ctx *pulumi.Context, args LookupGlobalPublicDelegatedPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalPublicDelegatedPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalPublicDelegatedPrefixResult, error) {
			args := v.(LookupGlobalPublicDelegatedPrefixArgs)
			r, err := LookupGlobalPublicDelegatedPrefix(ctx, &args, opts...)
			return *r, err
		}).(LookupGlobalPublicDelegatedPrefixResultOutput)
}

type LookupGlobalPublicDelegatedPrefixOutputArgs struct {
	Project               pulumi.StringPtrInput `pulumi:"project"`
	PublicDelegatedPrefix pulumi.StringInput    `pulumi:"publicDelegatedPrefix"`
}

func (LookupGlobalPublicDelegatedPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalPublicDelegatedPrefixArgs)(nil)).Elem()
}

type LookupGlobalPublicDelegatedPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalPublicDelegatedPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalPublicDelegatedPrefixResult)(nil)).Elem()
}

func (o LookupGlobalPublicDelegatedPrefixResultOutput) ToLookupGlobalPublicDelegatedPrefixResultOutput() LookupGlobalPublicDelegatedPrefixResultOutput {
	return o
}

func (o LookupGlobalPublicDelegatedPrefixResultOutput) ToLookupGlobalPublicDelegatedPrefixResultOutputWithContext(ctx context.Context) LookupGlobalPublicDelegatedPrefixResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.IpCidrRange }).(pulumi.StringOutput)
}

// If true, the prefix will be live migrated.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) IsLiveMigration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) bool { return v.IsLiveMigration }).(pulumi.BoolOutput)
}

// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) ParentPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.ParentPrefix }).(pulumi.StringOutput)
}

// The list of sub public delegated prefixes that exist for this public delegated prefix.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) PublicDelegatedSubPrefixs() PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) []PublicDelegatedPrefixPublicDelegatedSubPrefixResponse {
		return v.PublicDelegatedSubPrefixs
	}).(PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput)
}

// URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned.
func (o LookupGlobalPublicDelegatedPrefixResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalPublicDelegatedPrefixResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalPublicDelegatedPrefixResultOutput{})
}
