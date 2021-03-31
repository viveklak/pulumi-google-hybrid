// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type DatasetIamPolicy struct {
	pulumi.CustomResourceState
}

// NewDatasetIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamPolicy(ctx *pulumi.Context,
	name string, args *DatasetIamPolicyArgs, opts ...pulumi.ResourceOption) (*DatasetIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetsId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetsId'")
	}
	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource DatasetIamPolicy
	err := ctx.RegisterResource("google-cloud:healthcare/v1beta1:DatasetIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetIamPolicy gets an existing DatasetIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetIamPolicyState, opts ...pulumi.ResourceOption) (*DatasetIamPolicy, error) {
	var resource DatasetIamPolicy
	err := ctx.ReadResource("google-cloud:healthcare/v1beta1:DatasetIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetIamPolicy resources.
type datasetIamPolicyState struct {
}

type DatasetIamPolicyState struct {
}

func (DatasetIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamPolicyState)(nil)).Elem()
}

type datasetIamPolicyArgs struct {
	DatasetsId  string `pulumi:"datasetsId"`
	LocationsId string `pulumi:"locationsId"`
	// REQUIRED: The complete policy to be applied to the `resource`. The size of the policy is limited to a few 10s of KB. An empty policy is a valid policy but certain Cloud Platform services (such as Projects) might reject them.
	Policy     *Policy `pulumi:"policy"`
	ProjectsId string  `pulumi:"projectsId"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
}

// The set of arguments for constructing a DatasetIamPolicy resource.
type DatasetIamPolicyArgs struct {
	DatasetsId  pulumi.StringInput
	LocationsId pulumi.StringInput
	// REQUIRED: The complete policy to be applied to the `resource`. The size of the policy is limited to a few 10s of KB. An empty policy is a valid policy but certain Cloud Platform services (such as Projects) might reject them.
	Policy     PolicyPtrInput
	ProjectsId pulumi.StringInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
}

func (DatasetIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetIamPolicyArgs)(nil)).Elem()
}

type DatasetIamPolicyInput interface {
	pulumi.Input

	ToDatasetIamPolicyOutput() DatasetIamPolicyOutput
	ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput
}

func (*DatasetIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamPolicy)(nil))
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyOutput() DatasetIamPolicyOutput {
	return i.ToDatasetIamPolicyOutputWithContext(context.Background())
}

func (i *DatasetIamPolicy) ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetIamPolicyOutput)
}

type DatasetIamPolicyOutput struct {
	*pulumi.OutputState
}

func (DatasetIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetIamPolicy)(nil))
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyOutput() DatasetIamPolicyOutput {
	return o
}

func (o DatasetIamPolicyOutput) ToDatasetIamPolicyOutputWithContext(ctx context.Context) DatasetIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetIamPolicyOutput{})
}
