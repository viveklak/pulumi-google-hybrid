// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a `Membership`.
// Auto-naming is currently not supported for this resource.
type Membership struct {
	pulumi.CustomResourceState

	// The time when the `Membership` was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The [resource name](https://cloud.google.com/apis/design/resource_names) of the `Membership`. Shall be of the form `groups/{group}/memberships/{membership}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The `EntityKey` of the member.
	PreferredMemberKey EntityKeyResponseOutput `pulumi:"preferredMemberKey"`
	// The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
	Roles MembershipRoleResponseArrayOutput `pulumi:"roles"`
	// The type of the membership.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time when the `Membership` was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewMembership registers a new resource with the given unique name, arguments, and options.
func NewMembership(ctx *pulumi.Context,
	name string, args *MembershipArgs, opts ...pulumi.ResourceOption) (*Membership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.PreferredMemberKey == nil {
		return nil, errors.New("invalid value for required argument 'PreferredMemberKey'")
	}
	var resource Membership
	err := ctx.RegisterResource("google-native:cloudidentity/v1:Membership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembership gets an existing Membership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipState, opts ...pulumi.ResourceOption) (*Membership, error) {
	var resource Membership
	err := ctx.ReadResource("google-native:cloudidentity/v1:Membership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Membership resources.
type membershipState struct {
}

type MembershipState struct {
}

func (MembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipState)(nil)).Elem()
}

type membershipArgs struct {
	GroupId string `pulumi:"groupId"`
	// Immutable. The `EntityKey` of the member.
	PreferredMemberKey EntityKey `pulumi:"preferredMemberKey"`
	// The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
	Roles []MembershipRole `pulumi:"roles"`
}

// The set of arguments for constructing a Membership resource.
type MembershipArgs struct {
	GroupId pulumi.StringInput
	// Immutable. The `EntityKey` of the member.
	PreferredMemberKey EntityKeyInput
	// The `MembershipRole`s that apply to the `Membership`. If unspecified, defaults to a single `MembershipRole` with `name` `MEMBER`. Must not contain duplicate `MembershipRole`s with the same `name`.
	Roles MembershipRoleArrayInput
}

func (MembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipArgs)(nil)).Elem()
}

type MembershipInput interface {
	pulumi.Input

	ToMembershipOutput() MembershipOutput
	ToMembershipOutputWithContext(ctx context.Context) MembershipOutput
}

func (*Membership) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (i *Membership) ToMembershipOutput() MembershipOutput {
	return i.ToMembershipOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipOutput)
}

type MembershipOutput struct{ *pulumi.OutputState }

func (MembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Membership)(nil)).Elem()
}

func (o MembershipOutput) ToMembershipOutput() MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipInput)(nil)).Elem(), &Membership{})
	pulumi.RegisterOutputType(MembershipOutput{})
}
