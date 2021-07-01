// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new health dataset. Results are returned through the Operation interface which returns either an `Operation.response` which contains a Dataset or `Operation.error`. The metadata field type is OperationMetadata.
type Dataset struct {
	pulumi.CustomResourceState

	// Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Dataset
	err := ctx.RegisterResource("google-native:healthcare/v1beta1:Dataset", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:healthcare/v1beta1:Dataset", name, id, state, &resource, opts...)
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
	DatasetId *string `pulumi:"datasetId"`
	Location  string  `pulumi:"location"`
	// Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	DatasetId pulumi.StringPtrInput
	Location  pulumi.StringInput
	// Resource name of the dataset, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources, such as HL7 messages, where no explicit timezone is specified.
	TimeZone pulumi.StringPtrInput
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
	return reflect.TypeOf((*Dataset)(nil))
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

type DatasetOutput struct {
	*pulumi.OutputState
}

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dataset)(nil))
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetOutput{})
}
