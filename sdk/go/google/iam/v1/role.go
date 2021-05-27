// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new custom Role.
type Role struct {
	pulumi.CustomResourceState

	// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// Optional. A human-readable description for the role.
	Description pulumi.StringOutput `pulumi:"description"`
	// Used to perform a consistent read-modify-write.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The names of the permissions this role grants when bound in an IAM policy.
	IncludedPermissions pulumi.StringArrayOutput `pulumi:"includedPermissions"`
	// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
	Stage pulumi.StringOutput `pulumi:"stage"`
	// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Role
	err := ctx.RegisterResource("google-native:iam/v1:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("google-native:iam/v1:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
	Deleted *bool `pulumi:"deleted"`
	// Optional. A human-readable description for the role.
	Description *string `pulumi:"description"`
	// Used to perform a consistent read-modify-write.
	Etag *string `pulumi:"etag"`
	// The names of the permissions this role grants when bound in an IAM policy.
	IncludedPermissions []string `pulumi:"includedPermissions"`
	// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
	Name *string `pulumi:"name"`
	// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
	Stage *string `pulumi:"stage"`
	// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
	Title *string `pulumi:"title"`
}

type RoleState struct {
	// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
	Deleted pulumi.BoolPtrInput
	// Optional. A human-readable description for the role.
	Description pulumi.StringPtrInput
	// Used to perform a consistent read-modify-write.
	Etag pulumi.StringPtrInput
	// The names of the permissions this role grants when bound in an IAM policy.
	IncludedPermissions pulumi.StringArrayInput
	// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
	Name pulumi.StringPtrInput
	// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
	Stage pulumi.StringPtrInput
	// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
	Title pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
	Deleted *bool `pulumi:"deleted"`
	// Optional. A human-readable description for the role.
	Description *string `pulumi:"description"`
	// Used to perform a consistent read-modify-write.
	Etag *string `pulumi:"etag"`
	// The names of the permissions this role grants when bound in an IAM policy.
	IncludedPermissions []string `pulumi:"includedPermissions"`
	// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// The role ID to use for this role. A role ID may contain alphanumeric characters, underscores (`_`), and periods (`.`). It must contain a minimum of 3 characters and a maximum of 64 characters.
	RoleId *string `pulumi:"roleId"`
	// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
	Stage *string `pulumi:"stage"`
	// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
	Title *string `pulumi:"title"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
	Deleted pulumi.BoolPtrInput
	// Optional. A human-readable description for the role.
	Description pulumi.StringPtrInput
	// Used to perform a consistent read-modify-write.
	Etag pulumi.StringPtrInput
	// The names of the permissions this role grants when bound in an IAM policy.
	IncludedPermissions pulumi.StringArrayInput
	// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// The role ID to use for this role. A role ID may contain alphanumeric characters, underscores (`_`), and periods (`.`). It must contain a minimum of 3 characters and a maximum of 64 characters.
	RoleId pulumi.StringPtrInput
	// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
	Stage pulumi.StringPtrInput
	// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
	Title pulumi.StringPtrInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil))
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

type RoleOutput struct {
	*pulumi.OutputState
}

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil))
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoleOutput{})
}
