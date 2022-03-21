// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a page in the specified flow.
type Page struct {
	pulumi.CustomResourceState

	// The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment GoogleCloudDialogflowCxV3beta1FulfillmentResponseOutput `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers GoogleCloudDialogflowCxV3beta1EventHandlerResponseArrayOutput `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form GoogleCloudDialogflowCxV3beta1FormResponseOutput `pulumi:"form"`
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayOutput `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes GoogleCloudDialogflowCxV3beta1TransitionRouteResponseArrayOutput `pulumi:"transitionRoutes"`
}

// NewPage registers a new resource with the given unique name, arguments, and options.
func NewPage(ctx *pulumi.Context,
	name string, args *PageArgs, opts ...pulumi.ResourceOption) (*Page, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.FlowId == nil {
		return nil, errors.New("invalid value for required argument 'FlowId'")
	}
	var resource Page
	err := ctx.RegisterResource("google-native:dialogflow/v3beta1:Page", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPage gets an existing Page resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PageState, opts ...pulumi.ResourceOption) (*Page, error) {
	var resource Page
	err := ctx.ReadResource("google-native:dialogflow/v3beta1:Page", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Page resources.
type pageState struct {
}

type PageState struct {
}

func (PageState) ElementType() reflect.Type {
	return reflect.TypeOf((*pageState)(nil)).Elem()
}

type pageArgs struct {
	AgentId string `pulumi:"agentId"`
	// The human-readable name of the page, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment *GoogleCloudDialogflowCxV3beta1Fulfillment `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers []GoogleCloudDialogflowCxV3beta1EventHandler `pulumi:"eventHandlers"`
	FlowId        string                                       `pulumi:"flowId"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form *GoogleCloudDialogflowCxV3beta1Form `pulumi:"form"`
	// The language of the following fields in `page`: * `Page.entry_fulfillment.messages` * `Page.entry_fulfillment.conditional_cases` * `Page.event_handlers.trigger_fulfillment.messages` * `Page.event_handlers.trigger_fulfillment.conditional_cases` * `Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages` * `Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases` * `Page.form.parameters.fill_behavior.reprompt_event_handlers.messages` * `Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases` * `Page.transition_routes.trigger_fulfillment.messages` * `Page.transition_routes.trigger_fulfillment.conditional_cases` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	Location     *string `pulumi:"location"`
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes []GoogleCloudDialogflowCxV3beta1TransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a Page resource.
type PageArgs struct {
	AgentId pulumi.StringInput
	// The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringInput
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment GoogleCloudDialogflowCxV3beta1FulfillmentPtrInput
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers GoogleCloudDialogflowCxV3beta1EventHandlerArrayInput
	FlowId        pulumi.StringInput
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form GoogleCloudDialogflowCxV3beta1FormPtrInput
	// The language of the following fields in `page`: * `Page.entry_fulfillment.messages` * `Page.entry_fulfillment.conditional_cases` * `Page.event_handlers.trigger_fulfillment.messages` * `Page.event_handlers.trigger_fulfillment.conditional_cases` * `Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages` * `Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases` * `Page.form.parameters.fill_behavior.reprompt_event_handlers.messages` * `Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases` * `Page.transition_routes.trigger_fulfillment.messages` * `Page.transition_routes.trigger_fulfillment.conditional_cases` If not specified, the agent's default language is used. [Many languages](https://cloud.google.com/dialogflow/cx/docs/reference/language) are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
	TransitionRouteGroups pulumi.StringArrayInput
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
	TransitionRoutes GoogleCloudDialogflowCxV3beta1TransitionRouteArrayInput
}

func (PageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pageArgs)(nil)).Elem()
}

type PageInput interface {
	pulumi.Input

	ToPageOutput() PageOutput
	ToPageOutputWithContext(ctx context.Context) PageOutput
}

func (*Page) ElementType() reflect.Type {
	return reflect.TypeOf((**Page)(nil)).Elem()
}

func (i *Page) ToPageOutput() PageOutput {
	return i.ToPageOutputWithContext(context.Background())
}

func (i *Page) ToPageOutputWithContext(ctx context.Context) PageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PageOutput)
}

type PageOutput struct{ *pulumi.OutputState }

func (PageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Page)(nil)).Elem()
}

func (o PageOutput) ToPageOutput() PageOutput {
	return o
}

func (o PageOutput) ToPageOutputWithContext(ctx context.Context) PageOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PageInput)(nil)).Elem(), &Page{})
	pulumi.RegisterOutputType(PageOutput{})
}
