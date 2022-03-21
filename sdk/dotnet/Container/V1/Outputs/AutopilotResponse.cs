// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// Autopilot is the configuration for Autopilot settings on the cluster.
    /// </summary>
    [OutputType]
    public sealed class AutopilotResponse
    {
        /// <summary>
        /// Enable Autopilot
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private AutopilotResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
