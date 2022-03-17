// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    /// <summary>
    /// Additional info about the select item for when it is triggered in a dialog.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2IntentMessageSelectItemInfoResponse
    {
        /// <summary>
        /// A unique key that will be sent back to the agent if this response is given.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Optional. A list of synonyms that can also be used to trigger this item in dialog.
        /// </summary>
        public readonly ImmutableArray<string> Synonyms;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageSelectItemInfoResponse(
            string key,

            ImmutableArray<string> synonyms)
        {
            Key = key;
            Synonyms = synonyms;
        }
    }
}
