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
    /// InputDataset used to create model or do evaluation. NextID:5
    /// </summary>
    public sealed class GoogleCloudDialogflowV2InputDatasetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ConversationDataset resource name. Format: `projects//locations//conversationDatasets/`
        /// </summary>
        [Input("dataset", required: true)]
        public Input<string> Dataset { get; set; } = null!;

        public GoogleCloudDialogflowV2InputDatasetArgs()
        {
        }
    }
}
