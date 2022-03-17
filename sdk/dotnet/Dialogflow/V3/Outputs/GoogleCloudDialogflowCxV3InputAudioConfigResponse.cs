// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// Instructs the speech recognizer on how to process the audio content.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3InputAudioConfigResponse
    {
        /// <summary>
        /// Audio encoding of the audio content to process.
        /// </summary>
        public readonly string AudioEncoding;
        /// <summary>
        /// Optional. If `true`, Dialogflow returns SpeechWordInfo in StreamingRecognitionResult with information about the recognized speech words, e.g. start and end time offsets. If false or unspecified, Speech doesn't return any word-level information.
        /// </summary>
        public readonly bool EnableWordInfo;
        /// <summary>
        /// Optional. Which Speech model to select for the given request. Select the model best suited to your domain to get best results. If a model is not explicitly specified, then we auto-select a model based on the parameters in the InputAudioConfig. If enhanced speech model is enabled for the agent and an enhanced version of the specified model for the language does not exist, then the speech is recognized using the standard version of the specified model. Refer to [Cloud Speech API documentation](https://cloud.google.com/speech-to-text/docs/basics#select-model) for more details.
        /// </summary>
        public readonly string Model;
        /// <summary>
        /// Optional. Which variant of the Speech model to use.
        /// </summary>
        public readonly string ModelVariant;
        /// <summary>
        /// Optional. A list of strings containing words and phrases that the speech recognizer should recognize with higher likelihood. See [the Cloud Speech documentation](https://cloud.google.com/speech-to-text/docs/basics#phrase-hints) for more details.
        /// </summary>
        public readonly ImmutableArray<string> PhraseHints;
        /// <summary>
        /// Sample rate (in Hertz) of the audio content sent in the query. Refer to [Cloud Speech API documentation](https://cloud.google.com/speech-to-text/docs/basics) for more details.
        /// </summary>
        public readonly int SampleRateHertz;
        /// <summary>
        /// Optional. If `false` (default), recognition does not cease until the client closes the stream. If `true`, the recognizer will detect a single spoken utterance in input audio. Recognition ceases when it detects the audio's voice has stopped or paused. In this case, once a detected intent is received, the client should close the stream and start a new request with a new stream as needed. Note: This setting is relevant only for streaming methods.
        /// </summary>
        public readonly bool SingleUtterance;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3InputAudioConfigResponse(
            string audioEncoding,

            bool enableWordInfo,

            string model,

            string modelVariant,

            ImmutableArray<string> phraseHints,

            int sampleRateHertz,

            bool singleUtterance)
        {
            AudioEncoding = audioEncoding;
            EnableWordInfo = enableWordInfo;
            Model = model;
            ModelVariant = modelVariant;
            PhraseHints = phraseHints;
            SampleRateHertz = sampleRateHertz;
            SingleUtterance = singleUtterance;
        }
    }
}
