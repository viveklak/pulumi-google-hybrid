// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1.Inputs
{

    /// <summary>
    /// Network-level access control policy for the Airflow web server.
    /// </summary>
    public sealed class WebServerNetworkAccessControlArgs : Pulumi.ResourceArgs
    {
        [Input("allowedIpRanges")]
        private InputList<Inputs.AllowedIpRangeArgs>? _allowedIpRanges;

        /// <summary>
        /// A collection of allowed IP ranges with descriptions.
        /// </summary>
        public InputList<Inputs.AllowedIpRangeArgs> AllowedIpRanges
        {
            get => _allowedIpRanges ?? (_allowedIpRanges = new InputList<Inputs.AllowedIpRangeArgs>());
            set => _allowedIpRanges = value;
        }

        public WebServerNetworkAccessControlArgs()
        {
        }
    }
}
