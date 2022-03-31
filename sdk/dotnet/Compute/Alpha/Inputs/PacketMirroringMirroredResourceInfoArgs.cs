// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class PacketMirroringMirroredResourceInfoArgs : Pulumi.ResourceArgs
    {
        [Input("instances")]
        private InputList<Inputs.PacketMirroringMirroredResourceInfoInstanceInfoArgs>? _instances;

        /// <summary>
        /// A set of virtual machine instances that are being mirrored. They must live in zones contained in the same region as this packetMirroring. Note that this config will apply only to those network interfaces of the Instances that belong to the network specified in this packetMirroring. You may specify a maximum of 50 Instances.
        /// </summary>
        public InputList<Inputs.PacketMirroringMirroredResourceInfoInstanceInfoArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.PacketMirroringMirroredResourceInfoInstanceInfoArgs>());
            set => _instances = value;
        }

        [Input("subnetworks")]
        private InputList<Inputs.PacketMirroringMirroredResourceInfoSubnetInfoArgs>? _subnetworks;

        /// <summary>
        /// A set of subnetworks for which traffic from/to all VM instances will be mirrored. They must live in the same region as this packetMirroring. You may specify a maximum of 5 subnetworks.
        /// </summary>
        public InputList<Inputs.PacketMirroringMirroredResourceInfoSubnetInfoArgs> Subnetworks
        {
            get => _subnetworks ?? (_subnetworks = new InputList<Inputs.PacketMirroringMirroredResourceInfoSubnetInfoArgs>());
            set => _subnetworks = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of mirrored tags. Traffic from/to all VM instances that have one or more of these tags will be mirrored.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public PacketMirroringMirroredResourceInfoArgs()
        {
        }
    }
}
