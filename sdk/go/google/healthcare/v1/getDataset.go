// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets any metadata associated with a dataset.
func LookupDataset(ctx *pulumi.Context, args *LookupDatasetArgs, opts ...pulumi.InvokeOption) (*LookupDatasetResult, error) {
	var rv LookupDatasetResult
	err := ctx.Invoke("google-native:healthcare/v1:getDataset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatasetArgs struct {
	DatasetId string  `pulumi:"datasetId"`
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
}

type LookupDatasetResult struct {
	// Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
	Name string `pulumi:"name"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
	TimeZone string `pulumi:"timeZone"`
}

func LookupDatasetOutput(ctx *pulumi.Context, args LookupDatasetOutputArgs, opts ...pulumi.InvokeOption) LookupDatasetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatasetResult, error) {
			args := v.(LookupDatasetArgs)
			r, err := LookupDataset(ctx, &args, opts...)
			return *r, err
		}).(LookupDatasetResultOutput)
}

type LookupDatasetOutputArgs struct {
	DatasetId pulumi.StringInput    `pulumi:"datasetId"`
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDatasetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetArgs)(nil)).Elem()
}

type LookupDatasetResultOutput struct{ *pulumi.OutputState }

func (LookupDatasetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatasetResult)(nil)).Elem()
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutput() LookupDatasetResultOutput {
	return o
}

func (o LookupDatasetResultOutput) ToLookupDatasetResultOutputWithContext(ctx context.Context) LookupDatasetResultOutput {
	return o
}

// Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
func (o LookupDatasetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
func (o LookupDatasetResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatasetResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatasetResultOutput{})
}
