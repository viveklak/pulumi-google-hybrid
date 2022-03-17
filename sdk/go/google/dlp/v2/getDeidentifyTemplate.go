// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a DeidentifyTemplate. See https://cloud.google.com/dlp/docs/creating-templates-deid to learn more.
func LookupDeidentifyTemplate(ctx *pulumi.Context, args *LookupDeidentifyTemplateArgs, opts ...pulumi.InvokeOption) (*LookupDeidentifyTemplateResult, error) {
	var rv LookupDeidentifyTemplateResult
	err := ctx.Invoke("google-native:dlp/v2:getDeidentifyTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeidentifyTemplateArgs struct {
	DeidentifyTemplateId string  `pulumi:"deidentifyTemplateId"`
	Location             string  `pulumi:"location"`
	Project              *string `pulumi:"project"`
}

type LookupDeidentifyTemplateResult struct {
	// The creation timestamp of an inspectTemplate.
	CreateTime string `pulumi:"createTime"`
	// The core content of the template.
	DeidentifyConfig GooglePrivacyDlpV2DeidentifyConfigResponse `pulumi:"deidentifyConfig"`
	// Short description (max 256 chars).
	Description string `pulumi:"description"`
	// Display name (max 256 chars).
	DisplayName string `pulumi:"displayName"`
	// The template name. The template will have one of the following formats: `projects/PROJECT_ID/deidentifyTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/deidentifyTemplates/TEMPLATE_ID`
	Name string `pulumi:"name"`
	// The last update timestamp of an inspectTemplate.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupDeidentifyTemplateOutput(ctx *pulumi.Context, args LookupDeidentifyTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupDeidentifyTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeidentifyTemplateResult, error) {
			args := v.(LookupDeidentifyTemplateArgs)
			r, err := LookupDeidentifyTemplate(ctx, &args, opts...)
			return *r, err
		}).(LookupDeidentifyTemplateResultOutput)
}

type LookupDeidentifyTemplateOutputArgs struct {
	DeidentifyTemplateId pulumi.StringInput    `pulumi:"deidentifyTemplateId"`
	Location             pulumi.StringInput    `pulumi:"location"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDeidentifyTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeidentifyTemplateArgs)(nil)).Elem()
}

type LookupDeidentifyTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupDeidentifyTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeidentifyTemplateResult)(nil)).Elem()
}

func (o LookupDeidentifyTemplateResultOutput) ToLookupDeidentifyTemplateResultOutput() LookupDeidentifyTemplateResultOutput {
	return o
}

func (o LookupDeidentifyTemplateResultOutput) ToLookupDeidentifyTemplateResultOutputWithContext(ctx context.Context) LookupDeidentifyTemplateResultOutput {
	return o
}

// The creation timestamp of an inspectTemplate.
func (o LookupDeidentifyTemplateResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidentifyTemplateResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The core content of the template.
func (o LookupDeidentifyTemplateResultOutput) DeidentifyConfig() GooglePrivacyDlpV2DeidentifyConfigResponseOutput {
	return o.ApplyT(func(v LookupDeidentifyTemplateResult) GooglePrivacyDlpV2DeidentifyConfigResponse {
		return v.DeidentifyConfig
	}).(GooglePrivacyDlpV2DeidentifyConfigResponseOutput)
}

// Short description (max 256 chars).
func (o LookupDeidentifyTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidentifyTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

// Display name (max 256 chars).
func (o LookupDeidentifyTemplateResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidentifyTemplateResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The template name. The template will have one of the following formats: `projects/PROJECT_ID/deidentifyTemplates/TEMPLATE_ID` OR `organizations/ORGANIZATION_ID/deidentifyTemplates/TEMPLATE_ID`
func (o LookupDeidentifyTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidentifyTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// The last update timestamp of an inspectTemplate.
func (o LookupDeidentifyTemplateResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeidentifyTemplateResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeidentifyTemplateResultOutput{})
}
