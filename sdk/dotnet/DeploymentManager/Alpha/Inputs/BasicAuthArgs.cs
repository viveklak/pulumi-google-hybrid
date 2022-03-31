// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha.Inputs
{

    /// <summary>
    /// Basic Auth used as a credential.
    /// </summary>
    public sealed class BasicAuthArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("user")]
        public Input<string>? User { get; set; }

        public BasicAuthArgs()
        {
        }
    }
}
