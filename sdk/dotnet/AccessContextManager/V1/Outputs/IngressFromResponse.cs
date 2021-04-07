// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.AccessContextManager.V1.Outputs
{

    [OutputType]
    public sealed class IngressFromResponse
    {
        /// <summary>
        /// A list of identities that are allowed access through this ingress policy. Should be in the format of email address. The email address should represent individual user or service account only.
        /// </summary>
        public readonly ImmutableArray<string> Identities;
        /// <summary>
        /// Specifies the type of identities that are allowed access from outside the perimeter. If left unspecified, then members of `identities` field will be allowed access.
        /// </summary>
        public readonly string IdentityType;
        /// <summary>
        /// Sources that this IngressPolicy authorizes access from.
        /// </summary>
        public readonly ImmutableArray<Outputs.IngressSourceResponse> Sources;

        [OutputConstructor]
        private IngressFromResponse(
            ImmutableArray<string> identities,

            string identityType,

            ImmutableArray<Outputs.IngressSourceResponse> sources)
        {
            Identities = identities;
            IdentityType = identityType;
            Sources = sources;
        }
    }
}