// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dns.V1.Inputs
{

    /// <summary>
    /// Contains information about Service Directory-backed zones.
    /// </summary>
    public sealed class ManagedZoneServiceDirectoryConfigArgs : Pulumi.ResourceArgs
    {
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Contains information about the namespace associated with the zone.
        /// </summary>
        [Input("namespace")]
        public Input<Inputs.ManagedZoneServiceDirectoryConfigNamespaceArgs>? Namespace { get; set; }

        public ManagedZoneServiceDirectoryConfigArgs()
        {
        }
    }
}