// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class SecurityPolicyRuleHttpHeaderActionHttpHeaderOptionResponse
    {
        /// <summary>
        /// The name of the header to set.
        /// </summary>
        public readonly string HeaderName;
        /// <summary>
        /// The value to set the named header to.
        /// </summary>
        public readonly string HeaderValue;

        [OutputConstructor]
        private SecurityPolicyRuleHttpHeaderActionHttpHeaderOptionResponse(
            string headerName,

            string headerValue)
        {
            HeaderName = headerName;
            HeaderValue = headerValue;
        }
    }
}
