// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Apigee runtime instance. The instance is accessible from the authorized network configured on the organization. **Note:** Not supported for Apigee hybrid.
type OrganizationInstance struct {
	pulumi.CustomResourceState

	// Time the instance was created in milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Optional. Description of the instance.
	Description pulumi.StringOutput `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringOutput `pulumi:"diskEncryptionKeyName"`
	// Optional. Display name for the instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
	Host pulumi.StringOutput `pulumi:"host"`
	// Time the instance was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Required. Compute Engine location where the instance resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// Required. Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange pulumi.StringOutput `pulumi:"peeringCidrRange"`
	// Port number of the exposed Apigee endpoint.
	Port pulumi.StringOutput `pulumi:"port"`
	// State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewOrganizationInstance registers a new resource with the given unique name, arguments, and options.
func NewOrganizationInstance(ctx *pulumi.Context,
	name string, args *OrganizationInstanceArgs, opts ...pulumi.ResourceOption) (*OrganizationInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource OrganizationInstance
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationInstance gets an existing OrganizationInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationInstanceState, opts ...pulumi.ResourceOption) (*OrganizationInstance, error) {
	var resource OrganizationInstance
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationInstance resources.
type organizationInstanceState struct {
	// Time the instance was created in milliseconds since epoch.
	CreatedAt *string `pulumi:"createdAt"`
	// Optional. Description of the instance.
	Description *string `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName *string `pulumi:"diskEncryptionKeyName"`
	// Optional. Display name for the instance.
	DisplayName *string `pulumi:"displayName"`
	// Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
	Host *string `pulumi:"host"`
	// Time the instance was last modified in milliseconds since epoch.
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// Required. Compute Engine location where the instance resides.
	Location *string `pulumi:"location"`
	// Required. Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name *string `pulumi:"name"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange *string `pulumi:"peeringCidrRange"`
	// Port number of the exposed Apigee endpoint.
	Port *string `pulumi:"port"`
	// State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
	State *string `pulumi:"state"`
}

type OrganizationInstanceState struct {
	// Time the instance was created in milliseconds since epoch.
	CreatedAt pulumi.StringPtrInput
	// Optional. Description of the instance.
	Description pulumi.StringPtrInput
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringPtrInput
	// Optional. Display name for the instance.
	DisplayName pulumi.StringPtrInput
	// Internal hostname or IP address of the Apigee endpoint used by clients to connect to the service.
	Host pulumi.StringPtrInput
	// Time the instance was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringPtrInput
	// Required. Compute Engine location where the instance resides.
	Location pulumi.StringPtrInput
	// Required. Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name pulumi.StringPtrInput
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange pulumi.StringPtrInput
	// Port number of the exposed Apigee endpoint.
	Port pulumi.StringPtrInput
	// State of the instance. Values other than `ACTIVE` means the resource is not ready to use.
	State pulumi.StringPtrInput
}

func (OrganizationInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationInstanceState)(nil)).Elem()
}

type organizationInstanceArgs struct {
	// Optional. Description of the instance.
	Description *string `pulumi:"description"`
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName *string `pulumi:"diskEncryptionKeyName"`
	// Optional. Display name for the instance.
	DisplayName  *string `pulumi:"displayName"`
	Environments *string `pulumi:"environments"`
	// Required. Compute Engine location where the instance resides.
	Location *string `pulumi:"location"`
	// Required. Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange *string `pulumi:"peeringCidrRange"`
}

// The set of arguments for constructing a OrganizationInstance resource.
type OrganizationInstanceArgs struct {
	// Optional. Description of the instance.
	Description pulumi.StringPtrInput
	// Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only. Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
	DiskEncryptionKeyName pulumi.StringPtrInput
	// Optional. Display name for the instance.
	DisplayName  pulumi.StringPtrInput
	Environments pulumi.StringPtrInput
	// Required. Compute Engine location where the instance resides.
	Location pulumi.StringPtrInput
	// Required. Resource ID of the instance. Values must match the regular expression `^a-z{0,30}[a-z\d]$`.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
	PeeringCidrRange pulumi.StringPtrInput
}

func (OrganizationInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationInstanceArgs)(nil)).Elem()
}

type OrganizationInstanceInput interface {
	pulumi.Input

	ToOrganizationInstanceOutput() OrganizationInstanceOutput
	ToOrganizationInstanceOutputWithContext(ctx context.Context) OrganizationInstanceOutput
}

func (*OrganizationInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationInstance)(nil))
}

func (i *OrganizationInstance) ToOrganizationInstanceOutput() OrganizationInstanceOutput {
	return i.ToOrganizationInstanceOutputWithContext(context.Background())
}

func (i *OrganizationInstance) ToOrganizationInstanceOutputWithContext(ctx context.Context) OrganizationInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationInstanceOutput)
}

type OrganizationInstanceOutput struct {
	*pulumi.OutputState
}

func (OrganizationInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationInstance)(nil))
}

func (o OrganizationInstanceOutput) ToOrganizationInstanceOutput() OrganizationInstanceOutput {
	return o
}

func (o OrganizationInstanceOutput) ToOrganizationInstanceOutputWithContext(ctx context.Context) OrganizationInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationInstanceOutput{})
}
