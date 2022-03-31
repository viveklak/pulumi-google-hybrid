// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a glossary. Returns NOT_FOUND, if the glossary doesn't exist.
func LookupGlossary(ctx *pulumi.Context, args *LookupGlossaryArgs, opts ...pulumi.InvokeOption) (*LookupGlossaryResult, error) {
	var rv LookupGlossaryResult
	err := ctx.Invoke("google-native:translate/v3beta1:getGlossary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlossaryArgs struct {
	GlossaryId string  `pulumi:"glossaryId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupGlossaryResult struct {
	// When the glossary creation was finished.
	EndTime string `pulumi:"endTime"`
	// The number of entries defined in the glossary.
	EntryCount int `pulumi:"entryCount"`
	// Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
	InputConfig GlossaryInputConfigResponse `pulumi:"inputConfig"`
	// Used with equivalent term set glossaries.
	LanguageCodesSet LanguageCodesSetResponse `pulumi:"languageCodesSet"`
	// Used with unidirectional glossaries.
	LanguagePair LanguageCodePairResponse `pulumi:"languagePair"`
	// The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
	Name string `pulumi:"name"`
	// When CreateGlossary was called.
	SubmitTime string `pulumi:"submitTime"`
}

func LookupGlossaryOutput(ctx *pulumi.Context, args LookupGlossaryOutputArgs, opts ...pulumi.InvokeOption) LookupGlossaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlossaryResult, error) {
			args := v.(LookupGlossaryArgs)
			r, err := LookupGlossary(ctx, &args, opts...)
			return *r, err
		}).(LookupGlossaryResultOutput)
}

type LookupGlossaryOutputArgs struct {
	GlossaryId pulumi.StringInput    `pulumi:"glossaryId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGlossaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlossaryArgs)(nil)).Elem()
}

type LookupGlossaryResultOutput struct{ *pulumi.OutputState }

func (LookupGlossaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlossaryResult)(nil)).Elem()
}

func (o LookupGlossaryResultOutput) ToLookupGlossaryResultOutput() LookupGlossaryResultOutput {
	return o
}

func (o LookupGlossaryResultOutput) ToLookupGlossaryResultOutputWithContext(ctx context.Context) LookupGlossaryResultOutput {
	return o
}

// When the glossary creation was finished.
func (o LookupGlossaryResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlossaryResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The number of entries defined in the glossary.
func (o LookupGlossaryResultOutput) EntryCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGlossaryResult) int { return v.EntryCount }).(pulumi.IntOutput)
}

// Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
func (o LookupGlossaryResultOutput) InputConfig() GlossaryInputConfigResponseOutput {
	return o.ApplyT(func(v LookupGlossaryResult) GlossaryInputConfigResponse { return v.InputConfig }).(GlossaryInputConfigResponseOutput)
}

// Used with equivalent term set glossaries.
func (o LookupGlossaryResultOutput) LanguageCodesSet() LanguageCodesSetResponseOutput {
	return o.ApplyT(func(v LookupGlossaryResult) LanguageCodesSetResponse { return v.LanguageCodesSet }).(LanguageCodesSetResponseOutput)
}

// Used with unidirectional glossaries.
func (o LookupGlossaryResultOutput) LanguagePair() LanguageCodePairResponseOutput {
	return o.ApplyT(func(v LookupGlossaryResult) LanguageCodePairResponse { return v.LanguagePair }).(LanguageCodePairResponseOutput)
}

// The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
func (o LookupGlossaryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlossaryResult) string { return v.Name }).(pulumi.StringOutput)
}

// When CreateGlossary was called.
func (o LookupGlossaryResultOutput) SubmitTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlossaryResult) string { return v.SubmitTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlossaryResultOutput{})
}
