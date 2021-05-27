// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a session entity type.
type AgentEnvironmentSessionEntityType struct {
	pulumi.CustomResourceState

	// Required. The collection of entities to override or supplement the custom entity type.
	Entities GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayOutput `pulumi:"entities"`
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode pulumi.StringOutput `pulumi:"entityOverrideMode"`
	// Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAgentEnvironmentSessionEntityType registers a new resource with the given unique name, arguments, and options.
func NewAgentEnvironmentSessionEntityType(ctx *pulumi.Context,
	name string, args *AgentEnvironmentSessionEntityTypeArgs, opts ...pulumi.ResourceOption) (*AgentEnvironmentSessionEntityType, error) {
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
	if args.SessionId == nil {
		return nil, errors.New("invalid value for required argument 'SessionId'")
	}
	var resource AgentEnvironmentSessionEntityType
	err := ctx.RegisterResource("google-native:dialogflow/v3beta1:AgentEnvironmentSessionEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentEnvironmentSessionEntityType gets an existing AgentEnvironmentSessionEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentEnvironmentSessionEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentEnvironmentSessionEntityTypeState, opts ...pulumi.ResourceOption) (*AgentEnvironmentSessionEntityType, error) {
	var resource AgentEnvironmentSessionEntityType
	err := ctx.ReadResource("google-native:dialogflow/v3beta1:AgentEnvironmentSessionEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentEnvironmentSessionEntityType resources.
type agentEnvironmentSessionEntityTypeState struct {
	// Required. The collection of entities to override or supplement the custom entity type.
	Entities []GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse `pulumi:"entities"`
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode *string `pulumi:"entityOverrideMode"`
	// Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name *string `pulumi:"name"`
}

type AgentEnvironmentSessionEntityTypeState struct {
	// Required. The collection of entities to override or supplement the custom entity type.
	Entities GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponseArrayInput
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode pulumi.StringPtrInput
	// Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name pulumi.StringPtrInput
}

func (AgentEnvironmentSessionEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEnvironmentSessionEntityTypeState)(nil)).Elem()
}

type agentEnvironmentSessionEntityTypeArgs struct {
	AgentId string `pulumi:"agentId"`
	// Required. The collection of entities to override or supplement the custom entity type.
	Entities []GoogleCloudDialogflowCxV3beta1EntityTypeEntity `pulumi:"entities"`
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode *string `pulumi:"entityOverrideMode"`
	EnvironmentId      string  `pulumi:"environmentId"`
	Location           string  `pulumi:"location"`
	// Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	SessionId string  `pulumi:"sessionId"`
}

// The set of arguments for constructing a AgentEnvironmentSessionEntityType resource.
type AgentEnvironmentSessionEntityTypeArgs struct {
	AgentId pulumi.StringInput
	// Required. The collection of entities to override or supplement the custom entity type.
	Entities GoogleCloudDialogflowCxV3beta1EntityTypeEntityArrayInput
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode pulumi.StringPtrInput
	EnvironmentId      pulumi.StringInput
	Location           pulumi.StringInput
	// Required. The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	SessionId pulumi.StringInput
}

func (AgentEnvironmentSessionEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEnvironmentSessionEntityTypeArgs)(nil)).Elem()
}

type AgentEnvironmentSessionEntityTypeInput interface {
	pulumi.Input

	ToAgentEnvironmentSessionEntityTypeOutput() AgentEnvironmentSessionEntityTypeOutput
	ToAgentEnvironmentSessionEntityTypeOutputWithContext(ctx context.Context) AgentEnvironmentSessionEntityTypeOutput
}

func (*AgentEnvironmentSessionEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEnvironmentSessionEntityType)(nil))
}

func (i *AgentEnvironmentSessionEntityType) ToAgentEnvironmentSessionEntityTypeOutput() AgentEnvironmentSessionEntityTypeOutput {
	return i.ToAgentEnvironmentSessionEntityTypeOutputWithContext(context.Background())
}

func (i *AgentEnvironmentSessionEntityType) ToAgentEnvironmentSessionEntityTypeOutputWithContext(ctx context.Context) AgentEnvironmentSessionEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentEnvironmentSessionEntityTypeOutput)
}

type AgentEnvironmentSessionEntityTypeOutput struct {
	*pulumi.OutputState
}

func (AgentEnvironmentSessionEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEnvironmentSessionEntityType)(nil))
}

func (o AgentEnvironmentSessionEntityTypeOutput) ToAgentEnvironmentSessionEntityTypeOutput() AgentEnvironmentSessionEntityTypeOutput {
	return o
}

func (o AgentEnvironmentSessionEntityTypeOutput) ToAgentEnvironmentSessionEntityTypeOutputWithContext(ctx context.Context) AgentEnvironmentSessionEntityTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentEnvironmentSessionEntityTypeOutput{})
}
