// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new TunnelDestGroup.
type DestGroup struct {
	pulumi.CustomResourceState

	// null List of CIDRs that this group applies to.
	Cidrs pulumi.StringArrayOutput `pulumi:"cidrs"`
	// null List of FQDNs that this group applies to.
	Fqdns pulumi.StringArrayOutput `pulumi:"fqdns"`
	// Immutable. Identifier for the TunnelDestGroup. Must be unique within the project.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDestGroup registers a new resource with the given unique name, arguments, and options.
func NewDestGroup(ctx *pulumi.Context,
	name string, args *DestGroupArgs, opts ...pulumi.ResourceOption) (*DestGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TunnelDestGroupId == nil {
		return nil, errors.New("invalid value for required argument 'TunnelDestGroupId'")
	}
	var resource DestGroup
	err := ctx.RegisterResource("google-native:iap/v1:DestGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDestGroup gets an existing DestGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDestGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DestGroupState, opts ...pulumi.ResourceOption) (*DestGroup, error) {
	var resource DestGroup
	err := ctx.ReadResource("google-native:iap/v1:DestGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DestGroup resources.
type destGroupState struct {
}

type DestGroupState struct {
}

func (DestGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*destGroupState)(nil)).Elem()
}

type destGroupArgs struct {
	// null List of CIDRs that this group applies to.
	Cidrs []string `pulumi:"cidrs"`
	// null List of FQDNs that this group applies to.
	Fqdns    []string `pulumi:"fqdns"`
	Location *string  `pulumi:"location"`
	// Immutable. Identifier for the TunnelDestGroup. Must be unique within the project.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Required. The ID to use for the TunnelDestGroup, which will become the final component of the resource name. This value should be 4-63 characters, and valid characters are /a-z-/.
	TunnelDestGroupId string `pulumi:"tunnelDestGroupId"`
}

// The set of arguments for constructing a DestGroup resource.
type DestGroupArgs struct {
	// null List of CIDRs that this group applies to.
	Cidrs pulumi.StringArrayInput
	// null List of FQDNs that this group applies to.
	Fqdns    pulumi.StringArrayInput
	Location pulumi.StringPtrInput
	// Immutable. Identifier for the TunnelDestGroup. Must be unique within the project.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Required. The ID to use for the TunnelDestGroup, which will become the final component of the resource name. This value should be 4-63 characters, and valid characters are /a-z-/.
	TunnelDestGroupId pulumi.StringInput
}

func (DestGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*destGroupArgs)(nil)).Elem()
}

type DestGroupInput interface {
	pulumi.Input

	ToDestGroupOutput() DestGroupOutput
	ToDestGroupOutputWithContext(ctx context.Context) DestGroupOutput
}

func (*DestGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DestGroup)(nil)).Elem()
}

func (i *DestGroup) ToDestGroupOutput() DestGroupOutput {
	return i.ToDestGroupOutputWithContext(context.Background())
}

func (i *DestGroup) ToDestGroupOutputWithContext(ctx context.Context) DestGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestGroupOutput)
}

type DestGroupOutput struct{ *pulumi.OutputState }

func (DestGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestGroup)(nil)).Elem()
}

func (o DestGroupOutput) ToDestGroupOutput() DestGroupOutput {
	return o
}

func (o DestGroupOutput) ToDestGroupOutputWithContext(ctx context.Context) DestGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DestGroupInput)(nil)).Elem(), &DestGroup{})
	pulumi.RegisterOutputType(DestGroupOutput{})
}
