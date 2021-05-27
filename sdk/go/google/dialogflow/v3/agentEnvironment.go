// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Environment in the specified Agent.
type AgentEnvironment struct {
	pulumi.CustomResourceState

	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringOutput `pulumi:"description"`
	// Required. The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Update time of this environment.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Required. A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponseArrayOutput `pulumi:"versionConfigs"`
}

// NewAgentEnvironment registers a new resource with the given unique name, arguments, and options.
func NewAgentEnvironment(ctx *pulumi.Context,
	name string, args *AgentEnvironmentArgs, opts ...pulumi.ResourceOption) (*AgentEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource AgentEnvironment
	err := ctx.RegisterResource("google-native:dialogflow/v3:AgentEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentEnvironment gets an existing AgentEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentEnvironmentState, opts ...pulumi.ResourceOption) (*AgentEnvironment, error) {
	var resource AgentEnvironment
	err := ctx.ReadResource("google-native:dialogflow/v3:AgentEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentEnvironment resources.
type agentEnvironmentState struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// Required. The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name *string `pulumi:"name"`
	// Update time of this environment.
	UpdateTime *string `pulumi:"updateTime"`
	// Required. A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs []GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponse `pulumi:"versionConfigs"`
}

type AgentEnvironmentState struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// Required. The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringPtrInput
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name pulumi.StringPtrInput
	// Update time of this environment.
	UpdateTime pulumi.StringPtrInput
	// Required. A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponseArrayInput
}

func (AgentEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEnvironmentState)(nil)).Elem()
}

type agentEnvironmentArgs struct {
	AgentId string `pulumi:"agentId"`
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// Required. The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName *string `pulumi:"displayName"`
	Location    string  `pulumi:"location"`
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Required. A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs []GoogleCloudDialogflowCxV3EnvironmentVersionConfig `pulumi:"versionConfigs"`
}

// The set of arguments for constructing a AgentEnvironment resource.
type AgentEnvironmentArgs struct {
	AgentId pulumi.StringInput
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// Required. The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringPtrInput
	Location    pulumi.StringInput
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Required. A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs GoogleCloudDialogflowCxV3EnvironmentVersionConfigArrayInput
}

func (AgentEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentEnvironmentArgs)(nil)).Elem()
}

type AgentEnvironmentInput interface {
	pulumi.Input

	ToAgentEnvironmentOutput() AgentEnvironmentOutput
	ToAgentEnvironmentOutputWithContext(ctx context.Context) AgentEnvironmentOutput
}

func (*AgentEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEnvironment)(nil))
}

func (i *AgentEnvironment) ToAgentEnvironmentOutput() AgentEnvironmentOutput {
	return i.ToAgentEnvironmentOutputWithContext(context.Background())
}

func (i *AgentEnvironment) ToAgentEnvironmentOutputWithContext(ctx context.Context) AgentEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentEnvironmentOutput)
}

type AgentEnvironmentOutput struct {
	*pulumi.OutputState
}

func (AgentEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentEnvironment)(nil))
}

func (o AgentEnvironmentOutput) ToAgentEnvironmentOutput() AgentEnvironmentOutput {
	return o
}

func (o AgentEnvironmentOutput) ToAgentEnvironmentOutputWithContext(ctx context.Context) AgentEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentEnvironmentOutput{})
}
