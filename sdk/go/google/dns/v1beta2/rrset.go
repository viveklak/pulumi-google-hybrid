// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ResourceRecordSet.
type Rrset struct {
	pulumi.CustomResourceState

	Kind pulumi.StringOutput `pulumi:"kind"`
	// For example, www.example.com.
	Name pulumi.StringOutput `pulumi:"name"`
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
	Rrdatas pulumi.StringArrayOutput `pulumi:"rrdatas"`
	// As defined in RFC 4034 (section 3.2).
	SignatureRrdatas pulumi.StringArrayOutput `pulumi:"signatureRrdatas"`
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The identifier of a supported record type. See the list of Supported DNS record types.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRrset registers a new resource with the given unique name, arguments, and options.
func NewRrset(ctx *pulumi.Context,
	name string, args *RrsetArgs, opts ...pulumi.ResourceOption) (*Rrset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedZone == nil {
		return nil, errors.New("invalid value for required argument 'ManagedZone'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Rrset
	err := ctx.RegisterResource("google-native:dns/v1beta2:Rrset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRrset gets an existing Rrset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRrset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RrsetState, opts ...pulumi.ResourceOption) (*Rrset, error) {
	var resource Rrset
	err := ctx.ReadResource("google-native:dns/v1beta2:Rrset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rrset resources.
type rrsetState struct {
}

type RrsetState struct {
}

func (RrsetState) ElementType() reflect.Type {
	return reflect.TypeOf((*rrsetState)(nil)).Elem()
}

type rrsetArgs struct {
	ClientOperationId *string `pulumi:"clientOperationId"`
	Kind              *string `pulumi:"kind"`
	ManagedZone       string  `pulumi:"managedZone"`
	// For example, www.example.com.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
	Rrdatas []string `pulumi:"rrdatas"`
	// As defined in RFC 4034 (section 3.2).
	SignatureRrdatas []string `pulumi:"signatureRrdatas"`
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	Ttl *int `pulumi:"ttl"`
	// The identifier of a supported record type. See the list of Supported DNS record types.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Rrset resource.
type RrsetArgs struct {
	ClientOperationId pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	ManagedZone       pulumi.StringInput
	// For example, www.example.com.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
	Rrdatas pulumi.StringArrayInput
	// As defined in RFC 4034 (section 3.2).
	SignatureRrdatas pulumi.StringArrayInput
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	Ttl pulumi.IntPtrInput
	// The identifier of a supported record type. See the list of Supported DNS record types.
	Type pulumi.StringPtrInput
}

func (RrsetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rrsetArgs)(nil)).Elem()
}

type RrsetInput interface {
	pulumi.Input

	ToRrsetOutput() RrsetOutput
	ToRrsetOutputWithContext(ctx context.Context) RrsetOutput
}

func (*Rrset) ElementType() reflect.Type {
	return reflect.TypeOf((*Rrset)(nil))
}

func (i *Rrset) ToRrsetOutput() RrsetOutput {
	return i.ToRrsetOutputWithContext(context.Background())
}

func (i *Rrset) ToRrsetOutputWithContext(ctx context.Context) RrsetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RrsetOutput)
}

type RrsetOutput struct {
	*pulumi.OutputState
}

func (RrsetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rrset)(nil))
}

func (o RrsetOutput) ToRrsetOutput() RrsetOutput {
	return o
}

func (o RrsetOutput) ToRrsetOutputWithContext(ctx context.Context) RrsetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RrsetOutput{})
}
