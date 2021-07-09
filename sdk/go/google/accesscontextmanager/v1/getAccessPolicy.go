// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an AccessPolicy by name.
func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("google-native:accesscontextmanager/v1:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	AccessPolicyId string `pulumi:"accessPolicyId"`
}

type LookupAccessPolicyResult struct {
	// An opaque identifier for the current version of the `AccessPolicy`. This will always be a strongly validated etag, meaning that two Access Polices will be identical if and only if their etags are identical. Clients should not expect this to be in any specific format.
	Etag string `pulumi:"etag"`
	// Resource name of the `AccessPolicy`. Format: `accessPolicies/{policy_id}`
	Name string `pulumi:"name"`
	// The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Parent string `pulumi:"parent"`
	// Human readable title. Does not affect behavior.
	Title string `pulumi:"title"`
}
