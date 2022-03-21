// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a big query export.
func LookupFolderBigQueryExport(ctx *pulumi.Context, args *LookupFolderBigQueryExportArgs, opts ...pulumi.InvokeOption) (*LookupFolderBigQueryExportResult, error) {
	var rv LookupFolderBigQueryExportResult
	err := ctx.Invoke("google-native:securitycenter/v1:getFolderBigQueryExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderBigQueryExportArgs struct {
	BigQueryExportId string `pulumi:"bigQueryExportId"`
	FolderId         string `pulumi:"folderId"`
}

type LookupFolderBigQueryExportResult struct {
	// The time at which the big query export was created. This field is set by the server and will be ignored if provided on export on creation.
	CreateTime string `pulumi:"createTime"`
	// The dataset to write findings' updates to. Its format is "projects/[project_id]/datasets/[bigquery_dataset_id]". BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset string `pulumi:"dataset"`
	// The description of the export (max of 1024 characters).
	Description string `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or more restrictions combined via logical operators `AND` and `OR`. Parentheses are supported, and `OR` has higher precedence than `AND`. Restrictions have the form ` ` and may have a `-` character in front of them to indicate negation. The fields map to those defined in the corresponding resource. The supported operators are: * `=` for all value types. * `>`, `<`, `>=`, `<=` for integer values. * `:`, meaning substring matching, for strings. The supported value types are: * string literals in quotes. * integer literals without quotes. * boolean literals `true` and `false` without quotes.
	Filter string `pulumi:"filter"`
	// Email address of the user who last edited the big query export. This field is set by the server and will be ignored if provided on export creation or update.
	MostRecentEditor string `pulumi:"mostRecentEditor"`
	// The relative resource name of this export. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name. Example format: "organizations/{organization_id}/bigQueryExports/{export_id}" Example format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format: "projects/{project_id}/bigQueryExports/{export_id}" This field is provided in responses, and is ignored when provided in create requests.
	Name string `pulumi:"name"`
	// The service account that needs permission to create table, upload data to the big query dataset.
	Principal string `pulumi:"principal"`
	// The most recent time at which the big export was updated. This field is set by the server and will be ignored if provided on export creation or update.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupFolderBigQueryExportOutput(ctx *pulumi.Context, args LookupFolderBigQueryExportOutputArgs, opts ...pulumi.InvokeOption) LookupFolderBigQueryExportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderBigQueryExportResult, error) {
			args := v.(LookupFolderBigQueryExportArgs)
			r, err := LookupFolderBigQueryExport(ctx, &args, opts...)
			return *r, err
		}).(LookupFolderBigQueryExportResultOutput)
}

type LookupFolderBigQueryExportOutputArgs struct {
	BigQueryExportId pulumi.StringInput `pulumi:"bigQueryExportId"`
	FolderId         pulumi.StringInput `pulumi:"folderId"`
}

func (LookupFolderBigQueryExportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderBigQueryExportArgs)(nil)).Elem()
}

type LookupFolderBigQueryExportResultOutput struct{ *pulumi.OutputState }

func (LookupFolderBigQueryExportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderBigQueryExportResult)(nil)).Elem()
}

func (o LookupFolderBigQueryExportResultOutput) ToLookupFolderBigQueryExportResultOutput() LookupFolderBigQueryExportResultOutput {
	return o
}

func (o LookupFolderBigQueryExportResultOutput) ToLookupFolderBigQueryExportResultOutputWithContext(ctx context.Context) LookupFolderBigQueryExportResultOutput {
	return o
}

// The time at which the big query export was created. This field is set by the server and will be ignored if provided on export on creation.
func (o LookupFolderBigQueryExportResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The dataset to write findings' updates to. Its format is "projects/[project_id]/datasets/[bigquery_dataset_id]". BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
func (o LookupFolderBigQueryExportResultOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.Dataset }).(pulumi.StringOutput)
}

// The description of the export (max of 1024 characters).
func (o LookupFolderBigQueryExportResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.Description }).(pulumi.StringOutput)
}

// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or more restrictions combined via logical operators `AND` and `OR`. Parentheses are supported, and `OR` has higher precedence than `AND`. Restrictions have the form ` ` and may have a `-` character in front of them to indicate negation. The fields map to those defined in the corresponding resource. The supported operators are: * `=` for all value types. * `>`, `<`, `>=`, `<=` for integer values. * `:`, meaning substring matching, for strings. The supported value types are: * string literals in quotes. * integer literals without quotes. * boolean literals `true` and `false` without quotes.
func (o LookupFolderBigQueryExportResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.Filter }).(pulumi.StringOutput)
}

// Email address of the user who last edited the big query export. This field is set by the server and will be ignored if provided on export creation or update.
func (o LookupFolderBigQueryExportResultOutput) MostRecentEditor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.MostRecentEditor }).(pulumi.StringOutput)
}

// The relative resource name of this export. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name. Example format: "organizations/{organization_id}/bigQueryExports/{export_id}" Example format: "folders/{folder_id}/bigQueryExports/{export_id}" Example format: "projects/{project_id}/bigQueryExports/{export_id}" This field is provided in responses, and is ignored when provided in create requests.
func (o LookupFolderBigQueryExportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.Name }).(pulumi.StringOutput)
}

// The service account that needs permission to create table, upload data to the big query dataset.
func (o LookupFolderBigQueryExportResultOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.Principal }).(pulumi.StringOutput)
}

// The most recent time at which the big export was updated. This field is set by the server and will be ignored if provided on export creation or update.
func (o LookupFolderBigQueryExportResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderBigQueryExportResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderBigQueryExportResultOutput{})
}
