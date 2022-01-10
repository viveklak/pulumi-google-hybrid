// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an attestor, and returns a copy of the new attestor. Returns NOT_FOUND if the project does not exist, INVALID_ARGUMENT if the request is malformed, ALREADY_EXISTS if the attestor already exists.
type Attestor struct {
	pulumi.CustomResourceState

	// Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
	Name pulumi.StringOutput `pulumi:"name"`
	// Time when the attestor was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// A Drydock ATTESTATION_AUTHORITY Note, created by the user.
	UserOwnedDrydockNote UserOwnedDrydockNoteResponseOutput `pulumi:"userOwnedDrydockNote"`
}

// NewAttestor registers a new resource with the given unique name, arguments, and options.
func NewAttestor(ctx *pulumi.Context,
	name string, args *AttestorArgs, opts ...pulumi.ResourceOption) (*Attestor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttestorId == nil {
		return nil, errors.New("invalid value for required argument 'AttestorId'")
	}
	var resource Attestor
	err := ctx.RegisterResource("google-native:binaryauthorization/v1beta1:Attestor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestor gets an existing Attestor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorState, opts ...pulumi.ResourceOption) (*Attestor, error) {
	var resource Attestor
	err := ctx.ReadResource("google-native:binaryauthorization/v1beta1:Attestor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attestor resources.
type attestorState struct {
}

type AttestorState struct {
}

func (AttestorState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorState)(nil)).Elem()
}

type attestorArgs struct {
	AttestorId string `pulumi:"attestorId"`
	// Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description *string `pulumi:"description"`
	// The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A Drydock ATTESTATION_AUTHORITY Note, created by the user.
	UserOwnedDrydockNote *UserOwnedDrydockNote `pulumi:"userOwnedDrydockNote"`
}

// The set of arguments for constructing a Attestor resource.
type AttestorArgs struct {
	AttestorId pulumi.StringInput
	// Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description pulumi.StringPtrInput
	// The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A Drydock ATTESTATION_AUTHORITY Note, created by the user.
	UserOwnedDrydockNote UserOwnedDrydockNotePtrInput
}

func (AttestorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorArgs)(nil)).Elem()
}

type AttestorInput interface {
	pulumi.Input

	ToAttestorOutput() AttestorOutput
	ToAttestorOutputWithContext(ctx context.Context) AttestorOutput
}

func (*Attestor) ElementType() reflect.Type {
	return reflect.TypeOf((**Attestor)(nil)).Elem()
}

func (i *Attestor) ToAttestorOutput() AttestorOutput {
	return i.ToAttestorOutputWithContext(context.Background())
}

func (i *Attestor) ToAttestorOutputWithContext(ctx context.Context) AttestorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorOutput)
}

type AttestorOutput struct{ *pulumi.OutputState }

func (AttestorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Attestor)(nil)).Elem()
}

func (o AttestorOutput) ToAttestorOutput() AttestorOutput {
	return o
}

func (o AttestorOutput) ToAttestorOutputWithContext(ctx context.Context) AttestorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttestorInput)(nil)).Elem(), &Attestor{})
	pulumi.RegisterOutputType(AttestorOutput{})
}
