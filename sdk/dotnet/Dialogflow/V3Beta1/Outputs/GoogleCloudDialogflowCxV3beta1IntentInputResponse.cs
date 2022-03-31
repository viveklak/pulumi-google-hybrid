// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    /// <summary>
    /// Represents the intent to trigger programmatically rather than as a result of natural language processing.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1IntentInputResponse
    {
        /// <summary>
        /// The unique identifier of the intent. Format: `projects//locations//agents//intents/`.
        /// </summary>
        public readonly string Intent;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1IntentInputResponse(string intent)
        {
            Intent = intent;
        }
    }
}
