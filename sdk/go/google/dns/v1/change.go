// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Atomically updates the ResourceRecordSet collection.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Change struct {
	pulumi.CustomResourceState

	// Which ResourceRecordSets to add?
	Additions ResourceRecordSetResponseArrayOutput `pulumi:"additions"`
	// Which ResourceRecordSets to remove? Must match existing data exactly.
	Deletions ResourceRecordSetResponseArrayOutput `pulumi:"deletions"`
	// If the DNS queries for the zone will be served.
	IsServing pulumi.BoolOutput   `pulumi:"isServing"`
	Kind      pulumi.StringOutput `pulumi:"kind"`
	// The time that this operation was started by the server (output only). This is in RFC3339 text format.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewChange registers a new resource with the given unique name, arguments, and options.
func NewChange(ctx *pulumi.Context,
	name string, args *ChangeArgs, opts ...pulumi.ResourceOption) (*Change, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedZone == nil {
		return nil, errors.New("invalid value for required argument 'ManagedZone'")
	}
	var resource Change
	err := ctx.RegisterResource("google-native:dns/v1:Change", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChange gets an existing Change resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChangeState, opts ...pulumi.ResourceOption) (*Change, error) {
	var resource Change
	err := ctx.ReadResource("google-native:dns/v1:Change", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Change resources.
type changeState struct {
}

type ChangeState struct {
}

func (ChangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*changeState)(nil)).Elem()
}

type changeArgs struct {
	// Which ResourceRecordSets to add?
	Additions []ResourceRecordSetType `pulumi:"additions"`
	// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
	ClientOperationId *string `pulumi:"clientOperationId"`
	// Which ResourceRecordSets to remove? Must match existing data exactly.
	Deletions []ResourceRecordSetType `pulumi:"deletions"`
	// Unique identifier for the resource; defined by the server (output only).
	Id *string `pulumi:"id"`
	// If the DNS queries for the zone will be served.
	IsServing   *bool   `pulumi:"isServing"`
	Kind        *string `pulumi:"kind"`
	ManagedZone string  `pulumi:"managedZone"`
	Project     *string `pulumi:"project"`
	// The time that this operation was started by the server (output only). This is in RFC3339 text format.
	StartTime *string `pulumi:"startTime"`
	// Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
	Status *ChangeStatus `pulumi:"status"`
}

// The set of arguments for constructing a Change resource.
type ChangeArgs struct {
	// Which ResourceRecordSets to add?
	Additions ResourceRecordSetTypeArrayInput
	// For mutating operation requests only. An optional identifier specified by the client. Must be unique for operation resources in the Operations collection.
	ClientOperationId pulumi.StringPtrInput
	// Which ResourceRecordSets to remove? Must match existing data exactly.
	Deletions ResourceRecordSetTypeArrayInput
	// Unique identifier for the resource; defined by the server (output only).
	Id pulumi.StringPtrInput
	// If the DNS queries for the zone will be served.
	IsServing   pulumi.BoolPtrInput
	Kind        pulumi.StringPtrInput
	ManagedZone pulumi.StringInput
	Project     pulumi.StringPtrInput
	// The time that this operation was started by the server (output only). This is in RFC3339 text format.
	StartTime pulumi.StringPtrInput
	// Status of the operation (output only). A status of "done" means that the request to update the authoritative servers has been sent, but the servers might not be updated yet.
	Status ChangeStatusPtrInput
}

func (ChangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*changeArgs)(nil)).Elem()
}

type ChangeInput interface {
	pulumi.Input

	ToChangeOutput() ChangeOutput
	ToChangeOutputWithContext(ctx context.Context) ChangeOutput
}

func (*Change) ElementType() reflect.Type {
	return reflect.TypeOf((**Change)(nil)).Elem()
}

func (i *Change) ToChangeOutput() ChangeOutput {
	return i.ToChangeOutputWithContext(context.Background())
}

func (i *Change) ToChangeOutputWithContext(ctx context.Context) ChangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChangeOutput)
}

type ChangeOutput struct{ *pulumi.OutputState }

func (ChangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Change)(nil)).Elem()
}

func (o ChangeOutput) ToChangeOutput() ChangeOutput {
	return o
}

func (o ChangeOutput) ToChangeOutputWithContext(ctx context.Context) ChangeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChangeInput)(nil)).Elem(), &Change{})
	pulumi.RegisterOutputType(ChangeOutput{})
}
