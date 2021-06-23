// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Outputs
{

    [OutputType]
    public sealed class SubordinateConfigChainResponse
    {
        /// <summary>
        /// Required. Expected to be in leaf-to-root order according to RFC 5246.
        /// </summary>
        public readonly ImmutableArray<string> PemCertificates;

        [OutputConstructor]
        private SubordinateConfigChainResponse(ImmutableArray<string> pemCertificates)
        {
            PemCertificates = pemCertificates;
        }
    }
}
