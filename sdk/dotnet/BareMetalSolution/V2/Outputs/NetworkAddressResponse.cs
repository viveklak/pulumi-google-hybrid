// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Outputs
{

    /// <summary>
    /// A network.
    /// </summary>
    [OutputType]
    public sealed class NetworkAddressResponse
    {
        /// <summary>
        /// IPv4 address to be assigned to the server.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// Name of the existing network to use.
        /// </summary>
        public readonly string ExistingNetworkId;
        /// <summary>
        /// Id of the network to use, within the same ProvisioningConfig request.
        /// </summary>
        public readonly string NetworkId;

        [OutputConstructor]
        private NetworkAddressResponse(
            string address,

            string existingNetworkId,

            string networkId)
        {
            Address = address;
            ExistingNetworkId = existingNetworkId;
            NetworkId = networkId;
        }
    }
}
