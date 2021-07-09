// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an instruction for how data should be labeled.
type Instruction struct {
	pulumi.CustomResourceState

	// The names of any related resources that are blocking changes to the instruction.
	BlockingResources pulumi.StringArrayOutput `pulumi:"blockingResources"`
	// Creation time of instruction.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The data type of this instruction.
	DataType pulumi.StringOutput `pulumi:"dataType"`
	// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
	Description pulumi.StringOutput `pulumi:"description"`
	// The display name of the instruction. Maximum of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
	PdfInstruction GoogleCloudDatalabelingV1beta1PdfInstructionResponseOutput `pulumi:"pdfInstruction"`
	// Last update time of instruction.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewInstruction registers a new resource with the given unique name, arguments, and options.
func NewInstruction(ctx *pulumi.Context,
	name string, args *InstructionArgs, opts ...pulumi.ResourceOption) (*Instruction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Instruction
	err := ctx.RegisterResource("google-native:datalabeling/v1beta1:Instruction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstruction gets an existing Instruction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstruction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstructionState, opts ...pulumi.ResourceOption) (*Instruction, error) {
	var resource Instruction
	err := ctx.ReadResource("google-native:datalabeling/v1beta1:Instruction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instruction resources.
type instructionState struct {
	// The names of any related resources that are blocking changes to the instruction.
	BlockingResources []string `pulumi:"blockingResources"`
	// Creation time of instruction.
	CreateTime *string `pulumi:"createTime"`
	// The data type of this instruction.
	DataType *string `pulumi:"dataType"`
	// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
	Description *string `pulumi:"description"`
	// The display name of the instruction. Maximum of 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
	Name *string `pulumi:"name"`
	// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
	PdfInstruction *GoogleCloudDatalabelingV1beta1PdfInstructionResponse `pulumi:"pdfInstruction"`
	// Last update time of instruction.
	UpdateTime *string `pulumi:"updateTime"`
}

type InstructionState struct {
	// The names of any related resources that are blocking changes to the instruction.
	BlockingResources pulumi.StringArrayInput
	// Creation time of instruction.
	CreateTime pulumi.StringPtrInput
	// The data type of this instruction.
	DataType pulumi.StringPtrInput
	// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
	Description pulumi.StringPtrInput
	// The display name of the instruction. Maximum of 64 characters.
	DisplayName pulumi.StringPtrInput
	// Instruction resource name, format: projects/{project_id}/instructions/{instruction_id}
	Name pulumi.StringPtrInput
	// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
	PdfInstruction GoogleCloudDatalabelingV1beta1PdfInstructionResponsePtrInput
	// Last update time of instruction.
	UpdateTime pulumi.StringPtrInput
}

func (InstructionState) ElementType() reflect.Type {
	return reflect.TypeOf((*instructionState)(nil)).Elem()
}

type instructionArgs struct {
	// The data type of this instruction.
	DataType string `pulumi:"dataType"`
	// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
	Description *string `pulumi:"description"`
	// The display name of the instruction. Maximum of 64 characters.
	DisplayName string `pulumi:"displayName"`
	// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
	PdfInstruction *GoogleCloudDatalabelingV1beta1PdfInstruction `pulumi:"pdfInstruction"`
	Project        string                                        `pulumi:"project"`
}

// The set of arguments for constructing a Instruction resource.
type InstructionArgs struct {
	// The data type of this instruction.
	DataType InstructionDataType
	// Optional. User-provided description of the instruction. The description can be up to 10000 characters long.
	Description pulumi.StringPtrInput
	// The display name of the instruction. Maximum of 64 characters.
	DisplayName pulumi.StringInput
	// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
	PdfInstruction GoogleCloudDatalabelingV1beta1PdfInstructionPtrInput
	Project        pulumi.StringInput
}

func (InstructionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instructionArgs)(nil)).Elem()
}

type InstructionInput interface {
	pulumi.Input

	ToInstructionOutput() InstructionOutput
	ToInstructionOutputWithContext(ctx context.Context) InstructionOutput
}

func (*Instruction) ElementType() reflect.Type {
	return reflect.TypeOf((*Instruction)(nil))
}

func (i *Instruction) ToInstructionOutput() InstructionOutput {
	return i.ToInstructionOutputWithContext(context.Background())
}

func (i *Instruction) ToInstructionOutputWithContext(ctx context.Context) InstructionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstructionOutput)
}

type InstructionOutput struct {
	*pulumi.OutputState
}

func (InstructionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instruction)(nil))
}

func (o InstructionOutput) ToInstructionOutput() InstructionOutput {
	return o
}

func (o InstructionOutput) ToInstructionOutputWithContext(ctx context.Context) InstructionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstructionOutput{})
}
