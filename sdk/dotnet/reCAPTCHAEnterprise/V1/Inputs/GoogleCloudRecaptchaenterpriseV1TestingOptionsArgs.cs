// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.Inputs
{

    /// <summary>
    /// Options for user acceptance testing.
    /// </summary>
    public sealed class GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if CHALLENGE.
        /// </summary>
        [Input("testingChallenge")]
        public Input<Pulumi.GoogleNative.reCAPTCHAEnterprise.V1.GoogleCloudRecaptchaenterpriseV1TestingOptionsTestingChallenge>? TestingChallenge { get; set; }

        /// <summary>
        /// All assessments for this Key will return this score. Must be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive.
        /// </summary>
        [Input("testingScore")]
        public Input<double>? TestingScore { get; set; }

        public GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs()
        {
        }
    }
}
