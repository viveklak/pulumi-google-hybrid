// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.PolicySimulator.V1.Outputs
{

    /// <summary>
    /// The configuration used for a Replay.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudPolicysimulatorV1ReplayConfigResponse
    {
        /// <summary>
        /// The logs to use as input for the Replay.
        /// </summary>
        public readonly string LogSource;
        /// <summary>
        /// A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
        /// </summary>
        public readonly ImmutableDictionary<string, string> PolicyOverlay;

        [OutputConstructor]
        private GoogleCloudPolicysimulatorV1ReplayConfigResponse(
            string logSource,

            ImmutableDictionary<string, string> policyOverlay)
        {
            LogSource = logSource;
            PolicyOverlay = policyOverlay;
        }
    }
}
