// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Outputs
{

    /// <summary>
    /// The details about how to run the execution.
    /// </summary>
    [OutputType]
    public sealed class SpecificationResponse
    {
        /// <summary>
        /// An Android mobile test execution specification.
        /// </summary>
        public readonly Outputs.AndroidTestResponse AndroidTest;
        /// <summary>
        /// An iOS mobile test execution specification.
        /// </summary>
        public readonly Outputs.IosTestResponse IosTest;

        [OutputConstructor]
        private SpecificationResponse(
            Outputs.AndroidTestResponse androidTest,

            Outputs.IosTestResponse iosTest)
        {
            AndroidTest = androidTest;
            IosTest = iosTest;
        }
    }
}
