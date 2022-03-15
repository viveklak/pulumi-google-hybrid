// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// [Deprecated] Configuration for the mutual Tls mode for peer authentication. Configuration for the mutual Tls mode for peer authentication.
    /// </summary>
    [OutputType]
    public sealed class MutualTlsResponse
    {
        /// <summary>
        /// Specifies if the server TLS is configured to be strict or permissive. This field can be set to one of the following: STRICT: Client certificate must be presented, connection is in TLS. PERMISSIVE: Client certificate can be omitted, connection can be either plaintext or TLS.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private MutualTlsResponse(string mode)
        {
            Mode = mode;
        }
    }
}
