// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for NodeLocal DNSCache
    /// </summary>
    [OutputType]
    public sealed class DnsCacheConfigResponse
    {
        /// <summary>
        /// Whether NodeLocal DNSCache is enabled for this cluster.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private DnsCacheConfigResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
