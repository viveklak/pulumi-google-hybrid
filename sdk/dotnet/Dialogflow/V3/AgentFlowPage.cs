// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    /// <summary>
    /// Creates a page in the specified flow.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3:AgentFlowPage")]
    public partial class AgentFlowPage : Pulumi.CustomResource
    {
        /// <summary>
        /// Required. The human-readable name of the page, unique within the agent.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The fulfillment to call when the session is entering the page.
        /// </summary>
        [Output("entryFulfillment")]
        public Output<Outputs.GoogleCloudDialogflowCxV3FulfillmentResponse> EntryFulfillment { get; private set; } = null!;

        /// <summary>
        /// Handlers associated with the page to handle events such as webhook errors, no match or no input.
        /// </summary>
        [Output("eventHandlers")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3EventHandlerResponse>> EventHandlers { get; private set; } = null!;

        /// <summary>
        /// The form associated with the page, used for collecting parameters relevant to the page.
        /// </summary>
        [Output("form")]
        public Output<Outputs.GoogleCloudDialogflowCxV3FormResponse> Form { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -&gt; page's transition route group -&gt; flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        [Output("transitionRouteGroups")]
        public Output<ImmutableArray<string>> TransitionRouteGroups { get; private set; } = null!;

        /// <summary>
        /// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
        /// </summary>
        [Output("transitionRoutes")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3TransitionRouteResponse>> TransitionRoutes { get; private set; } = null!;


        /// <summary>
        /// Create a AgentFlowPage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AgentFlowPage(string name, AgentFlowPageArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:AgentFlowPage", name, args ?? new AgentFlowPageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AgentFlowPage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:AgentFlowPage", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AgentFlowPage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AgentFlowPage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AgentFlowPage(name, id, options);
        }
    }

    public sealed class AgentFlowPageArgs : Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        /// <summary>
        /// Required. The human-readable name of the page, unique within the agent.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The fulfillment to call when the session is entering the page.
        /// </summary>
        [Input("entryFulfillment")]
        public Input<Inputs.GoogleCloudDialogflowCxV3FulfillmentArgs>? EntryFulfillment { get; set; }

        [Input("eventHandlers")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3EventHandlerArgs>? _eventHandlers;

        /// <summary>
        /// Handlers associated with the page to handle events such as webhook errors, no match or no input.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3EventHandlerArgs> EventHandlers
        {
            get => _eventHandlers ?? (_eventHandlers = new InputList<Inputs.GoogleCloudDialogflowCxV3EventHandlerArgs>());
            set => _eventHandlers = value;
        }

        [Input("flowId", required: true)]
        public Input<string> FlowId { get; set; } = null!;

        /// <summary>
        /// The form associated with the page, used for collecting parameters relevant to the page.
        /// </summary>
        [Input("form")]
        public Input<Inputs.GoogleCloudDialogflowCxV3FormArgs>? Form { get; set; }

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("transitionRouteGroups")]
        private InputList<string>? _transitionRouteGroups;

        /// <summary>
        /// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -&gt; page's transition route group -&gt; flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        public InputList<string> TransitionRouteGroups
        {
            get => _transitionRouteGroups ?? (_transitionRouteGroups = new InputList<string>());
            set => _transitionRouteGroups = value;
        }

        [Input("transitionRoutes")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3TransitionRouteArgs>? _transitionRoutes;

        /// <summary>
        /// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3TransitionRouteArgs> TransitionRoutes
        {
            get => _transitionRoutes ?? (_transitionRoutes = new InputList<Inputs.GoogleCloudDialogflowCxV3TransitionRouteArgs>());
            set => _transitionRoutes = value;
        }

        public AgentFlowPageArgs()
        {
        }
    }
}
