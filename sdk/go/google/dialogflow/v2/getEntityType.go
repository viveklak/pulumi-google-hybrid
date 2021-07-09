// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified entity type.
func LookupEntityType(ctx *pulumi.Context, args *LookupEntityTypeArgs, opts ...pulumi.InvokeOption) (*LookupEntityTypeResult, error) {
	var rv LookupEntityTypeResult
	err := ctx.Invoke("google-native:dialogflow/v2:getEntityType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityTypeArgs struct {
	EntityTypeId string  `pulumi:"entityTypeId"`
	LanguageCode *string `pulumi:"languageCode"`
	Location     string  `pulumi:"location"`
	Project      string  `pulumi:"project"`
}

type LookupEntityTypeResult struct {
	// Optional. Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode string `pulumi:"autoExpansionMode"`
	// The name of the entity type.
	DisplayName string `pulumi:"displayName"`
	// Optional. Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction bool `pulumi:"enableFuzzyExtraction"`
	// Optional. The collection of entity entries associated with the entity type.
	Entities []GoogleCloudDialogflowV2EntityTypeEntityResponse `pulumi:"entities"`
	// Indicates the kind of entity type.
	Kind string `pulumi:"kind"`
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType and EntityTypes.BatchUpdateEntityTypes methods. Format: `projects//agent/entityTypes/`.
	Name string `pulumi:"name"`
}
