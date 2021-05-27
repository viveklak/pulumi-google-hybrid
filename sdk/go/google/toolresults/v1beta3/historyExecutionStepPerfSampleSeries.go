// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a PerfSampleSeries. May return any of the following error code(s): - ALREADY_EXISTS - PerfMetricSummary already exists for the given Step - NOT_FOUND - The containing Step does not exist
type HistoryExecutionStepPerfSampleSeries struct {
	pulumi.CustomResourceState

	// Basic series represented by a line chart
	BasicPerfSampleSeries BasicPerfSampleSeriesResponseOutput `pulumi:"basicPerfSampleSeries"`
	// A tool results execution ID. @OutputOnly
	ExecutionId pulumi.StringOutput `pulumi:"executionId"`
	// A tool results history ID. @OutputOnly
	HistoryId pulumi.StringOutput `pulumi:"historyId"`
	// The cloud project @OutputOnly
	Project pulumi.StringOutput `pulumi:"project"`
	// A sample series id @OutputOnly
	SampleSeriesId pulumi.StringOutput `pulumi:"sampleSeriesId"`
	// A tool results step ID. @OutputOnly
	StepId pulumi.StringOutput `pulumi:"stepId"`
}

// NewHistoryExecutionStepPerfSampleSeries registers a new resource with the given unique name, arguments, and options.
func NewHistoryExecutionStepPerfSampleSeries(ctx *pulumi.Context,
	name string, args *HistoryExecutionStepPerfSampleSeriesArgs, opts ...pulumi.ResourceOption) (*HistoryExecutionStepPerfSampleSeries, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExecutionId == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionId'")
	}
	if args.HistoryId == nil {
		return nil, errors.New("invalid value for required argument 'HistoryId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.StepId == nil {
		return nil, errors.New("invalid value for required argument 'StepId'")
	}
	var resource HistoryExecutionStepPerfSampleSeries
	err := ctx.RegisterResource("google-native:toolresults/v1beta3:HistoryExecutionStepPerfSampleSeries", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHistoryExecutionStepPerfSampleSeries gets an existing HistoryExecutionStepPerfSampleSeries resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHistoryExecutionStepPerfSampleSeries(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HistoryExecutionStepPerfSampleSeriesState, opts ...pulumi.ResourceOption) (*HistoryExecutionStepPerfSampleSeries, error) {
	var resource HistoryExecutionStepPerfSampleSeries
	err := ctx.ReadResource("google-native:toolresults/v1beta3:HistoryExecutionStepPerfSampleSeries", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HistoryExecutionStepPerfSampleSeries resources.
type historyExecutionStepPerfSampleSeriesState struct {
	// Basic series represented by a line chart
	BasicPerfSampleSeries *BasicPerfSampleSeriesResponse `pulumi:"basicPerfSampleSeries"`
	// A tool results execution ID. @OutputOnly
	ExecutionId *string `pulumi:"executionId"`
	// A tool results history ID. @OutputOnly
	HistoryId *string `pulumi:"historyId"`
	// The cloud project @OutputOnly
	Project *string `pulumi:"project"`
	// A sample series id @OutputOnly
	SampleSeriesId *string `pulumi:"sampleSeriesId"`
	// A tool results step ID. @OutputOnly
	StepId *string `pulumi:"stepId"`
}

type HistoryExecutionStepPerfSampleSeriesState struct {
	// Basic series represented by a line chart
	BasicPerfSampleSeries BasicPerfSampleSeriesResponsePtrInput
	// A tool results execution ID. @OutputOnly
	ExecutionId pulumi.StringPtrInput
	// A tool results history ID. @OutputOnly
	HistoryId pulumi.StringPtrInput
	// The cloud project @OutputOnly
	Project pulumi.StringPtrInput
	// A sample series id @OutputOnly
	SampleSeriesId pulumi.StringPtrInput
	// A tool results step ID. @OutputOnly
	StepId pulumi.StringPtrInput
}

func (HistoryExecutionStepPerfSampleSeriesState) ElementType() reflect.Type {
	return reflect.TypeOf((*historyExecutionStepPerfSampleSeriesState)(nil)).Elem()
}

type historyExecutionStepPerfSampleSeriesArgs struct {
	// Basic series represented by a line chart
	BasicPerfSampleSeries *BasicPerfSampleSeries `pulumi:"basicPerfSampleSeries"`
	// A tool results execution ID. @OutputOnly
	ExecutionId string `pulumi:"executionId"`
	// A tool results history ID. @OutputOnly
	HistoryId string `pulumi:"historyId"`
	// The cloud project @OutputOnly
	Project string `pulumi:"project"`
	// A sample series id @OutputOnly
	SampleSeriesId *string `pulumi:"sampleSeriesId"`
	// A tool results step ID. @OutputOnly
	StepId string `pulumi:"stepId"`
}

// The set of arguments for constructing a HistoryExecutionStepPerfSampleSeries resource.
type HistoryExecutionStepPerfSampleSeriesArgs struct {
	// Basic series represented by a line chart
	BasicPerfSampleSeries BasicPerfSampleSeriesPtrInput
	// A tool results execution ID. @OutputOnly
	ExecutionId pulumi.StringInput
	// A tool results history ID. @OutputOnly
	HistoryId pulumi.StringInput
	// The cloud project @OutputOnly
	Project pulumi.StringInput
	// A sample series id @OutputOnly
	SampleSeriesId pulumi.StringPtrInput
	// A tool results step ID. @OutputOnly
	StepId pulumi.StringInput
}

func (HistoryExecutionStepPerfSampleSeriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*historyExecutionStepPerfSampleSeriesArgs)(nil)).Elem()
}

type HistoryExecutionStepPerfSampleSeriesInput interface {
	pulumi.Input

	ToHistoryExecutionStepPerfSampleSeriesOutput() HistoryExecutionStepPerfSampleSeriesOutput
	ToHistoryExecutionStepPerfSampleSeriesOutputWithContext(ctx context.Context) HistoryExecutionStepPerfSampleSeriesOutput
}

func (*HistoryExecutionStepPerfSampleSeries) ElementType() reflect.Type {
	return reflect.TypeOf((*HistoryExecutionStepPerfSampleSeries)(nil))
}

func (i *HistoryExecutionStepPerfSampleSeries) ToHistoryExecutionStepPerfSampleSeriesOutput() HistoryExecutionStepPerfSampleSeriesOutput {
	return i.ToHistoryExecutionStepPerfSampleSeriesOutputWithContext(context.Background())
}

func (i *HistoryExecutionStepPerfSampleSeries) ToHistoryExecutionStepPerfSampleSeriesOutputWithContext(ctx context.Context) HistoryExecutionStepPerfSampleSeriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HistoryExecutionStepPerfSampleSeriesOutput)
}

type HistoryExecutionStepPerfSampleSeriesOutput struct {
	*pulumi.OutputState
}

func (HistoryExecutionStepPerfSampleSeriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HistoryExecutionStepPerfSampleSeries)(nil))
}

func (o HistoryExecutionStepPerfSampleSeriesOutput) ToHistoryExecutionStepPerfSampleSeriesOutput() HistoryExecutionStepPerfSampleSeriesOutput {
	return o
}

func (o HistoryExecutionStepPerfSampleSeriesOutput) ToHistoryExecutionStepPerfSampleSeriesOutputWithContext(ctx context.Context) HistoryExecutionStepPerfSampleSeriesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HistoryExecutionStepPerfSampleSeriesOutput{})
}
