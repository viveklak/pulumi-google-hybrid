// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Reference in the specified environment.
type Reference struct {
	pulumi.CustomResourceState

	// Optional. A human-readable description of this reference.
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource id of this reference. Values must match the regular expression [\w\s\-.]+.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resource_type.
	Refers pulumi.StringOutput `pulumi:"refers"`
	// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
}

// NewReference registers a new resource with the given unique name, arguments, and options.
func NewReference(ctx *pulumi.Context,
	name string, args *ReferenceArgs, opts ...pulumi.ResourceOption) (*Reference, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.Refers == nil {
		return nil, errors.New("invalid value for required argument 'Refers'")
	}
	var resource Reference
	err := ctx.RegisterResource("google-native:apigee/v1:Reference", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReference gets an existing Reference resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReference(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceState, opts ...pulumi.ResourceOption) (*Reference, error) {
	var resource Reference
	err := ctx.ReadResource("google-native:apigee/v1:Reference", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Reference resources.
type referenceState struct {
	// Optional. A human-readable description of this reference.
	Description *string `pulumi:"description"`
	// The resource id of this reference. Values must match the regular expression [\w\s\-.]+.
	Name *string `pulumi:"name"`
	// The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resource_type.
	Refers *string `pulumi:"refers"`
	// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
	ResourceType *string `pulumi:"resourceType"`
}

type ReferenceState struct {
	// Optional. A human-readable description of this reference.
	Description pulumi.StringPtrInput
	// The resource id of this reference. Values must match the regular expression [\w\s\-.]+.
	Name pulumi.StringPtrInput
	// The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resource_type.
	Refers pulumi.StringPtrInput
	// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
	ResourceType pulumi.StringPtrInput
}

func (ReferenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceState)(nil)).Elem()
}

type referenceArgs struct {
	// Optional. A human-readable description of this reference.
	Description   *string `pulumi:"description"`
	EnvironmentId string  `pulumi:"environmentId"`
	// The resource id of this reference. Values must match the regular expression [\w\s\-.]+.
	Name           string `pulumi:"name"`
	OrganizationId string `pulumi:"organizationId"`
	// The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resource_type.
	Refers string `pulumi:"refers"`
	// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
	ResourceType *string `pulumi:"resourceType"`
}

// The set of arguments for constructing a Reference resource.
type ReferenceArgs struct {
	// Optional. A human-readable description of this reference.
	Description   pulumi.StringPtrInput
	EnvironmentId pulumi.StringInput
	// The resource id of this reference. Values must match the regular expression [\w\s\-.]+.
	Name           pulumi.StringInput
	OrganizationId pulumi.StringInput
	// The id of the resource to which this reference refers. Must be the id of a resource that exists in the parent environment and is of the given resource_type.
	Refers pulumi.StringInput
	// The type of resource referred to by this reference. Valid values are 'KeyStore' or 'TrustStore'.
	ResourceType pulumi.StringPtrInput
}

func (ReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceArgs)(nil)).Elem()
}

type ReferenceInput interface {
	pulumi.Input

	ToReferenceOutput() ReferenceOutput
	ToReferenceOutputWithContext(ctx context.Context) ReferenceOutput
}

func (*Reference) ElementType() reflect.Type {
	return reflect.TypeOf((*Reference)(nil))
}

func (i *Reference) ToReferenceOutput() ReferenceOutput {
	return i.ToReferenceOutputWithContext(context.Background())
}

func (i *Reference) ToReferenceOutputWithContext(ctx context.Context) ReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceOutput)
}

type ReferenceOutput struct {
	*pulumi.OutputState
}

func (ReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Reference)(nil))
}

func (o ReferenceOutput) ToReferenceOutput() ReferenceOutput {
	return o
}

func (o ReferenceOutput) ToReferenceOutputWithContext(ctx context.Context) ReferenceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReferenceOutput{})
}
