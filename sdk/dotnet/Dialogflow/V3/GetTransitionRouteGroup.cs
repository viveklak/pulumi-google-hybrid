// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    public static class GetTransitionRouteGroup
    {
        /// <summary>
        /// Retrieves the specified TransitionRouteGroup.
        /// </summary>
        public static Task<GetTransitionRouteGroupResult> InvokeAsync(GetTransitionRouteGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTransitionRouteGroupResult>("google-native:dialogflow/v3:getTransitionRouteGroup", args ?? new GetTransitionRouteGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified TransitionRouteGroup.
        /// </summary>
        public static Output<GetTransitionRouteGroupResult> Invoke(GetTransitionRouteGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTransitionRouteGroupResult>("google-native:dialogflow/v3:getTransitionRouteGroup", args ?? new GetTransitionRouteGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTransitionRouteGroupArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public string AgentId { get; set; } = null!;

        [Input("flowId", required: true)]
        public string FlowId { get; set; } = null!;

        [Input("languageCode")]
        public string? LanguageCode { get; set; }

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("transitionRouteGroupId", required: true)]
        public string TransitionRouteGroupId { get; set; } = null!;

        public GetTransitionRouteGroupArgs()
        {
        }
    }

    public sealed class GetTransitionRouteGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("flowId", required: true)]
        public Input<string> FlowId { get; set; } = null!;

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("transitionRouteGroupId", required: true)]
        public Input<string> TransitionRouteGroupId { get; set; } = null!;

        public GetTransitionRouteGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTransitionRouteGroupResult
    {
        /// <summary>
        /// The human-readable name of the transition route group, unique within the Agent. The display name can be no longer than 30 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Transition routes associated with the TransitionRouteGroup.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3TransitionRouteResponse> TransitionRoutes;

        [OutputConstructor]
        private GetTransitionRouteGroupResult(
            string displayName,

            string name,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3TransitionRouteResponse> transitionRoutes)
        {
            DisplayName = displayName;
            Name = name;
            TransitionRoutes = transitionRoutes;
        }
    }
}
