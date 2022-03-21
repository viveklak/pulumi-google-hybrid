// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the specified entity type.
func LookupEntityType(ctx *pulumi.Context, args *LookupEntityTypeArgs, opts ...pulumi.InvokeOption) (*LookupEntityTypeResult, error) {
	var rv LookupEntityTypeResult
	err := ctx.Invoke("google-native:dialogflow/v3beta1:getEntityType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntityTypeArgs struct {
	AgentId      string  `pulumi:"agentId"`
	EntityTypeId string  `pulumi:"entityTypeId"`
	LanguageCode *string `pulumi:"languageCode"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupEntityTypeResult struct {
	// Indicates whether the entity type can be automatically expanded.
	AutoExpansionMode string `pulumi:"autoExpansionMode"`
	// The human-readable name of the entity type, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction bool `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	Entities []GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse `pulumi:"entities"`
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	ExcludedPhrases []GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponse `pulumi:"excludedPhrases"`
	// Indicates the kind of entity type.
	Kind string `pulumi:"kind"`
	// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
	Name string `pulumi:"name"`
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
	Redact bool `pulumi:"redact"`
}

func LookupEntityTypeOutput(ctx *pulumi.Context, args LookupEntityTypeOutputArgs, opts ...pulumi.InvokeOption) LookupEntityTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntityTypeResult, error) {
			args := v.(LookupEntityTypeArgs)
			r, err := LookupEntityType(ctx, &args, opts...)
			return *r, err
		}).(LookupEntityTypeResultOutput)
}

type LookupEntityTypeOutputArgs struct {
	AgentId      pulumi.StringInput    `pulumi:"agentId"`
	EntityTypeId pulumi.StringInput    `pulumi:"entityTypeId"`
	LanguageCode pulumi.StringPtrInput `pulumi:"languageCode"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupEntityTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityTypeArgs)(nil)).Elem()
}

type LookupEntityTypeResultOutput struct{ *pulumi.OutputState }

func (LookupEntityTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntityTypeResult)(nil)).Elem()
}

func (o LookupEntityTypeResultOutput) ToLookupEntityTypeResultOutput() LookupEntityTypeResultOutput {
	return o
}

func (o LookupEntityTypeResultOutput) ToLookupEntityTypeResultOutputWithContext(ctx context.Context) LookupEntityTypeResultOutput {
	return o
}

// Indicates whether the entity type can be automatically expanded.
func (o LookupEntityTypeResultOutput) AutoExpansionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) string { return v.AutoExpansionMode }).(pulumi.StringOutput)
}

// The human-readable name of the entity type, unique within the agent.
func (o LookupEntityTypeResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Enables fuzzy entity extraction during classification.
func (o LookupEntityTypeResultOutput) EnableFuzzyExtraction() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) bool { return v.EnableFuzzyExtraction }).(pulumi.BoolOutput)
}

// The collection of entity entries associated with the entity type.
func (o LookupEntityTypeResultOutput) Entities() GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) []GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse {
		return v.Entities
	}).(GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayOutput)
}

// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry `giant`(an adjective), you might consider adding `giants`(a noun) as an exclusion. If the kind of entity type is `KIND_MAP`, then the phrases specified by entities and excluded phrases should be mutually exclusive.
func (o LookupEntityTypeResultOutput) ExcludedPhrases() GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponseArrayOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) []GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponse {
		return v.ExcludedPhrases
	}).(GoogleCloudDialogflowCxV3beta1EntityTypeExcludedPhraseResponseArrayOutput)
}

// Indicates the kind of entity type.
func (o LookupEntityTypeResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The unique identifier of the entity type. Required for EntityTypes.UpdateEntityType. Format: `projects//locations//agents//entityTypes/`.
func (o LookupEntityTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name during logging.
func (o LookupEntityTypeResultOutput) Redact() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEntityTypeResult) bool { return v.Redact }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntityTypeResultOutput{})
}
