// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new environment group.
type Envgroup struct {
	pulumi.CustomResourceState

	// The time at which the environment group was created as milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Required. Host names for this environment group.
	Hostnames pulumi.StringArrayOutput `pulumi:"hostnames"`
	// The time at which the environment group was last updated as milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// ID of the environment group.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the environment group. Values other than ACTIVE means the resource is not ready to use.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewEnvgroup registers a new resource with the given unique name, arguments, and options.
func NewEnvgroup(ctx *pulumi.Context,
	name string, args *EnvgroupArgs, opts ...pulumi.ResourceOption) (*Envgroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource Envgroup
	err := ctx.RegisterResource("google-native:apigee/v1:Envgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvgroup gets an existing Envgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvgroupState, opts ...pulumi.ResourceOption) (*Envgroup, error) {
	var resource Envgroup
	err := ctx.ReadResource("google-native:apigee/v1:Envgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Envgroup resources.
type envgroupState struct {
}

type EnvgroupState struct {
}

func (EnvgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*envgroupState)(nil)).Elem()
}

type envgroupArgs struct {
	// Required. Host names for this environment group.
	Hostnames []string `pulumi:"hostnames"`
	// ID of the environment group.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
}

// The set of arguments for constructing a Envgroup resource.
type EnvgroupArgs struct {
	// Required. Host names for this environment group.
	Hostnames pulumi.StringArrayInput
	// ID of the environment group.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
}

func (EnvgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*envgroupArgs)(nil)).Elem()
}

type EnvgroupInput interface {
	pulumi.Input

	ToEnvgroupOutput() EnvgroupOutput
	ToEnvgroupOutputWithContext(ctx context.Context) EnvgroupOutput
}

func (*Envgroup) ElementType() reflect.Type {
	return reflect.TypeOf((*Envgroup)(nil))
}

func (i *Envgroup) ToEnvgroupOutput() EnvgroupOutput {
	return i.ToEnvgroupOutputWithContext(context.Background())
}

func (i *Envgroup) ToEnvgroupOutputWithContext(ctx context.Context) EnvgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvgroupOutput)
}

type EnvgroupOutput struct {
	*pulumi.OutputState
}

func (EnvgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Envgroup)(nil))
}

func (o EnvgroupOutput) ToEnvgroupOutput() EnvgroupOutput {
	return o
}

func (o EnvgroupOutput) ToEnvgroupOutputWithContext(ctx context.Context) EnvgroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EnvgroupOutput{})
}
