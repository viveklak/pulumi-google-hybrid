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
    /// Configuration of logging on a NAT.
    /// </summary>
    [OutputType]
    public sealed class RouterNatLogConfigResponse
    {
        /// <summary>
        /// Indicates whether or not to export logs. This is false by default.
        /// </summary>
        public readonly bool Enable;
        /// <summary>
        /// Specify the desired filtering of logs on this NAT. If unspecified, logs are exported for all connections handled by this NAT. This option can take one of the following values: - ERRORS_ONLY: Export logs only for connection failures. - TRANSLATIONS_ONLY: Export logs only for successful connections. - ALL: Export logs for all connections, successful and unsuccessful. 
        /// </summary>
        public readonly string Filter;

        [OutputConstructor]
        private RouterNatLogConfigResponse(
            bool enable,

            string filter)
        {
            Enable = enable;
            Filter = filter;
        }
    }
}
