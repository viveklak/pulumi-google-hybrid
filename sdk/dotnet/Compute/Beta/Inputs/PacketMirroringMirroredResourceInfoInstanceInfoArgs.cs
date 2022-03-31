// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class PacketMirroringMirroredResourceInfoInstanceInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource URL to the virtual machine instance which is being mirrored.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public PacketMirroringMirroredResourceInfoInstanceInfoArgs()
        {
        }
    }
}
