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
    /// An excluded entity phrase that should not be matched.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3EntityTypeExcludedPhraseResponse
    {
        /// <summary>
        /// The word or phrase to be excluded.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3EntityTypeExcludedPhraseResponse(string value)
        {
            Value = value;
        }
    }
}
