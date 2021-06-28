// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new resource policy.
type ResourcePolicy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	Description       pulumi.StringOutput `pulumi:"description"`
	// Resource policy for instances for placement configuration.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyResponseOutput `pulumi:"groupPlacementPolicy"`
	// Resource policy for scheduling instance operations.
	InstanceSchedulePolicy ResourcePolicyInstanceSchedulePolicyResponseOutput `pulumi:"instanceSchedulePolicy"`
	// Type of the resource. Always compute#resource_policies for resource policies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name   pulumi.StringOutput `pulumi:"name"`
	Region pulumi.StringOutput `pulumi:"region"`
	// The system status of the resource policy.
	ResourceStatus ResourcePolicyResourceStatusResponseOutput `pulumi:"resourceStatus"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Resource policy for persistent disks for creating snapshots.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyResponseOutput `pulumi:"snapshotSchedulePolicy"`
	// The status of resource policy creation.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicy(ctx *pulumi.Context,
	name string, args *ResourcePolicyArgs, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource ResourcePolicy
	err := ctx.RegisterResource("google-native:compute/beta:ResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicy gets an existing ResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyState, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	var resource ResourcePolicy
	err := ctx.ReadResource("google-native:compute/beta:ResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicy resources.
type resourcePolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	Description       *string `pulumi:"description"`
	// Resource policy for instances for placement configuration.
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicyResponse `pulumi:"groupPlacementPolicy"`
	// Resource policy for scheduling instance operations.
	InstanceSchedulePolicy *ResourcePolicyInstanceSchedulePolicyResponse `pulumi:"instanceSchedulePolicy"`
	// Type of the resource. Always compute#resource_policies for resource policies.
	Kind *string `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
	// The system status of the resource policy.
	ResourceStatus *ResourcePolicyResourceStatusResponse `pulumi:"resourceStatus"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink *string `pulumi:"selfLink"`
	// Resource policy for persistent disks for creating snapshots.
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicyResponse `pulumi:"snapshotSchedulePolicy"`
	// The status of resource policy creation.
	Status *string `pulumi:"status"`
}

type ResourcePolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	// Resource policy for instances for placement configuration.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyResponsePtrInput
	// Resource policy for scheduling instance operations.
	InstanceSchedulePolicy ResourcePolicyInstanceSchedulePolicyResponsePtrInput
	// Type of the resource. Always compute#resource_policies for resource policies.
	Kind pulumi.StringPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name   pulumi.StringPtrInput
	Region pulumi.StringPtrInput
	// The system status of the resource policy.
	ResourceStatus ResourcePolicyResourceStatusResponsePtrInput
	// Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringPtrInput
	// Resource policy for persistent disks for creating snapshots.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyResponsePtrInput
	// The status of resource policy creation.
	Status pulumi.StringPtrInput
}

func (ResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyState)(nil)).Elem()
}

type resourcePolicyArgs struct {
	Description *string `pulumi:"description"`
	// Resource policy for instances for placement configuration.
	GroupPlacementPolicy *ResourcePolicyGroupPlacementPolicy `pulumi:"groupPlacementPolicy"`
	// Resource policy for scheduling instance operations.
	InstanceSchedulePolicy *ResourcePolicyInstanceSchedulePolicy `pulumi:"instanceSchedulePolicy"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
	// Resource policy for persistent disks for creating snapshots.
	SnapshotSchedulePolicy *ResourcePolicySnapshotSchedulePolicy `pulumi:"snapshotSchedulePolicy"`
}

// The set of arguments for constructing a ResourcePolicy resource.
type ResourcePolicyArgs struct {
	Description pulumi.StringPtrInput
	// Resource policy for instances for placement configuration.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyPtrInput
	// Resource policy for scheduling instance operations.
	InstanceSchedulePolicy ResourcePolicyInstanceSchedulePolicyPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// Resource policy for persistent disks for creating snapshots.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyPtrInput
}

func (ResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyArgs)(nil)).Elem()
}

type ResourcePolicyInput interface {
	pulumi.Input

	ToResourcePolicyOutput() ResourcePolicyOutput
	ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput
}

func (*ResourcePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicy)(nil))
}

func (i *ResourcePolicy) ToResourcePolicyOutput() ResourcePolicyOutput {
	return i.ToResourcePolicyOutputWithContext(context.Background())
}

func (i *ResourcePolicy) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyOutput)
}

type ResourcePolicyOutput struct {
	*pulumi.OutputState
}

func (ResourcePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourcePolicy)(nil))
}

func (o ResourcePolicyOutput) ToResourcePolicyOutput() ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourcePolicyOutput{})
}
