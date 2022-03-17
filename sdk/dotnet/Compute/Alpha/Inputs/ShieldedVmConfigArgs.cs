// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// A set of Shielded VM options.
    /// </summary>
    public sealed class ShieldedVmConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether the instance has integrity monitoring enabled.
        /// </summary>
        [Input("enableIntegrityMonitoring")]
        public Input<bool>? EnableIntegrityMonitoring { get; set; }

        /// <summary>
        /// Defines whether the instance has Secure Boot enabled.
        /// </summary>
        [Input("enableSecureBoot")]
        public Input<bool>? EnableSecureBoot { get; set; }

        /// <summary>
        /// Defines whether the instance has the vTPM enabled.
        /// </summary>
        [Input("enableVtpm")]
        public Input<bool>? EnableVtpm { get; set; }

        public ShieldedVmConfigArgs()
        {
        }
    }
}
