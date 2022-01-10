// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a `Group`.
// Auto-naming is currently not supported for this resource.
type Group struct {
	pulumi.CustomResourceState

	// Additional entity key aliases for a Group.
	AdditionalGroupKeys EntityKeyResponseArrayOutput `pulumi:"additionalGroupKeys"`
	// The time when the `Group` was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the `Group`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Dynamic group metadata like queries and status.
	DynamicGroupMetadata DynamicGroupMetadataResponseOutput `pulumi:"dynamicGroupMetadata"`
	// Immutable. The `EntityKey` of the `Group`.
	GroupKey EntityKeyResponseOutput `pulumi:"groupKey"`
	// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Group`. Shall be of the form `groups/{group_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups. The `customer_id` must begin with "C" (for example, 'C046psxkn').
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Optional. The POSIX groups associated with the `Group`.
	PosixGroups PosixGroupResponseArrayOutput `pulumi:"posixGroups"`
	// The time when the `Group` was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupKey == nil {
		return nil, errors.New("invalid value for required argument 'GroupKey'")
	}
	if args.InitialGroupConfig == nil {
		return nil, errors.New("invalid value for required argument 'InitialGroupConfig'")
	}
	if args.Labels == nil {
		return nil, errors.New("invalid value for required argument 'Labels'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource Group
	err := ctx.RegisterResource("google-native:cloudidentity/v1beta1:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("google-native:cloudidentity/v1beta1:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
}

type GroupState struct {
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Additional entity key aliases for a Group.
	AdditionalGroupKeys []EntityKey `pulumi:"additionalGroupKeys"`
	// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
	Description *string `pulumi:"description"`
	// The display name of the `Group`.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Dynamic group metadata like queries and status.
	DynamicGroupMetadata *DynamicGroupMetadata `pulumi:"dynamicGroupMetadata"`
	// Immutable. The `EntityKey` of the `Group`.
	GroupKey           EntityKey `pulumi:"groupKey"`
	InitialGroupConfig string    `pulumi:"initialGroupConfig"`
	// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups. The `customer_id` must begin with "C" (for example, 'C046psxkn').
	Parent string `pulumi:"parent"`
	// Optional. The POSIX groups associated with the `Group`.
	PosixGroups []PosixGroup `pulumi:"posixGroups"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Additional entity key aliases for a Group.
	AdditionalGroupKeys EntityKeyArrayInput
	// An extended description to help users determine the purpose of a `Group`. Must not be longer than 4,096 characters.
	Description pulumi.StringPtrInput
	// The display name of the `Group`.
	DisplayName pulumi.StringPtrInput
	// Optional. Dynamic group metadata like queries and status.
	DynamicGroupMetadata DynamicGroupMetadataPtrInput
	// Immutable. The `EntityKey` of the `Group`.
	GroupKey           EntityKeyInput
	InitialGroupConfig pulumi.StringInput
	// One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value. Google Groups are the default type of group and have a label with a key of `cloudidentity.googleapis.com/groups.discussion_forum` and an empty value. Existing Google Groups can have an additional label with a key of `cloudidentity.googleapis.com/groups.security` and an empty value added to them. **This is an immutable change and the security label cannot be removed once added.** Dynamic groups have a label with a key of `cloudidentity.googleapis.com/groups.dynamic`. Identity-mapped groups for Cloud Search have a label with a key of `system/groups/external` and an empty value.
	Labels pulumi.StringMapInput
	// Immutable. The resource name of the entity under which this `Group` resides in the Cloud Identity resource hierarchy. Must be of the form `identitysources/{identity_source_id}` for external- identity-mapped groups or `customers/{customer_id}` for Google Groups. The `customer_id` must begin with "C" (for example, 'C046psxkn').
	Parent pulumi.StringInput
	// Optional. The POSIX groups associated with the `Group`.
	PosixGroups PosixGroupArrayInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterOutputType(GroupOutput{})
}
