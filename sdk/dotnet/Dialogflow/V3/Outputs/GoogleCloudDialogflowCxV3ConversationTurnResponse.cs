// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dialogflow.V3.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3ConversationTurnResponse
    {
        /// <summary>
        /// The user input.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3ConversationTurnUserInputResponse UserInput;
        /// <summary>
        /// The virtual agent output.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3ConversationTurnVirtualAgentOutputResponse VirtualAgentOutput;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3ConversationTurnResponse(
            Outputs.GoogleCloudDialogflowCxV3ConversationTurnUserInputResponse userInput,

            Outputs.GoogleCloudDialogflowCxV3ConversationTurnVirtualAgentOutputResponse virtualAgentOutput)
        {
            UserInput = userInput;
            VirtualAgentOutput = virtualAgentOutput;
        }
    }
}