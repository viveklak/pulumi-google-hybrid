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
    /// The text response message.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1ResponseMessageTextResponse
    {
        /// <summary>
        /// Whether the playback of this message can be interrupted by the end user's speech and the client can then starts the next Dialogflow request.
        /// </summary>
        public readonly bool AllowPlaybackInterruption;
        /// <summary>
        /// A collection of text responses.
        /// </summary>
        public readonly ImmutableArray<string> Text;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1ResponseMessageTextResponse(
            bool allowPlaybackInterruption,

            ImmutableArray<string> text)
        {
            AllowPlaybackInterruption = allowPlaybackInterruption;
            Text = text;
        }
    }
}
