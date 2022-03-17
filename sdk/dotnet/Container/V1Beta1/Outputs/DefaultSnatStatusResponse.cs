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
    /// DefaultSnatStatus contains the desired state of whether default sNAT should be disabled on the cluster.
    /// </summary>
    [OutputType]
    public sealed class DefaultSnatStatusResponse
    {
        /// <summary>
        /// Disables cluster default sNAT rules.
        /// </summary>
        public readonly bool Disabled;

        [OutputConstructor]
        private DefaultSnatStatusResponse(bool disabled)
        {
            Disabled = disabled;
        }
    }
}
