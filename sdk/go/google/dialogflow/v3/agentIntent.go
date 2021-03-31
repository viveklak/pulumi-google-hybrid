// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an intent in the specified agent.
type AgentIntent struct {
	pulumi.CustomResourceState
}

// NewAgentIntent registers a new resource with the given unique name, arguments, and options.
func NewAgentIntent(ctx *pulumi.Context,
	name string, args *AgentIntentArgs, opts ...pulumi.ResourceOption) (*AgentIntent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentsId == nil {
		return nil, errors.New("invalid value for required argument 'AgentsId'")
	}
	if args.IntentsId == nil {
		return nil, errors.New("invalid value for required argument 'IntentsId'")
	}
	if args.LocationsId == nil {
		return nil, errors.New("invalid value for required argument 'LocationsId'")
	}
	if args.ProjectsId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectsId'")
	}
	var resource AgentIntent
	err := ctx.RegisterResource("google-cloud:dialogflow/v3:AgentIntent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentIntent gets an existing AgentIntent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentIntentState, opts ...pulumi.ResourceOption) (*AgentIntent, error) {
	var resource AgentIntent
	err := ctx.ReadResource("google-cloud:dialogflow/v3:AgentIntent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentIntent resources.
type agentIntentState struct {
}

type AgentIntentState struct {
}

func (AgentIntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentIntentState)(nil)).Elem()
}

type agentIntentArgs struct {
	AgentsId string `pulumi:"agentsId"`
	// Optional. Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description *string `pulumi:"description"`
	// Required. The human-readable name of the intent, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	IntentsId   string  `pulumi:"intentsId"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback *bool `pulumi:"isFallback"`
	// Optional. The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	Labels      map[string]string `pulumi:"labels"`
	LocationsId string            `pulumi:"locationsId"`
	// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
	Name *string `pulumi:"name"`
	// The collection of parameters associated with the intent.
	Parameters []GoogleCloudDialogflowCxV3IntentParameter `pulumi:"parameters"`
	// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority   *int   `pulumi:"priority"`
	ProjectsId string `pulumi:"projectsId"`
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases []GoogleCloudDialogflowCxV3IntentTrainingPhrase `pulumi:"trainingPhrases"`
}

// The set of arguments for constructing a AgentIntent resource.
type AgentIntentArgs struct {
	AgentsId pulumi.StringInput
	// Optional. Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description pulumi.StringPtrInput
	// Required. The human-readable name of the intent, unique within the agent.
	DisplayName pulumi.StringPtrInput
	IntentsId   pulumi.StringInput
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback pulumi.BoolPtrInput
	// Optional. The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	Labels      pulumi.StringMapInput
	LocationsId pulumi.StringInput
	// The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
	Name pulumi.StringPtrInput
	// The collection of parameters associated with the intent.
	Parameters GoogleCloudDialogflowCxV3IntentParameterArrayInput
	// The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority   pulumi.IntPtrInput
	ProjectsId pulumi.StringInput
	// The collection of training phrases the agent is trained on to identify the intent.
	TrainingPhrases GoogleCloudDialogflowCxV3IntentTrainingPhraseArrayInput
}

func (AgentIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentIntentArgs)(nil)).Elem()
}

type AgentIntentInput interface {
	pulumi.Input

	ToAgentIntentOutput() AgentIntentOutput
	ToAgentIntentOutputWithContext(ctx context.Context) AgentIntentOutput
}

func (*AgentIntent) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentIntent)(nil))
}

func (i *AgentIntent) ToAgentIntentOutput() AgentIntentOutput {
	return i.ToAgentIntentOutputWithContext(context.Background())
}

func (i *AgentIntent) ToAgentIntentOutputWithContext(ctx context.Context) AgentIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentIntentOutput)
}

type AgentIntentOutput struct {
	*pulumi.OutputState
}

func (AgentIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentIntent)(nil))
}

func (o AgentIntentOutput) ToAgentIntentOutput() AgentIntentOutput {
	return o
}

func (o AgentIntentOutput) ToAgentIntentOutputWithContext(ctx context.Context) AgentIntentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentIntentOutput{})
}
