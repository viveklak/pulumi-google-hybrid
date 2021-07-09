// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1PageResponse
    {
        /// <summary>
        /// The human-readable name of the page, unique within the agent.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The fulfillment to call when the session is entering the page.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3beta1FulfillmentResponse EntryFulfillment;
        /// <summary>
        /// Handlers associated with the page to handle events such as webhook errors, no match or no input.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EventHandlerResponse> EventHandlers;
        /// <summary>
        /// The form associated with the page, used for collecting parameters relevant to the page.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3beta1FormResponse Form;
        /// <summary>
        /// The unique identifier of the page. Required for the Pages.UpdatePage method. Pages.CreatePage populates the name automatically. Format: `projects//locations//agents//flows//pages/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Ordered list of `TransitionRouteGroups` associated with the page. Transition route groups must be unique within a page. * If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -&gt; page's transition route group -&gt; flow's transition routes. * If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence. Format:`projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        public readonly ImmutableArray<string> TransitionRouteGroups;
        /// <summary>
        /// A list of transitions for the transition rules of this page. They route the conversation to another page in the same flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order: * TransitionRoutes defined in the page with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in flow with intent specified. * TransitionRoutes defined in the transition route groups with intent specified. * TransitionRoutes defined in the page with only condition specified. * TransitionRoutes defined in the transition route groups with only condition specified.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse> TransitionRoutes;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1PageResponse(
            string displayName,

            Outputs.GoogleCloudDialogflowCxV3beta1FulfillmentResponse entryFulfillment,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EventHandlerResponse> eventHandlers,

            Outputs.GoogleCloudDialogflowCxV3beta1FormResponse form,

            string name,

            ImmutableArray<string> transitionRouteGroups,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse> transitionRoutes)
        {
            DisplayName = displayName;
            EntryFulfillment = entryFulfillment;
            EventHandlers = eventHandlers;
            Form = form;
            Name = name;
            TransitionRouteGroups = transitionRouteGroups;
            TransitionRoutes = transitionRoutes;
        }
    }
}
