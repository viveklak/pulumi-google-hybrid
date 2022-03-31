// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified PublicDelegatedPrefix resource in the given region.
func LookupPublicDelegatedPrefix(ctx *pulumi.Context, args *LookupPublicDelegatedPrefixArgs, opts ...pulumi.InvokeOption) (*LookupPublicDelegatedPrefixResult, error) {
	var rv LookupPublicDelegatedPrefixResult
	err := ctx.Invoke("google-native:compute/alpha:getPublicDelegatedPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicDelegatedPrefixArgs struct {
	Project               *string `pulumi:"project"`
	PublicDelegatedPrefix string  `pulumi:"publicDelegatedPrefix"`
	Region                string  `pulumi:"region"`
}

type LookupPublicDelegatedPrefixResult struct {
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
	// Server-defined URL with id for the resource.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned.
	Status string `pulumi:"status"`
}

func LookupPublicDelegatedPrefixOutput(ctx *pulumi.Context, args LookupPublicDelegatedPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupPublicDelegatedPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicDelegatedPrefixResult, error) {
			args := v.(LookupPublicDelegatedPrefixArgs)
			r, err := LookupPublicDelegatedPrefix(ctx, &args, opts...)
			return *r, err
		}).(LookupPublicDelegatedPrefixResultOutput)
}

type LookupPublicDelegatedPrefixOutputArgs struct {
	Project               pulumi.StringPtrInput `pulumi:"project"`
	PublicDelegatedPrefix pulumi.StringInput    `pulumi:"publicDelegatedPrefix"`
	Region                pulumi.StringInput    `pulumi:"region"`
}

func (LookupPublicDelegatedPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicDelegatedPrefixArgs)(nil)).Elem()
}

type LookupPublicDelegatedPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupPublicDelegatedPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicDelegatedPrefixResult)(nil)).Elem()
}

func (o LookupPublicDelegatedPrefixResultOutput) ToLookupPublicDelegatedPrefixResultOutput() LookupPublicDelegatedPrefixResultOutput {
	return o
}

func (o LookupPublicDelegatedPrefixResultOutput) ToLookupPublicDelegatedPrefixResultOutputWithContext(ctx context.Context) LookupPublicDelegatedPrefixResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupPublicDelegatedPrefixResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupPublicDelegatedPrefixResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
func (o LookupPublicDelegatedPrefixResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
func (o LookupPublicDelegatedPrefixResultOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.IpCidrRange }).(pulumi.StringOutput)
}

// If true, the prefix will be live migrated.
func (o LookupPublicDelegatedPrefixResultOutput) IsLiveMigration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) bool { return v.IsLiveMigration }).(pulumi.BoolOutput)
}

// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
func (o LookupPublicDelegatedPrefixResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupPublicDelegatedPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
func (o LookupPublicDelegatedPrefixResultOutput) ParentPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.ParentPrefix }).(pulumi.StringOutput)
}

// The list of sub public delegated prefixes that exist for this public delegated prefix.
func (o LookupPublicDelegatedPrefixResultOutput) PublicDelegatedSubPrefixs() PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) []PublicDelegatedPrefixPublicDelegatedSubPrefixResponse {
		return v.PublicDelegatedSubPrefixs
	}).(PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput)
}

// URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o LookupPublicDelegatedPrefixResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupPublicDelegatedPrefixResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL with id for the resource.
func (o LookupPublicDelegatedPrefixResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned.
func (o LookupPublicDelegatedPrefixResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicDelegatedPrefixResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicDelegatedPrefixResultOutput{})
}
