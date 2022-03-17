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
    /// A file or directory to install on the device before the test starts.
    /// </summary>
    public sealed class IosDeviceFileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The bundle id of the app where this file lives. iOS apps sandbox their own filesystem, so app files must specify which app installed on the device.
        /// </summary>
        [Input("bundleId")]
        public Input<string>? BundleId { get; set; }

        /// <summary>
        /// The source file
        /// </summary>
        [Input("content")]
        public Input<Inputs.FileReferenceArgs>? Content { get; set; }

        /// <summary>
        /// Location of the file on the device, inside the app's sandboxed filesystem
        /// </summary>
        [Input("devicePath")]
        public Input<string>? DevicePath { get; set; }

        public IosDeviceFileArgs()
        {
        }
    }
}
