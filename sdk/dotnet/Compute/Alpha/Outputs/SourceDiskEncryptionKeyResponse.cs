// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class SourceDiskEncryptionKeyResponse
    {
        /// <summary>
        /// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse DiskEncryptionKey;
        /// <summary>
        /// URL of the disk attached to the source instance. This can be a full or valid partial URL. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /disks/disk - projects/project/zones/zone/disks/disk - zones/zone/disks/disk 
        /// </summary>
        public readonly string SourceDisk;

        [OutputConstructor]
        private SourceDiskEncryptionKeyResponse(
            Outputs.CustomerEncryptionKeyResponse diskEncryptionKey,

            string sourceDisk)
        {
            DiskEncryptionKey = diskEncryptionKey;
            SourceDisk = sourceDisk;
        }
    }
}
