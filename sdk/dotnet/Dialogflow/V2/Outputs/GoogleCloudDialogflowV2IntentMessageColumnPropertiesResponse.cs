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
    public sealed class GoogleCloudDialogflowV2IntentMessageColumnPropertiesResponse
    {
        /// <summary>
        /// Column heading.
        /// </summary>
        public readonly string Header;
        /// <summary>
        /// Optional. Defines text alignment for all cells in this column.
        /// </summary>
        public readonly string HorizontalAlignment;

        [OutputConstructor]
        private GoogleCloudDialogflowV2IntentMessageColumnPropertiesResponse(
            string header,

            string horizontalAlignment)
        {
            Header = header;
            HorizontalAlignment = horizontalAlignment;
        }
    }
}
