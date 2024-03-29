// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Outputs
{

    /// <summary>
    /// A point in a conversation that marks the start or the end of an annotation.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1AnnotationBoundaryResponse
    {
        /// <summary>
        /// The index in the sequence of transcribed pieces of the conversation where the boundary is located. This index starts at zero.
        /// </summary>
        public readonly int TranscriptIndex;
        /// <summary>
        /// The word index of this boundary with respect to the first word in the transcript piece. This index starts at zero.
        /// </summary>
        public readonly int WordIndex;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1AnnotationBoundaryResponse(
            int transcriptIndex,

            int wordIndex)
        {
            TranscriptIndex = transcriptIndex;
            WordIndex = wordIndex;
        }
    }
}
