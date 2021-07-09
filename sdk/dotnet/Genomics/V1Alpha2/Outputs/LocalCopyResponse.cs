// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Genomics.V1Alpha2.Outputs
{

    [OutputType]
    public sealed class LocalCopyResponse
    {
        /// <summary>
        /// The name of the disk where this parameter is located. Can be the name of one of the disks specified in the Resources field, or "boot", which represents the Docker instance's boot disk and has a mount point of `/`.
        /// </summary>
        public readonly string Disk;
        /// <summary>
        /// The path within the user's docker container where this input should be localized to and from, relative to the specified disk's mount point. For example: file.txt,
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private LocalCopyResponse(
            string disk,

            string path)
        {
            Disk = disk;
            Path = path;
        }
    }
}
