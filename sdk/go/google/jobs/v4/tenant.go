// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v4

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new tenant entity.
type Tenant struct {
	pulumi.CustomResourceState

	// Required. Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTenant registers a new resource with the given unique name, arguments, and options.
func NewTenant(ctx *pulumi.Context,
	name string, args *TenantArgs, opts ...pulumi.ResourceOption) (*Tenant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Tenant
	err := ctx.RegisterResource("google-native:jobs/v4:Tenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenant gets an existing Tenant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantState, opts ...pulumi.ResourceOption) (*Tenant, error) {
	var resource Tenant
	err := ctx.ReadResource("google-native:jobs/v4:Tenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tenant resources.
type tenantState struct {
}

type TenantState struct {
}

func (TenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantState)(nil)).Elem()
}

type tenantArgs struct {
	// Required. Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
	ExternalId *string `pulumi:"externalId"`
	// Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
}

// The set of arguments for constructing a Tenant resource.
type TenantArgs struct {
	// Required. Client side tenant identifier, used to uniquely identify the tenant. The maximum number of allowed characters is 255.
	ExternalId pulumi.StringPtrInput
	// Required during tenant update. The resource name for a tenant. This is generated by the service when a tenant is created. The format is "projects/{project_id}/tenants/{tenant_id}", for example, "projects/foo/tenants/bar".
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
}

func (TenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantArgs)(nil)).Elem()
}

type TenantInput interface {
	pulumi.Input

	ToTenantOutput() TenantOutput
	ToTenantOutputWithContext(ctx context.Context) TenantOutput
}

func (*Tenant) ElementType() reflect.Type {
	return reflect.TypeOf((*Tenant)(nil))
}

func (i *Tenant) ToTenantOutput() TenantOutput {
	return i.ToTenantOutputWithContext(context.Background())
}

func (i *Tenant) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantOutput)
}

type TenantOutput struct {
	*pulumi.OutputState
}

func (TenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tenant)(nil))
}

func (o TenantOutput) ToTenantOutput() TenantOutput {
	return o
}

func (o TenantOutput) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TenantOutput{})
}
