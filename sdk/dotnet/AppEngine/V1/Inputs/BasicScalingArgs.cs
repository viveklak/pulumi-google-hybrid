// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1.Inputs
{

    /// <summary>
    /// A service with basic scaling will create an instance when the application receives a request. The instance will be turned down when the app becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
    /// </summary>
    public sealed class BasicScalingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Duration of time after the last request that an instance must wait before the instance is shut down.
        /// </summary>
        [Input("idleTimeout")]
        public Input<string>? IdleTimeout { get; set; }

        /// <summary>
        /// Maximum number of instances to create for this version.
        /// </summary>
        [Input("maxInstances")]
        public Input<int>? MaxInstances { get; set; }

        public BasicScalingArgs()
        {
        }
    }
}
