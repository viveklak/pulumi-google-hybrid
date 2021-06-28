// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified global PublicDelegatedPrefix resource.
func LookupGlobalPublicDelegatedPrefix(ctx *pulumi.Context, args *LookupGlobalPublicDelegatedPrefixArgs, opts ...pulumi.InvokeOption) (*LookupGlobalPublicDelegatedPrefixResult, error) {
	var rv LookupGlobalPublicDelegatedPrefixResult
	err := ctx.Invoke("google-native:compute/alpha:getGlobalPublicDelegatedPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalPublicDelegatedPrefixArgs struct {
	Project               string `pulumi:"project"`
	PublicDelegatedPrefix string `pulumi:"publicDelegatedPrefix"`
}

type LookupGlobalPublicDelegatedPrefixResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet.
	//
	// To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
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
	// The status of the public delegated prefix.
	Status string `pulumi:"status"`
}
