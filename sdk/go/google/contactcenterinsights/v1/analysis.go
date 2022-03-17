// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an analysis. The long running operation is done when the analysis has completed.
type Analysis struct {
	pulumi.CustomResourceState

	// The result of the analysis, which is populated when the analysis finishes.
	AnalysisResult GoogleCloudContactcenterinsightsV1AnalysisResultResponseOutput `pulumi:"analysisResult"`
	// The time at which the analysis was created, which occurs when the long-running operation completes.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Immutable. The resource name of the analysis. Format: projects/{project}/locations/{location}/conversations/{conversation}/analyses/{analysis}
	Name pulumi.StringOutput `pulumi:"name"`
	// The time at which the analysis was requested.
	RequestTime pulumi.StringOutput `pulumi:"requestTime"`
}

// NewAnalysis registers a new resource with the given unique name, arguments, and options.
func NewAnalysis(ctx *pulumi.Context,
	name string, args *AnalysisArgs, opts ...pulumi.ResourceOption) (*Analysis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConversationId == nil {
		return nil, errors.New("invalid value for required argument 'ConversationId'")
	}
	var resource Analysis
	err := ctx.RegisterResource("google-native:contactcenterinsights/v1:Analysis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalysis gets an existing Analysis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalysis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalysisState, opts ...pulumi.ResourceOption) (*Analysis, error) {
	var resource Analysis
	err := ctx.ReadResource("google-native:contactcenterinsights/v1:Analysis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Analysis resources.
type analysisState struct {
}

type AnalysisState struct {
}

func (AnalysisState) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisState)(nil)).Elem()
}

type analysisArgs struct {
	ConversationId string  `pulumi:"conversationId"`
	Location       *string `pulumi:"location"`
	// Immutable. The resource name of the analysis. Format: projects/{project}/locations/{location}/conversations/{conversation}/analyses/{analysis}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Analysis resource.
type AnalysisArgs struct {
	ConversationId pulumi.StringInput
	Location       pulumi.StringPtrInput
	// Immutable. The resource name of the analysis. Format: projects/{project}/locations/{location}/conversations/{conversation}/analyses/{analysis}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (AnalysisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisArgs)(nil)).Elem()
}

type AnalysisInput interface {
	pulumi.Input

	ToAnalysisOutput() AnalysisOutput
	ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput
}

func (*Analysis) ElementType() reflect.Type {
	return reflect.TypeOf((**Analysis)(nil)).Elem()
}

func (i *Analysis) ToAnalysisOutput() AnalysisOutput {
	return i.ToAnalysisOutputWithContext(context.Background())
}

func (i *Analysis) ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalysisOutput)
}

type AnalysisOutput struct{ *pulumi.OutputState }

func (AnalysisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Analysis)(nil)).Elem()
}

func (o AnalysisOutput) ToAnalysisOutput() AnalysisOutput {
	return o
}

func (o AnalysisOutput) ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalysisInput)(nil)).Elem(), &Analysis{})
	pulumi.RegisterOutputType(AnalysisOutput{})
}
