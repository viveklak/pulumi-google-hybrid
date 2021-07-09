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
    public sealed class GoogleCloudDialogflowCxV3beta1FormParameterFillBehaviorResponse
    {
        /// <summary>
        /// The fulfillment to provide the initial prompt that the agent can present to the user in order to fill the parameter.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3beta1FulfillmentResponse InitialPromptFulfillment;
        /// <summary>
        /// The handlers for parameter-level events, used to provide reprompt for the parameter or transition to a different page/flow. The supported events are: * `sys.no-match-`, where N can be from 1 to 6 * `sys.no-match-default` * `sys.no-input-`, where N can be from 1 to 6 * `sys.no-input-default` * `sys.invalid-parameter` `initial_prompt_fulfillment` provides the first prompt for the parameter. If the user's response does not fill the parameter, a no-match/no-input event will be triggered, and the fulfillment associated with the `sys.no-match-1`/`sys.no-input-1` handler (if defined) will be called to provide a prompt. The `sys.no-match-2`/`sys.no-input-2` handler (if defined) will respond to the next no-match/no-input event, and so on. A `sys.no-match-default` or `sys.no-input-default` handler will be used to handle all following no-match/no-input events after all numbered no-match/no-input handlers for the parameter are consumed. A `sys.invalid-parameter` handler can be defined to handle the case where the parameter values have been `invalidated` by webhook. For example, if the user's response fill the parameter, however the parameter was invalidated by webhook, the fulfillment associated with the `sys.invalid-parameter` handler (if defined) will be called to provide a prompt. If the event handler for the corresponding event can't be found on the parameter, `initial_prompt_fulfillment` will be re-prompted.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EventHandlerResponse> RepromptEventHandlers;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1FormParameterFillBehaviorResponse(
            Outputs.GoogleCloudDialogflowCxV3beta1FulfillmentResponse initialPromptFulfillment,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3beta1EventHandlerResponse> repromptEventHandlers)
        {
            InitialPromptFulfillment = initialPromptFulfillment;
            RepromptEventHandlers = repromptEventHandlers;
        }
    }
}
