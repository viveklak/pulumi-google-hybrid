// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a sink that exports trace spans to a destination. The export of newly-ingested traces begins immediately, unless the sink's `writer_identity` is not permitted to write to the destination. A sink can export traces only from the resource owning the sink (the 'parent').
type TraceSink struct {
	pulumi.CustomResourceState

	// The canonical sink resource name, unique within the project. Must be of the form: project/[PROJECT_NUMBER]/traceSinks/[SINK_ID]. E.g.: `"projects/12345/traceSinks/my-project-trace-sink"`. Sink identifiers are limited to 256 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods.
	Name pulumi.StringOutput `pulumi:"name"`
	// The export destination.
	OutputConfig OutputConfigResponseOutput `pulumi:"outputConfig"`
	// A service account name for exporting the data. This field is set by sinks.create and sinks.update. The service account will need to be granted write access to the destination specified in the output configuration, see [Granting access for a resource](/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). To create tables and write data this account will need the dataEditor role. Read more about roles in the [BigQuery documentation](https://cloud.google.com/bigquery/docs/access-control). E.g.: "service-00000001@00000002.iam.gserviceaccount.com"
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewTraceSink registers a new resource with the given unique name, arguments, and options.
func NewTraceSink(ctx *pulumi.Context,
	name string, args *TraceSinkArgs, opts ...pulumi.ResourceOption) (*TraceSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OutputConfig == nil {
		return nil, errors.New("invalid value for required argument 'OutputConfig'")
	}
	var resource TraceSink
	err := ctx.RegisterResource("google-native:cloudtrace/v2beta1:TraceSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTraceSink gets an existing TraceSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTraceSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TraceSinkState, opts ...pulumi.ResourceOption) (*TraceSink, error) {
	var resource TraceSink
	err := ctx.ReadResource("google-native:cloudtrace/v2beta1:TraceSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TraceSink resources.
type traceSinkState struct {
}

type TraceSinkState struct {
}

func (TraceSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*traceSinkState)(nil)).Elem()
}

type traceSinkArgs struct {
	// The canonical sink resource name, unique within the project. Must be of the form: project/[PROJECT_NUMBER]/traceSinks/[SINK_ID]. E.g.: `"projects/12345/traceSinks/my-project-trace-sink"`. Sink identifiers are limited to 256 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods.
	Name *string `pulumi:"name"`
	// The export destination.
	OutputConfig OutputConfig `pulumi:"outputConfig"`
	Project      *string      `pulumi:"project"`
}

// The set of arguments for constructing a TraceSink resource.
type TraceSinkArgs struct {
	// The canonical sink resource name, unique within the project. Must be of the form: project/[PROJECT_NUMBER]/traceSinks/[SINK_ID]. E.g.: `"projects/12345/traceSinks/my-project-trace-sink"`. Sink identifiers are limited to 256 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods.
	Name pulumi.StringPtrInput
	// The export destination.
	OutputConfig OutputConfigInput
	Project      pulumi.StringPtrInput
}

func (TraceSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*traceSinkArgs)(nil)).Elem()
}

type TraceSinkInput interface {
	pulumi.Input

	ToTraceSinkOutput() TraceSinkOutput
	ToTraceSinkOutputWithContext(ctx context.Context) TraceSinkOutput
}

func (*TraceSink) ElementType() reflect.Type {
	return reflect.TypeOf((**TraceSink)(nil)).Elem()
}

func (i *TraceSink) ToTraceSinkOutput() TraceSinkOutput {
	return i.ToTraceSinkOutputWithContext(context.Background())
}

func (i *TraceSink) ToTraceSinkOutputWithContext(ctx context.Context) TraceSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TraceSinkOutput)
}

type TraceSinkOutput struct{ *pulumi.OutputState }

func (TraceSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TraceSink)(nil)).Elem()
}

func (o TraceSinkOutput) ToTraceSinkOutput() TraceSinkOutput {
	return o
}

func (o TraceSinkOutput) ToTraceSinkOutputWithContext(ctx context.Context) TraceSinkOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TraceSinkInput)(nil)).Elem(), &TraceSink{})
	pulumi.RegisterOutputType(TraceSinkOutput{})
}
