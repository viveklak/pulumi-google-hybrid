// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a session entity type. If the specified session entity type already exists, overrides the session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.
type AgentSessionEntityType struct {
	pulumi.CustomResourceState

	// Required. The collection of entities associated with this session entity type.
	Entities GoogleCloudDialogflowV2beta1EntityTypeEntityResponseArrayOutput `pulumi:"entities"`
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode pulumi.StringOutput `pulumi:"entityOverrideMode"`
	// Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAgentSessionEntityType registers a new resource with the given unique name, arguments, and options.
func NewAgentSessionEntityType(ctx *pulumi.Context,
	name string, args *AgentSessionEntityTypeArgs, opts ...pulumi.ResourceOption) (*AgentSessionEntityType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource AgentSessionEntityType
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:AgentSessionEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentSessionEntityType gets an existing AgentSessionEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentSessionEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentSessionEntityTypeState, opts ...pulumi.ResourceOption) (*AgentSessionEntityType, error) {
	var resource AgentSessionEntityType
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:AgentSessionEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentSessionEntityType resources.
type agentSessionEntityTypeState struct {
	// Required. The collection of entities associated with this session entity type.
	Entities []GoogleCloudDialogflowV2beta1EntityTypeEntityResponse `pulumi:"entities"`
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode *string `pulumi:"entityOverrideMode"`
	// Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
	Name *string `pulumi:"name"`
}

type AgentSessionEntityTypeState struct {
	// Required. The collection of entities associated with this session entity type.
	Entities GoogleCloudDialogflowV2beta1EntityTypeEntityResponseArrayInput
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode pulumi.StringPtrInput
	// Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
	Name pulumi.StringPtrInput
}

func (AgentSessionEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentSessionEntityTypeState)(nil)).Elem()
}

type agentSessionEntityTypeArgs struct {
	// Required. The collection of entities associated with this session entity type.
	Entities []GoogleCloudDialogflowV2beta1EntityTypeEntity `pulumi:"entities"`
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode *string `pulumi:"entityOverrideMode"`
	Location           string  `pulumi:"location"`
	// Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
	Name      *string `pulumi:"name"`
	Project   string  `pulumi:"project"`
	SessionId string  `pulumi:"sessionId"`
}

// The set of arguments for constructing a AgentSessionEntityType resource.
type AgentSessionEntityTypeArgs struct {
	// Required. The collection of entities associated with this session entity type.
	Entities GoogleCloudDialogflowV2beta1EntityTypeEntityArrayInput
	// Required. Indicates whether the additional data should override or supplement the custom entity type definition.
	EntityOverrideMode pulumi.StringPtrInput
	Location           pulumi.StringInput
	// Required. The unique identifier of this session entity type. Supported formats: - `projects//agent/sessions//entityTypes/` - `projects//locations//agent/sessions//entityTypes/` - `projects//agent/environments//users//sessions//entityTypes/` - `projects//locations//agent/environments/ /users//sessions//entityTypes/` If `Location ID` is not specified we assume default 'us' location. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringInput
	SessionId pulumi.StringInput
}

func (AgentSessionEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentSessionEntityTypeArgs)(nil)).Elem()
}

type AgentSessionEntityTypeInput interface {
	pulumi.Input

	ToAgentSessionEntityTypeOutput() AgentSessionEntityTypeOutput
	ToAgentSessionEntityTypeOutputWithContext(ctx context.Context) AgentSessionEntityTypeOutput
}

func (*AgentSessionEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentSessionEntityType)(nil))
}

func (i *AgentSessionEntityType) ToAgentSessionEntityTypeOutput() AgentSessionEntityTypeOutput {
	return i.ToAgentSessionEntityTypeOutputWithContext(context.Background())
}

func (i *AgentSessionEntityType) ToAgentSessionEntityTypeOutputWithContext(ctx context.Context) AgentSessionEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentSessionEntityTypeOutput)
}

type AgentSessionEntityTypeOutput struct {
	*pulumi.OutputState
}

func (AgentSessionEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentSessionEntityType)(nil))
}

func (o AgentSessionEntityTypeOutput) ToAgentSessionEntityTypeOutput() AgentSessionEntityTypeOutput {
	return o
}

func (o AgentSessionEntityTypeOutput) ToAgentSessionEntityTypeOutputWithContext(ctx context.Context) AgentSessionEntityTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentSessionEntityTypeOutput{})
}
