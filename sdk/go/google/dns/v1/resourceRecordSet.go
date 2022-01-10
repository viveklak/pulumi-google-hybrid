// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ResourceRecordSet.
type ResourceRecordSet struct {
	pulumi.CustomResourceState

	Kind pulumi.StringOutput `pulumi:"kind"`
	// For example, www.example.com.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
	RoutingPolicy RRSetRoutingPolicyResponseOutput `pulumi:"routingPolicy"`
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
	Rrdatas pulumi.StringArrayOutput `pulumi:"rrdatas"`
	// As defined in RFC 4034 (section 3.2).
	SignatureRrdatas pulumi.StringArrayOutput `pulumi:"signatureRrdatas"`
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The identifier of a supported record type. See the list of Supported DNS record types.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResourceRecordSet registers a new resource with the given unique name, arguments, and options.
func NewResourceRecordSet(ctx *pulumi.Context,
	name string, args *ResourceRecordSetArgs, opts ...pulumi.ResourceOption) (*ResourceRecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedZone == nil {
		return nil, errors.New("invalid value for required argument 'ManagedZone'")
	}
	var resource ResourceRecordSet
	err := ctx.RegisterResource("google-native:dns/v1:ResourceRecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceRecordSet gets an existing ResourceRecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceRecordSetState, opts ...pulumi.ResourceOption) (*ResourceRecordSet, error) {
	var resource ResourceRecordSet
	err := ctx.ReadResource("google-native:dns/v1:ResourceRecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceRecordSet resources.
type resourceRecordSetState struct {
}

type ResourceRecordSetState struct {
}

func (ResourceRecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceRecordSetState)(nil)).Elem()
}

type resourceRecordSetArgs struct {
	ClientOperationId *string `pulumi:"clientOperationId"`
	Kind              *string `pulumi:"kind"`
	ManagedZone       string  `pulumi:"managedZone"`
	// For example, www.example.com.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
	RoutingPolicy *RRSetRoutingPolicy `pulumi:"routingPolicy"`
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
	Rrdatas []string `pulumi:"rrdatas"`
	// As defined in RFC 4034 (section 3.2).
	SignatureRrdatas []string `pulumi:"signatureRrdatas"`
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	Ttl *int `pulumi:"ttl"`
	// The identifier of a supported record type. See the list of Supported DNS record types.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ResourceRecordSet resource.
type ResourceRecordSetArgs struct {
	ClientOperationId pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	ManagedZone       pulumi.StringInput
	// For example, www.example.com.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
	RoutingPolicy RRSetRoutingPolicyPtrInput
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
	Rrdatas pulumi.StringArrayInput
	// As defined in RFC 4034 (section 3.2).
	SignatureRrdatas pulumi.StringArrayInput
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	Ttl pulumi.IntPtrInput
	// The identifier of a supported record type. See the list of Supported DNS record types.
	Type pulumi.StringPtrInput
}

func (ResourceRecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceRecordSetArgs)(nil)).Elem()
}

type ResourceRecordSetInput interface {
	pulumi.Input

	ToResourceRecordSetOutput() ResourceRecordSetOutput
	ToResourceRecordSetOutputWithContext(ctx context.Context) ResourceRecordSetOutput
}

func (*ResourceRecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRecordSet)(nil)).Elem()
}

func (i *ResourceRecordSet) ToResourceRecordSetOutput() ResourceRecordSetOutput {
	return i.ToResourceRecordSetOutputWithContext(context.Background())
}

func (i *ResourceRecordSet) ToResourceRecordSetOutputWithContext(ctx context.Context) ResourceRecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRecordSetOutput)
}

type ResourceRecordSetOutput struct{ *pulumi.OutputState }

func (ResourceRecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRecordSet)(nil)).Elem()
}

func (o ResourceRecordSetOutput) ToResourceRecordSetOutput() ResourceRecordSetOutput {
	return o
}

func (o ResourceRecordSetOutput) ToResourceRecordSetOutputWithContext(ctx context.Context) ResourceRecordSetOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceRecordSetInput)(nil)).Elem(), &ResourceRecordSet{})
	pulumi.RegisterOutputType(ResourceRecordSetOutput{})
}
