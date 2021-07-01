// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new domain mapping.
type DomainMapping struct {
	pulumi.CustomResourceState

	// The API version for this call such as "domains.cloudrun.com/v1".
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// The kind of resource, in this case "DomainMapping".
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Metadata associated with this BuildTemplate.
	Metadata ObjectMetaResponseOutput `pulumi:"metadata"`
	// The spec for this DomainMapping.
	Spec DomainMappingSpecResponseOutput `pulumi:"spec"`
	// The current status of the DomainMapping.
	Status DomainMappingStatusResponseOutput `pulumi:"status"`
}

// NewDomainMapping registers a new resource with the given unique name, arguments, and options.
func NewDomainMapping(ctx *pulumi.Context,
	name string, args *DomainMappingArgs, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource DomainMapping
	err := ctx.RegisterResource("google-native:run/v1:DomainMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainMapping gets an existing DomainMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainMappingState, opts ...pulumi.ResourceOption) (*DomainMapping, error) {
	var resource DomainMapping
	err := ctx.ReadResource("google-native:run/v1:DomainMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainMapping resources.
type domainMappingState struct {
}

type DomainMappingState struct {
}

func (DomainMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingState)(nil)).Elem()
}

type domainMappingArgs struct {
	// The API version for this call such as "domains.cloudrun.com/v1".
	ApiVersion *string `pulumi:"apiVersion"`
	DryRun     *string `pulumi:"dryRun"`
	// The kind of resource, in this case "DomainMapping".
	Kind     *string `pulumi:"kind"`
	Location string  `pulumi:"location"`
	// Metadata associated with this BuildTemplate.
	Metadata *ObjectMeta `pulumi:"metadata"`
	Project  string      `pulumi:"project"`
	// The spec for this DomainMapping.
	Spec *DomainMappingSpec `pulumi:"spec"`
	// The current status of the DomainMapping.
	Status *DomainMappingStatus `pulumi:"status"`
}

// The set of arguments for constructing a DomainMapping resource.
type DomainMappingArgs struct {
	// The API version for this call such as "domains.cloudrun.com/v1".
	ApiVersion pulumi.StringPtrInput
	DryRun     pulumi.StringPtrInput
	// The kind of resource, in this case "DomainMapping".
	Kind     pulumi.StringPtrInput
	Location pulumi.StringInput
	// Metadata associated with this BuildTemplate.
	Metadata ObjectMetaPtrInput
	Project  pulumi.StringInput
	// The spec for this DomainMapping.
	Spec DomainMappingSpecPtrInput
	// The current status of the DomainMapping.
	Status DomainMappingStatusPtrInput
}

func (DomainMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainMappingArgs)(nil)).Elem()
}

type DomainMappingInput interface {
	pulumi.Input

	ToDomainMappingOutput() DomainMappingOutput
	ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput
}

func (*DomainMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainMapping)(nil))
}

func (i *DomainMapping) ToDomainMappingOutput() DomainMappingOutput {
	return i.ToDomainMappingOutputWithContext(context.Background())
}

func (i *DomainMapping) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainMappingOutput)
}

type DomainMappingOutput struct {
	*pulumi.OutputState
}

func (DomainMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainMapping)(nil))
}

func (o DomainMappingOutput) ToDomainMappingOutput() DomainMappingOutput {
	return o
}

func (o DomainMappingOutput) ToDomainMappingOutputWithContext(ctx context.Context) DomainMappingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainMappingOutput{})
}
