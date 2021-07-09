// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Outputs
{

    [OutputType]
    public sealed class GoogleCloudDialogflowV2IntentMessageTableCardCellResponse
    {
        /// <summary>
        /// Text in this cell.
        /// </summary>
        public readonly string Text;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageTableCardCellResponse(string text)
        {
            Text = text;
        }
    }
}
