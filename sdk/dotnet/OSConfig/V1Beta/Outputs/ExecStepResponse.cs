// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// A step that runs an executable for a PatchJob.
    /// </summary>
    [OutputType]
    public sealed class ExecStepResponse
    {
        /// <summary>
        /// The ExecStepConfig for all Linux VMs targeted by the PatchJob.
        /// </summary>
        public readonly Outputs.ExecStepConfigResponse LinuxExecStepConfig;
        /// <summary>
        /// The ExecStepConfig for all Windows VMs targeted by the PatchJob.
        /// </summary>
        public readonly Outputs.ExecStepConfigResponse WindowsExecStepConfig;

        [OutputConstructor]
        private ExecStepResponse(
            Outputs.ExecStepConfigResponse linuxExecStepConfig,

            Outputs.ExecStepConfigResponse windowsExecStepConfig)
        {
            LinuxExecStepConfig = linuxExecStepConfig;
            WindowsExecStepConfig = windowsExecStepConfig;
        }
    }
}
