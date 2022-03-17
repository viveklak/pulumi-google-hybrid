// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Inputs
{

    /// <summary>
    /// A single device file description.
    /// </summary>
    public sealed class DeviceFileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A reference to an opaque binary blob file.
        /// </summary>
        [Input("obbFile")]
        public Input<Inputs.ObbFileArgs>? ObbFile { get; set; }

        /// <summary>
        /// A reference to a regular file.
        /// </summary>
        [Input("regularFile")]
        public Input<Inputs.RegularFileArgs>? RegularFile { get; set; }

        public DeviceFileArgs()
        {
        }
    }
}
