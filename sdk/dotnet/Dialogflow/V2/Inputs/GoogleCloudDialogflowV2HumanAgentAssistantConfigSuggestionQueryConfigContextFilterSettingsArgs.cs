// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Inputs
{

    /// <summary>
    /// Settings that determine how to filter recent conversation context when generating suggestions.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigContextFilterSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to true, the last message from virtual agent (hand off message) and the message before it (trigger message of hand off) are dropped.
        /// </summary>
        [Input("dropHandoffMessages")]
        public Input<bool>? DropHandoffMessages { get; set; }

        /// <summary>
        /// If set to true, all messages from ivr stage are dropped.
        /// </summary>
        [Input("dropIvrMessages")]
        public Input<bool>? DropIvrMessages { get; set; }

        /// <summary>
        /// If set to true, all messages from virtual agent are dropped.
        /// </summary>
        [Input("dropVirtualAgentMessages")]
        public Input<bool>? DropVirtualAgentMessages { get; set; }

        public GoogleCloudDialogflowV2HumanAgentAssistantConfigSuggestionQueryConfigContextFilterSettingsArgs()
        {
        }
    }
}
