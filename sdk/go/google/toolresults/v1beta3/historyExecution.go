// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Execution. The returned Execution will have the id set. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the containing History does not exist
type HistoryExecution struct {
	pulumi.CustomResourceState

	// The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
	CompletionTime TimestampResponseOutput `pulumi:"completionTime"`
	// The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
	CreationTime TimestampResponseOutput `pulumi:"creationTime"`
	// The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
	DimensionDefinitions MatrixDimensionDefinitionResponseArrayOutput `pulumi:"dimensionDefinitions"`
	// A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
	ExecutionId pulumi.StringOutput `pulumi:"executionId"`
	// Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
	Outcome OutcomeResponseOutput `pulumi:"outcome"`
	// Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
	Specification SpecificationResponseOutput `pulumi:"specification"`
	// The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
	State pulumi.StringOutput `pulumi:"state"`
	// TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
	TestExecutionMatrixId pulumi.StringOutput `pulumi:"testExecutionMatrixId"`
}

// NewHistoryExecution registers a new resource with the given unique name, arguments, and options.
func NewHistoryExecution(ctx *pulumi.Context,
	name string, args *HistoryExecutionArgs, opts ...pulumi.ResourceOption) (*HistoryExecution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HistoryId == nil {
		return nil, errors.New("invalid value for required argument 'HistoryId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource HistoryExecution
	err := ctx.RegisterResource("google-native:toolresults/v1beta3:HistoryExecution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHistoryExecution gets an existing HistoryExecution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHistoryExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HistoryExecutionState, opts ...pulumi.ResourceOption) (*HistoryExecution, error) {
	var resource HistoryExecution
	err := ctx.ReadResource("google-native:toolresults/v1beta3:HistoryExecution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HistoryExecution resources.
type historyExecutionState struct {
	// The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
	CompletionTime *TimestampResponse `pulumi:"completionTime"`
	// The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
	CreationTime *TimestampResponse `pulumi:"creationTime"`
	// The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
	DimensionDefinitions []MatrixDimensionDefinitionResponse `pulumi:"dimensionDefinitions"`
	// A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
	ExecutionId *string `pulumi:"executionId"`
	// Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
	Outcome *OutcomeResponse `pulumi:"outcome"`
	// Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
	Specification *SpecificationResponse `pulumi:"specification"`
	// The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
	State *string `pulumi:"state"`
	// TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
	TestExecutionMatrixId *string `pulumi:"testExecutionMatrixId"`
}

type HistoryExecutionState struct {
	// The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
	CompletionTime TimestampResponsePtrInput
	// The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
	CreationTime TimestampResponsePtrInput
	// The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
	DimensionDefinitions MatrixDimensionDefinitionResponseArrayInput
	// A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
	ExecutionId pulumi.StringPtrInput
	// Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
	Outcome OutcomeResponsePtrInput
	// Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
	Specification SpecificationResponsePtrInput
	// The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
	State pulumi.StringPtrInput
	// TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
	TestExecutionMatrixId pulumi.StringPtrInput
}

func (HistoryExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*historyExecutionState)(nil)).Elem()
}

type historyExecutionArgs struct {
	// The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
	CompletionTime *Timestamp `pulumi:"completionTime"`
	// The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
	CreationTime *Timestamp `pulumi:"creationTime"`
	// The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
	DimensionDefinitions []MatrixDimensionDefinition `pulumi:"dimensionDefinitions"`
	// A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
	ExecutionId *string `pulumi:"executionId"`
	HistoryId   string  `pulumi:"historyId"`
	// Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
	Outcome   *Outcome `pulumi:"outcome"`
	Project   string   `pulumi:"project"`
	RequestId *string  `pulumi:"requestId"`
	// Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
	Specification *Specification `pulumi:"specification"`
	// The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
	State *string `pulumi:"state"`
	// TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
	TestExecutionMatrixId *string `pulumi:"testExecutionMatrixId"`
}

// The set of arguments for constructing a HistoryExecution resource.
type HistoryExecutionArgs struct {
	// The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
	CompletionTime TimestampPtrInput
	// The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
	CreationTime TimestampPtrInput
	// The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
	DimensionDefinitions MatrixDimensionDefinitionArrayInput
	// A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
	ExecutionId pulumi.StringPtrInput
	HistoryId   pulumi.StringInput
	// Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
	Outcome   OutcomePtrInput
	Project   pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
	Specification SpecificationPtrInput
	// The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
	State pulumi.StringPtrInput
	// TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
	TestExecutionMatrixId pulumi.StringPtrInput
}

func (HistoryExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*historyExecutionArgs)(nil)).Elem()
}

type HistoryExecutionInput interface {
	pulumi.Input

	ToHistoryExecutionOutput() HistoryExecutionOutput
	ToHistoryExecutionOutputWithContext(ctx context.Context) HistoryExecutionOutput
}

func (*HistoryExecution) ElementType() reflect.Type {
	return reflect.TypeOf((*HistoryExecution)(nil))
}

func (i *HistoryExecution) ToHistoryExecutionOutput() HistoryExecutionOutput {
	return i.ToHistoryExecutionOutputWithContext(context.Background())
}

func (i *HistoryExecution) ToHistoryExecutionOutputWithContext(ctx context.Context) HistoryExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HistoryExecutionOutput)
}

type HistoryExecutionOutput struct {
	*pulumi.OutputState
}

func (HistoryExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HistoryExecution)(nil))
}

func (o HistoryExecutionOutput) ToHistoryExecutionOutput() HistoryExecutionOutput {
	return o
}

func (o HistoryExecutionOutput) ToHistoryExecutionOutputWithContext(ctx context.Context) HistoryExecutionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HistoryExecutionOutput{})
}
