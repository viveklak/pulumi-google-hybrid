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
type Environment struct {
	pulumi.CustomResourceState

	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Update time of this environment.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponseArrayOutput `pulumi:"versionConfigs"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.VersionConfigs == nil {
		return nil, errors.New("invalid value for required argument 'VersionConfigs'")
	}
	var resource Environment
	err := ctx.RegisterResource("google-native:dialogflow/v3:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("google-native:dialogflow/v3:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name *string `pulumi:"name"`
	// Update time of this environment.
	UpdateTime *string `pulumi:"updateTime"`
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs []GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponse `pulumi:"versionConfigs"`
}

type EnvironmentState struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringPtrInput
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name pulumi.StringPtrInput
	// Update time of this environment.
	UpdateTime pulumi.StringPtrInput
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs GoogleCloudDialogflowCxV3EnvironmentVersionConfigResponseArrayInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	AgentId string `pulumi:"agentId"`
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName string `pulumi:"displayName"`
	Location    string `pulumi:"location"`
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs []GoogleCloudDialogflowCxV3EnvironmentVersionConfig `pulumi:"versionConfigs"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	AgentId pulumi.StringInput
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringInput
	Location    pulumi.StringInput
	// The name of the environment. Format: `projects//locations//agents//environments/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from `Start Flow` in the agent. Otherwise, an error will be returned.
	VersionConfigs GoogleCloudDialogflowCxV3EnvironmentVersionConfigArrayInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct {
	*pulumi.OutputState
}

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Environment)(nil))
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EnvironmentOutput{})
}
