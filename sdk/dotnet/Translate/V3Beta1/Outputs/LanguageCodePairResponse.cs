// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3Beta1.Outputs
{

    [OutputType]
    public sealed class LanguageCodePairResponse
    {
        /// <summary>
        /// The BCP-47 language code of the input text, for example, "en-US". Expected to be an exact match for GlossaryTerm.language_code.
        /// </summary>
        public readonly string SourceLanguageCode;
        /// <summary>
        /// The BCP-47 language code for translation output, for example, "zh-CN". Expected to be an exact match for GlossaryTerm.language_code.
        /// </summary>
        public readonly string TargetLanguageCode;

        [OutputConstructor]
        private LanguageCodePairResponse(
            string sourceLanguageCode,

            string targetLanguageCode)
        {
            SourceLanguageCode = sourceLanguageCode;
            TargetLanguageCode = targetLanguageCode;
        }
    }
}
