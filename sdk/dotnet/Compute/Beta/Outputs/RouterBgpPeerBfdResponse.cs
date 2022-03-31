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
    public sealed class RouterBgpPeerBfdResponse
    {
        /// <summary>
        /// The minimum interval, in milliseconds, between BFD control packets received from the peer router. The actual value is negotiated between the two routers and is equal to the greater of this value and the transmit interval of the other router. If set, this value must be between 1000 and 30000. The default is 1000.
        /// </summary>
        public readonly int MinReceiveInterval;
        /// <summary>
        /// The minimum interval, in milliseconds, between BFD control packets transmitted to the peer router. The actual value is negotiated between the two routers and is equal to the greater of this value and the corresponding receive interval of the other router. If set, this value must be between 1000 and 30000. The default is 1000.
        /// </summary>
        public readonly int MinTransmitInterval;
        /// <summary>
        /// The number of consecutive BFD packets that must be missed before BFD declares that a peer is unavailable. If set, the value must be a value between 5 and 16. The default is 5.
        /// </summary>
        public readonly int Multiplier;
        /// <summary>
        /// The BFD session initialization mode for this BGP peer. If set to ACTIVE, the Cloud Router will initiate the BFD session for this BGP peer. If set to PASSIVE, the Cloud Router will wait for the peer router to initiate the BFD session for this BGP peer. If set to DISABLED, BFD is disabled for this BGP peer. The default is DISABLED.
        /// </summary>
        public readonly string SessionInitializationMode;

        [OutputConstructor]
        private RouterBgpPeerBfdResponse(
            int minReceiveInterval,

            int minTransmitInterval,

            int multiplier,

            string sessionInitializationMode)
        {
            MinReceiveInterval = minReceiveInterval;
            MinTransmitInterval = minTransmitInterval;
            Multiplier = multiplier;
            SessionInitializationMode = sessionInitializationMode;
        }
    }
}
