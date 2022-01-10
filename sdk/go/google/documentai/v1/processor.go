// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a processor from the type processor that the user chose. The processor will be at "ENABLED" state by default after its creation.
// Auto-naming is currently not supported for this resource.
type Processor struct {
	pulumi.CustomResourceState

	// The time the processor was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The default processor version.
	DefaultProcessorVersion pulumi.StringOutput `pulumi:"defaultProcessorVersion"`
	// The display name of the processor.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringOutput `pulumi:"kmsKeyName"`
	// Immutable. The resource name of the processor. Format: `projects/{project}/locations/{location}/processors/{processor}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The http endpoint that can be called to invoke processing.
	ProcessEndpoint pulumi.StringOutput `pulumi:"processEndpoint"`
	// The state of the processor.
	State pulumi.StringOutput `pulumi:"state"`
	// The processor type, e.g., INVOICE_PARSING, W2_PARSING, etc.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProcessor registers a new resource with the given unique name, arguments, and options.
func NewProcessor(ctx *pulumi.Context,
	name string, args *ProcessorArgs, opts ...pulumi.ResourceOption) (*Processor, error) {
	if args == nil {
		args = &ProcessorArgs{}
	}

	var resource Processor
	err := ctx.RegisterResource("google-native:documentai/v1:Processor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProcessor gets an existing Processor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProcessor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProcessorState, opts ...pulumi.ResourceOption) (*Processor, error) {
	var resource Processor
	err := ctx.ReadResource("google-native:documentai/v1:Processor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Processor resources.
type processorState struct {
}

type ProcessorState struct {
}

func (ProcessorState) ElementType() reflect.Type {
	return reflect.TypeOf((*processorState)(nil)).Elem()
}

type processorArgs struct {
	// The time the processor was created.
	CreateTime *string `pulumi:"createTime"`
	// The default processor version.
	DefaultProcessorVersion *string `pulumi:"defaultProcessorVersion"`
	// The display name of the processor.
	DisplayName *string `pulumi:"displayName"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	Location   *string `pulumi:"location"`
	Project    *string `pulumi:"project"`
	// The processor type, e.g., INVOICE_PARSING, W2_PARSING, etc.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Processor resource.
type ProcessorArgs struct {
	// The time the processor was created.
	CreateTime pulumi.StringPtrInput
	// The default processor version.
	DefaultProcessorVersion pulumi.StringPtrInput
	// The display name of the processor.
	DisplayName pulumi.StringPtrInput
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringPtrInput
	Location   pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	// The processor type, e.g., INVOICE_PARSING, W2_PARSING, etc.
	Type pulumi.StringPtrInput
}

func (ProcessorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*processorArgs)(nil)).Elem()
}

type ProcessorInput interface {
	pulumi.Input

	ToProcessorOutput() ProcessorOutput
	ToProcessorOutputWithContext(ctx context.Context) ProcessorOutput
}

func (*Processor) ElementType() reflect.Type {
	return reflect.TypeOf((**Processor)(nil)).Elem()
}

func (i *Processor) ToProcessorOutput() ProcessorOutput {
	return i.ToProcessorOutputWithContext(context.Background())
}

func (i *Processor) ToProcessorOutputWithContext(ctx context.Context) ProcessorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProcessorOutput)
}

type ProcessorOutput struct{ *pulumi.OutputState }

func (ProcessorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Processor)(nil)).Elem()
}

func (o ProcessorOutput) ToProcessorOutput() ProcessorOutput {
	return o
}

func (o ProcessorOutput) ToProcessorOutputWithContext(ctx context.Context) ProcessorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProcessorInput)(nil)).Elem(), &Processor{})
	pulumi.RegisterOutputType(ProcessorOutput{})
}
