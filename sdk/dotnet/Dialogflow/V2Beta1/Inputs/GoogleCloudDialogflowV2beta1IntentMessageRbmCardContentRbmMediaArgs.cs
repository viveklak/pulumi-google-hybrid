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
    /// Rich Business Messaging (RBM) Media displayed in Cards The following media-types are currently supported: Image Types * image/jpeg * image/jpg' * image/gif * image/png Video Types * video/h263 * video/m4v * video/mp4 * video/mpeg * video/mpeg4 * video/webm
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageRbmCardContentRbmMediaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Publicly reachable URI of the file. The RBM platform determines the MIME type of the file from the content-type field in the HTTP headers when the platform fetches the file. The content-type field must be present and accurate in the HTTP response from the URL.
        /// </summary>
        [Input("fileUri", required: true)]
        public Input<string> FileUri { get; set; } = null!;

        /// <summary>
        /// Required for cards with vertical orientation. The height of the media within a rich card with a vertical layout. For a standalone card with horizontal layout, height is not customizable, and this field is ignored.
        /// </summary>
        [Input("height")]
        public Input<Pulumi.GoogleNative.Dialogflow.V2Beta1.GoogleCloudDialogflowV2beta1IntentMessageRbmCardContentRbmMediaHeight>? Height { get; set; }

        /// <summary>
        /// Optional. Publicly reachable URI of the thumbnail.If you don't provide a thumbnail URI, the RBM platform displays a blank placeholder thumbnail until the user's device downloads the file. Depending on the user's setting, the file may not download automatically and may require the user to tap a download button.
        /// </summary>
        [Input("thumbnailUri")]
        public Input<string>? ThumbnailUri { get; set; }

        public GoogleCloudDialogflowV2beta1IntentMessageRbmCardContentRbmMediaArgs()
        {
        }
    }
}
