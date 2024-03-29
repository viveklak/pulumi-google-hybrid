// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// A iOS mobile test specification
    /// </summary>
    public sealed class IosTestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information about the application under test.
        /// </summary>
        [Input("iosAppInfo")]
        public Input<Inputs.IosAppInfoArgs>? IosAppInfo { get; set; }

        /// <summary>
        /// An iOS Robo test.
        /// </summary>
        [Input("iosRoboTest")]
        public Input<Inputs.IosRoboTestArgs>? IosRoboTest { get; set; }

        /// <summary>
        /// An iOS test loop.
        /// </summary>
        [Input("iosTestLoop")]
        public Input<Inputs.IosTestLoopArgs>? IosTestLoop { get; set; }

        /// <summary>
        /// An iOS XCTest.
        /// </summary>
        [Input("iosXcTest")]
        public Input<Inputs.IosXcTestArgs>? IosXcTest { get; set; }

        /// <summary>
        /// Max time a test is allowed to run before it is automatically cancelled.
        /// </summary>
        [Input("testTimeout")]
        public Input<Inputs.DurationArgs>? TestTimeout { get; set; }

        public IosTestArgs()
        {
        }
    }
}
