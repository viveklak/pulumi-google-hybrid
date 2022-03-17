// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an agent in the specified location. Note: You should always train flows prior to sending them queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
type Agent struct {
	pulumi.CustomResourceState

	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	AdvancedSettings GoogleCloudDialogflowCxV3AdvancedSettingsResponseOutput `pulumi:"advancedSettings"`
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted [Web Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo) integration.
	AvatarUri pulumi.StringOutput `pulumi:"avatarUri"`
	// Immutable. The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be set by the Agents.UpdateAgent method.
	DefaultLanguageCode pulumi.StringOutput `pulumi:"defaultLanguageCode"`
	// The description of the agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human-readable name of the agent, unique within the location.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolOutput `pulumi:"enableSpellCorrection"`
	// Indicates if stackdriver logging is enabled for the agent. Please use agent.advanced_settings instead.
	EnableStackdriverLogging pulumi.BoolOutput `pulumi:"enableStackdriverLogging"`
	// The unique identifier of the agent. Required for the Agents.UpdateAgent method. Agents.CreateAgent populates the name automatically. Format: `projects//locations//agents/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings pulumi.StringOutput `pulumi:"securitySettings"`
	// Speech recognition related settings.
	SpeechToTextSettings GoogleCloudDialogflowCxV3SpeechToTextSettingsResponseOutput `pulumi:"speechToTextSettings"`
	// Immutable. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: `projects//locations//agents//flows/`.
	StartFlow pulumi.StringOutput `pulumi:"startFlow"`
	// The list of all languages supported by the agent (except for the `default_language_code`).
	SupportedLanguageCodes pulumi.StringArrayOutput `pulumi:"supportedLanguageCodes"`
	// The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewAgent registers a new resource with the given unique name, arguments, and options.
func NewAgent(ctx *pulumi.Context,
	name string, args *AgentArgs, opts ...pulumi.ResourceOption) (*Agent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultLanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'DefaultLanguageCode'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	var resource Agent
	err := ctx.RegisterResource("google-native:dialogflow/v3:Agent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgent gets an existing Agent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentState, opts ...pulumi.ResourceOption) (*Agent, error) {
	var resource Agent
	err := ctx.ReadResource("google-native:dialogflow/v3:Agent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agent resources.
type agentState struct {
}

type AgentState struct {
}

func (AgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentState)(nil)).Elem()
}

type agentArgs struct {
	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	AdvancedSettings *GoogleCloudDialogflowCxV3AdvancedSettings `pulumi:"advancedSettings"`
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted [Web Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo) integration.
	AvatarUri *string `pulumi:"avatarUri"`
	// Immutable. The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be set by the Agents.UpdateAgent method.
	DefaultLanguageCode string `pulumi:"defaultLanguageCode"`
	// The description of the agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The human-readable name of the agent, unique within the location.
	DisplayName string `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `pulumi:"enableSpellCorrection"`
	// Indicates if stackdriver logging is enabled for the agent. Please use agent.advanced_settings instead.
	EnableStackdriverLogging *bool   `pulumi:"enableStackdriverLogging"`
	Location                 *string `pulumi:"location"`
	// The unique identifier of the agent. Required for the Agents.UpdateAgent method. Agents.CreateAgent populates the name automatically. Format: `projects//locations//agents/`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Name of the SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings *string `pulumi:"securitySettings"`
	// Speech recognition related settings.
	SpeechToTextSettings *GoogleCloudDialogflowCxV3SpeechToTextSettings `pulumi:"speechToTextSettings"`
	// Immutable. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: `projects//locations//agents//flows/`.
	StartFlow *string `pulumi:"startFlow"`
	// The list of all languages supported by the agent (except for the `default_language_code`).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	// The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	TimeZone string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Agent resource.
type AgentArgs struct {
	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	AdvancedSettings GoogleCloudDialogflowCxV3AdvancedSettingsPtrInput
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted [Web Demo](https://cloud.google.com/dialogflow/docs/integrations/web-demo) integration.
	AvatarUri pulumi.StringPtrInput
	// Immutable. The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes. This field cannot be set by the Agents.UpdateAgent method.
	DefaultLanguageCode pulumi.StringInput
	// The description of the agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The human-readable name of the agent, unique within the location.
	DisplayName pulumi.StringInput
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolPtrInput
	// Indicates if stackdriver logging is enabled for the agent. Please use agent.advanced_settings instead.
	EnableStackdriverLogging pulumi.BoolPtrInput
	Location                 pulumi.StringPtrInput
	// The unique identifier of the agent. Required for the Agents.UpdateAgent method. Agents.CreateAgent populates the name automatically. Format: `projects//locations//agents/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Name of the SecuritySettings reference for the agent. Format: `projects//locations//securitySettings/`.
	SecuritySettings pulumi.StringPtrInput
	// Speech recognition related settings.
	SpeechToTextSettings GoogleCloudDialogflowCxV3SpeechToTextSettingsPtrInput
	// Immutable. Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: `projects//locations//agents//flows/`.
	StartFlow pulumi.StringPtrInput
	// The list of all languages supported by the agent (except for the `default_language_code`).
	SupportedLanguageCodes pulumi.StringArrayInput
	// The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	TimeZone pulumi.StringInput
}

func (AgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentArgs)(nil)).Elem()
}

type AgentInput interface {
	pulumi.Input

	ToAgentOutput() AgentOutput
	ToAgentOutputWithContext(ctx context.Context) AgentOutput
}

func (*Agent) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (i *Agent) ToAgentOutput() AgentOutput {
	return i.ToAgentOutputWithContext(context.Background())
}

func (i *Agent) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentOutput)
}

type AgentOutput struct{ *pulumi.OutputState }

func (AgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (o AgentOutput) ToAgentOutput() AgentOutput {
	return o
}

func (o AgentOutput) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentInput)(nil)).Elem(), &Agent{})
	pulumi.RegisterOutputType(AgentOutput{})
}
