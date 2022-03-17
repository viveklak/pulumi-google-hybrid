// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Outputs
{

    /// <summary>
    /// A single device file description.
    /// </summary>
    [OutputType]
    public sealed class DeviceFileResponse
    {
        /// <summary>
        /// A reference to an opaque binary blob file.
        /// </summary>
        public readonly Outputs.ObbFileResponse ObbFile;
        /// <summary>
        /// A reference to a regular file.
        /// </summary>
        public readonly Outputs.RegularFileResponse RegularFile;

        [OutputConstructor]
        private DeviceFileResponse(
            Outputs.ObbFileResponse obbFile,

            Outputs.RegularFileResponse regularFile)
        {
            ObbFile = obbFile;
            RegularFile = regularFile;
        }
    }
}
