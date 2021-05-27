// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a flow in the specified agent.
type AgentFlow struct {
	pulumi.CustomResourceState

	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringOutput `pulumi:"description"`
	// Required. The human-readable name of the flow.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers GoogleCloudDialogflowCxV3EventHandlerResponseArrayOutput `pulumi:"eventHandlers"`
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// NLU related settings of the flow.
	NluSettings GoogleCloudDialogflowCxV3NluSettingsResponseOutput `pulumi:"nluSettings"`
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayOutput `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteResponseArrayOutput `pulumi:"transitionRoutes"`
}

// NewAgentFlow registers a new resource with the given unique name, arguments, and options.
func NewAgentFlow(ctx *pulumi.Context,
	name string, args *AgentFlowArgs, opts ...pulumi.ResourceOption) (*AgentFlow, error) {
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
	var resource AgentFlow
	err := ctx.RegisterResource("google-native:dialogflow/v3:AgentFlow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentFlow gets an existing AgentFlow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentFlowState, opts ...pulumi.ResourceOption) (*AgentFlow, error) {
	var resource AgentFlow
	err := ctx.ReadResource("google-native:dialogflow/v3:AgentFlow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentFlow resources.
type agentFlowState struct {
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// Required. The human-readable name of the flow.
	DisplayName *string `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers []GoogleCloudDialogflowCxV3EventHandlerResponse `pulumi:"eventHandlers"`
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name *string `pulumi:"name"`
	// NLU related settings of the flow.
	NluSettings *GoogleCloudDialogflowCxV3NluSettingsResponse `pulumi:"nluSettings"`
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes []GoogleCloudDialogflowCxV3TransitionRouteResponse `pulumi:"transitionRoutes"`
}

type AgentFlowState struct {
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// Required. The human-readable name of the flow.
	DisplayName pulumi.StringPtrInput
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers GoogleCloudDialogflowCxV3EventHandlerResponseArrayInput
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name pulumi.StringPtrInput
	// NLU related settings of the flow.
	NluSettings GoogleCloudDialogflowCxV3NluSettingsResponsePtrInput
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayInput
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteResponseArrayInput
}

func (AgentFlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentFlowState)(nil)).Elem()
}

type agentFlowArgs struct {
	AgentId string `pulumi:"agentId"`
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// Required. The human-readable name of the flow.
	DisplayName *string `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers []GoogleCloudDialogflowCxV3EventHandler `pulumi:"eventHandlers"`
	LanguageCode  *string                                 `pulumi:"languageCode"`
	Location      string                                  `pulumi:"location"`
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name *string `pulumi:"name"`
	// NLU related settings of the flow.
	NluSettings *GoogleCloudDialogflowCxV3NluSettings `pulumi:"nluSettings"`
	Project     string                                `pulumi:"project"`
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes []GoogleCloudDialogflowCxV3TransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a AgentFlow resource.
type AgentFlowArgs struct {
	AgentId pulumi.StringInput
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// Required. The human-readable name of the flow.
	DisplayName pulumi.StringPtrInput
	// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
	EventHandlers GoogleCloudDialogflowCxV3EventHandlerArrayInput
	LanguageCode  pulumi.StringPtrInput
	Location      pulumi.StringInput
	// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
	Name pulumi.StringPtrInput
	// NLU related settings of the flow.
	NluSettings GoogleCloudDialogflowCxV3NluSettingsPtrInput
	Project     pulumi.StringInput
	// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayInput
	// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteArrayInput
}

func (AgentFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentFlowArgs)(nil)).Elem()
}

type AgentFlowInput interface {
	pulumi.Input

	ToAgentFlowOutput() AgentFlowOutput
	ToAgentFlowOutputWithContext(ctx context.Context) AgentFlowOutput
}

func (*AgentFlow) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentFlow)(nil))
}

func (i *AgentFlow) ToAgentFlowOutput() AgentFlowOutput {
	return i.ToAgentFlowOutputWithContext(context.Background())
}

func (i *AgentFlow) ToAgentFlowOutputWithContext(ctx context.Context) AgentFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentFlowOutput)
}

type AgentFlowOutput struct {
	*pulumi.OutputState
}

func (AgentFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentFlow)(nil))
}

func (o AgentFlowOutput) ToAgentFlowOutput() AgentFlowOutput {
	return o
}

func (o AgentFlowOutput) ToAgentFlowOutputWithContext(ctx context.Context) AgentFlowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentFlowOutput{})
}
