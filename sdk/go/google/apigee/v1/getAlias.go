// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an alias.
func LookupAlias(ctx *pulumi.Context, args *LookupAliasArgs, opts ...pulumi.InvokeOption) (*LookupAliasResult, error) {
	var rv LookupAliasResult
	err := ctx.Invoke("google-native:apigee/v1:getAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAliasArgs struct {
	AliasId        string `pulumi:"aliasId"`
	EnvironmentId  string `pulumi:"environmentId"`
	KeystoreId     string `pulumi:"keystoreId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupAliasResult struct {
	// Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
	Alias string `pulumi:"alias"`
	// Chain of certificates under this alias.
	CertsInfo GoogleCloudApigeeV1CertificateResponse `pulumi:"certsInfo"`
	// Type of alias.
	Type string `pulumi:"type"`
}