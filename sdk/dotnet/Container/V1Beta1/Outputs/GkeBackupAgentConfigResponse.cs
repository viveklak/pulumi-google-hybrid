// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for the Backup for GKE Agent.
    /// </summary>
    [OutputType]
    public sealed class GkeBackupAgentConfigResponse
    {
        /// <summary>
        /// Whether the Backup for GKE agent is enabled for this cluster.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private GkeBackupAgentConfigResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
