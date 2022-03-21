// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1Alpha2.Outputs
{

    /// <summary>
    /// WorkerConfig defines the configuration to be used for a creating workers in the pool.
    /// </summary>
    [OutputType]
    public sealed class WorkerConfigResponse
    {
        /// <summary>
        /// Size of the disk attached to the worker, in GB. See https://cloud.google.com/compute/docs/disks/ If `0` is specified, Cloud Build will use a standard disk size.
        /// </summary>
        public readonly string DiskSizeGb;
        /// <summary>
        /// Machine Type of the worker, such as n1-standard-1. See https://cloud.google.com/compute/docs/machine-types. If left blank, Cloud Build will use a standard unspecified machine to create the worker pool.
        /// </summary>
        public readonly string MachineType;

        [OutputConstructor]
        private WorkerConfigResponse(
            string diskSizeGb,

            string machineType)
        {
            DiskSizeGb = diskSizeGb;
            MachineType = machineType;
        }
    }
}
