// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this method to create a stream.
// Auto-naming is currently not supported for this resource.
type Stream struct {
	pulumi.CustomResourceState

	// Automatically backfill objects included in the stream source configuration. Specific objects can be excluded.
	BackfillAll BackfillAllStrategyResponseOutput `pulumi:"backfillAll"`
	// Do not automatically backfill any objects.
	BackfillNone BackfillNoneStrategyResponseOutput `pulumi:"backfillNone"`
	// The creation time of the stream.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey pulumi.StringOutput `pulumi:"customerManagedEncryptionKey"`
	// Destination connection profile configuration.
	DestinationConfig DestinationConfigResponseOutput `pulumi:"destinationConfig"`
	// Display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Errors on the Stream.
	Errors ErrorResponseArrayOutput `pulumi:"errors"`
	// Labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The stream's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Source connection profile configuration.
	SourceConfig SourceConfigResponseOutput `pulumi:"sourceConfig"`
	// The state of the stream.
	State pulumi.StringOutput `pulumi:"state"`
	// The last update time of the stream.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewStream registers a new resource with the given unique name, arguments, and options.
func NewStream(ctx *pulumi.Context,
	name string, args *StreamArgs, opts ...pulumi.ResourceOption) (*Stream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationConfig == nil {
		return nil, errors.New("invalid value for required argument 'DestinationConfig'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.SourceConfig == nil {
		return nil, errors.New("invalid value for required argument 'SourceConfig'")
	}
	if args.StreamId == nil {
		return nil, errors.New("invalid value for required argument 'StreamId'")
	}
	var resource Stream
	err := ctx.RegisterResource("google-native:datastream/v1:Stream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStream gets an existing Stream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamState, opts ...pulumi.ResourceOption) (*Stream, error) {
	var resource Stream
	err := ctx.ReadResource("google-native:datastream/v1:Stream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stream resources.
type streamState struct {
}

type StreamState struct {
}

func (StreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamState)(nil)).Elem()
}

type streamArgs struct {
	// Automatically backfill objects included in the stream source configuration. Specific objects can be excluded.
	BackfillAll *BackfillAllStrategy `pulumi:"backfillAll"`
	// Do not automatically backfill any objects.
	BackfillNone *BackfillNoneStrategy `pulumi:"backfillNone"`
	// Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey *string `pulumi:"customerManagedEncryptionKey"`
	// Destination connection profile configuration.
	DestinationConfig DestinationConfig `pulumi:"destinationConfig"`
	// Display name.
	DisplayName string `pulumi:"displayName"`
	// Optional. Create the stream without validating it.
	Force *string `pulumi:"force"`
	// Labels.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Source connection profile configuration.
	SourceConfig SourceConfig `pulumi:"sourceConfig"`
	// The state of the stream.
	State *StreamStateEnum `pulumi:"state"`
	// Required. The stream identifier.
	StreamId string `pulumi:"streamId"`
	// Optional. Only validate the stream, but don't create any resources. The default is false.
	ValidateOnly *string `pulumi:"validateOnly"`
}

// The set of arguments for constructing a Stream resource.
type StreamArgs struct {
	// Automatically backfill objects included in the stream source configuration. Specific objects can be excluded.
	BackfillAll BackfillAllStrategyPtrInput
	// Do not automatically backfill any objects.
	BackfillNone BackfillNoneStrategyPtrInput
	// Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
	CustomerManagedEncryptionKey pulumi.StringPtrInput
	// Destination connection profile configuration.
	DestinationConfig DestinationConfigInput
	// Display name.
	DisplayName pulumi.StringInput
	// Optional. Create the stream without validating it.
	Force pulumi.StringPtrInput
	// Labels.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Source connection profile configuration.
	SourceConfig SourceConfigInput
	// The state of the stream.
	State StreamStateEnumPtrInput
	// Required. The stream identifier.
	StreamId pulumi.StringInput
	// Optional. Only validate the stream, but don't create any resources. The default is false.
	ValidateOnly pulumi.StringPtrInput
}

func (StreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamArgs)(nil)).Elem()
}

type StreamInput interface {
	pulumi.Input

	ToStreamOutput() StreamOutput
	ToStreamOutputWithContext(ctx context.Context) StreamOutput
}

func (*Stream) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (i *Stream) ToStreamOutput() StreamOutput {
	return i.ToStreamOutputWithContext(context.Background())
}

func (i *Stream) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamOutput)
}

type StreamOutput struct{ *pulumi.OutputState }

func (StreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (o StreamOutput) ToStreamOutput() StreamOutput {
	return o
}

func (o StreamOutput) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInput)(nil)).Elem(), &Stream{})
	pulumi.RegisterOutputType(StreamOutput{})
}