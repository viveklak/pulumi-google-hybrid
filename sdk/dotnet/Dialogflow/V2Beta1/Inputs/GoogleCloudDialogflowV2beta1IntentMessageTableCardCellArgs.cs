// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// Cell of TableCardRow.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageTableCardCellArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text in this cell.
        /// </summary>
        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1IntentMessageTableCardCellArgs()
        {
        }
    }
}
