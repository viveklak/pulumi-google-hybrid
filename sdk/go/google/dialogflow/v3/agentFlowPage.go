// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a page in the specified flow.
type AgentFlowPage struct {
	pulumi.CustomResourceState

	// Required. The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment GoogleCloudDialogflowCxV3FulfillmentResponseOutput `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers GoogleCloudDialogflowCxV3EventHandlerResponseArrayOutput `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form GoogleCloudDialogflowCxV3FormResponseOutput `pulumi:"form"`
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayOutput `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteResponseArrayOutput `pulumi:"transitionRoutes"`
}

// NewAgentFlowPage registers a new resource with the given unique name, arguments, and options.
func NewAgentFlowPage(ctx *pulumi.Context,
	name string, args *AgentFlowPageArgs, opts ...pulumi.ResourceOption) (*AgentFlowPage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.FlowId == nil {
		return nil, errors.New("invalid value for required argument 'FlowId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource AgentFlowPage
	err := ctx.RegisterResource("google-native:dialogflow/v3:AgentFlowPage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentFlowPage gets an existing AgentFlowPage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentFlowPage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentFlowPageState, opts ...pulumi.ResourceOption) (*AgentFlowPage, error) {
	var resource AgentFlowPage
	err := ctx.ReadResource("google-native:dialogflow/v3:AgentFlowPage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentFlowPage resources.
type agentFlowPageState struct {
	// Required. The human-readable name of the page, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment *GoogleCloudDialogflowCxV3FulfillmentResponse `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers []GoogleCloudDialogflowCxV3EventHandlerResponse `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form *GoogleCloudDialogflowCxV3FormResponse `pulumi:"form"`
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name *string `pulumi:"name"`
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes []GoogleCloudDialogflowCxV3TransitionRouteResponse `pulumi:"transitionRoutes"`
}

type AgentFlowPageState struct {
	// Required. The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment GoogleCloudDialogflowCxV3FulfillmentResponsePtrInput
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers GoogleCloudDialogflowCxV3EventHandlerResponseArrayInput
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form GoogleCloudDialogflowCxV3FormResponsePtrInput
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name pulumi.StringPtrInput
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayInput
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteResponseArrayInput
}

func (AgentFlowPageState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentFlowPageState)(nil)).Elem()
}

type agentFlowPageArgs struct {
	AgentId string `pulumi:"agentId"`
	// Required. The human-readable name of the page, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment *GoogleCloudDialogflowCxV3Fulfillment `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers []GoogleCloudDialogflowCxV3EventHandler `pulumi:"eventHandlers"`
	FlowId        string                                  `pulumi:"flowId"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form         *GoogleCloudDialogflowCxV3Form `pulumi:"form"`
	LanguageCode *string                        `pulumi:"languageCode"`
	Location     string                         `pulumi:"location"`
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes []GoogleCloudDialogflowCxV3TransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a AgentFlowPage resource.
type AgentFlowPageArgs struct {
	AgentId pulumi.StringInput
	// Required. The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment GoogleCloudDialogflowCxV3FulfillmentPtrInput
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers GoogleCloudDialogflowCxV3EventHandlerArrayInput
	FlowId        pulumi.StringInput
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form         GoogleCloudDialogflowCxV3FormPtrInput
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringInput
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayInput
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes GoogleCloudDialogflowCxV3TransitionRouteArrayInput
}

func (AgentFlowPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentFlowPageArgs)(nil)).Elem()
}

type AgentFlowPageInput interface {
	pulumi.Input

	ToAgentFlowPageOutput() AgentFlowPageOutput
	ToAgentFlowPageOutputWithContext(ctx context.Context) AgentFlowPageOutput
}

func (*AgentFlowPage) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentFlowPage)(nil))
}

func (i *AgentFlowPage) ToAgentFlowPageOutput() AgentFlowPageOutput {
	return i.ToAgentFlowPageOutputWithContext(context.Background())
}

func (i *AgentFlowPage) ToAgentFlowPageOutputWithContext(ctx context.Context) AgentFlowPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentFlowPageOutput)
}

type AgentFlowPageOutput struct {
	*pulumi.OutputState
}

func (AgentFlowPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentFlowPage)(nil))
}

func (o AgentFlowPageOutput) ToAgentFlowPageOutput() AgentFlowPageOutput {
	return o
}

func (o AgentFlowPageOutput) ToAgentFlowPageOutputWithContext(ctx context.Context) AgentFlowPageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AgentFlowPageOutput{})
}
