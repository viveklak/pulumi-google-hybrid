// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    [OutputType]
    public sealed class IdentityConfigResponse
    {
        /// <summary>
        /// Map of user to service account.
        /// </summary>
        public readonly ImmutableDictionary<string, string> UserServiceAccountMapping;

        [OutputConstructor]
        private IdentityConfigResponse(ImmutableDictionary<string, string> userServiceAccountMapping)
        {
            UserServiceAccountMapping = userServiceAccountMapping;
        }
    }
}
