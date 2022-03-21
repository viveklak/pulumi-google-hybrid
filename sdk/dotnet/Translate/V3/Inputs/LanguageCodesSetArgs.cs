// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3.Inputs
{

    /// <summary>
    /// Used with equivalent term set glossaries.
    /// </summary>
    public sealed class LanguageCodesSetArgs : Pulumi.ResourceArgs
    {
        [Input("languageCodes")]
        private InputList<string>? _languageCodes;

        /// <summary>
        /// The BCP-47 language code(s) for terms defined in the glossary. All entries are unique. The list contains at least two entries. Expected to be an exact match for GlossaryTerm.language_code.
        /// </summary>
        public InputList<string> LanguageCodes
        {
            get => _languageCodes ?? (_languageCodes = new InputList<string>());
            set => _languageCodes = value;
        }

        public LanguageCodesSetArgs()
        {
        }
    }
}
