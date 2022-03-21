// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Specifies the config of disk options for a group of VM instances.
    /// </summary>
    public sealed class DiskConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Size in GB of the boot disk (default is 500GB).
        /// </summary>
        [Input("bootDiskSizeGb")]
        public Input<int>? BootDiskSizeGb { get; set; }

        /// <summary>
        /// Optional. Type of the boot disk (default is "pd-standard"). Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive), "pd-ssd" (Persistent Disk Solid State Drive), or "pd-standard" (Persistent Disk Hard Disk Drive). See Disk types (https://cloud.google.com/compute/docs/disks#disk-types).
        /// </summary>
        [Input("bootDiskType")]
        public Input<string>? BootDiskType { get; set; }

        /// <summary>
        /// Optional. Interface type of local SSDs (default is "scsi"). Valid values: "scsi" (Small Computer System Interface), "nvme" (Non-Volatile Memory Express). See local SSD performance (https://cloud.google.com/compute/docs/disks/local-ssd#performance).
        /// </summary>
        [Input("localSsdInterface")]
        public Input<string>? LocalSsdInterface { get; set; }

        /// <summary>
        /// Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and HDFS (https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries.
        /// </summary>
        [Input("numLocalSsds")]
        public Input<int>? NumLocalSsds { get; set; }

        public DiskConfigArgs()
        {
        }
    }
}
