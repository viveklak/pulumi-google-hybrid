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
    /// Configuration for Binary Authorization.
    /// </summary>
    [OutputType]
    public sealed class BinaryAuthorizationResponse
    {
        /// <summary>
        /// Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binauthz.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private BinaryAuthorizationResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
