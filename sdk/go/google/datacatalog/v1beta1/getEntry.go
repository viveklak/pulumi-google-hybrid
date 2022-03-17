// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an entry.
func LookupEntry(ctx *pulumi.Context, args *LookupEntryArgs, opts ...pulumi.InvokeOption) (*LookupEntryResult, error) {
	var rv LookupEntryResult
	err := ctx.Invoke("google-native:datacatalog/v1beta1:getEntry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEntryArgs struct {
	EntryGroupId string  `pulumi:"entryGroupId"`
	EntryId      string  `pulumi:"entryId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupEntryResult struct {
	// Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
	BigqueryDateShardedSpec GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponse `pulumi:"bigqueryDateShardedSpec"`
	// Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
	BigqueryTableSpec GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponse `pulumi:"bigqueryTableSpec"`
	// Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
	Description string `pulumi:"description"`
	// Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
	DisplayName string `pulumi:"displayName"`
	// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
	GcsFilesetSpec GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponse `pulumi:"gcsFilesetSpec"`
	// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
	IntegratedSystem string `pulumi:"integratedSystem"`
	// The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
	LinkedResource string `pulumi:"linkedResource"`
	// The Data Catalog resource name of the entry in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id}/entries/{entry_id} Note that this Entry and its child resources may not actually be stored in the location in this name.
	Name string `pulumi:"name"`
	// Schema of the entry. An entry might not have any schema attached to it.
	Schema GoogleCloudDatacatalogV1beta1SchemaResponse `pulumi:"schema"`
	// Timestamps about the underlying resource, not about this Data Catalog entry. Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty timestamp.
	SourceSystemTimestamps GoogleCloudDatacatalogV1beta1SystemTimestampsResponse `pulumi:"sourceSystemTimestamps"`
	// The type of the entry. Only used for Entries with types in the EntryType enum.
	Type string `pulumi:"type"`
	// Statistics on the usage level of the resource.
	UsageSignal GoogleCloudDatacatalogV1beta1UsageSignalResponse `pulumi:"usageSignal"`
	// This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem string `pulumi:"userSpecifiedSystem"`
	// Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
	UserSpecifiedType string `pulumi:"userSpecifiedType"`
}

func LookupEntryOutput(ctx *pulumi.Context, args LookupEntryOutputArgs, opts ...pulumi.InvokeOption) LookupEntryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEntryResult, error) {
			args := v.(LookupEntryArgs)
			r, err := LookupEntry(ctx, &args, opts...)
			return *r, err
		}).(LookupEntryResultOutput)
}

type LookupEntryOutputArgs struct {
	EntryGroupId pulumi.StringInput    `pulumi:"entryGroupId"`
	EntryId      pulumi.StringInput    `pulumi:"entryId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupEntryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntryArgs)(nil)).Elem()
}

type LookupEntryResultOutput struct{ *pulumi.OutputState }

func (LookupEntryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEntryResult)(nil)).Elem()
}

func (o LookupEntryResultOutput) ToLookupEntryResultOutput() LookupEntryResultOutput {
	return o
}

func (o LookupEntryResultOutput) ToLookupEntryResultOutputWithContext(ctx context.Context) LookupEntryResultOutput {
	return o
}

// Specification for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`. Context: https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
func (o LookupEntryResultOutput) BigqueryDateShardedSpec() GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponseOutput {
	return o.ApplyT(func(v LookupEntryResult) GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponse {
		return v.BigqueryDateShardedSpec
	}).(GoogleCloudDatacatalogV1beta1BigQueryDateShardedSpecResponseOutput)
}

// Specification that applies to a BigQuery table. This is only valid on entries of type `TABLE`.
func (o LookupEntryResultOutput) BigqueryTableSpec() GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponseOutput {
	return o.ApplyT(func(v LookupEntryResult) GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponse {
		return v.BigqueryTableSpec
	}).(GoogleCloudDatacatalogV1beta1BigQueryTableSpecResponseOutput)
}

// Entry description, which can consist of several sentences or paragraphs that describe entry contents. Default value is an empty string.
func (o LookupEntryResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.Description }).(pulumi.StringOutput)
}

// Display information such as title and description. A short name to identify the entry, for example, "Analytics Data - Jan 2011". Default value is an empty string.
func (o LookupEntryResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
func (o LookupEntryResultOutput) GcsFilesetSpec() GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponseOutput {
	return o.ApplyT(func(v LookupEntryResult) GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponse { return v.GcsFilesetSpec }).(GoogleCloudDatacatalogV1beta1GcsFilesetSpecResponseOutput)
}

// This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
func (o LookupEntryResultOutput) IntegratedSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.IntegratedSystem }).(pulumi.StringOutput)
}

// The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [full name of the resource](https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty string.
func (o LookupEntryResultOutput) LinkedResource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.LinkedResource }).(pulumi.StringOutput)
}

// The Data Catalog resource name of the entry in URL format. Example: * projects/{project_id}/locations/{location}/entryGroups/{entry_group_id}/entries/{entry_id} Note that this Entry and its child resources may not actually be stored in the location in this name.
func (o LookupEntryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Schema of the entry. An entry might not have any schema attached to it.
func (o LookupEntryResultOutput) Schema() GoogleCloudDatacatalogV1beta1SchemaResponseOutput {
	return o.ApplyT(func(v LookupEntryResult) GoogleCloudDatacatalogV1beta1SchemaResponse { return v.Schema }).(GoogleCloudDatacatalogV1beta1SchemaResponseOutput)
}

// Timestamps about the underlying resource, not about this Data Catalog entry. Output only when Entry is of type in the EntryType enum. For entries with user_specified_type, this field is optional and defaults to an empty timestamp.
func (o LookupEntryResultOutput) SourceSystemTimestamps() GoogleCloudDatacatalogV1beta1SystemTimestampsResponseOutput {
	return o.ApplyT(func(v LookupEntryResult) GoogleCloudDatacatalogV1beta1SystemTimestampsResponse {
		return v.SourceSystemTimestamps
	}).(GoogleCloudDatacatalogV1beta1SystemTimestampsResponseOutput)
}

// The type of the entry. Only used for Entries with types in the EntryType enum.
func (o LookupEntryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.Type }).(pulumi.StringOutput)
}

// Statistics on the usage level of the resource.
func (o LookupEntryResultOutput) UsageSignal() GoogleCloudDatacatalogV1beta1UsageSignalResponseOutput {
	return o.ApplyT(func(v LookupEntryResult) GoogleCloudDatacatalogV1beta1UsageSignalResponse { return v.UsageSignal }).(GoogleCloudDatacatalogV1beta1UsageSignalResponseOutput)
}

// This field indicates the entry's source system that Data Catalog does not integrate with. `user_specified_system` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
func (o LookupEntryResultOutput) UserSpecifiedSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.UserSpecifiedSystem }).(pulumi.StringOutput)
}

// Entry type if it does not fit any of the input-allowed values listed in `EntryType` enum above. When creating an entry, users should check the enum values first, if nothing matches the entry to be created, then provide a custom value, for example "my_special_type". `user_specified_type` strings must begin with a letter or underscore and can only contain letters, numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long. Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use `user_specified_type`.
func (o LookupEntryResultOutput) UserSpecifiedType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEntryResult) string { return v.UserSpecifiedType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEntryResultOutput{})
}
