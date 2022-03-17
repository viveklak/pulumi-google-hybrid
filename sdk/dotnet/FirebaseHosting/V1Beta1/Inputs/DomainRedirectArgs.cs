// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseHosting.V1Beta1.Inputs
{

    /// <summary>
    /// Defines the behavior of a domain-level redirect. Domain redirects preserve the path of the redirect but replace the requested domain with the one specified in the redirect configuration.
    /// </summary>
    public sealed class DomainRedirectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name to redirect to.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The redirect status code.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.FirebaseHosting.V1Beta1.DomainRedirectType> Type { get; set; } = null!;

        public DomainRedirectArgs()
        {
        }
    }
}
