// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1Beta1.Outputs
{

    [OutputType]
    public sealed class TextStreamResponse
    {
        /// <summary>
        /// The codec for this text stream. The default is `"webvtt"`. Supported text codecs: - 'srt' - 'ttml' - 'cea608' - 'cea708' - 'webvtt'
        /// </summary>
        public readonly string Codec;
        /// <summary>
        /// The BCP-47 language code, such as `"en-US"` or `"sr-Latn"`. For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        /// </summary>
        public readonly string LanguageCode;
        /// <summary>
        /// The mapping for the `Job.edit_list` atoms with text `EditAtom.inputs`.
        /// </summary>
        public readonly ImmutableArray<Outputs.TextAtomResponse> Mapping;

        [OutputConstructor]
        private TextStreamResponse(
            string codec,

            string languageCode,

            ImmutableArray<Outputs.TextAtomResponse> mapping)
        {
            Codec = codec;
            LanguageCode = languageCode;
            Mapping = mapping;
        }
    }
}
