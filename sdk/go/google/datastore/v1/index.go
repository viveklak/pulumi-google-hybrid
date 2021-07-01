// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates the specified index. A newly created index's initial state is `CREATING`. On completion of the returned google.longrunning.Operation, the state will be `READY`. If the index already exists, the call will return an `ALREADY_EXISTS` status. During index creation, the process could result in an error, in which case the index will move to the `ERROR` state. The process can be recovered by fixing the data that caused the error, removing the index with delete, then re-creating the index with create. Indexes with a single property cannot be created.
type Index struct {
	pulumi.CustomResourceState

	// Required. The index's ancestor mode. Must not be ANCESTOR_MODE_UNSPECIFIED.
	Ancestor pulumi.StringOutput `pulumi:"ancestor"`
	// The resource ID of the index.
	IndexId pulumi.StringOutput `pulumi:"indexId"`
	// Required. The entity kind to which this index applies.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Project ID.
	Project pulumi.StringOutput `pulumi:"project"`
	// Required. An ordered sequence of property names and their index attributes.
	Properties GoogleDatastoreAdminV1IndexedPropertyResponseArrayOutput `pulumi:"properties"`
	// The state of the index.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Index
	err := ctx.RegisterResource("google-native:datastore/v1:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("google-native:datastore/v1:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
}

type IndexState struct {
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// Required. The index's ancestor mode. Must not be ANCESTOR_MODE_UNSPECIFIED.
	Ancestor *string `pulumi:"ancestor"`
	// Required. The entity kind to which this index applies.
	Kind    *string `pulumi:"kind"`
	Project string  `pulumi:"project"`
	// Required. An ordered sequence of property names and their index attributes.
	Properties []GoogleDatastoreAdminV1IndexedProperty `pulumi:"properties"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// Required. The index's ancestor mode. Must not be ANCESTOR_MODE_UNSPECIFIED.
	Ancestor *IndexAncestor
	// Required. The entity kind to which this index applies.
	Kind    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Required. An ordered sequence of property names and their index attributes.
	Properties GoogleDatastoreAdminV1IndexedPropertyArrayInput
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexInput interface {
	pulumi.Input

	ToIndexOutput() IndexOutput
	ToIndexOutputWithContext(ctx context.Context) IndexOutput
}

func (*Index) ElementType() reflect.Type {
	return reflect.TypeOf((*Index)(nil))
}

func (i *Index) ToIndexOutput() IndexOutput {
	return i.ToIndexOutputWithContext(context.Background())
}

func (i *Index) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexOutput)
}

type IndexOutput struct {
	*pulumi.OutputState
}

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Index)(nil))
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IndexOutput{})
}
