// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// The available logging options for a firewall rule.
    /// </summary>
    [OutputType]
    public sealed class FirewallLogConfigResponse
    {
        /// <summary>
        /// This field denotes whether to enable logging for a particular firewall rule.
        /// </summary>
        public readonly bool Enable;
        /// <summary>
        /// This field can only be specified for a particular firewall rule if logging is enabled for that rule. This field denotes whether to include or exclude metadata for firewall logs.
        /// </summary>
        public readonly string Metadata;

        [OutputConstructor]
        private FirewallLogConfigResponse(
            bool enable,

            string metadata)
        {
            Enable = enable;
            Metadata = metadata;
        }
    }
}
