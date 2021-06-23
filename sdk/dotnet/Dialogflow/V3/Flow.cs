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
    /// Creates a flow in the specified agent. Note: You should always train a flow prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/cx/docs/concept/training).
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3:Flow")]
    public partial class Flow : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Required. The human-readable name of the flow.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
        /// </summary>
        [Output("eventHandlers")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3EventHandlerResponse>> EventHandlers { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// NLU related settings of the flow.
        /// </summary>
        [Output("nluSettings")]
        public Output<Outputs.GoogleCloudDialogflowCxV3NluSettingsResponse> NluSettings { get; private set; } = null!;

        /// <summary>
        /// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        [Output("transitionRouteGroups")]
        public Output<ImmutableArray<string>> TransitionRouteGroups { get; private set; } = null!;

        /// <summary>
        /// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
        /// </summary>
        [Output("transitionRoutes")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3TransitionRouteResponse>> TransitionRoutes { get; private set; } = null!;


        /// <summary>
        /// Create a Flow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Flow(string name, FlowArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:Flow", name, args ?? new FlowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Flow(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:Flow", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Flow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Flow Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Flow(name, id, options);
        }
    }

    public sealed class FlowArgs : Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        /// <summary>
        /// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. The human-readable name of the flow.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("eventHandlers")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3EventHandlerArgs>? _eventHandlers;

        /// <summary>
        /// A flow's event handlers serve two purposes: * They are responsible for handling events (e.g. no match, webhook errors) in the flow. * They are inherited by every page's event handlers, which can be used to handle common events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the flow. Unlike transition_routes, these handlers are evaluated on a first-match basis. The first one that matches the event get executed, with the rest being ignored.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3EventHandlerArgs> EventHandlers
        {
            get => _eventHandlers ?? (_eventHandlers = new InputList<Inputs.GoogleCloudDialogflowCxV3EventHandlerArgs>());
            set => _eventHandlers = value;
        }

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The unique identifier of the flow. Format: `projects//locations//agents//flows/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NLU related settings of the flow.
        /// </summary>
        [Input("nluSettings")]
        public Input<Inputs.GoogleCloudDialogflowCxV3NluSettingsArgs>? NluSettings { get; set; }

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        [Input("transitionRouteGroups")]
        private InputList<string>? _transitionRouteGroups;

        /// <summary>
        /// A flow's transition route group serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition route groups. Transition route groups defined in the page have higher priority than those defined in the flow. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        public InputList<string> TransitionRouteGroups
        {
            get => _transitionRouteGroups ?? (_transitionRouteGroups = new InputList<string>());
            set => _transitionRouteGroups = value;
        }

        [Input("transitionRoutes")]
        private InputList<Inputs.GoogleCloudDialogflowCxV3TransitionRouteArgs>? _transitionRoutes;

        /// <summary>
        /// A flow's transition routes serve two purposes: * They are responsible for matching the user's first utterances in the flow. * They are inherited by every page's transition routes and can support use cases such as the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page. Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are evalauted in the following order: * TransitionRoutes with intent specified.. * TransitionRoutes with only condition specified. TransitionRoutes with intent specified are inherited by pages in the flow.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3TransitionRouteArgs> TransitionRoutes
        {
            get => _transitionRoutes ?? (_transitionRoutes = new InputList<Inputs.GoogleCloudDialogflowCxV3TransitionRouteArgs>());
            set => _transitionRoutes = value;
        }

        public FlowArgs()
        {
        }
    }
}
