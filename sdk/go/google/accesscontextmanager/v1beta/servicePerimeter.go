// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a Service Perimeter. The longrunning operation from this RPC will have a successful status once the Service Perimeter has propagated to long-lasting storage. Service Perimeters containing errors will result in an error response for the first error encountered.
type ServicePerimeter struct {
	pulumi.CustomResourceState

	// Description of the `ServicePerimeter` and its use. Does not affect behavior.
	Description pulumi.StringOutput `pulumi:"description"`
	// Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
	PerimeterType pulumi.StringOutput `pulumi:"perimeterType"`
	// Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
	Status ServicePerimeterConfigResponseOutput `pulumi:"status"`
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewServicePerimeter registers a new resource with the given unique name, arguments, and options.
func NewServicePerimeter(ctx *pulumi.Context,
	name string, args *ServicePerimeterArgs, opts ...pulumi.ResourceOption) (*ServicePerimeter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'AccessPolicyId'")
	}
	var resource ServicePerimeter
	err := ctx.RegisterResource("google-native:accesscontextmanager/v1beta:ServicePerimeter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePerimeter gets an existing ServicePerimeter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePerimeter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePerimeterState, opts ...pulumi.ResourceOption) (*ServicePerimeter, error) {
	var resource ServicePerimeter
	err := ctx.ReadResource("google-native:accesscontextmanager/v1beta:ServicePerimeter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePerimeter resources.
type servicePerimeterState struct {
}

type ServicePerimeterState struct {
}

func (ServicePerimeterState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterState)(nil)).Elem()
}

type servicePerimeterArgs struct {
	AccessPolicyId string `pulumi:"accessPolicyId"`
	// Description of the `ServicePerimeter` and its use. Does not affect behavior.
	Description *string `pulumi:"description"`
	// Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
	Name *string `pulumi:"name"`
	// Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
	PerimeterType *ServicePerimeterPerimeterType `pulumi:"perimeterType"`
	// Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
	Status *ServicePerimeterConfig `pulumi:"status"`
	// Human readable title. Must be unique within the Policy.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a ServicePerimeter resource.
type ServicePerimeterArgs struct {
	AccessPolicyId pulumi.StringInput
	// Description of the `ServicePerimeter` and its use. Does not affect behavior.
	Description pulumi.StringPtrInput
	// Resource name for the ServicePerimeter. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/servicePerimeters/{short_name}`
	Name pulumi.StringPtrInput
	// Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
	PerimeterType ServicePerimeterPerimeterTypePtrInput
	// Current ServicePerimeter configuration. Specifies sets of resources, restricted/unrestricted services and access levels that determine perimeter content and boundaries.
	Status ServicePerimeterConfigPtrInput
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringPtrInput
}

func (ServicePerimeterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePerimeterArgs)(nil)).Elem()
}

type ServicePerimeterInput interface {
	pulumi.Input

	ToServicePerimeterOutput() ServicePerimeterOutput
	ToServicePerimeterOutputWithContext(ctx context.Context) ServicePerimeterOutput
}

func (*ServicePerimeter) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePerimeter)(nil)).Elem()
}

func (i *ServicePerimeter) ToServicePerimeterOutput() ServicePerimeterOutput {
	return i.ToServicePerimeterOutputWithContext(context.Background())
}

func (i *ServicePerimeter) ToServicePerimeterOutputWithContext(ctx context.Context) ServicePerimeterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePerimeterOutput)
}

type ServicePerimeterOutput struct{ *pulumi.OutputState }

func (ServicePerimeterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePerimeter)(nil)).Elem()
}

func (o ServicePerimeterOutput) ToServicePerimeterOutput() ServicePerimeterOutput {
	return o
}

func (o ServicePerimeterOutput) ToServicePerimeterOutputWithContext(ctx context.Context) ServicePerimeterOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePerimeterInput)(nil)).Elem(), &ServicePerimeter{})
	pulumi.RegisterOutputType(ServicePerimeterOutput{})
}
