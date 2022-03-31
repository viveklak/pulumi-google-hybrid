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
    /// The matrix of environments in which the test is to be executed.
    /// </summary>
    public sealed class EnvironmentMatrixArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A list of Android devices; the test will be run only on the specified devices.
        /// </summary>
        [Input("androidDeviceList")]
        public Input<Inputs.AndroidDeviceListArgs>? AndroidDeviceList { get; set; }

        /// <summary>
        /// A matrix of Android devices.
        /// </summary>
        [Input("androidMatrix")]
        public Input<Inputs.AndroidMatrixArgs>? AndroidMatrix { get; set; }

        /// <summary>
        /// A list of iOS devices.
        /// </summary>
        [Input("iosDeviceList")]
        public Input<Inputs.IosDeviceListArgs>? IosDeviceList { get; set; }

        public EnvironmentMatrixArgs()
        {
        }
    }
}
