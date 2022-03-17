// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Annotation record. It is valid to create Annotation objects for the same source more than once since a unique ID is assigned to each record by this service.
type Annotation struct {
	pulumi.CustomResourceState

	// Details of the source.
	AnnotationSource AnnotationSourceResponseOutput `pulumi:"annotationSource"`
	// Additional information for this annotation record, such as annotator and verifier information or study campaign.
	CustomData pulumi.StringMapOutput `pulumi:"customData"`
	// Annotations for images. For example, bounding polygons.
	ImageAnnotation ImageAnnotationResponseOutput `pulumi:"imageAnnotation"`
	// Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Annotations for resource. For example, classification tags.
	ResourceAnnotation ResourceAnnotationResponseOutput `pulumi:"resourceAnnotation"`
	// Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
	TextAnnotation SensitiveTextAnnotationResponseOutput `pulumi:"textAnnotation"`
}

// NewAnnotation registers a new resource with the given unique name, arguments, and options.
func NewAnnotation(ctx *pulumi.Context,
	name string, args *AnnotationArgs, opts ...pulumi.ResourceOption) (*Annotation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnnotationStoreId == nil {
		return nil, errors.New("invalid value for required argument 'AnnotationStoreId'")
	}
	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	var resource Annotation
	err := ctx.RegisterResource("google-native:healthcare/v1beta1:Annotation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnnotation gets an existing Annotation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnnotation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnnotationState, opts ...pulumi.ResourceOption) (*Annotation, error) {
	var resource Annotation
	err := ctx.ReadResource("google-native:healthcare/v1beta1:Annotation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Annotation resources.
type annotationState struct {
}

type AnnotationState struct {
}

func (AnnotationState) ElementType() reflect.Type {
	return reflect.TypeOf((*annotationState)(nil)).Elem()
}

type annotationArgs struct {
	// Details of the source.
	AnnotationSource  *AnnotationSource `pulumi:"annotationSource"`
	AnnotationStoreId string            `pulumi:"annotationStoreId"`
	// Additional information for this annotation record, such as annotator and verifier information or study campaign.
	CustomData map[string]string `pulumi:"customData"`
	DatasetId  string            `pulumi:"datasetId"`
	// Annotations for images. For example, bounding polygons.
	ImageAnnotation *ImageAnnotation `pulumi:"imageAnnotation"`
	Location        *string          `pulumi:"location"`
	// Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Annotations for resource. For example, classification tags.
	ResourceAnnotation *ResourceAnnotation `pulumi:"resourceAnnotation"`
	// Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
	TextAnnotation *SensitiveTextAnnotation `pulumi:"textAnnotation"`
}

// The set of arguments for constructing a Annotation resource.
type AnnotationArgs struct {
	// Details of the source.
	AnnotationSource  AnnotationSourcePtrInput
	AnnotationStoreId pulumi.StringInput
	// Additional information for this annotation record, such as annotator and verifier information or study campaign.
	CustomData pulumi.StringMapInput
	DatasetId  pulumi.StringInput
	// Annotations for images. For example, bounding polygons.
	ImageAnnotation ImageAnnotationPtrInput
	Location        pulumi.StringPtrInput
	// Resource name of the Annotation, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationStores/{annotation_store_id}/annotations/{annotation_id}`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Annotations for resource. For example, classification tags.
	ResourceAnnotation ResourceAnnotationPtrInput
	// Annotations for sensitive texts. For example, a range that describes the location of sensitive text.
	TextAnnotation SensitiveTextAnnotationPtrInput
}

func (AnnotationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*annotationArgs)(nil)).Elem()
}

type AnnotationInput interface {
	pulumi.Input

	ToAnnotationOutput() AnnotationOutput
	ToAnnotationOutputWithContext(ctx context.Context) AnnotationOutput
}

func (*Annotation) ElementType() reflect.Type {
	return reflect.TypeOf((**Annotation)(nil)).Elem()
}

func (i *Annotation) ToAnnotationOutput() AnnotationOutput {
	return i.ToAnnotationOutputWithContext(context.Background())
}

func (i *Annotation) ToAnnotationOutputWithContext(ctx context.Context) AnnotationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnnotationOutput)
}

type AnnotationOutput struct{ *pulumi.OutputState }

func (AnnotationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Annotation)(nil)).Elem()
}

func (o AnnotationOutput) ToAnnotationOutput() AnnotationOutput {
	return o
}

func (o AnnotationOutput) ToAnnotationOutputWithContext(ctx context.Context) AnnotationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnnotationInput)(nil)).Elem(), &Annotation{})
	pulumi.RegisterOutputType(AnnotationOutput{})
}
