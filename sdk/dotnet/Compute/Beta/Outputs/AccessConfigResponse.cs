// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class AccessConfigResponse
    {
        /// <summary>
        /// Type of the resource. Always compute#accessConfig for access configs.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of this access configuration. The default and recommended name is External NAT, but you can use any arbitrary string, such as My external IP or Network Access.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An external IP address associated with this instance. Specify an unused static external IP address available to the project or leave this field undefined to use an IP from a shared ephemeral IP address pool. If you specify a static external IP address, it must live in the same region as the zone of the instance.
        /// </summary>
        public readonly string NatIP;
        /// <summary>
        /// This signifies the networking tier used for configuring this access configuration and can only take the following values: PREMIUM, STANDARD.
        /// 
        /// If an AccessConfig is specified without a valid external IP address, an ephemeral IP will be created with this networkTier.
        /// 
        /// If an AccessConfig with a valid external IP address is specified, it must match that of the networkTier associated with the Address resource owning that IP.
        /// </summary>
        public readonly string NetworkTier;
        /// <summary>
        /// The DNS domain name for the public PTR record. You can set this field only if the `setPublicPtr` field is enabled.
        /// </summary>
        public readonly string PublicPtrDomainName;
        /// <summary>
        /// Specifies whether a public DNS 'PTR' record should be created to map the external IP address of the instance to a DNS domain name.
        /// </summary>
        public readonly bool SetPublicPtr;
        /// <summary>
        /// The type of configuration. The default and only option is ONE_TO_ONE_NAT.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AccessConfigResponse(
            string kind,

            string name,

            string natIP,

            string networkTier,

            string publicPtrDomainName,

            bool setPublicPtr,

            string type)
        {
            Kind = kind;
            Name = name;
            NatIP = natIP;
            NetworkTier = networkTier;
            PublicPtrDomainName = publicPtrDomainName;
            SetPublicPtr = setPublicPtr;
            Type = type;
        }
    }
}
