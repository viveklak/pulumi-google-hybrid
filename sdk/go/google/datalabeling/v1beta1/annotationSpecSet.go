// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an annotation spec set by providing a set of labels.
// Auto-naming is currently not supported for this resource.
type AnnotationSpecSet struct {
	pulumi.CustomResourceState

	// The array of AnnotationSpecs that you define when you create the AnnotationSpecSet. These are the possible labels for the labeling task.
	AnnotationSpecs GoogleCloudDatalabelingV1beta1AnnotationSpecResponseArrayOutput `pulumi:"annotationSpecs"`
	// The names of any related resources that are blocking changes to the annotation spec set.
	BlockingResources pulumi.StringArrayOutput `pulumi:"blockingResources"`
	// Optional. User-provided description of the annotation specification set. The description can be up to 10,000 characters long.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name for AnnotationSpecSet that you define when you create it. Maximum of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The AnnotationSpecSet resource name in the following format: "projects/{project_id}/annotationSpecSets/{annotation_spec_set_id}"
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAnnotationSpecSet registers a new resource with the given unique name, arguments, and options.
func NewAnnotationSpecSet(ctx *pulumi.Context,
	name string, args *AnnotationSpecSetArgs, opts ...pulumi.ResourceOption) (*AnnotationSpecSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnnotationSpecs == nil {
		return nil, errors.New("invalid value for required argument 'AnnotationSpecs'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource AnnotationSpecSet
	err := ctx.RegisterResource("google-native:datalabeling/v1beta1:AnnotationSpecSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnnotationSpecSet gets an existing AnnotationSpecSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnnotationSpecSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnnotationSpecSetState, opts ...pulumi.ResourceOption) (*AnnotationSpecSet, error) {
	var resource AnnotationSpecSet
	err := ctx.ReadResource("google-native:datalabeling/v1beta1:AnnotationSpecSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnnotationSpecSet resources.
type annotationSpecSetState struct {
}

type AnnotationSpecSetState struct {
}

func (AnnotationSpecSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*annotationSpecSetState)(nil)).Elem()
}

type annotationSpecSetArgs struct {
	// The array of AnnotationSpecs that you define when you create the AnnotationSpecSet. These are the possible labels for the labeling task.
	AnnotationSpecs []GoogleCloudDatalabelingV1beta1AnnotationSpec `pulumi:"annotationSpecs"`
	// Optional. User-provided description of the annotation specification set. The description can be up to 10,000 characters long.
	Description *string `pulumi:"description"`
	// The display name for AnnotationSpecSet that you define when you create it. Maximum of 64 characters.
	DisplayName string  `pulumi:"displayName"`
	Project     *string `pulumi:"project"`
}

// The set of arguments for constructing a AnnotationSpecSet resource.
type AnnotationSpecSetArgs struct {
	// The array of AnnotationSpecs that you define when you create the AnnotationSpecSet. These are the possible labels for the labeling task.
	AnnotationSpecs GoogleCloudDatalabelingV1beta1AnnotationSpecArrayInput
	// Optional. User-provided description of the annotation specification set. The description can be up to 10,000 characters long.
	Description pulumi.StringPtrInput
	// The display name for AnnotationSpecSet that you define when you create it. Maximum of 64 characters.
	DisplayName pulumi.StringInput
	Project     pulumi.StringPtrInput
}

func (AnnotationSpecSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*annotationSpecSetArgs)(nil)).Elem()
}

type AnnotationSpecSetInput interface {
	pulumi.Input

	ToAnnotationSpecSetOutput() AnnotationSpecSetOutput
	ToAnnotationSpecSetOutputWithContext(ctx context.Context) AnnotationSpecSetOutput
}

func (*AnnotationSpecSet) ElementType() reflect.Type {
	return reflect.TypeOf((**AnnotationSpecSet)(nil)).Elem()
}

func (i *AnnotationSpecSet) ToAnnotationSpecSetOutput() AnnotationSpecSetOutput {
	return i.ToAnnotationSpecSetOutputWithContext(context.Background())
}

func (i *AnnotationSpecSet) ToAnnotationSpecSetOutputWithContext(ctx context.Context) AnnotationSpecSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnnotationSpecSetOutput)
}

type AnnotationSpecSetOutput struct{ *pulumi.OutputState }

func (AnnotationSpecSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnnotationSpecSet)(nil)).Elem()
}

func (o AnnotationSpecSetOutput) ToAnnotationSpecSetOutput() AnnotationSpecSetOutput {
	return o
}

func (o AnnotationSpecSetOutput) ToAnnotationSpecSetOutputWithContext(ctx context.Context) AnnotationSpecSetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnnotationSpecSetInput)(nil)).Elem(), &AnnotationSpecSet{})
	pulumi.RegisterOutputType(AnnotationSpecSetOutput{})
}
