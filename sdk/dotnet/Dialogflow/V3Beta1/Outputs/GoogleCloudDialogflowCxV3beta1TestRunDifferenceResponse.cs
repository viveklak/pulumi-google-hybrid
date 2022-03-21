// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Outputs
{

    /// <summary>
    /// The description of differences between original and replayed agent output.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3beta1TestRunDifferenceResponse
    {
        /// <summary>
        /// A description of the diff, showing the actual output vs expected output.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The type of diff.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3beta1TestRunDifferenceResponse(
            string description,

            string type)
        {
            Description = description;
            Type = type;
        }
    }
}
