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
    /// Configuration options for Istio addon.
    /// </summary>
    [OutputType]
    public sealed class IstioConfigResponse
    {
        /// <summary>
        /// The specified Istio auth mode, either none, or mutual TLS.
        /// </summary>
        public readonly string Auth;
        /// <summary>
        /// Whether Istio is enabled for this cluster.
        /// </summary>
        public readonly bool Disabled;

        [OutputConstructor]
        private IstioConfigResponse(
            string auth,

            bool disabled)
        {
            Auth = auth;
            Disabled = disabled;
        }
    }
}
