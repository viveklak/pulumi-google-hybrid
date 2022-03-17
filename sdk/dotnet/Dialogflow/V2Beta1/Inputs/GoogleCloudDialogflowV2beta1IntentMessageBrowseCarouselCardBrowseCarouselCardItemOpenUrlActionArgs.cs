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
    /// Actions on Google action to open a given url.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardBrowseCarouselCardItemOpenUrlActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Optional. Specifies the type of viewer that is used when opening the URL. Defaults to opening via web browser.
        /// </summary>
        [Input("urlTypeHint")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2Beta1.GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardBrowseCarouselCardItemOpenUrlActionUrlTypeHint>? UrlTypeHint { get; set; }

        public GoogleCloudDialogflowV2beta1IntentMessageBrowseCarouselCardBrowseCarouselCardItemOpenUrlActionArgs()
        {
        }
    }
}
