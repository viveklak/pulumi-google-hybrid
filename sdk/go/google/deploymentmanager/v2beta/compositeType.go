// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a composite type.
type CompositeType struct {
	pulumi.CustomResourceState

	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringOutput `pulumi:"description"`
	// Creation timestamp in RFC3339 text format.
	InsertTime pulumi.StringOutput `pulumi:"insertTime"`
	// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	Labels CompositeTypeLabelEntryResponseArrayOutput `pulumi:"labels"`
	// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Operation that most recently ran, or is currently running, on this composite type.
	Operation OperationResponseOutput `pulumi:"operation"`
	// Server defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	Status   pulumi.StringOutput `pulumi:"status"`
	// Files for the template type.
	TemplateContents TemplateContentsResponseOutput `pulumi:"templateContents"`
}

// NewCompositeType registers a new resource with the given unique name, arguments, and options.
func NewCompositeType(ctx *pulumi.Context,
	name string, args *CompositeTypeArgs, opts ...pulumi.ResourceOption) (*CompositeType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource CompositeType
	err := ctx.RegisterResource("google-native:deploymentmanager/v2beta:CompositeType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompositeType gets an existing CompositeType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompositeType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CompositeTypeState, opts ...pulumi.ResourceOption) (*CompositeType, error) {
	var resource CompositeType
	err := ctx.ReadResource("google-native:deploymentmanager/v2beta:CompositeType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CompositeType resources.
type compositeTypeState struct {
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	// Creation timestamp in RFC3339 text format.
	InsertTime *string `pulumi:"insertTime"`
	// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	Labels []CompositeTypeLabelEntryResponse `pulumi:"labels"`
	// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
	Name *string `pulumi:"name"`
	// The Operation that most recently ran, or is currently running, on this composite type.
	Operation *OperationResponse `pulumi:"operation"`
	// Server defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	Status   *string `pulumi:"status"`
	// Files for the template type.
	TemplateContents *TemplateContentsResponse `pulumi:"templateContents"`
}

type CompositeTypeState struct {
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	InsertTime pulumi.StringPtrInput
	// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	Labels CompositeTypeLabelEntryResponseArrayInput
	// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
	Name pulumi.StringPtrInput
	// The Operation that most recently ran, or is currently running, on this composite type.
	Operation OperationResponsePtrInput
	// Server defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	Status   pulumi.StringPtrInput
	// Files for the template type.
	TemplateContents TemplateContentsResponsePtrInput
}

func (CompositeTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*compositeTypeState)(nil)).Elem()
}

type compositeTypeArgs struct {
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description *string `pulumi:"description"`
	Id          *string `pulumi:"id"`
	// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	Labels []CompositeTypeLabelEntry `pulumi:"labels"`
	// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	Status  *string `pulumi:"status"`
	// Files for the template type.
	TemplateContents *TemplateContents `pulumi:"templateContents"`
}

// The set of arguments for constructing a CompositeType resource.
type CompositeTypeArgs struct {
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description pulumi.StringPtrInput
	Id          pulumi.StringPtrInput
	// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	Labels CompositeTypeLabelEntryArrayInput
	// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	Status  *CompositeTypeStatus
	// Files for the template type.
	TemplateContents TemplateContentsPtrInput
}

func (CompositeTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*compositeTypeArgs)(nil)).Elem()
}

type CompositeTypeInput interface {
	pulumi.Input

	ToCompositeTypeOutput() CompositeTypeOutput
	ToCompositeTypeOutputWithContext(ctx context.Context) CompositeTypeOutput
}

func (*CompositeType) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositeType)(nil))
}

func (i *CompositeType) ToCompositeTypeOutput() CompositeTypeOutput {
	return i.ToCompositeTypeOutputWithContext(context.Background())
}

func (i *CompositeType) ToCompositeTypeOutputWithContext(ctx context.Context) CompositeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositeTypeOutput)
}

type CompositeTypeOutput struct {
	*pulumi.OutputState
}

func (CompositeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositeType)(nil))
}

func (o CompositeTypeOutput) ToCompositeTypeOutput() CompositeTypeOutput {
	return o
}

func (o CompositeTypeOutput) ToCompositeTypeOutputWithContext(ctx context.Context) CompositeTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CompositeTypeOutput{})
}
