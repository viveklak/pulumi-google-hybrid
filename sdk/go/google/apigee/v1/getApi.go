// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an API proxy including a list of existing revisions.
func LookupApi(ctx *pulumi.Context, args *LookupApiArgs, opts ...pulumi.InvokeOption) (*LookupApiResult, error) {
	var rv LookupApiResult
	err := ctx.Invoke("google-native:apigee/v1:getApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiArgs struct {
	ApiId          string `pulumi:"apiId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupApiResult struct {
	// The id of the most recently created revision for this api proxy.
	LatestRevisionId string `pulumi:"latestRevisionId"`
	// Metadata describing the API proxy.
	MetaData GoogleCloudApigeeV1EntityMetadataResponse `pulumi:"metaData"`
	// Name of the API proxy.
	Name string `pulumi:"name"`
	// List of revisons defined for the API proxy.
	Revision []string `pulumi:"revision"`
}