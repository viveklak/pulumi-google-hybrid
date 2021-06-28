// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    [OutputType]
    public sealed class PublicDelegatedPrefixPublicDelegatedSubPrefixResponse
    {
        /// <summary>
        /// Name of the project scoping this PublicDelegatedSubPrefix.
        /// </summary>
        public readonly string DelegateeProject;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The IPv4 address range, in CIDR format, represented by this sub public delegated prefix.
        /// </summary>
        public readonly string IpCidrRange;
        /// <summary>
        /// Whether the sub prefix is delegated to create Address resources in the delegatee project.
        /// </summary>
        public readonly bool IsAddress;
        /// <summary>
        /// The name of the sub public delegated prefix.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The region of the sub public delegated prefix if it is regional. If absent, the sub prefix is global.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The status of the sub public delegated prefix.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private PublicDelegatedPrefixPublicDelegatedSubPrefixResponse(
            string delegateeProject,

            string description,

            string ipCidrRange,

            bool isAddress,

            string name,

            string region,

            string status)
        {
            DelegateeProject = delegateeProject;
            Description = description;
            IpCidrRange = ipCidrRange;
            IsAddress = isAddress;
            Name = name;
            Region = region;
            Status = status;
        }
    }
}
