// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Experiment in the specified Environment.
type AgentEnvironmentExperiment struct {
	pulumi.CustomResourceState

	// Creation time of this experiment.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The definition of the experiment.
	Definition GoogleCloudDialogflowCxV3beta1ExperimentDefinitionResponseOutput `pulumi:"definition"`
	// The human-readable description of the experiment.
	Description pulumi.StringOutput `pulumi:"description"`
	// Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// End time of this experiment.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength pulumi.StringOutput `pulumi:"experimentLength"`
	// Last update time of this experiment.
	LastUpdateTime pulumi.StringOutput `pulumi:"lastUpdateTime"`
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name pulumi.StringOutput `pulumi:"name"`
	// Inference result of the experiment.
	Result GoogleCloudDialogflowCxV3beta1ExperimentResultResponseOutput `pulumi:"result"`
	// Start time of this experiment.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
	State pulumi.StringOutput `pulumi:"state"`
	// The history of updates to the experiment variants.
	VariantsHistory GoogleCloudDialogflowCxV3beta1VariantsHistoryResponseArrayOutput `pulumi:"variantsHistory"`
}

// NewAgentEnvironmentExperiment registers a new resource with the given unique name, arguments, and options.
func NewAgentEnvironmentExperiment(ctx *pulumi.Context,
	name string, args *AgentEnvironmentExperimentArgs, opts ...pulumi.ResourceOption) (*AgentEnvironmentExperiment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource AgentEnvironmentExperiment
	err := ctx.RegisterResource("google-native:dialogflow/v3beta1:AgentEnvironmentExperiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentEnvironmentExperiment gets an existing AgentEnvironmentExperiment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentEnvironmentExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentEnvironmentExperimentState, opts ...pulumi.ResourceOption) (*AgentEnvironmentExperiment, error) {
	var resource AgentEnvironmentExperiment
	err := ctx.ReadResource("google-native:dialogflow/v3beta1:AgentEnvironmentExperiment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentEnvironmentExperiment resources.
type agentEnvironmentExperimentState struct {
	// Creation time of this experiment.
	CreateTime *string `pulumi:"createTime"`
	// The definition of the experiment.
	Definition *GoogleCloudDialogflowCxV3beta1ExperimentDefinitionResponse `pulumi:"definition"`
	// The human-readable description of the experiment.
	Description *string `pulumi:"description"`
	// Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// End time of this experiment.
	EndTime *string `pulumi:"endTime"`
	// Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength *string `pulumi:"experimentLength"`
	// Last update time of this experiment.
	LastUpdateTime *string `pulumi:"lastUpdateTime"`
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name *string `pulumi:"name"`
	// Inference result of the experiment.
	Result *GoogleCloudDialogflowCxV3beta1ExperimentResultResponse `pulumi:"result"`
	// Start time of this experiment.
	StartTime *string `pulumi:"startTime"`
	// The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
	State *string `pulumi:"state"`
	// The history of updates to the experiment variants.
	VariantsHistory []GoogleCloudDialogflowCxV3beta1VariantsHistoryResponse `pulumi:"variantsHistory"`
}

type AgentEnvironmentExperimentState struct {
	// Creation time of this experiment.
	CreateTime pulumi.StringPtrInput
	// The definition of the experiment.
	Definition GoogleCloudDialogflowCxV3beta1ExperimentDefinitionResponsePtrInput
	// The human-readable description of the experiment.
	Description pulumi.StringPtrInput
	// Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName pulumi.StringPtrInput
	// End time of this experiment.
	EndTime pulumi.StringPtrInput
	// Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength pulumi.StringPtrInput
	// Last update time of this experiment.
	LastUpdateTime pulumi.StringPtrInput
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name pulumi.StringPtrInput
	// Inference result of the experiment.
	Result GoogleCloudDialogflowCxV3beta1ExperimentResultResponsePtrInput
	// Start time of this experiment.
	StartTime pulumi.StringPtrInput
	// The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
	State pulumi.StringPtrInput
	// The history of updates to the experiment variants.
	VariantsHistory GoogleCloudDialogflowCxV3beta1VariantsHistoryResponseArrayInput
}

func (AgentEnvironmentExperimentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEnvironmentExperimentState)(nil)).Elem()
}

type agentEnvironmentExperimentArgs struct {
	AgentId string `pulumi:"agentId"`
	// Creation time of this experiment.
	CreateTime *string `pulumi:"createTime"`
	// The definition of the experiment.
	Definition *GoogleCloudDialogflowCxV3beta1ExperimentDefinition `pulumi:"definition"`
	// The human-readable description of the experiment.
	Description *string `pulumi:"description"`
	// Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// End time of this experiment.
	EndTime       *string `pulumi:"endTime"`
	EnvironmentId string  `pulumi:"environmentId"`
	// Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength *string `pulumi:"experimentLength"`
	// Last update time of this experiment.
	LastUpdateTime *string `pulumi:"lastUpdateTime"`
	Location       string  `pulumi:"location"`
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Inference result of the experiment.
	Result *GoogleCloudDialogflowCxV3beta1ExperimentResult `pulumi:"result"`
	// Start time of this experiment.
	StartTime *string `pulumi:"startTime"`
	// The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
	State *string `pulumi:"state"`
	// The history of updates to the experiment variants.
	VariantsHistory []GoogleCloudDialogflowCxV3beta1VariantsHistory `pulumi:"variantsHistory"`
}

// The set of arguments for constructing a AgentEnvironmentExperiment resource.
type AgentEnvironmentExperimentArgs struct {
	AgentId pulumi.StringInput
	// Creation time of this experiment.
	CreateTime pulumi.StringPtrInput
	// The definition of the experiment.
	Definition GoogleCloudDialogflowCxV3beta1ExperimentDefinitionPtrInput
	// The human-readable description of the experiment.
	Description pulumi.StringPtrInput
	// Required. The human-readable name of the experiment (unique in an environment). Limit of 64 characters.
	DisplayName pulumi.StringPtrInput
	// End time of this experiment.
	EndTime       pulumi.StringPtrInput
	EnvironmentId pulumi.StringInput
	// Maximum number of days to run the experiment. If auto-rollout is not enabled, default value and maximum will be 30 days. If auto-rollout is enabled, default value and maximum will be 6 days.
	ExperimentLength pulumi.StringPtrInput
	// Last update time of this experiment.
	LastUpdateTime pulumi.StringPtrInput
	Location       pulumi.StringInput
	// The name of the experiment. Format: projects//locations//agents//environments//experiments/..
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Inference result of the experiment.
	Result GoogleCloudDialogflowCxV3beta1ExperimentResultPtrInput
	// Start time of this experiment.
	StartTime pulumi.StringPtrInput
	// The current state of the experiment. Transition triggered by Expriments.StartExperiment: PENDING->RUNNING. Transition triggered by Expriments.CancelExperiment: PENDING->CANCELLED or RUNNING->CANCELLED.
	State pulumi.StringPtrInput
	// The history of updates to the experiment variants.
	VariantsHistory GoogleCloudDialogflowCxV3beta1VariantsHistoryArrayInput
}

func (AgentEnvironmentExperimentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEnvironmentExperimentArgs)(nil)).Elem()
}

type AgentEnvironmentExperimentInput interface {
	pulumi.Input

	ToAgentEnvironmentExperimentOutput() AgentEnvironmentExperimentOutput
	ToAgentEnvironmentExperimentOutputWithContext(ctx context.Context) AgentEnvironmentExperimentOutput
}

func (*AgentEnvironmentExperiment) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEnvironmentExperiment)(nil))
}

func (i *AgentEnvironmentExperiment) ToAgentEnvironmentExperimentOutput() AgentEnvironmentExperimentOutput {
	return i.ToAgentEnvironmentExperimentOutputWithContext(context.Background())
}

func (i *AgentEnvironmentExperiment) ToAgentEnvironmentExperimentOutputWithContext(ctx context.Context) AgentEnvironmentExperimentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentEnvironmentExperimentOutput)
}

type AgentEnvironmentExperimentOutput struct {
	*pulumi.OutputState
}

func (AgentEnvironmentExperimentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEnvironmentExperiment)(nil))
}

func (o AgentEnvironmentExperimentOutput) ToAgentEnvironmentExperimentOutput() AgentEnvironmentExperimentOutput {
	return o
}

func (o AgentEnvironmentExperimentOutput) ToAgentEnvironmentExperimentOutputWithContext(ctx context.Context) AgentEnvironmentExperimentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentEnvironmentExperimentOutput{})
}
