// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.FirebaseHosting.V1Beta1.Outputs
{

    [OutputType]
    public sealed class DomainRedirectResponse
    {
        /// <summary>
        /// Required. The domain name to redirect to.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// Required. The redirect status code.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DomainRedirectResponse(
            string domainName,

            string type)
        {
            DomainName = domainName;
            Type = type;
        }
    }
}