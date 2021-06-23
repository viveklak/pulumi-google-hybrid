// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class ImageRawDiskResponse
    {
        /// <summary>
        /// The format used to encode and transmit the block device, which should be TAR. This is just a container and transmission format and not a runtime format. Provided by the client when the disk image is created.
        /// </summary>
        public readonly string ContainerType;
        /// <summary>
        /// The full Google Cloud Storage URL where the disk image is stored. In order to create an image, you must provide the full or partial URL of one of the following: - The rawDisk.source URL - The sourceDisk URL - The sourceImage URL - The sourceSnapshot URL 
        /// </summary>
        public readonly string Source;

        [OutputConstructor]
        private ImageRawDiskResponse(
            string containerType,

            string source)
        {
            ContainerType = containerType;
            Source = source;
        }
    }
}
