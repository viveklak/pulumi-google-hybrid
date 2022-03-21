// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates dataset. If success return a Dataset resource.
// Auto-naming is currently not supported for this resource.
type Dataset struct {
	pulumi.CustomResourceState

	// The names of any related resources that are blocking changes to the dataset.
	BlockingResources pulumi.StringArrayOutput `pulumi:"blockingResources"`
	// Time the dataset is created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The number of data items in the dataset.
	DataItemCount pulumi.StringOutput `pulumi:"dataItemCount"`
	// Optional. User-provided description of the annotation specification set. The description can be up to 10000 characters long.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the dataset. Maximum of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// This is populated with the original input configs where ImportData is called. It is available only after the clients import data to this dataset.
	InputConfigs GoogleCloudDatalabelingV1beta1InputConfigResponseArrayOutput `pulumi:"inputConfigs"`
	// Last time that the Dataset is migrated to AI Platform V2. If any of the AnnotatedDataset is migrated, the last_migration_time in Dataset is also updated.
	LastMigrateTime pulumi.StringOutput `pulumi:"lastMigrateTime"`
	// Dataset resource name, format is: projects/{project_id}/datasets/{dataset_id}
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Dataset
	err := ctx.RegisterResource("google-native:datalabeling/v1beta1:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("google-native:datalabeling/v1beta1:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
}

type DatasetState struct {
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// Optional. User-provided description of the annotation specification set. The description can be up to 10000 characters long.
	Description *string `pulumi:"description"`
	// The display name of the dataset. Maximum of 64 characters.
	DisplayName string `pulumi:"displayName"`
	// Last time that the Dataset is migrated to AI Platform V2. If any of the AnnotatedDataset is migrated, the last_migration_time in Dataset is also updated.
	LastMigrateTime *string `pulumi:"lastMigrateTime"`
	Project         *string `pulumi:"project"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// Optional. User-provided description of the annotation specification set. The description can be up to 10000 characters long.
	Description pulumi.StringPtrInput
	// The display name of the dataset. Maximum of 64 characters.
	DisplayName pulumi.StringInput
	// Last time that the Dataset is migrated to AI Platform V2. If any of the AnnotatedDataset is migrated, the last_migration_time in Dataset is also updated.
	LastMigrateTime pulumi.StringPtrInput
	Project         pulumi.StringPtrInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetInput)(nil)).Elem(), &Dataset{})
	pulumi.RegisterOutputType(DatasetOutput{})
}
