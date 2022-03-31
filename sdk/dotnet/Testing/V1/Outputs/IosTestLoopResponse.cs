// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Outputs
{

    /// <summary>
    /// A test of an iOS application that implements one or more game loop scenarios. This test type accepts an archived application (.ipa file) and a list of integer scenarios that will be executed on the app sequentially.
    /// </summary>
    [OutputType]
    public sealed class IosTestLoopResponse
    {
        /// <summary>
        /// The bundle id for the application under test.
        /// </summary>
        public readonly string AppBundleId;
        /// <summary>
        /// The .ipa of the application to test.
        /// </summary>
        public readonly Outputs.FileReferenceResponse AppIpa;
        /// <summary>
        /// The list of scenarios that should be run during the test. Defaults to the single scenario 0 if unspecified.
        /// </summary>
        public readonly ImmutableArray<int> Scenarios;

        [OutputConstructor]
        private IosTestLoopResponse(
            string appBundleId,

            Outputs.FileReferenceResponse appIpa,

            ImmutableArray<int> scenarios)
        {
            AppBundleId = appBundleId;
            AppIpa = appIpa;
            Scenarios = scenarios;
        }
    }
}
