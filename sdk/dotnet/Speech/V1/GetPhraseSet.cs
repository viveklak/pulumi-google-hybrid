// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.GoogleNative.Speech.V1
{
    public static class GetPhraseSet
    {
        /// <summary>
        /// Get a phrase set.
        /// </summary>
        public static Task<GetPhraseSetResult> InvokeAsync(GetPhraseSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPhraseSetResult>("google-native:speech/v1:getPhraseSet", args ?? new GetPhraseSetArgs(), options.WithVersion());

        /// <summary>
        /// Get a phrase set.
        /// </summary>
        public static Output<GetPhraseSetResult> Invoke(GetPhraseSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPhraseSetResult>("google-native:speech/v1:getPhraseSet", args ?? new GetPhraseSetInvokeArgs(), options.WithVersion());
    }


    public sealed class GetPhraseSetArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("phraseSetId", required: true)]
        public string PhraseSetId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetPhraseSetArgs()
        {
        }
    }

    public sealed class GetPhraseSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("phraseSetId", required: true)]
        public Input<string> PhraseSetId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetPhraseSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPhraseSetResult
    {
        /// <summary>
        /// Hint Boost. Positive value will increase the probability that a specific phrase will be recognized over other similar sounding phrases. The higher the boost, the higher the chance of false positive recognition as well. Negative boost values would correspond to anti-biasing. Anti-biasing is not enabled, so negative boost will simply be ignored. Though `boost` can accept a wide range of positive values, most use cases are best served with values between 0 (exclusive) and 20. We recommend using a binary search approach to finding the optimal value for your use case. Speech recognition will skip PhraseSets with a boost value of 0.
        /// </summary>
        public readonly double Boost;
        /// <summary>
        /// The resource name of the phrase set.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of word and phrases.
        /// </summary>
        public readonly ImmutableArray<Outputs.PhraseResponse> Phrases;

        [OutputConstructor]
        private GetPhraseSetResult(
            double boost,

            string name,

            ImmutableArray<Outputs.PhraseResponse> phrases)
        {
            Boost = boost;
            Name = name;
            Phrases = phrases;
        }
    }
}